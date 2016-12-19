package auth

import (
	"github.com/facebookgo/inject"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine  auth engine
type Engine struct {
	Cache    *web.Cache    `inject:""`
	Job      *web.Job      `inject:""`
	I18n     *web.I18n     `inject:""`
	Jwt      *Jwt          `inject:""`
	Dao      *Dao          `inject:""`
	Security *web.Security `inject:""`
}

// Map map objects
func (p *Engine) Map(inj *inject.Graph) error {
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
