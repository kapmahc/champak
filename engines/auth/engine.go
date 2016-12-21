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
	return func(c *gin.Context) []web.Dropdown {
		var items []web.Dropdown
		if _, ok := c.Get(CurrentUser); ok {
			items = append(
				items,
				web.Dropdown{
					Label: "auth.personal.title",
					Links: []*web.Link{
						&web.Link{
							Href:  "/personal/profile",
							Label: "auth.personal.profile.title",
						},
						&web.Link{
							Href:  "/personal/change-password",
							Label: "auth.personal.change-password.title",
						},
						&web.Link{
							Href:  "/personal/logs",
							Label: "auth.personal.logs.title",
						},
					},
				},
			)
		}

		return items
	}
}

// -----------------------------------------------------------------------------

func init() {

	web.Register(&Engine{})
}
