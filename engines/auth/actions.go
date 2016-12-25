package auth

import (
	"crypto/aes"
	"fmt"

	"github.com/SermoDigital/jose/crypto"
	"github.com/facebookgo/inject"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

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

		if err := inj.Provide(
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
