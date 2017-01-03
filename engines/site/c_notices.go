package site

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/kapmahc/champak/engines/auth"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) newNotice(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "site.notices.new.title")
	fm := web.NewForm(c, "new-notices", title, "/notices")

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
	if err := p.Db.Create(&Notice{Body: fm.Body}).Error; err != nil {
		return err
	}

	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "success"), web.NOTICE)
	ss.Save()

	return nil
}

func (p *Engine) editNotice(c *gin.Context) error {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	var n Notice
	if err := p.Db.Where("id = ?", c.Param("id")).First(&n).Error; err != nil {
		return err
	}

	title := p.I18n.T(lng, "site.notices.edit.title", n.ID)
	fm := web.NewForm(c, "new-notices", title, fmt.Sprintf("/notices/%d", n.ID))

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

	if err := p.Db.Model(&Notice{}).Where("id = ?", c.Param("id")).
		Update("body", fm.Body).Error; err != nil {
		return err
	}

	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "success"), web.NOTICE)
	ss.Save()

	return nil
}

func (p *Engine) indexNotices(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	data["title"] = p.I18n.T(lng, "site.notices.index.title")

	if user, ok := c.Get(auth.CurrentUser); ok && p.Dao.Is(user.(*auth.User).ID, auth.RoleAdmin) {
		data["can"] = true
	}

	var items []Notice
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["items"] = items
	c.HTML(http.StatusOK, "notices", data)
}

func (p *Engine) destoryNotice(c *gin.Context) (interface{}, error) {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Notice{}).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}
