package ops

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getSiteInfo(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "ops.site.info.title")
	fm := web.NewForm(c, "site-info", title, "/ops/site/info")
	for _, k := range []string{"title", "subTitle", "keywords", "copyright"} {
		fm.AddFields(web.NewTextField(
			k,
			p.I18n.T(lng, fmt.Sprintf("ops.attributes.site.%s", k)),
			p.I18n.T(lng, fmt.Sprintf("site.%s", k)),
		))
	}
	fm.AddFields(
		web.NewTextArea(
			"description",
			p.I18n.T(lng, "ops.attributes.site.description"),
			p.I18n.T(lng, "site.description"),
		),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/form", data)
}

type fmSiteInfo struct {
	Title       string `form:"title" binding:"required,max=255"`
	SubTitle    string `form:"subTitle" binding:"required,max=32"`
	Keywords    string `form:"keywords" binding:"required,max=255"`
	Description string `form:"description" binding:"required,max=500"`
	Copyright   string `form:"copyright" binding:"required,max=255"`
}

func (p *Engine) postSiteInfo(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmSiteInfo)

	for k, v := range map[string]string{
		"title":       fm.Title,
		"subTitle":    fm.SubTitle,
		"keywords":    fm.Keywords,
		"description": fm.Description,
		"copyright":   fm.Copyright,
	} {
		p.I18n.Set(lng, fmt.Sprintf("site.%s", k), v)
	}

	c.Redirect(http.StatusFound, "/ops/site/info")
	return nil
}

func (p *Engine) getSiteAuthor(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "ops.site.author.title")
	fm := web.NewForm(c, "site-author", title, "/ops/site/author")
	for _, k := range []string{"email", "name"} {
		var v string
		if err := p.Settings.Get(fmt.Sprintf("site.author.%s", k), &v); err != nil {
			log.Error(err)
		}
		fm.AddFields(web.NewTextField(
			k,
			p.I18n.T(lng, fmt.Sprintf("ops.attributes.site.author.%s", k)),
			v,
		))
	}

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/form", data)
}

type fmSiteAuthor struct {
	Email string `form:"email" binding:"email,max=255"`
	Name  string `form:"name" binding:"required,max=32"`
}

func (p *Engine) postSiteAuthor(c *gin.Context, o interface{}) error {
	fm := o.(*fmSiteAuthor)

	for k, v := range map[string]string{
		"name":  fm.Name,
		"email": fm.Email,
	} {
		if err := p.Settings.Set(fmt.Sprintf("site.author.%s", k), v, false); err != nil {
			log.Error(err)
		}
	}

	c.Redirect(http.StatusFound, "/ops/site/author")
	return nil
}
