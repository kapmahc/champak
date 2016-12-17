package auth

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/facebookgo/inject"
	"github.com/fvbock/endless"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{

		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start the app server",
			Action: IocAction(func(*cli.Context, *inject.Graph) error {
				if web.IsProduction() {
					gin.SetMode(gin.ReleaseMode)
				}
				rt := gin.Default()

				theme := viper.GetString("server.theme")
				tpl, err := template.
					New("").
					Funcs(template.FuncMap{
					// TODO cards links t
					// "T": p.I18n.T,
					}).
					ParseGlob(
						fmt.Sprintf("themes/%s/templates/**/*", theme),
					)
				if err != nil {
					return err
				}
				rt.SetHTMLTemplate(tpl)
				rt.Static("/assets", path.Join("themes", theme, "assets"))

				//TODO
				// rt.Use(sessions.Sessions(
				// 	"_session_",
				// 	sessions.NewCookieStore([]byte(viper.GetString("secrets.session"))),
				// ))
				// rt.Use(i18n.LocaleHandler(p.Logger))

				web.Loop(func(en web.Engine) error {
					en.Mount(rt)
					return nil
				})

				adr := fmt.Sprintf(":%d", viper.GetInt("server.port"))

				// hnd := cors.New(cors.Options{
				// 	AllowCredentials: true,
				// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
				// 	AllowedHeaders:   []string{"*"},
				// 	Debug:            !IsProduction(),
				// }).Handler(rt)

				if web.IsProduction() {
					return endless.ListenAndServe(adr, rt)
				}
				return http.ListenAndServe(adr, rt)
			}),
		},
		{
			Name:    "worker",
			Aliases: []string{"w"},
			Usage:   "start the worker progress",
			Action: IocAction(func(_ *cli.Context, inj *inject.Graph) error {
				forever := make(chan bool)

				web.Loop(func(en web.Engine) error {
					go en.Worker()
					return nil
				})
				log.Printf(" [*] Waiting for logs. To exit press CTRL+C")

				<-forever
				return nil
			}),
		},
		{
			Name:    "redis",
			Aliases: []string{"re"},
			Usage:   "open redis connection",
			Action: Action(func(*cli.Context) error {
				return web.Shell(
					"redis-cli",
					"-h", viper.GetString("redis.host"),
					"-p", viper.GetString("redis.port"),
					"-n", viper.GetString("redis.db"),
				)
			}),
		},
		{
			Name:    "cache",
			Aliases: []string{"c"},
			Usage:   "cache operations",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Usage:   "list all cache keys",
					Aliases: []string{"l"},
					Action: IocAction(func(*cli.Context, *inject.Graph) error {
						keys, err := p.Cache.Keys()
						if err != nil {
							return err
						}
						for _, k := range keys {
							fmt.Println(k)
						}
						return nil
					}),
				},
				{
					Name:    "clear",
					Usage:   "clear cache items",
					Aliases: []string{"c"},
					Action: IocAction(func(*cli.Context, *inject.Graph) error {
						return p.Cache.Flush()
					}),
				},
			},
		},
	}
}
