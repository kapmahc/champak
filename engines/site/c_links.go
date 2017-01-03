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

func (p *Engine) linkLocSelect(lng, value string) *web.Select {
	var options []web.Option
	for _, v := range []string{"top", "bottom"} {
		options = append(options, web.Option{Label: v, Value: v, Selected: value == v})
	}
	return web.NewSelect("loc", p.I18n.T(lng, "site.attributes.link.loc"), value, options)
}

func (p *Engine) newLink(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "site.links.new.title")
	fm := web.NewForm(c, "new-links", title, "/links")

	fm.AddFields(
		p.linkLocSelect(lng, ""),
		web.NewTextField("href", p.I18n.T(lng, "site.attributes.link.href"), ""),
		web.NewTextField("label", p.I18n.T(lng, "site.attributes.link.label"), ""),
		web.NewOrderSelect("sortOrder", p.I18n.T(lng, "attributes.sortOrder"), 0, -5, 5),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, auth.TplForm, data)
}

type fmLink struct {
	Href      string `form:"href" binding:"required,max=255"`
	Label     string `form:"label" binding:"required,max=255"`
	Loc       string `form:"loc" binding:"required,max=16"`
	SortOrder int    `form:"sortOrder"`
}

func (p *Engine) createLink(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmLink)
	if err := p.Db.Create(
		&Link{
			Href:      fm.Href,
			Label:     fm.Label,
			SortOrder: fm.SortOrder,
			Loc:       fm.Loc,
		}).Error; err != nil {
		return err
	}

	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "success"), web.NOTICE)
	ss.Save()

	return nil
}

func (p *Engine) editLink(c *gin.Context) (tpl string, err error) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	id := c.Param("id")
	tpl = auth.TplForm

	var n Link
	if err = p.Db.Where("id = ?", id).First(&n).Error; err != nil {
		return
	}

	title := p.I18n.T(lng, "site.links.edit.title", n.ID)

	fm := web.NewForm(c, "edit-links", title, fmt.Sprintf("/links/%d", n.ID))
	fm.AddFields(
		p.linkLocSelect(lng, n.Loc),
		web.NewTextField("href", p.I18n.T(lng, "site.attributes.link.href"), n.Href),
		web.NewTextField("label", p.I18n.T(lng, "site.attributes.link.label"), n.Label),
		web.NewOrderSelect("sortOrder", p.I18n.T(lng, "attributes.sortOrder"), n.SortOrder, -5, 5),
	)

	data["title"] = title
	data["form"] = fm

	return
}

func (p *Engine) updateLink(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmLink)

	if err := p.Db.Model(&Link{}).Where("id = ?", c.Param("id")).
		Updates(map[string]interface{}{
			"label":      fm.Label,
			"href":       fm.Href,
			"sort_order": fm.SortOrder,
			"loc":        fm.Loc,
		}).Error; err != nil {
		return err
	}

	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "success"), web.NOTICE)
	ss.Save()

	return nil
}

func (p *Engine) indexLinks(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	data["title"] = p.I18n.T(lng, "site.links.index.title")

	var items []Link
	if err := p.Db.Order("loc ASC, sort_order ASC").Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["items"] = items
	c.HTML(http.StatusOK, "links", data)
}

func (p *Engine) destoryLink(c *gin.Context) (interface{}, error) {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Link{}).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}
