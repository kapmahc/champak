package web

import (
	"net/http"

	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/kapmahc/champak/web/mux"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Link link
type Link struct {
	Label string
	Href  string
}

// Dropdown drop-down
type Dropdown struct {
	Label string
	Links []*Link
}

// DashboardHandler dashboard handler
type DashboardHandler func(req *http.Request) []Dropdown

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

// Walk walk engines
func Walk(fn func(Engine) error) error {
	for _, en := range engines {
		if err := fn(en); err != nil {
			return err
		}
	}
	return nil
}
