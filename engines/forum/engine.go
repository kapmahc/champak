package forum

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/engines/auth"
	"github.com/kapmahc/champak/web"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine forum engine
type Engine struct {
	Db      *gorm.DB      `inject:""`
	I18n    *web.I18n     `inject:""`
	Session *auth.Session `inject:""`
}

// Atom rss-atom
func (p *Engine) Atom() ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

// Dashboard dashboard links
func (p *Engine) Dashboard(c *gin.Context) []web.Dropdown {
	var items []web.Dropdown
	return items
}

// Do background job
func (p *Engine) Do() {

}

// Home home
func (p *Engine) Home(c *gin.Context) {
}

// Mount web points
func (p *Engine) Mount(rt *gin.Engine) {

}

// Shell console commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	web.Register(&Engine{})
}
