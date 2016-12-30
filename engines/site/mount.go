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

	rt.POST("/leave-words", p.W.Form(&fmBody{}, p.createLeaveWord))
	rt.DELETE("/leave-words/:id", p.Jwt.MustAdmin(p.W.JSON(p.destroyLeaveWord)))
	rt.GET("/leave-words", p.Jwt.MustAdmin(p.W.JSON(p.indexLeaveWords)))

	// -------------------
	rt.GET("/admin/site/info", p.Jwt.MustAdmin(p.W.JSON(p.getAdminSiteInfo)))
	rt.POST("/admin/site/info", p.Jwt.MustAdmin(p.W.Form(&fmSiteInfo{}, p.postAdminSiteInfo)))
	rt.GET("/admin/site/author", p.Jwt.MustAdmin(p.W.JSON(p.getAdminSiteAuthor)))
	rt.POST("/admin/site/author", p.Jwt.MustAdmin(p.W.Form(&fmSiteAuthor{}, p.postAdminSiteAuthor)))
	rt.GET("/admin/site/seo", p.Jwt.MustAdmin(p.W.JSON(p.getAdminSiteSeo)))
	rt.POST("/admin/site/seo", p.Jwt.MustAdmin(p.W.Form(&fmSiteSeo{}, p.postAdminSiteSeo)))
	rt.GET("/admin/site/smtp", p.Jwt.MustAdmin(p.W.JSON(p.getAdminSiteSMTP)))
	rt.POST("/admin/site/smtp", p.Jwt.MustAdmin(p.W.Form(&fmSiteSMTP{}, p.postAdminSiteSMTP)))
	rt.GET("/admin/site/status", p.Jwt.MustAdmin(p.W.JSON(p.getAdminSiteStatus)))
}
