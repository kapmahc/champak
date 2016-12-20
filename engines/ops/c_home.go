package ops

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getHome(c *gin.Context) {
	data := c.MustGet("data").(gin.H)
	c.HTML(http.StatusOK, "home", data)
}
