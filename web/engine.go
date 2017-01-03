package web

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine engine
type Engine interface {
	Mount(*gin.Engine)
	Do()
	Shell() []cli.Command
	Atom() ([]*atom.Entry, error)
	Sitemap() ([]stm.URL, error)
	Dashboard(*gin.Context) []Dropdown
}

// -----------------------------------------------------------------------------

var engines []Engine

// Register register engines
func Register(ens ...Engine) {
	engines = append(engines, ens...)
}

// Walk walk engines
func Walk(fn func(Engine) error) error {
	for _, en := range engines {
		if err := fn(en); err != nil {
			return err
		}
	}
	return nil
}
