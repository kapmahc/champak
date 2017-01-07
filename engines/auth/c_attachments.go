package auth

import (
	"net/http"

	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) attachmentsUpload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (p *Engine) attachmentsIndex(c *gin.Context) {
	data := c.MustGet(web.DATA).(gin.H)
	lng := c.MustGet(web.LOCALE).(string)
	title := p.I18n.T(lng, "auth.attachments.index.title")
	data["title"] = title
	fm := web.NewForm(c, "upload", title, "/attachments/upload")
	fm.AddFields(
		web.NewHiddenField("type", c.Query("type")),
		web.NewHiddenField("id", c.Query("id")),
		web.NewFileField("files", p.I18n.T(lng, "attributes.files")),
	)
	data["form"] = fm

	c.HTML(http.StatusOK, "attachments", data)
}
