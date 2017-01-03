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
	Dao      *auth.Dao         `inject:""`
}

// Do background job
func (p *Engine) Do() {}

// Atom rss-atom
func (p *Engine) Atom() ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{
		{"loc": "/notices", "changefreq": "monthly"},
	}, nil
}

// Dashboard dashboard links
func (p *Engine) Dashboard(c *gin.Context) []web.Dropdown {
	var items []web.Dropdown
	user := c.MustGet(auth.CurrentUser).(*auth.User)
	if p.Dao.Is(user.ID, auth.RoleAdmin) {
		items = append(items, web.Dropdown{
			Label: "site.profile",
			Links: []web.Link{
				{
					Label: "site.admin.info.title",
					Href:  "/admin/site/info",
				},
				{
					Label: "site.admin.author.title",
					Href:  "/admin/site/author",
				},
				{
					Label: "site.admin.seo.title",
					Href:  "/admin/site/seo",
				},
				{
					Label: "site.admin.smtp.title",
					Href:  "/admin/site/smtp",
				},
				{
					Label: "site.admin.status.title",
					Href:  "/admin/site/status",
				},
				{
					Label: "site.notices.index.title",
					Href:  "/notices",
				},
				{
					Label: "site.leave-words.index.title",
					Href:  "/leave-words",
				},
				{
					Label: "site.locales.index.title",
					Href:  "/locales",
				},
				{
					Label: "site.links.index.title",
					Href:  "/links",
				},
				{
					Label: "site.admin.users.title",
					Href:  "/admin/users",
				},
			},
		})
	}
	return items
}

func init() {
	web.Register(&Engine{})
}
