package web

import (
	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
)

// Engine engine
type Engine interface {
	Map(*inject.Graph) error
	Mount(*mux.Router)
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
