package site

import "github.com/kapmahc/champak/web"

// Mount mount web points
func (p *Engine) Mount(rt web.Router) {
	rt.GET("/locales/:lang", p.W.JSON(p.getLocales))
	rt.GET("/site/info", p.W.JSON(p.getSiteInfo))

	p.W.Rest(
		rt,
		"/notices",
		p.W.Form(&fmNotice{}, p.createNotice),
		p.W.Form(&fmNotice{}, p.updateNotice),
		p.W.JSON(p.showNotice),
		p.W.JSON(p.destroyNotice),
		p.W.JSON(p.indexNotices),
	)
	p.W.Rest(
		rt,
		"/leave_words",
		p.W.Form(&fmLeaveWord{}, p.createLeaveWord),
		nil,
		p.W.JSON(p.showLeaveWord),
		p.W.JSON(p.destroyLeaveWord),
		p.W.JSON(p.indexLeaveWords),
	)

	// -------------------
	rt.GET("/admin/site/info", p.Jwt.MustAdmin(p.W.JSON(p.getAdminSiteInfo)))
	rt.POST("/admin/site/info", p.Jwt.MustAdmin(p.W.Form(&fmSiteInfo{}, p.postAdminSiteInfo)))
	rt.GET("/admin/site/author", p.Jwt.MustAdmin(p.W.JSON(p.getAdminSiteAuthor)))
	rt.POST("/admin/site/author", p.Jwt.MustAdmin(p.W.Form(&fmSiteAuthor{}, p.postAdminSiteAuthor)))
	rt.GET("/admin/site/seo", p.Jwt.MustAdmin(p.W.JSON(p.getAdminSiteSeo)))
	rt.POST("/admin/site/seo", p.Jwt.MustAdmin(p.W.Form(&fmSiteSeo{}, p.postAdminSiteSeo)))
	rt.GET("/admin/site/smtp", p.Jwt.MustAdmin(p.W.JSON(p.getAdminSiteSMTP)))
	rt.POST("/admin/site/smtp", p.Jwt.MustAdmin(p.W.Form(&fmSiteSMTP{}, p.postAdminSiteSMTP)))
	rt.GET("/admin/status/db", p.Jwt.MustAdmin(p.W.JSON(p.getAdminDbStatus)))
	rt.GET("/admin/status/redis", p.Jwt.MustAdmin(p.W.JSON(p.getAdminRedisStatus)))
	rt.GET("/admin/status/os", p.Jwt.MustAdmin(p.W.JSON(p.getAdminRuntimeStatus)))
}
