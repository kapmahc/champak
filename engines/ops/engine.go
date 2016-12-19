package ops

import (
	"github.com/facebookgo/inject"
	"github.com/kapmahc/champak/web"
	"github.com/kapmahc/champak/web/cache"
)

// Engine ops engine
type Engine struct {
	Cache cache.Store `inject:""`
	Job   web.Job     `inject:""`
}

// Map inject objects
func (p *Engine) Map(*inject.Graph) error {
	return nil
}

// Worker background workers
func (p *Engine) Worker() {}

// -----------------------------------------------------------------------------

func init() {
	web.Register(&Engine{})
}
