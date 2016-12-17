package auth

import (
	"github.com/facebookgo/inject"
	"github.com/kapmahc/champak/web"
	"github.com/kapmahc/champak/web/cache"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine  auth engine
type Engine struct {
	Cache cache.Store `inject:""`
}

// Map map objects
func (p *Engine) Map(*inject.Graph) error {
	return nil
}

// Mount mount web points
func (p *Engine) Mount(*gin.Engine) {}

// Worker background job
func (p *Engine) Worker() {}

// -----------------------------------------------------------------------------

func init() {
	web.Register(&Engine{})
}
