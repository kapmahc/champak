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
}
