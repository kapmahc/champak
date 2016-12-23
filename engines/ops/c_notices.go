package ops

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/kapmahc/champak/engines/site"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) newNotice(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "ops.notices.new.title")
	fm := web.NewForm(c, "new-notices", title, "/ops/notices")

	fm.AddFields(
		web.NewTextArea("body", p.I18n.T(lng, "attributes.body"), ""),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/form", data)
}

type fmNotice struct {
	Body string `form:"body" binding:"required,max=500"`
}

func (p *Engine) createNotice(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmNotice)
	if err := p.Db.Create(&site.Notice{Body: fm.Body}).Error; err != nil {
		return err
	}

	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "success"), web.NOTICE)
	ss.Save()
	c.Redirect(http.StatusFound, "/ops/notices")
	return nil
}

func (p *Engine) editNotice(c *gin.Context) error {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	var n site.Notice
	if err := p.Db.Where("id = ?", c.Param("id")).First(&n).Error; err != nil {
		return err
	}

	title := p.I18n.T(lng, "ops.notices.edit.title", n.ID)
	fm := web.NewForm(c, "new-notices", title, fmt.Sprintf("/ops/notices/%d", n.ID))

	fm.AddFields(
		web.NewTextArea("body", p.I18n.T(lng, "attributes.body"), n.Body),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/form", data)
	return nil
}

func (p *Engine) updateNotice(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmNotice)

	if err := p.Db.Model(&site.Notice{}).Where("id = ?", c.Param("id")).
		Update("body", fm.Body).Error; err != nil {
		return err
	}

	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "success"), web.NOTICE)
	ss.Save()
	c.Redirect(http.StatusFound, "/ops/notices")
	return nil
}

func (p *Engine) indexNotices(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	data["title"] = p.I18n.T(lng, "ops.notices.index.title")
	var items []site.Notice
	if err := p.Db.Order("updated DESC").Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["items"] = items
	c.HTML(http.StatusOK, "ops/notices", data)
}

func (p *Engine) destoryNotice(c *gin.Context) error {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(site.Notice{}).Error; err != nil {
		return err
	}
	c.JSON(http.StatusOK, gin.H{web.TO: "/ops/notices"})
	return nil
}
