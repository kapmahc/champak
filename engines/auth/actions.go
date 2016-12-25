package auth

import (
	"fmt"
	"html/template"
	"path"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	"github.com/kapmahc/champak/web"
	"github.com/kapmahc/champak/web/cache"
	"github.com/kapmahc/champak/web/cache/redis"
	_crypto "github.com/kapmahc/champak/web/crypto"
	"github.com/kapmahc/champak/web/i18n"
	_gorm "github.com/kapmahc/champak/web/i18n/gorm"
	"github.com/kapmahc/champak/web/mux"
	"github.com/kapmahc/champak/web/settings"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
	"github.com/urfave/cli"
)

// InjectAction  inject action
func InjectAction(fn func(*cli.Context) error) cli.ActionFunc {
	return Action(func(ctx *cli.Context) error {
		inj := inject.Graph{Logger: &injectLogger{}}
		// ----------------
		if err := _crypto.Use(
			[]byte(viper.GetString("secrets.aes")),
			[]byte(viper.GetString("secrets.hmac")),
		); err != nil {
			return err
		}
		// ----------------
		db, err := OpenDatabase()
		if err != nil {
			return err
		}
		settings.Use(db, false)
		i18n.Use(_gorm.New(db, false))
		// ------------------
		namespace := viper.GetString("app.name")
		// ------------------
		rep := OpenRedis()
		cache.Use(
			redis.Use(rep, namespace),
		)
		// ----------------------
		rdr := render.New(render.Options{
			Directory:  path.Join("themes", viper.GetString("server.theme"), "views"),
			Layout:     "application",
			Extensions: []string{".html"},
			Funcs: []template.FuncMap{
				{
					"url": mux.URL,
					"t":   i18n.T,
					"cards": func(loc string) []Card {
						key := fmt.Sprintf("cards/%s", loc)
						var items []Card
						if err := cache.Get(key, &items); err == nil {
							return items
						}
						if err := db.
							Select([]string{"title", "summary", "logo", "href"}).
							Where("loc = ?", loc).
							Order("sort_order ASC").
							Find(&items).Error; err != nil {
							log.Error(err)
						}
						if len(items) > 0 {
							cache.Set(key, items, 60*60*24)
						}
						return items
					},
					"links": func(loc string) []Link {
						key := fmt.Sprintf("links/%s", loc)
						var items []Link
						if err := cache.Get(key, &items); err == nil {
							return items
						}
						if err := db.
							Select([]string{"label", "href"}).
							Where("loc = ?", loc).
							Order("sort_order ASC").
							Find(&items).Error; err != nil {
							log.Error(err)
						}
						if len(items) > 0 {
							cache.Set(key, items, 60*60*24)
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
		// ----------------------

		if err := inj.Provide(
			&inject.Object{Value: rdr},
			&inject.Object{Value: db},
			&inject.Object{Value: rep},
			&inject.Object{Value: namespace, Name: "namespace"},
		); err != nil {
			return err
		}
		// -----------------
		web.Walk(func(en web.Engine) error {
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
