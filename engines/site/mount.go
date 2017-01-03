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

	rt.GET("/leave-words/new", p.newLeaveWord)
	rt.POST(
		"/leave-words",
		web.PostFormHandler("/leave-words/new", &fmLeaveWord{}, p.createLeaveWord),
	)
	rt.GET(
		"/leave-words",
		p.Session.MustSignInHandler(), p.Session.MustAdminHandler(),
		p.indexLeaveWords,
	)
	rt.DELETE(
		"/leave-words/:id",
		p.Session.MustSignInHandler(), p.Session.MustAdminHandler(),
		web.FlashHandler("/admin/leave-words", p.destoryLeaveWord),
	)

	rt.GET("/notices", p.indexNotices)
	rt.GET("/notices/new",
		p.Session.MustSignInHandler(), p.Session.MustAdminHandler(),
		p.newNotice,
	)
	rt.POST(
		"/notices",
		p.Session.MustSignInHandler(), p.Session.MustAdminHandler(),
		web.PostFormHandler("/notices/new", &fmNotice{}, p.createNotice),
	)
	rt.POST(
		"/notices/:id",
		p.Session.MustSignInHandler(), p.Session.MustAdminHandler(),
		web.PostFormHandler("/notices", &fmNotice{}, p.updateNotice),
	)
	rt.DELETE(
		"/notices/:id",
		p.Session.MustSignInHandler(), p.Session.MustAdminHandler(),
		web.FlashHandler("/notices", p.destoryNotice),
	)
	rt.GET(
		"/notices/edit/:id",
		p.Session.MustSignInHandler(), p.Session.MustAdminHandler(),
		web.FlashHandler("/notices", p.editNotice),
	)
}
