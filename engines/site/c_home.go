package site

import (
	"net/http"

	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getHome(c *gin.Context) {
	data := c.MustGet("data").(gin.H)
	c.HTML(http.StatusOK, "site/home", data)
}

func (p *Engine) getDashboard(c *gin.Context) {
	data := c.MustGet(web.DATA).(gin.H)
	lng := c.MustGet(web.LOCALE).(string)
	data["title"] = p.I18n.T(lng, "header.dashboard")
	c.HTML(http.StatusOK, "site/dashboard", data)
}