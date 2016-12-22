package ops

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/kapmahc/champak/engines/site"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexLeaveWords(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	data["title"] = p.I18n.T(lng, "ops.leave-words.title")
	var items []site.LeaveWord
	if err := p.Db.Order("created_at DESC").Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["items"] = items
	c.HTML(http.StatusOK, "ops/leave-words", data)
}

func (p *Engine) destoryLeaveWord(c *gin.Context) error {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(site.LeaveWord{}).Error; err != nil {
		return err
	}
	c.JSON(http.StatusOK, gin.H{web.TO: "/ops/leave-words"})
	return nil
}
