package site

import (
	"fmt"
	"html/template"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/gorilla/csrf"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) authorHandler(c *gin.Context) {
	var ae, an string
	if err := p.Settings.Get("site.author.email", &ae); err != nil {
		log.Error(err)
	}
	if err := p.Settings.Get("site.author.name", &an); err != nil {
		log.Error(err)
	}
	data := c.MustGet(web.DATA).(gin.H)
	data["author"] = gin.H{"name": an, "email": ae}
	c.Set(web.DATA, data)
}
func flashsHandler(c *gin.Context) {
	data := c.MustGet(web.DATA).(gin.H)
	ss := sessions.Default(c)
	for _, k := range []string{web.ALERT, web.NOTICE} {
		data[k] = ss.Flashes(k)
	}
	ss.Save()
	c.Set(web.DATA, data)
}

func csrfHandler(c *gin.Context) {
	tkn := csrf.Token(c.Request)
	c.Writer.Header().Set("X-CSRF-Token", tkn)
	data := c.MustGet(web.DATA).(gin.H)
	data["csrf"] = tkn
	c.Set(web.DATA, data)
}

func (p *Engine) loadTemplates(theme string) (*template.Template, error) {
	return template.
		New("").
		Funcs(template.FuncMap{
			"t": p.I18n.T,
			// "links": p.Layout.Links,
			// "cards": p.Layout.Cards,
			"fmt": fmt.Sprintf,
			"eq": func(arg1, arg2 interface{}) bool {
				return arg1 == arg2
			},
			"str2htm": func(s string) template.HTML {
				return template.HTML(s)
			},
			"dtf": func(t time.Time) string {
				return t.Format("Mon Jan _2 15:04:05 2006")
			},
		}).
		ParseGlob(
			fmt.Sprintf("themes/%s/views/**/*", theme),
		)

}
