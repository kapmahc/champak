package ops

import (
	"fmt"
	"html/template"
)

func (p *Engine) loadTemplates(theme string) (*template.Template, error) {
	return template.
		New("").
		Funcs(template.FuncMap{
			"t": p.I18n.T,
		}).
		ParseGlob(
			fmt.Sprintf("themes/%s/views/**/*", theme),
		)

}
