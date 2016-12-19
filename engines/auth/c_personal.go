package auth

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getSignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "sign-in.html", gin.H{})
}
func (p *Engine) postSignIn(c *gin.Context) {

}

func (p *Engine) getSignUp(c *gin.Context) {

}
func (p *Engine) postSignUp(c *gin.Context) {

}

func (p *Engine) getConfirm(c *gin.Context) {

}
func (p *Engine) postConfirm(c *gin.Context) {

}

func (p *Engine) getForgotPassword(c *gin.Context) {

}
func (p *Engine) postForgotPassword(c *gin.Context) {

}

func (p *Engine) getChangePassword(c *gin.Context) {

}
func (p *Engine) postChangePassword(c *gin.Context) {

}

func (p *Engine) getUnlock(c *gin.Context) {

}
func (p *Engine) postUnlock(c *gin.Context) {

}

func (p *Engine) getProfile(c *gin.Context) {

}
func (p *Engine) postProfile(c *gin.Context) {

}

func (p *Engine) deleteSignOut(c *gin.Context) {

}
