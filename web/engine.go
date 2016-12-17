package web

import (
	"github.com/facebookgo/inject"
	"github.com/urfave/cli"
	gin "gopkg.in/gin-gonic/gin.v1"
)

type Engine interface {
	Map(*inject.Graph) error
	Mount(*gin.Engine)
	Worker()
	Shell() []cli.Command
}

var engines []Engine

func Loop(fn func(Engine) error) error {
	for _, en := range engines {
		if err := fn(en); err != nil {
			return err
		}
	}
	return nil
}
