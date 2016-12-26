package site

import "github.com/kapmahc/champak/web"

// Mount mount web points
func (p *Engine) Mount(rt web.Router) {
	rt.GET("/site/info", p.getSiteInfo)
	p.W.Rest(
		rt,
		"/notices",
		p.W.Form(&fmNotice{}, p.createNotice),
		p.W.Form(&fmNotice{}, p.updateNotice),
		p.showNotice,
		p.destroyNotice,
		p.indexNotices,
	)
	p.W.Rest(
		rt,
		"/leave_words",
		p.W.Form(&fmLeaveWord{}, p.createLeaveWord),
		nil,
		p.showLeaveWord,
		p.destroyLeaveWord,
		p.indexLeaveWords,
	)
}
