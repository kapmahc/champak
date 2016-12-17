package auth

import (
	"github.com/facebookgo/inject"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func IocAction(fn func(*cli.Context, *inject.Graph) error) cli.ActionFunc {
	return Action(func(ctx *cli.Context) error {
		var inj inject.Graph
		web.Loop(func(en web.Engine) error {
			if e := en.Map(&inj); e != nil {
				return e
			}
			return inj.Provide(&inject.Object{Value: en})
		})
		if err := inj.Populate(); err != nil {
			return err
		}
		return fn(ctx, &inj)
	})
}

func Action(f cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		return f(c)
	}
}
