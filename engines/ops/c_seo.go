package ops

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getSiteSeo(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "ops.site.seo.title")
	fm := web.NewForm(c, "site-seo", title, "/ops/site/seo")
	for _, k := range []string{"baiduVerifyID", "googleVerifyID"} {
		var v string
		if err := p.Settings.Get(fmt.Sprintf("site.%s", k), &v); err != nil {
			log.Error(err)
		}
		data[k] = v
		fm.AddFields(web.NewTextField(
			k,
			p.I18n.T(lng, fmt.Sprintf("ops.attributes.site.%s", k)),
			v,
		))
	}

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "ops/site/seo", data)
}

type fmSiteSeo struct {
	BaiduVerifyID  string `form:"baiduVerifyID"`
	GoogleVerifyID string `form:"googleVerifyID"`
}

func (p *Engine) postSiteSeo(c *gin.Context, o interface{}) error {
	fm := o.(*fmSiteSeo)
	for k, v := range map[string]string{
		"googleVerifyID": fm.GoogleVerifyID,
		"baiduVerifyID":  fm.BaiduVerifyID,
	} {
		if err := p.Settings.Set(fmt.Sprintf("site.%s", k), v, false); err != nil {
			log.Error(err)
		}
	}

	c.Redirect(http.StatusFound, "/ops/site/seo")
	return nil
}
