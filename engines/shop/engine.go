package shop

import (
	"net/http"

	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/kapmahc/champak/web"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Engine shop engine
type Engine struct {
}

// Map map objects
func (*Engine) Map(*inject.Graph) error {
	return nil
}

// Mount mount web points
func (*Engine) Mount(*mux.Router) {}

// Worker background jobs
func (*Engine) Worker() {}

// Dashboard dashboard links(by user)
func (*Engine) Dashboard(req *http.Request) []web.Dropdown {
	return []web.Dropdown{}
}

// Shell command lines
func (*Engine) Shell() []cli.Command {
	return []cli.Command{}
}

// Atom rss-atom
func (*Engine) Atom() ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml
func (*Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

func init() {
	web.Register(&Engine{})
}
