package site

import "github.com/kapmahc/champak/web"

// Mount mount web points
func (p *Engine) Mount(rt web.Router) {
	rt.GET("/site/info", p.W.JSON(p.getSiteInfo, true))
	p.W.Rest(
		rt,
		"/notices",
		p.W.Form(&fmNotice{}, p.createNotice),
		p.W.Form(&fmNotice{}, p.updateNotice),
		p.W.JSON(p.showNotice, true),
		p.W.JSON(p.destroyNotice, false),
		p.W.JSON(p.indexNotices, true),
	)
	p.W.Rest(
		rt,
		"/leave_words",
		p.W.Form(&fmLeaveWord{}, p.createLeaveWord),
		nil,
		p.W.JSON(p.showLeaveWord, true),
		p.W.JSON(p.destroyLeaveWord, false),
		p.W.JSON(p.indexLeaveWords, true),
	)
}
