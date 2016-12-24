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
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
	"github.com/urfave/cli"
)

type injectLogger struct {
}

func (p *injectLogger) Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

// InjectAction  inject action
func InjectAction(fn func(*cli.Context) error) cli.ActionFunc {
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

		rmq := viper.GetStringMap("rabbitmq")
		rdr := render.New(render.Options{
			Directory:  path.Join("themes", viper.GetString("server.theme"), "views"),
			Layout:     "application",
			Extensions: []string{".html"},
			Funcs: []template.FuncMap{
				{
					"t": func(lng, code string, args ...interface{}) string {
						var l web.Locale
						if err := db.
							Select("message").
							Where("lang = ? AND code = ?", lng, code).
							First(&l).Error; err != nil {
							log.Error(err)
							return code
						}

						return fmt.Sprintf(l.Message, args...)
					},
					"cards": func(loc string) []Card {
						var items []Card
						if err := db.
							Select([]string{"title", "summary", "logo", "href"}).
							Where("loc = ?", loc).
							Order("sort_order ASC").
							Find(&items).Error; err != nil {
							log.Error(err)
						}
						return items
					},
					"links": func(loc string) []Link {
						var items []Link
						if err := db.
							Select([]string{"label", "href"}).
							Where("loc = ?", loc).
							Order("sort_order ASC").
							Find(&items).Error; err != nil {
							log.Error(err)
						}
						return items
					},
					"fmt": fmt.Sprintf,
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
