package auth

import (
	"github.com/facebookgo/inject"
	"github.com/jinzhu/gorm"
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
	Db       *gorm.DB      `inject:""`
	Security *web.Security `inject:""`
	Session  *Session      `inject:""`
}

// Map map objects
func (p *Engine) Map(inj *inject.Graph) error {
	return nil
}

// Dashboard dashboard links
func (p *Engine) Dashboard() web.DashboardHandler {
	return func(*gin.Context) []web.Link {
		return []web.Link{}
	}
}

// -----------------------------------------------------------------------------

func init() {

	web.Register(&Engine{})
}
