package web

import (
	"github.com/facebookgo/inject"
	"github.com/urfave/cli"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine engine
type Engine interface {
	Map(*inject.Graph) error
	Mount(*gin.Engine)
	Worker()
	Shell() []cli.Command
}

// -----------------------------------------------------------------------------

var engines []Engine

// Register register engines
func Register(ens ...Engine) {
	engines = append(engines, ens...)
}

// Loop loop engines
func Loop(fn func(Engine) error) error {
	for _, en := range engines {
		if err := fn(en); err != nil {
			return err
		}
	}
	return nil
}
