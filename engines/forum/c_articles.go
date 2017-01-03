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

func (p *Engine) tagsSelect(lng string, art *Article) *web.Select {
	var options []web.Option
	var tags []Tag
	if err := p.Db.
		Select([]string{"id", "name"}).
		Order("name ASC").
		Find(&tags).Error; err != nil {
		log.Error(err)
	}
	if art != nil {
		if err := p.Db.Model(art).Related(&art.Tags, "Tags").Error; err != nil {
			log.Error(err)
		}
	}
	for _, i := range tags {
		o := web.Option{Label: i.Name, Value: i.ID}
		if art != nil {
			for _, j := range art.Tags {
				if i.ID == j.ID {
					o.Selected = true
					break
				}
			}
		}
		options = append(options, o)
	}
	sel := web.NewSelect("tags", p.I18n.T(lng, "forum.attributes.articles.tags"), options)
	sel.Multiple = true
	return sel
}

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
		p.tagsSelect(lng, nil),
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
	Tags    []uint `form:"tags"`
}

func (p *Engine) createArticle(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	user := c.MustGet(auth.CurrentUser).(*auth.User)
	fm := o.(*fmArticle)

	var tags []Tag
	for _, i := range fm.Tags {
		var t Tag
		if err := p.Db.Where("id = ?", i).First(&t).Error; err != nil {
			return err
		}
		tags = append(tags, t)
	}

	if err := p.Db.Create(&Article{
		Title:   fm.Title,
		Summary: fm.Summary,
		Body:    fm.Body,
		UserID:  user.ID,
		Tags:    tags,
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
		p.tagsSelect(lng, &n),
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
	if err := p.Db.Where("id = ?", id).First(&n).Error; err != nil {
		return err
	}
	if err := p.check(c, n.UserID); err != nil {
		return err
	}

	var tags []Tag
	for _, i := range fm.Tags {
		var t Tag
		if err := p.Db.Where("id = ?", i).First(&t).Error; err != nil {
			return err
		}
		tags = append(tags, t)
	}

	if err := p.Db.Model(&Article{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"title":   fm.Title,
			"summary": fm.Summary,
			"body":    fm.Body,
		}).Error; err != nil {
		return err
	}

	if err := p.Db.Model(&n).Association("Tags").Replace(tags).Error; err != nil {
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
	db := p.Db
	if !p.Dao.Is(user.ID, auth.RoleAdmin) {
		db = p.Db.Where("user_id = ?", user.ID)
	}
	if err := db.
		Select([]string{"id", "title", "updated_at"}).
		Order("updated_at DESC").
		Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["items"] = items
	c.HTML(http.StatusOK, "forum/articles", data)
}

func (p *Engine) destoryArticle(c *gin.Context) (interface{}, error) {
	var a Article
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		First(&a).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Model(&a).Association("Tags").Clear().Error; err != nil {
		return nil, err
	}
	if err := p.Db.Where("article_id = ?", a.ID).Delete(Comment{}).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Delete(&a).Error; err != nil {
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
		Select([]string{"id", "title", "summary", "updated_at"}).
		Order("updated_at DESC").
		Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["articles"] = items
	c.HTML(http.StatusOK, "forum/articles/list", data)
}

func (p *Engine) showArticle(c *gin.Context) (tpl string, err error) {
	id := c.Param("id")
	tpl = "forum/articles/show"
	var art Article
	if err = p.Db.Where("id = ?", id).First(&art).Error; err != nil {
		return
	}
	if err = p.Db.Model(&art).Related(&art.Tags, "Tags").Error; err != nil {
		return
	}
	if err = p.Db.Model(&art).Related(&art.Comments).Error; err != nil {
		return
	}
	if err = p.Db.Model(&art).Related(&art.User).Error; err != nil {
		return
	}
	data := c.MustGet(web.DATA).(gin.H)
	data["article"] = art
	data["title"] = art.Title
	c.Set(web.DATA, data)
	return
}
