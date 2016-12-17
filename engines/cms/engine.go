package cms

import (
	"github.com/facebookgo/inject"
	"github.com/kapmahc/champak/web"
	"github.com/urfave/cli"
	gin "gopkg.in/gin-gonic/gin.v1"
)

type Engine struct {
}

func (p *Engine) Map(*inject.Graph) error {
	return nil
}

func (p *Engine) Mount(*gin.Engine) {}

func (p *Engine) Worker() {}

func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

// -----------------------------------------------------------------------------

func init() {
	web.Register(&Engine{})
}
