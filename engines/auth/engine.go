package auth

import (
	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/web"
	"github.com/unrolled/render"
	"golang.org/x/tools/blog/atom"
)

// Engine auth engine
type Engine struct {
	R   *render.Render `inject:""`
	Dao *Dao           `inject:""`
	Db  *gorm.DB       `inject:""`
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
