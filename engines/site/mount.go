package site

import gin "gopkg.in/gin-gonic/gin.v1"

// Mount mount
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/", p.getHome)
	rt.GET("/sitemap.xml.gz", p.getSitemap)
	rt.GET("/rss.atom", p.getRss)

	rt.GET("/dashboard", p.Session.MustSignInHandler(), p.getDashboard)
}
