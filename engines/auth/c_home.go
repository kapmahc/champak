package auth

import gin "gopkg.in/gin-gonic/gin.v1"

// Home home
func (p *Engine) Home() gin.HandlerFunc {
	return p.getUsersSignIn
}
