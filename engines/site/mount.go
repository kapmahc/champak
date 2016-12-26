package site

import "github.com/kapmahc/champak/web"

// Mount mount web points
func (p *Engine) Mount(rt web.Router) {
	rt.GET("/site/info", p.getSiteInfo)
}
