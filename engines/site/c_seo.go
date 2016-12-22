package site

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/google/uuid"
	"github.com/kapmahc/champak/web"
	"github.com/kapmahc/champak/web/sitemap"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getRobots(c *gin.Context) {

}
func (p *Engine) getSitemap(c *gin.Context) {

	sm := sitemap.New()
	web.Loop(func(en web.Engine) error {
		if items, err := en.Sitemap(); err == nil {
			sm.Items = append(sm.Items, items...)
		} else {
			log.Error(err)
		}
		return nil
	})
	home := web.HostURL()
	for k := range sm.Items {
		sm.Items[k].Link = home + sm.Items[k].Link
	}
	c.XML(http.StatusOK, sm)
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
		Link: []atom.Link{{Href: web.HostURL()}},
		Author: &atom.Person{
			Email: ae,
			Name:  an,
		},
		Title:   p.I18n.T(lng, "site.title"),
		Updated: atom.Time(time.Now()),
		Entry:   make([]*atom.Entry, 0),
	}
	web.Loop(func(en web.Engine) error {
		if items, err := en.Atom(); err == nil {
			feed.Entry = append(feed.Entry, items...)
		} else {
			log.Error(err)
		}
		return nil
	})
	c.XML(http.StatusOK, feed)
}
