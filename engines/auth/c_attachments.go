package auth

import (
	"net/http"

	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexAttachments(c *gin.Context) {
	data := c.MustGet(web.DATA).(gin.H)
	lng := c.MustGet(web.LOCALE).(string)
	title := p.I18n.T(lng, "auth.attachments.index.title")
	data["title"] = title
	c.HTML(http.StatusOK, "attachments", data)
}
