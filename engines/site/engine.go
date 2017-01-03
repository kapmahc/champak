package site

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/garyburd/redigo/redis"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/engines/auth"
	"github.com/kapmahc/champak/web"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine site engine
type Engine struct {
	Cache    *web.Cache        `inject:""`
	I18n     *web.I18n         `inject:""`
	Settings *web.Settings     `inject:""`
	Server   *machinery.Server `inject:""`
	Jwt      *auth.Jwt         `inject:""`
	Db       *gorm.DB          `inject:""`
	Redis    *redis.Pool       `inject:""`
	Session  *auth.Session     `inject:""`
}

// Do background job
func (p *Engine) Do() {}

// Atom rss-atom
func (p *Engine) Atom() ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

// Dashboard dashboard links
func (p *Engine) Dashboard(*gin.Context) []web.Dropdown {
	return []web.Dropdown{}
}

func init() {
	web.Register(&Engine{})
}
