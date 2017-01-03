package site

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/google/uuid"
	"github.com/kapmahc/champak/web"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

//http://www.robotstxt.org/robotstxt.html
func (p *Engine) getRobots(c *gin.Context) {
	tpl := `
User-agent: *
Disallow:
Sitemap: %s/sitemap.xml.gz
`
	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(fmt.Sprintf(tpl, web.Home())))

}

func (p *Engine) getRss(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	var ae, an string
	if err := p.Settings.Get("site.author.email", &ae); err != nil {
		log.Error(err)
	}
	if err := p.Settings.Get("site.author.name", &an); err != nil {
		log.Error(err)
	}
	feed := &atom.Feed{
		ID:   uuid.New().String(),
		Link: []atom.Link{{Href: web.Home()}},
		Author: &atom.Person{
			Email: ae,
			Name:  an,
		},
		Title:   p.I18n.T(lng, "site.title"),
		Updated: atom.Time(time.Now()),
		Entry:   make([]*atom.Entry, 0),
	}
	web.Walk(func(en web.Engine) error {
		if items, err := en.Atom(); err == nil {
			feed.Entry = append(feed.Entry, items...)
		} else {
			log.Error(err)
		}
		return nil
	})
	c.XML(http.StatusOK, feed)
}
