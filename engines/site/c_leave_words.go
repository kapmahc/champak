package site

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) newLeaveWord(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "site.leave-words.new.title")
	fm := web.NewForm(c, "new-leave-words", title, "/leave-words")

	fm.AddFields(
		web.NewTextArea("body", p.I18n.T(lng, "attributes.body"), ""),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/non-sign-in", data)
}

type fmLeaveWord struct {
	Body string `form:"body" binding:"required,max=500"`
}

func (p *Engine) createLeaveWord(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmLeaveWord)
	if err := p.Db.Create(&LeaveWord{Body: fm.Body}).Error; err != nil {
		return err
	}

	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "success"), web.NOTICE)
	ss.Save()
	c.Redirect(http.StatusFound, "/leave-words/new")
	return nil
}
