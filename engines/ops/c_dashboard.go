package ops

import (
	"net/http"

	"github.com/kapmahc/champak/web"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getDashboard(c *gin.Context) {
	data := c.MustGet(web.DATA).(gin.H)
	lng := c.MustGet(web.LOCALE).(string)
	data["title"] = p.I18n.T(lng, "header.dashboard")
	c.HTML(http.StatusOK, "ops/dashboard", data)
}
