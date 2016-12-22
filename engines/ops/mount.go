package ops

import (
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	og := rt.Group("/ops", p.Session.MustSignInHandler(), p.Session.MustAdminHandler())

	og.GET("/site/info", p.getSiteInfo)
	og.POST(
		"/site/info",
		web.PostFormHandler("/ops/site/info", &fmSiteInfo{}, p.postSiteInfo),
	)

}
