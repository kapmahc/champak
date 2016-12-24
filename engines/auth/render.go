package auth

import (
	"fmt"
	"html/template"
	"path"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
)

// Render render
type Render struct {
	*render.Render
	Db    *gorm.DB   `inject:""`
	I18n  *web.I18n  `inject:""`
	Cache *web.Cache `inject:""`
}

// Open init config
func (p *Render) Open() {
	p.Render = render.New(render.Options{
		Directory:  path.Join("themes", viper.GetString("server.theme"), "views"),
		Layout:     "application",
		Extensions: []string{".html"},
		Funcs: []template.FuncMap{
			{
				"t": p.I18n.T,
				"cards": func(loc string) []Card {
					key := fmt.Sprintf("cards/%s", loc)
					var items []Card
					if err := p.Cache.Get(key, &items); err == nil {
						return items
					}
					if err := p.Db.
						Select([]string{"title", "summary", "logo", "href"}).
						Where("loc = ?", loc).
						Order("sort_order ASC").
						Find(&items).Error; err != nil {
						log.Error(err)
					}
					p.Cache.Set(key, items, 60*60*24)
					return items
				},
				"links": func(loc string) []Link {
					key := fmt.Sprintf("links/%s", loc)
					var items []Link
					if err := p.Cache.Get(key, &items); err == nil {
						return items
					}
					if err := p.Db.
						Select([]string{"label", "href"}).
						Where("loc = ?", loc).
						Order("sort_order ASC").
						Find(&items).Error; err != nil {
						log.Error(err)
					}
					p.Cache.Set(key, items, 60*60*24)
					return items
				},
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
			},
		},
		IndentJSON:    !IsProduction(),
		IndentXML:     !IsProduction(),
		IsDevelopment: !IsProduction(),
	})
}
