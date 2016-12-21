package ops

import gin "gopkg.in/gin-gonic/gin.v1"

// Mount mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/", p.getHome)
	rt.GET("/dashboard", p.getDashboard)
}
