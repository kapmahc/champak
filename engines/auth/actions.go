package auth

import (
	log "github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
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

// Action cfg action
func Action(f cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		return f(c)
	}
}
