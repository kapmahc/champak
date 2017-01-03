package blog

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/kapmahc/champak/web"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine blog engine
type Engine struct {
}

// Mount web points
func (p *Engine) Mount(*gin.Engine) {}

// Do background job
func (p *Engine) Do() {}

// Shell shell commands
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

// Dashboard dashboard links
func (p *Engine) Dashboard(*gin.Context) []web.Dropdown {
	return []web.Dropdown{}
}

func init() {
	web.Register(&Engine{})
}
