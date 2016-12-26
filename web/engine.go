package web

import (
	"net/http"

	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Link link
type Link struct {
	Link string
	Href string
}

// Engine engine
type Engine interface {
	Map(*inject.Graph) error
	Mount(*httprouter.Router)
	Worker()
	Dashboard(*http.Request) []Link
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
