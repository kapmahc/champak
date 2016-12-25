package auth

import (
	"net/http"

	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/web"
	"golang.org/x/tools/blog/atom"
)

// Engine auth engine
type Engine struct {
	Dao *Dao     `inject:""`
	Db  *gorm.DB `inject:""`
}

// Map inject objects
func (p *Engine) Map(*inject.Graph) error {
	return nil
}

// Worker register background jobs
func (p *Engine) Worker() {}

// Dashboard dashboard links
func (p *Engine) Dashboard() web.DashboardHandler {
	return func(req *http.Request) []web.Dropdown {
		return []web.Dropdown{}
	}
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
