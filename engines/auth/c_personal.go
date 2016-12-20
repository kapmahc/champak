package auth

import (
	"net/http"

	"github.com/kapmahc/champak/web"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getSignIn(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	c.HTML(http.StatusOK, "auth/sign-in", gin.H{
		"locale":    lng,
		"languages": []string{"en-US", "zh-Hans", "zh-Hant"},
	})
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
