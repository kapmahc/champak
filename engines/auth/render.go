package auth

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"path"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
)

// Render render
type Render struct {
	*render.Render
	router *mux.Router
	Db     *gorm.DB   `inject:""`
	I18n   *web.I18n  `inject:""`
	Cache  *web.Cache `inject:""`
}

// Redirect redirect to
func (p *Render) Redirect(wrt http.ResponseWriter, req *http.Request, name string, args ...interface{}) {
	http.Redirect(wrt, req, p.URLFor(name, args...), http.StatusFound)
}

// ClientIP get client ip
func (p *Render) ClientIP(req *http.Request) string {
	if proxy := req.Header.Get("X-FORWARDED-FOR"); len(proxy) > 0 {
		return proxy
	}
	ip, _, _ := net.SplitHostPort(req.RemoteAddr)
	return ip
}

// URLFor get url by name
func (p *Render) URLFor(name string, args ...interface{}) string {
	var pairs []string
	for _, v := range args {
		switch v.(type) {
		case string:
			pairs = append(pairs, v.(string))
		default:
			pairs = append(pairs, fmt.Sprintf("%v", v))
		}
	}
	if r := p.router.Get(name); r != nil {
		u, e := r.URL(pairs...)
		if e == nil {
			return u.String()
		}
		log.Error(e)
	}
	return "not-found"
}

// Open init config
func (p *Render) Open(rt *mux.Router) {
	p.router = rt
	p.Render = render.New(render.Options{
		Directory:  path.Join("themes", viper.GetString("server.theme"), "views"),
		Layout:     "application",
		Extensions: []string{".html"},
		Funcs: []template.FuncMap{
			{
				"url": p.URLFor,
				"t":   p.I18n.T,
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
					if len(items) > 0 {
						p.Cache.Set(key, items, 60*60*24)
					}
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
					if len(items) > 0 {
						p.Cache.Set(key, items, 60*60*24)
					}
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
