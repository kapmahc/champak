package reading

import (
	"github.com/facebookgo/inject"
	"github.com/kapmahc/champak/web"
	"github.com/kapmahc/champak/web/sitemap"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine engine
type Engine struct {
}

// Map map
func (p *Engine) Map(*inject.Graph) error {
	return nil
}

// Mount mount
func (p *Engine) Mount(*gin.Engine) {}

// Worker worker
func (p *Engine) Worker() {}

// Shell shell
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

// Dashboard dashboard links
func (p *Engine) Dashboard() web.DashboardHandler {
	return func(*gin.Context) []web.Dropdown {
		return []web.Dropdown{}
	}
}

// Atom atom entry
func (p *Engine) Atom() ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap entry
func (p *Engine) Sitemap() ([]*sitemap.Item, error) {
	return []*sitemap.Item{}, nil
}

// -----------------------------------------------------------------------------

func init() {
	web.Register(&Engine{})
}
