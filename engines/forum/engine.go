package forum

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/engines/auth"
	"github.com/kapmahc/champak/web"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine forum engine
type Engine struct {
	Dao     *auth.Dao     `inject:""`
	Db      *gorm.DB      `inject:""`
	I18n    *web.I18n     `inject:""`
	Session *auth.Session `inject:""`
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
func (p *Engine) Dashboard(c *gin.Context) []web.Dropdown {

	user := c.MustGet(auth.CurrentUser).(*auth.User)
	item := web.Dropdown{
		Label: "forum.profile",
		Links: []web.Link{
			{
				Label: "forum.articles.new.title",
				Href:  "/forum/articles/new",
			},
			{
				Label: "forum.articles.index.title",
				Href:  "/forum/articles",
			},
			{
				Label: "forum.comments.index.title",
				Href:  "/forum/comments",
			},
		},
	}
	if p.Dao.Is(user.ID, auth.RoleAdmin) {
		item.Links = append(
			item.Links,
			web.Link{
				Label: "forum.tags.index.title",
				Href:  "/forum/tags",
			},
		)
	}
	return []web.Dropdown{item}
}

// Do background job
func (p *Engine) Do() {

}

// Home home
func (p *Engine) Home(c *gin.Context) {
}

// Shell console commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	web.Register(&Engine{})
}
