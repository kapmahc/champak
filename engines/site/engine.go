package site

import (
	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/kapmahc/champak/web"
	"golang.org/x/tools/blog/atom"
)

// Engine site engine
type Engine struct {
	Cache    *web.Cache    `inject:""`
	I18n     *web.I18n     `inject:""`
	Settings *web.Settings `inject:""`
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
