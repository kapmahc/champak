package site

import gin "gopkg.in/gin-gonic/gin.v1"

// Mount mount
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/", p.getHome)

	rt.GET("/dashboard", p.Session.MustSignInHandler(), p.getDashboard)
}
