package ops

import (
	"html/template"
	"path"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/kapmahc/multitemplate"
)

func (p *Engine) loadTemplates(dir string) (multitemplate.Render, error) {
	r := multitemplate.New()

	layouts, err := filepath.Glob(path.Join(dir, "layouts", "*.html"))
	if err != nil {
		return nil, err
	}

	includes, err := filepath.Glob(path.Join(dir, "includes", "*.html"))
	if err != nil {
		return nil, err
	}

	log.Debug(layouts)
	log.Debug(includes)

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		name := filepath.Base(layout)
		log.Debugf("add %s %+v", name, files)
		r.Add(name, template.Must(template.ParseFiles(files...)))
	}
	return r, nil
}
