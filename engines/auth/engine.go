package auth

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/web"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine auth engine
type Engine struct {
	Dao      *Dao              `inject:""`
	Db       *gorm.DB          `inject:""`
	I18n     *web.I18n         `inject:""`
	Security *web.Security     `inject:""`
	Jwt      *Jwt              `inject:""`
	Server   *machinery.Server `inject:""`
	Session  *Session          `inject:""`
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
func (p *Engine) Dashboard(c *gin.Context) []web.Dropdown {
	var items []web.Dropdown
	if _, ok := c.Get(CurrentUser); ok {
		items = append(
			items,
			web.Dropdown{
				Label: "auth.profile",
				Links: []*web.Link{
					&web.Link{
						Href:  "/users/info",
						Label: "auth.users.info.title",
					},
					&web.Link{
						Href:  "/users/change-password",
						Label: "auth.users.change-password.title",
					},
					&web.Link{
						Href:  "/users/logs",
						Label: "auth.users.logs.title",
					},
				},
			},
		)
	}

	return items
}

func init() {
	web.Register(&Engine{})
}
