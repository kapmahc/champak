package ops

import (
	"crypto/x509/pkix"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/facebookgo/inject"
	"github.com/fvbock/endless"
	"github.com/gin-contrib/sessions"
	"github.com/gorilla/csrf"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
	"github.com/steinbacher/goose"
	"github.com/urfave/cli"
	"golang.org/x/text/language"
	gin "gopkg.in/gin-gonic/gin.v1"
)

//Shell command options
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{

		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start the app server",
			Action: web.IocAction(func(*cli.Context, *inject.Graph) error {
				if web.IsProduction() {
					gin.SetMode(gin.ReleaseMode)
				}
				rt := gin.Default()

				theme := viper.GetString("server.theme")
				if tpl, err := p.loadTemplates(theme); err == nil {
					rt.SetHTMLTemplate(tpl)
				} else {
					return err
				}
				rt.Static("/assets", path.Join("themes", theme, "assets"))

				rt.Use(
					sessions.Sessions(
						"_session_",
						sessions.NewCookieStore([]byte(viper.GetString("secrets.session"))),
					),
					p.I18n.Handler(),
					csrfHandler,
					p.Jwt.CurrentUserHandler,
				)

				web.Loop(func(en web.Engine) error {
					en.Mount(rt)
					return nil
				})

				adr := fmt.Sprintf(":%d", viper.GetInt("server.port"))
				hnd := csrf.Protect([]byte(viper.GetString("secrets.csrf")))(rt)

				if web.IsProduction() {
					return endless.ListenAndServe(adr, hnd)
				}
				return http.ListenAndServe(adr, hnd)
			}),
		},
		{
			Name:    "worker",
			Aliases: []string{"w"},
			Usage:   "start the worker progress",
			Action: web.IocAction(func(_ *cli.Context, inj *inject.Graph) error {
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
			Action: web.Action(func(*cli.Context) error {
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
					Action: web.IocAction(func(*cli.Context, *inject.Graph) error {
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
					Action: web.IocAction(func(*cli.Context, *inject.Graph) error {
						return p.Cache.Flush()
					}),
				},
			},
		},
		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "database operations",
			Subcommands: []cli.Command{
				{
					Name:    "example",
					Usage:   "scripts example for create database and user",
					Aliases: []string{"e"},
					Action: web.Action(func(*cli.Context) error {
						drv := viper.GetString("database.driver")
						args := viper.GetStringMapString("database.args")
						var err error
						switch drv {
						case "postgres":
							fmt.Printf("CREATE USER %s WITH PASSWORD '%s';\n", args["user"], args["password"])
							fmt.Printf("CREATE DATABASE %s WITH ENCODING='UTF8';\n", args["dbname"])
							fmt.Printf("GRANT ALL PRIVILEGES ON DATABASE %s TO %s;\n", args["dbname"], args["user"])
						default:
							err = fmt.Errorf("unknown driver %s", drv)
						}
						return err
					}),
				},
				{
					Name:    "migrate",
					Usage:   "migrate the DB to the most recent version available",
					Aliases: []string{"m"},
					Action: web.Action(func(*cli.Context) error {
						conf, err := dbConf()
						if err != nil {
							return err
						}

						target, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
						if err != nil {
							return err
						}

						return goose.RunMigrations(conf, conf.MigrationsDir, target)
					}),
				},
				{
					Name:    "rollback",
					Usage:   "roll back the version by 1",
					Aliases: []string{"r"},
					Action: web.Action(func(*cli.Context) error {
						conf, err := dbConf()
						if err != nil {
							return err
						}

						current, err := goose.GetDBVersion(conf)
						if err != nil {
							return err
						}

						previous, err := goose.GetPreviousDBVersion(conf.MigrationsDir, current)
						if err != nil {
							return err
						}

						return goose.RunMigrations(conf, conf.MigrationsDir, previous)
					}),
				},
				{
					Name:    "version",
					Usage:   "dump the migration status for the current DB",
					Aliases: []string{"v"},
					Action: web.Action(func(*cli.Context) error {
						conf, err := dbConf()
						if err != nil {
							return err
						}

						// collect all migrations
						migrations, err := goose.CollectMigrations(conf.MigrationsDir)
						if err != nil {
							return err
						}

						db, err := goose.OpenDBFromDBConf(conf)
						if err != nil {
							return err
						}
						defer db.Close()

						// must ensure that the version table exists if we're running on a pristine DB
						if _, err = goose.EnsureDBVersion(conf, db); err != nil {
							return err
						}

						fmt.Println("    Applied At                  Migration")
						fmt.Println("    =======================================")
						for _, m := range migrations {
							if err = printMigrationStatus(db, m.Version, filepath.Base(m.Source)); err != nil {
								return err
							}
						}
						return nil
					}),
				},
				{
					Name:    "connect",
					Usage:   "connect database",
					Aliases: []string{"c"},
					Action: web.Action(func(*cli.Context) error {
						drv := viper.GetString("database.driver")
						args := viper.GetStringMapString("database.args")
						var err error
						switch drv {
						case "postgres":
							err = web.Shell("psql",
								"-h", args["host"],
								"-p", args["port"],
								"-U", args["user"],
								args["dbname"],
							)
						default:
							err = fmt.Errorf("unknown driver %s", drv)
						}
						return err
					}),
				},
				{
					Name:    "create",
					Usage:   "create database",
					Aliases: []string{"n"},
					Action: web.Action(func(*cli.Context) error {
						drv := viper.GetString("database.driver")
						args := viper.GetStringMapString("database.args")
						var err error
						switch drv {
						case "postgres":
							err = web.Shell("psql",
								"-h", args["host"],
								"-p", args["port"],
								"-U", "postgres",
								"-c", fmt.Sprintf(
									"CREATE DATABASE %s WITH ENCODING='UTF8'",
									args["dbname"],
								),
							)
						default:
							err = fmt.Errorf("unknown driver %s", drv)
						}
						return err
					}),
				},
				{
					Name:    "drop",
					Usage:   "drop database",
					Aliases: []string{"d"},
					Action: web.Action(func(*cli.Context) error {
						drv := viper.GetString("database.driver")
						args := viper.GetStringMapString("database.args")
						var err error
						switch drv {
						case "postgres":
							err = web.Shell("psql",
								"-h", args["host"],
								"-p", args["port"],
								"-U", "postgres",
								"-c", fmt.Sprintf("DROP DATABASE %s", args["dbname"]),
							)
						default:
							err = fmt.Errorf("unknown driver %s", drv)
						}
						return err
					}),
				},
			},
		},
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate file template",
			Subcommands: []cli.Command{
				{
					Name:    "config",
					Aliases: []string{"c"},
					Usage:   "generate config file",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "environment, e",
							Value: "development",
							Usage: "environment, like: development, production, stage, test...",
						},
					},
					Action: func(c *cli.Context) error {
						const fn = "config.toml"
						if _, err := os.Stat(fn); err == nil {
							return fmt.Errorf("file %s already exists", fn)
						}
						fmt.Printf("generate file %s\n", fn)

						viper.Set("env", c.String("environment"))
						args := viper.AllSettings()
						fd, err := os.Create(fn)
						if err != nil {
							return err
						}
						defer fd.Close()
						end := toml.NewEncoder(fd)
						err = end.Encode(args)

						return err

					},
				},
				{
					Name:    "nginx",
					Aliases: []string{"ng"},
					Usage:   "generate nginx.conf",
					Action: web.Action(func(*cli.Context) error {
						const tpl = `
		server {
		  listen 80;
		  server_name {{.Name}};
		  rewrite ^(.*) https://$host$1 permanent;
		}

		upstream {{.Name}}_prod {
		  server localhost:{{.Port}} fail_timeout=0;
		}

		server {
		  listen 443;

		  ssl  on;
		  ssl_certificate  /etc/ssl/certs/{{.Name}}.crt;
		  ssl_certificate_key  /etc/ssl/private/{{.Name}}.key;
		  ssl_session_timeout  5m;
		  ssl_protocols  SSLv2 SSLv3 TLSv1;
		  ssl_ciphers  RC4:HIGH:!aNULL:!MD5;
		  ssl_prefer_server_ciphers  on;

		  client_max_body_size 4G;
		  keepalive_timeout 10;
		  proxy_buffers 16 64k;
		  proxy_buffer_size 128k;

		  server_name {{.Name}};
		  root {{.Root}}/public;
		  index index.html;
		  access_log /var/log/nginx/{{.Name}}.access.log;
		  error_log /var/log/nginx/{{.Name}}.error.log;
		  location ~* \.(?:css|js)$ {
		    gzip_static on;
		    expires max;
		    access_log off;
		    add_header Cache-Control "public";
		  }
		  location ~* \.(?:jpg|jpeg|gif|png|ico|cur|gz|svg|svgz|mp4|ogg|ogv|webm|htc)$ {
		    expires 1M;
		    expires max;
		    access_log off;
		    add_header Cache-Control "public";
		  }
		  location ~* \.(?:rss|atom)$ {
		    expires 12h;
		    access_log off;
		    add_header Cache-Control "public";
		  }
		  location / {
		    proxy_set_header X-Forwarded-Proto https;
		    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		    proxy_set_header Host $http_host;
		    proxy_set_header X-Real-IP $remote_addr;
		    proxy_redirect off;
		    proxy_pass http://{{.Name}}_prod;
		    # limit_req zone=one;
		  }
		}
		`
						t, err := template.New("").Parse(tpl)
						if err != nil {
							return err
						}
						pwd, err := os.Getwd()
						if err != nil {
							return err
						}

						name := viper.GetString("server.name")
						fn := path.Join("etc", "nginx", "sites-enabled", name+".conf")
						if err = os.MkdirAll(path.Dir(fn), 0700); err != nil {
							return err
						}
						fmt.Printf("generate file %s\n", fn)
						fd, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
						if err != nil {
							return err
						}
						defer fd.Close()

						return t.Execute(fd, struct {
							Name    string
							Port    int
							Root    string
							Version string
						}{
							Name:    name,
							Port:    viper.GetInt("http.port"),
							Root:    pwd,
							Version: "v1",
						})
					}),
				},

				{
					Name:    "openssl",
					Aliases: []string{"ssl"},
					Usage:   "generate ssl certificates",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name, n",
							Usage: "name",
						},
						cli.StringFlag{
							Name:  "country, c",
							Value: "Earth",
							Usage: "country",
						},
						cli.StringFlag{
							Name:  "organization, o",
							Value: "Mother Nature",
							Usage: "organization",
						},
						cli.IntFlag{
							Name:  "years, y",
							Value: 1,
							Usage: "years",
						},
					},
					Action: web.Action(func(c *cli.Context) error {
						name := c.String("name")
						if len(name) == 0 {
							cli.ShowCommandHelp(c, "openssl")
							return nil
						}
						root := path.Join("etc", "ssl", name)

						key, crt, err := CreateCertificate(
							true,
							pkix.Name{
								Country:      []string{c.String("country")},
								Organization: []string{c.String("organization")},
							},
							c.Int("years"),
						)
						if err != nil {
							return err
						}

						fnk := path.Join(root, "key.pem")
						fnc := path.Join(root, "crt.pem")

						fmt.Printf("generate pem file %s\n", fnk)
						err = WritePemFile(fnk, "RSA PRIVATE KEY", key, 0600)
						fmt.Printf("test: openssl rsa -noout -text -in %s\n", fnk)

						if err == nil {
							fmt.Printf("generate pem file %s\n", fnc)
							err = WritePemFile(fnc, "CERTIFICATE", crt, 0444)
							fmt.Printf("test: openssl x509 -noout -text -in %s\n", fnc)
						}
						if err == nil {
							fmt.Printf(
								"verify: diff <(openssl rsa -noout -modulus -in %s) <(openssl x509 -noout -modulus -in %s)",
								fnk,
								fnc,
							)
						}
						fmt.Println()
						return err
					}),
				},

				{
					Name:    "migration",
					Usage:   "generate migration file",
					Aliases: []string{"m"},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name, n",
							Usage: "name",
						},
					},
					Action: web.Action(func(c *cli.Context) error {
						name := c.String("name")
						if len(name) == 0 {
							cli.ShowCommandHelp(c, "migration")
							return nil
						}
						cfg, err := dbConf()
						if err != nil {
							return err
						}
						if err = os.MkdirAll(cfg.MigrationsDir, 0700); err != nil {
							return err
						}
						file, err := goose.CreateMigration(name, "sql", cfg.MigrationsDir, time.Now())
						if err != nil {
							return err
						}

						fmt.Printf("generate file %s\n", file)
						return nil
					}),
				},

				{
					Name:    "locale",
					Usage:   "generate locale file",
					Aliases: []string{"l"},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name, n",
							Usage: "locale name",
						},
					},
					Action: web.Action(func(c *cli.Context) error {
						name := c.String("name")
						if len(name) == 0 {
							cli.ShowCommandHelp(c, "locale")
							return nil
						}
						lng, err := language.Parse(name)
						if err != nil {
							return err
						}
						const root = "locales"
						if err = os.MkdirAll(root, 0700); err != nil {
							return err
						}
						file := path.Join(root, fmt.Sprintf("%s.ini", lng.String()))
						fmt.Printf("generate file %s\n", file)
						fd, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
						if err != nil {
							return err
						}
						defer fd.Close()
						return err
					}),
				},
			},
		},
		{
			Name:    "routes",
			Aliases: []string{"rt"},
			Usage:   "print out all defined routes",
			Action: func(*cli.Context) error {
				gin.SetMode(gin.ReleaseMode)
				rt := gin.New()
				web.Loop(func(en web.Engine) error {
					en.Mount(rt)
					return nil
				})
				fmt.Println("METHOD\tPATH")
				for _, r := range rt.Routes() {
					fmt.Printf("%s\t%s\n", r.Method, r.Path)
				}
				return nil
			},
		},
		{
			Name:  "i18n",
			Usage: "i18n operations",
			Subcommands: []cli.Command{
				{
					Name:    "sync",
					Aliases: []string{"s"},
					Usage:   "sync locales from files",
					Action: web.IocAction(func(*cli.Context, *inject.Graph) error {
						return p.I18n.Sync("locales")
					}),
				},
			},
		},
	}
}
