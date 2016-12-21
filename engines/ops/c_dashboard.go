package ops

import (
	"net/http"

	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getDashboard(c *gin.Context) {
	data := c.MustGet(web.DATA).(gin.H)
	lng := c.MustGet(web.LOCALE).(string)
	var links []web.Dropdown
	web.Loop(func(en web.Engine) error {
		items := en.Dashboard()(c)
		links = append(links, items...)
		return nil
	})
	data["links"] = links
	data["title"] = p.I18n.T(lng, "header.dashboard")
	c.HTML(http.StatusOK, "ops/dashboard", data)
}
