package ops

import (
	"fmt"
	"html/template"

	"github.com/gorilla/csrf"
	"github.com/kapmahc/champak/web"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func csrfHandler(c *gin.Context) {
	data := c.MustGet(web.DATA).(gin.H)
	data[csrf.TemplateTag] = csrf.TemplateField(c.Request)
	c.Writer.Header().Set("X-CSRF-Token", csrf.Token(c.Request))
	c.Set(web.DATA, data)
}

func (p *Engine) loadTemplates(theme string) (*template.Template, error) {
	return template.
		New("").
		Funcs(template.FuncMap{
			"t":     p.I18n.T,
			"links": p.Layout.Links,
			"cards": p.Layout.Cards,
			"fmt":   fmt.Sprintf,
		}).
		ParseGlob(
			fmt.Sprintf("themes/%s/views/**/*", theme),
		)

}
