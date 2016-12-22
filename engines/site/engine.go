package site

import (
	"github.com/facebookgo/inject"
	"github.com/kapmahc/champak/engines/auth"
	"github.com/kapmahc/champak/web"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine engine
type Engine struct {
	Session  *auth.Session `inject:""`
	I18n     *web.I18n     `inject:""`
	Settings *web.Settings `inject:""`
}

// Map map
func (p *Engine) Map(*inject.Graph) error {
	return nil
}

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

// -----------------------------------------------------------------------------

func init() {
	web.Register(&Engine{})
}