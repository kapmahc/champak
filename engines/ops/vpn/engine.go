package vpn

import (
	"github.com/facebookgo/inject"
	"github.com/kapmahc/champak/web"
	"github.com/urfave/cli"
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
	return func(*gin.Context) []*web.Link {
		return []*web.Link{}
	}
}

// -----------------------------------------------------------------------------

func init() {
	web.Register(&Engine{})
}
