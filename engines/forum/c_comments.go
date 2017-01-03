package forum

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/kapmahc/champak/engines/auth"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) newComment(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "forum.comments.new.title")
	fm := web.NewForm(c, "new-comment", title, "/forum/comments")
	body := web.NewTextArea("body", p.I18n.T(lng, "attributes.body"), "")
	body.Help = p.I18n.T(lng, "helps.markdown")
	fm.AddFields(
		web.NewHiddenField("articleID", c.Query("articleID")),
		body,
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, auth.TplForm, data)
}

type fmComment struct {
	Body      string `form:"body" binding:"required"`
	ArticleID uint   `form:"articleID"`
}

func (p *Engine) createComment(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	user := c.MustGet(auth.CurrentUser).(*auth.User)

	fm := o.(*fmComment)
	if err := p.Db.Create(&Comment{
		Body:      fm.Body,
		UserID:    user.ID,
		ArticleID: fm.ArticleID,
		Type:      web.TypeMARKDOWN,
	}).Error; err != nil {
		return err
	}

	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "success"), web.NOTICE)
	ss.Save()

	c.Set("next", fmt.Sprintf("/forum/articles/show/%d", fm.ArticleID))
	return nil
}

func (p *Engine) editComment(c *gin.Context) (tpl string, err error) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	id := c.Param("id")
	tpl = auth.TplForm

	var n Comment
	if err = p.Db.Where("id = ?", id).First(&n).Error; err != nil {
		return
	}
	if err = p.check(c, n.UserID); err != nil {
		return
	}

	title := p.I18n.T(lng, "forum.comments.edit.title", n.ID)

	fm := web.NewForm(c, "edit-comment", title, fmt.Sprintf("/forum/comments/%d", n.ID))
	body := web.NewTextArea("body", p.I18n.T(lng, "attributes.body"), n.Body)
	body.Help = p.I18n.T(lng, "helps.markdown")
	fm.AddFields(
		body,
	)

	data["title"] = title
	data["form"] = fm

	return
}

func (p *Engine) updateComment(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)

	fm := o.(*fmComment)
	var n Comment
	id := c.Param("id")
	if err := p.Db.Where("id = ?", id).First(&n).Error; err != nil {
		return err
	}
	if err := p.check(c, n.UserID); err != nil {
		return err
	}

	if err := p.Db.Model(&Comment{}).Where("id = ?", id).
		Update("body", fm.Body).Error; err != nil {
		return err
	}

	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "success"), web.NOTICE)
	ss.Save()

	c.Set("next", fmt.Sprintf("/forum/articles/show/%d", n.ArticleID))
	return nil
}

func (p *Engine) indexComments(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	data["title"] = p.I18n.T(lng, "forum.comments.index.title")
	user := c.MustGet(auth.CurrentUser).(*auth.User)

	var items []Comment
	db := p.Db
	if !p.Dao.Is(user.ID, auth.RoleAdmin) {
		db = p.Db.Where("user_id = ?", user.ID)
	}
	if err := db.Order("updated_at DESC").Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["items"] = items
	c.HTML(http.StatusOK, "forum/comments", data)
}

func (p *Engine) destoryComment(c *gin.Context) (interface{}, error) {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Comment{}).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Engine) latestComments(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	data["title"] = p.I18n.T(lng, "forum.comments.latest.title")

	var items []Comment
	if err := p.Db.
		Order("updated_at DESC").
		Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["items"] = items
	c.HTML(http.StatusOK, "forum/comments/list", data)
}
