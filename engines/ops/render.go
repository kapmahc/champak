package ops

import (
	"fmt"
	"html/template"

	"github.com/gorilla/csrf"

	gin "gopkg.in/gin-gonic/gin.v1"
)

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
		}).
		ParseGlob(
			fmt.Sprintf("themes/%s/views/**/*", theme),
		)

}
