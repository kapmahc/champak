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

	// ---------------------

	rt.GET("/dashboard", p.Session.MustSignInHandler(), p.getDashboard)

	// ---------------------------
	ag := rt.Group("/admin", p.Session.MustSignInHandler(), p.Session.MustAdminHandler())

	ag.GET("/site/info", p.getAdminSiteInfo)
	ag.POST(
		"/site/info",
		web.PostFormHandler("/admin/site/info", &fmSiteInfo{}, p.postAdminSiteInfo),
	)
	ag.GET("/site/author", p.getAdminSiteAuthor)
	ag.POST(
		"/site/author",
		web.PostFormHandler("/admin/site/author", &fmSiteAuthor{}, p.postAdminSiteAuthor),
	)
	ag.GET("/site/seo", p.getAdminSiteSeo)
	ag.POST(
		"/site/seo",
		web.PostFormHandler("/admin/site/seo", &fmSiteSeo{}, p.postAdminSiteSeo),
	)
	ag.GET("/site/smtp", p.getAdminSiteSMTP)
	ag.POST(
		"/site/smtp",
		web.PostFormHandler("/admin/site/smtp", &SMTP{}, p.postAdminSiteSMTP),
	)
	ag.GET("/site/status", p.getAdminSiteStatus)

	ag.GET("/leave-words", p.indexLeaveWords)
	ag.DELETE("/leave-words/:id", web.FlashHandler("/ops/leave-words", p.destoryLeaveWord))

	// ag.GET("/notices", p.indexNotices)
	// ag.GET("/notices/new", p.newNotice)
	// ag.POST("/notices", web.PostFormHandler("/ops/notices", &fmNotice{}, p.createNotice))
	// ag.POST("/notices/:id", web.PostFormHandler("/ops/notices", &fmNotice{}, p.updateNotice))
	// ag.DELETE("/notices/:id", web.FlashHandler("/ops/notices", p.destoryNotice))
	// ag.GET("/notices/:id/edit", web.FlashHandler("/ops/notices", p.editNotice))
}
