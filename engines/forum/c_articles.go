package forum

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/engines/auth"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) newArticle(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "forum.articles.new.title")
	fm := web.NewForm(c, "new-article", title, "/forum/articles")
	body := web.NewTextArea("body", p.I18n.T(lng, "attributes.body"), "")
	body.Help = p.I18n.T(lng, "helps.markdown")
	fm.AddFields(
		web.NewTextField("title", p.I18n.T(lng, "attributes.title"), ""),
		web.NewTextArea("summary", p.I18n.T(lng, "attributes.summary"), ""),
		body,
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, auth.TplForm, data)
}

type fmArticle struct {
	Title   string `form:"title" binding:"required,max=255"`
	Summary string `form:"summary" binding:"required,max=500"`
	Body    string `form:"body" binding:"required"`
}

func (p *Engine) createArticle(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	user := c.MustGet(auth.CurrentUser).(*auth.User)
	fm := o.(*fmArticle)
	if err := p.Db.Create(&Article{
		Title:   fm.Title,
		Summary: fm.Summary,
		Body:    fm.Body,
		UserID:  user.ID,
	}).Error; err != nil {
		return err
	}

	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "success"), web.NOTICE)
	ss.Save()

	return nil
}

func (p *Engine) editArticle(c *gin.Context) (tpl string, err error) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	id := c.Param("id")
	tpl = auth.TplForm

	var n Article
	if err = p.Db.Where("id = ?", id).First(&n).Error; err != nil {
		return
	}
	if err = p.check(c, n.UserID); err != nil {
		return
	}

	title := p.I18n.T(lng, "forum.articles.edit.title", n.ID)

	fm := web.NewForm(c, "edit-article", title, fmt.Sprintf("/forum/articles/%d", n.ID))
	body := web.NewTextArea("body", p.I18n.T(lng, "attributes.body"), n.Body)
	body.Help = p.I18n.T(lng, "helps.markdown")
	fm.AddFields(
		web.NewTextField("title", p.I18n.T(lng, "attributes.title"), n.Title),
		web.NewTextArea("summary", p.I18n.T(lng, "attributes.summary"), n.Summary),
		body,
	)

	data["title"] = title
	data["form"] = fm

	return
}

func (p *Engine) updateArticle(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmArticle)

	var n Article
	id := c.Param("id")
	if err := p.Db.Select([]string{"user_id"}).Where("id = ?", id).First(&n).Error; err != nil {
		return err
	}
	if err := p.check(c, n.UserID); err != nil {
		return err
	}

	if err := p.Db.Model(&Article{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"title":   fm.Title,
			"summary": fm.Summary,
			"body":    fm.Body,
		}).Error; err != nil {
		return err
	}

	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "success"), web.NOTICE)
	ss.Save()

	return nil
}

func (p *Engine) indexArticles(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	data["title"] = p.I18n.T(lng, "forum.articles.index.title")

	user := c.MustGet(auth.CurrentUser).(*auth.User)
	var items []Article
	var db *gorm.DB
	if !p.Dao.Is(user.ID, auth.RoleAdmin) {
		db = p.Db.Where("user_id = ?", user.ID)
	}
	if err := db.
		Select([]string{"id", "title", "summary"}).
		Order("updated_at DESC").
		Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["items"] = items
	c.HTML(http.StatusOK, "forum/articles", data)
}

func (p *Engine) destoryArticle(c *gin.Context) (interface{}, error) {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Article{}).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Engine) latestArticles(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	data["title"] = p.I18n.T(lng, "forum.articles.latest.title")

	var items []Article
	if err := p.Db.
		Select([]string{"id", "title", "summary"}).
		Order("updated_at DESC").
		Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["items"] = items
	c.HTML(http.StatusOK, "forum/articles/list", data)
}
