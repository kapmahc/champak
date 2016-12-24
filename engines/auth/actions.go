package auth

import (
	"crypto/aes"
	"fmt"
	"html/template"
	"path"
	"time"

	"github.com/SermoDigital/jose/crypto"
	log "github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
	"github.com/urfave/cli"
	"golang.org/x/text/language"
)

type injectLogger struct {
}

func (p *injectLogger) Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

// Inject inject contatiner
type Inject struct {
	I18n *web.I18n `inject:""`
	Db   *gorm.DB  `inject:""`
}

func (p *Inject) t(lng, code string, args ...interface{}) string {
	return p.I18n.T(lng, code, args...)
}

func (p *Inject) links(loc string) []Link {
	var items []Link
	if err := p.Db.
		Select([]string{"label", "href"}).
		Where("loc = ?", loc).
		Order("sort_order ASC").
		Find(&items).Error; err != nil {
		log.Error(err)
	}
	return items
}

func (p *Inject) cards(loc string) []Card {
	var items []Card
	if err := p.Db.
		Select([]string{"title", "summary", "logo", "href"}).
		Where("loc = ?", loc).
		Order("sort_order ASC").
		Find(&items).Error; err != nil {
		log.Error(err)
	}
	return items
}

// Action  action
func (p *Inject) Action(fn func(*cli.Context) error) cli.ActionFunc {
	return Action(func(ctx *cli.Context) error {
		inj := inject.Graph{Logger: &injectLogger{}}
		// ----------------
		db, err := OpenDatabase()
		if err != nil {
			return err
		}
		rep := OpenRedis()
		cip, err := aes.NewCipher([]byte(viper.GetString("secrets.aes")))
		if err != nil {
			return err
		}

		var langs []language.Tag
		for _, l := range viper.GetStringSlice("languages") {
			if lng, err := language.Parse(l); err == nil {
				langs = append(langs, lng)
			} else {
				return err
			}
		}

		rmq := viper.GetStringMap("rabbitmq")
		rdr := render.New(render.Options{
			Directory:  path.Join("themes", viper.GetString("server.theme"), "views"),
			Layout:     "application",
			Extensions: []string{".html"},
			Funcs: []template.FuncMap{
				{
					"t":     p.t,
					"links": p.links,
					"cards": p.cards,
					"fmt":   fmt.Sprintf,
					"eq": func(arg1, arg2 interface{}) bool {
						return arg1 == arg2
					},
					"str2htm": func(s string) template.HTML {
						return template.HTML(s)
					},
					"dtf": func(t time.Time) string {
						return t.Format("Mon Jan _2 15:04:05 2006")
					},
				},
			},
			IndentJSON:    !IsProduction(),
			IndentXML:     !IsProduction(),
			IsDevelopment: !IsProduction(),
		})

		if err := inj.Provide(
			&inject.Object{Value: rdr},
			&inject.Object{Value: db},
			&inject.Object{Value: rep},
			&inject.Object{Value: cip},
			&inject.Object{Value: cip, Name: "aes.cip"},
			&inject.Object{Value: &web.I18n{Items: make(map[string]map[string]string)}},
			&inject.Object{
				Value: fmt.Sprintf(
					"amqp://%v:%v@%v:%v/%v",
					rmq["user"],
					rmq["password"],
					rmq["host"],
					rmq["port"],
					rmq["virtual"],
				),
				Name: "rabbitmq.url",
			},
			&inject.Object{Value: []byte(viper.GetString("secrets.hmac")), Name: "hmac.key"},
			&inject.Object{Value: []byte(viper.GetString("secrets.jwt")), Name: "jwt.key"},
			&inject.Object{Value: viper.GetString("app.name"), Name: "namespace"},
			&inject.Object{Value: language.NewMatcher(langs), Name: "language.matcher"},
			&inject.Object{Value: crypto.SigningMethodHS512, Name: "jwt.method"},
		); err != nil {
			return err
		}
		// -----------------

		web.Loop(func(en web.Engine) error {
			if err := en.Map(&inj); err != nil {
				return err
			}
			return inj.Provide(&inject.Object{Value: en})
		})
		if err := inj.Populate(); err != nil {
			return err
		}
		return fn(ctx)
	})
}

// Action cfg action
func Action(f cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		return f(c)
	}
}
