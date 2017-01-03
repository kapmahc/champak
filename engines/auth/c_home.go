package auth

import gin "gopkg.in/gin-gonic/gin.v1"

// Home home
func (p *Engine) Home(c *gin.Context) {
	p.getUsersSignIn(c)
}
