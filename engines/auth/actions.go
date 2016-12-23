package auth

import (
	"crypto/aes"
	"fmt"

	"github.com/SermoDigital/jose/crypto"
	log "github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"golang.org/x/text/language"
)

type injectLogger struct {
}

func (p *injectLogger) Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

// IocAction ioc action
func IocAction(fn func(*cli.Context, *inject.Graph) error) cli.ActionFunc {
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

		if err := inj.Provide(
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
			&inject.Object{Value: viper.GetStringSlice("languages"), Name: "language.tags"},
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
		return fn(ctx, &inj)
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
