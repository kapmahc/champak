package site

import (
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/", p.getHome)

	rt.GET("/rss.atom", p.getRss)
	rt.GET("/robots.txt", p.getRobots)

	rt.GET("/leave-words/new", p.newLeaveWord)
	rt.POST(
		"/leave-words",
		web.PostFormHandler("/leave-words/new", &fmLeaveWord{}, p.createLeaveWord),
	)

	rt.GET("/dashboard", p.Session.MustSignInHandler(), p.getDashboard)
}
