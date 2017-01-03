package auth

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/web"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine auth engine
type Engine struct {
	Dao      *Dao              `inject:""`
	Db       *gorm.DB          `inject:""`
	I18n     *web.I18n         `inject:""`
	Security *web.Security     `inject:""`
	Jwt      *Jwt              `inject:""`
	Server   *machinery.Server `inject:""`
	Session  *Session          `inject:""`
}

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
