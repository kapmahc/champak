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

func (p *Engine) newTag(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "forum.tags.new.title")
	fm := web.NewForm(c, "new-tag", title, "/forum/tags")
	fm.AddFields(
		web.NewTextField("name", p.I18n.T(lng, "attributes.name"), ""),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, auth.TplForm, data)
}

type fmTag struct {
	Name string `form:"name" binding:"required,max=32"`
}

func (p *Engine) createTag(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmTag)
	if err := p.Db.Create(&Tag{Name: fm.Name}).Error; err != nil {
		return err
	}

	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "success"), web.NOTICE)
	ss.Save()

	return nil
}

func (p *Engine) editTag(c *gin.Context) (tpl string, err error) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	id := c.Param("id")
	tpl = auth.TplForm

	var n Tag
	if err = p.Db.Where("id = ?", id).First(&n).Error; err != nil {
		return
	}

	title := p.I18n.T(lng, "forum.tags.edit.title", n.ID)

	fm := web.NewForm(c, "edit-tag", title, fmt.Sprintf("/forum/tags/%d", n.ID))

	fm.AddFields(
		web.NewTextField("name", p.I18n.T(lng, "attributes.name"), n.Name),
	)

	data["title"] = title
	data["form"] = fm

	return
}

func (p *Engine) updateTag(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmTag)

	if err := p.Db.Model(&Tag{}).Where("id = ?", c.Param("id")).
		Update("name", fm.Name).Error; err != nil {
		return err
	}

	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "success"), web.NOTICE)
	ss.Save()

	return nil
}

func (p *Engine) indexTags(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	data["title"] = p.I18n.T(lng, "forum.tags.index.title")

	if user, ok := c.Get(auth.CurrentUser); ok && p.Dao.Is(user.(*auth.User).ID, auth.RoleAdmin) {
		data["can"] = true
	}

	var items []Tag
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["items"] = items
	c.HTML(http.StatusOK, "forum/tags", data)
}

func (p *Engine) destoryTag(c *gin.Context) (interface{}, error) {
	var t Tag
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		First(&t).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Model(&t).Association("Articles").Clear().Error; err != nil {
		return nil, err
	}
	if err := p.Db.Delete(&t).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Engine) showTag(c *gin.Context) (tpl string, err error) {
	id := c.Param("id")
	tpl = "forum/articles/list"
	var tag Tag
	if err = p.Db.Where("id = ?", id).First(&tag).Error; err != nil {
		return
	}
	if err = p.Db.Model(&tag).Related(&tag.Articles, "Articles").Error; err != nil {
		return
	}
	data := c.MustGet(web.DATA).(gin.H)
	data["articles"] = tag.Articles
	data["title"] = tag.Name
	c.Set(web.DATA, data)
	return
}
