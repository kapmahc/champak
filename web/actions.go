package web

import (
	"crypto/aes"

	"github.com/SermoDigital/jose/crypto"
	log "github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

// IocAction ioc action
func IocAction(fn func(*cli.Context, *inject.Graph) error) cli.ActionFunc {
	return Action(func(ctx *cli.Context) error {
		inj := inject.Graph{Logger: &injectLogger{}}
		// -------------------
		db, err := OpenDatabase()
		if err != nil {
			return err
		}
		// -------------------
		rep := OpenRedis()
		// -------------------
		bws, err := NewWorkerServer()
		if err != nil {
			return err
		}
		// -------------------
		cip, err := aes.NewCipher([]byte(viper.GetString("secrets.aes")))
		if err != nil {
			return err
		}
		// -------------------

		if err := inj.Provide(
			&inject.Object{Value: db},
			&inject.Object{Value: bws},
			&inject.Object{Value: rep},
			&inject.Object{Value: cip},
			&inject.Object{Value: cip, Name: "aes.cip"},
			&inject.Object{Value: []byte(viper.GetString("secrets.hmac")), Name: "hmac.key"},
			&inject.Object{Value: []byte(viper.GetString("secrets.jwt")), Name: "jwt.key"},
			&inject.Object{Value: viper.GetString("app.name"), Name: "namespace"},
			&inject.Object{Value: crypto.SigningMethodHS512, Name: "jwt.method"},
		); err != nil {
			return err
		}
		// -----------------
		if err := Walk(func(en Engine) error {
			return inj.Provide(&inject.Object{Value: en})
		}); err != nil {
			return err
		}

		if err := inj.Populate(); err != nil {
			return err
		}
		return fn(ctx, &inj)
	})
}

// Action cfg action
func Action(f cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		log.Infof("read config from config.toml")
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		return f(c)
	}
}
