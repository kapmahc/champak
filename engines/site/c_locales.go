package site

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) newLocale(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "site.locales.new.title")
	fm := web.NewForm(c, "new-locales", title, "/locales")

	fm.AddFields(
		web.NewTextField("code", p.I18n.T(lng, "site.attributes.locale.code"), ""),
		web.NewTextArea("message", p.I18n.T(lng, "site.attributes.locale.message"), ""),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/form", data)
}

type fmLocale struct {
	Code    string `form:"code" binding:"required,max=255"`
	Message string `form:"message" binding:"required"`
}

func (p *Engine) saveLocale(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmLocale)
	if err := p.I18n.Set(lng, fm.Code, fm.Message); err != nil {
		return err
	}

	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "success"), web.NOTICE)
	ss.Save()

	return nil
}

func (p *Engine) editLocale(c *gin.Context) (tpl string, err error) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	id := c.Param("id")
	tpl = "auth/form"

	var l web.Locale
	if err = p.Db.Where("id = ?", id).First(&l).Error; err != nil {
		return
	}

	title := p.I18n.T(lng, "site.locales.edit.title", l.ID)

	fm := web.NewForm(c, "edit-locales", title, "/locales")
	fm.AddFields(
		web.NewTextField("code", p.I18n.T(lng, "site.attributes.locale.code"), l.Code),
		web.NewTextArea("message", p.I18n.T(lng, "site.attributes.locale.message"), l.Message),
	)

	data["title"] = title
	data["form"] = fm

	return
}

func (p *Engine) indexLocales(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	data["title"] = p.I18n.T(lng, "site.locales.index.title")

	var items []web.Locale
	if err := p.Db.
		Select([]string{"id", "code", "message"}).
		Where("lang = ?", lng).
		Order("code ASC").Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["items"] = items
	c.HTML(http.StatusOK, "locales", data)
}

func (p *Engine) destoryLocale(c *gin.Context) (interface{}, error) {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(web.Locale{}).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}
