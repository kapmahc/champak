package web

import (
	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Dropdown drop-down
type Dropdown struct {
	Label string
	Links []*Link
}

// DashboardHandler dashboard handler
type DashboardHandler func(*gin.Context) []Dropdown

// Engine engine
type Engine interface {
	Map(*inject.Graph) error
	Mount(*mux.Router)
	Worker()
	Dashboard() DashboardHandler
	Shell() []cli.Command
	Atom() ([]*atom.Entry, error)
	Sitemap() ([]stm.URL, error)
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
