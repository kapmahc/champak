package shop

import (
	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/kapmahc/champak/web"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Engine shop engine
type Engine struct {
}

// Map inject objects
func (p *Engine) Map(*inject.Graph) error {
	return nil
}

// Mount mount web points
func (p *Engine) Mount(web.Router) {

}

// Worker background jobs
func (p *Engine) Worker() {

}

// Shell console commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

// Atom rss-atom
func (p *Engine) Atom() ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

func init() {
	web.Register(&Engine{})
}
