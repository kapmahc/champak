package auth

import (
	"net/http"

	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

type fmVote struct {
	Type  string `form:"type" binding:"required,max=255"`
	ID    uint   `form:"id"`
	Point int    `form:"point"`
}

func (p *Engine) postVotes(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmVote)
	var v Vote
	null := p.Db.
		Where("resource_type = ? AND resource_id = ?", fm.Type, fm.ID).
		First(&v).RecordNotFound()
	if null {
		return p.Db.Create(&Vote{
			ResourceType: fm.Type,
			ResourceID:   fm.ID,
			Point:        fm.Point,
		}).Error
	}
	err := p.Db.Model(&Vote{}).
		Where("resource_type = ? AND resource_id = ?", fm.Type, fm.ID).
		Update("point", v.Point+fm.Point).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": p.I18n.T(lng, "success")})
	}
	return err
}
