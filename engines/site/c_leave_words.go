package site

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
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

func (p *Engine) indexLeaveWords(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	data["title"] = p.I18n.T(lng, "site.leave-words.index.title")
	var items []LeaveWord
	if err := p.Db.Order("created_at DESC").Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["items"] = items
	c.HTML(http.StatusOK, "leave-words", data)
}

func (p *Engine) destoryLeaveWord(c *gin.Context) error {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(LeaveWord{}).Error; err != nil {
		return err
	}
	c.JSON(http.StatusOK, gin.H{})
	return nil
}
