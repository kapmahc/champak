package ops

import (
	"fmt"
	"html/template"

	"github.com/gin-contrib/sessions"
	"github.com/gorilla/csrf"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

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
	c.Writer.Header().Set("X-CSRF-Token", csrf.Token(c.Request))
}

func (p *Engine) loadTemplates(theme string) (*template.Template, error) {
	return template.
		New("").
		Funcs(template.FuncMap{
			"t":     p.I18n.T,
			"links": p.Layout.Links,
			"cards": p.Layout.Cards,
			"fmt":   fmt.Sprintf,
			"eq": func(arg1, arg2 interface{}) bool {
				return arg1 == arg2
			},
			"str2htm": func(s string) template.HTML {
				return template.HTML(s)
			},
		}).
		ParseGlob(
			fmt.Sprintf("themes/%s/views/**/*", theme),
		)

}
