package site

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getAdminSiteSMTP(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "site.admin.smtp.title")

	var v SMTP
	if err := p.Settings.Get("site.smtp", &v); err != nil {
		log.Error(err)
		v.Port = 25
		v.Host = "localhost"
	}
	fm := web.NewForm(c, "site-smtp", title, "/admin/site/smtp")
	pwd := web.NewPasswordField("password", p.I18n.T(lng, "attributes.password"))
	pwd.Help = p.I18n.T(lng, "helps.password")
	pwc := web.NewPasswordField("passwordConfirmation", p.I18n.T(lng, "attributes.passwordConfirmation"))
	pwc.Help = p.I18n.T(lng, "helps.passwordConfirmation")

	ports := []interface{}{25, 465, 587}
	fm.AddFields(
		web.NewTextField("host", p.I18n.T(lng, "attributes.host"), v.Host),
		web.NewSelect("port", p.I18n.T(lng, "attributes.port"), v.Port, ports...),
		web.NewEmailField("user", p.I18n.T(lng, "attributes.user"), v.User),
		pwd, pwc,
		web.NewCheckbox("ssl", p.I18n.T(lng, "attributes.ssl"), v.Ssl),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/form", data)
}

// SMTP smtp config
type SMTP struct {
	Host                 string `form:"host" binding:"required"`
	User                 string `form:"user" binding:"required,email"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
	Port                 int    `form:"port"`
	Ssl                  bool   `form:"ssl"`
}

func (p *Engine) postAdminSiteSMTP(c *gin.Context, o interface{}) error {
	fm := o.(*SMTP)
	if err := p.Settings.Set("site.smtp", fm, true); err != nil {
		log.Error(err)
	}
	c.Redirect(http.StatusFound, "/admin/site/smtp")
	return nil
}

func (p *Engine) getAdminSiteSeo(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "site.admin.seo.title")
	fm := web.NewForm(c, "site-seo", title, "/admin/site/seo")
	for _, k := range []string{"baiduVerifyID", "googleVerifyID"} {
		var v string
		if err := p.Settings.Get(fmt.Sprintf("site.%s", k), &v); err != nil {
			log.Error(err)
		}
		data[k] = v
		fm.AddFields(web.NewTextField(
			k,
			p.I18n.T(lng, fmt.Sprintf("site.attributes.%s", k)),
			v,
		))
	}

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "admin/site/seo", data)
}

type fmSiteSeo struct {
	BaiduVerifyID  string `form:"baiduVerifyID"`
	GoogleVerifyID string `form:"googleVerifyID"`
}

func (p *Engine) postAdminSiteSeo(c *gin.Context, o interface{}) error {
	fm := o.(*fmSiteSeo)
	for k, v := range map[string]string{
		"googleVerifyID": fm.GoogleVerifyID,
		"baiduVerifyID":  fm.BaiduVerifyID,
	} {
		if err := p.Settings.Set(fmt.Sprintf("site.%s", k), v, false); err != nil {
			log.Error(err)
		}
	}

	c.Redirect(http.StatusFound, "/admin/site/seo")
	return nil
}

func (p *Engine) databaseStatus() gin.H {
	type Status struct {
		Version string
	}
	var sts Status
	p.Db.Raw("SELECT VERSION() AS version").Scan(&sts)
	return gin.H{
		"version": sts.Version,
	}
}
func (p *Engine) runtimeStatus() gin.H {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return gin.H{
		"GO":           runtime.Version(),
		"OS":           runtime.GOOS,
		"ARCH":         runtime.GOARCH,
		"CPUS":         runtime.NumCPU(),
		"MEMORY USAGE": fmt.Sprintf("%d/%d MB", mem.Alloc/(1024*1024), mem.Sys/(1024*1024)),
		"LAST GC":      time.Unix(0, int64(mem.LastGC)),
	}
}
func (p *Engine) redisStatus() string {
	c := p.Redis.Get()
	defer c.Close()

	sts, err := redis.String(c.Do("INFO"))
	if err != nil {
		log.Error(err)
	}
	return sts
}
func (p *Engine) cacheStatus() gin.H {
	return gin.H{}
}
func (p *Engine) jobStatus() gin.H {
	return gin.H{}
}

func (p *Engine) getAdminSiteStatus(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	title := p.I18n.T(lng, "site.admin.status.title")
	status := gin.H{
		"site.admin.status.os":       p.runtimeStatus(),
		"site.admin.status.database": p.databaseStatus(),
		"site.admin.status.cache":    p.cacheStatus(),
		"site.admin.status.jobs":     p.jobStatus(),
	}
	data["title"] = title
	data["status"] = status
	data["redis"] = p.redisStatus()
	c.HTML(http.StatusOK, "admin/site/status", data)
}

func (p *Engine) getAdminSiteInfo(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "site.admin.info.title")
	fm := web.NewForm(c, "site-info", title, "/admin/site/info")
	for _, k := range []string{"title", "subTitle", "keywords", "copyright"} {
		fm.AddFields(web.NewTextField(
			k,
			p.I18n.T(lng, fmt.Sprintf("site.attributes.%s", k)),
			p.I18n.T(lng, fmt.Sprintf("site.%s", k)),
		))
	}
	fm.AddFields(
		web.NewTextArea(
			"description",
			p.I18n.T(lng, "site.attributes.description"),
			p.I18n.T(lng, "site.description"),
		),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/form", data)
}

type fmSiteInfo struct {
	Title       string `form:"title" binding:"required,max=255"`
	SubTitle    string `form:"subTitle" binding:"required,max=32"`
	Keywords    string `form:"keywords" binding:"required,max=255"`
	Description string `form:"description" binding:"required,max=500"`
	Copyright   string `form:"copyright" binding:"required,max=255"`
}

func (p *Engine) postAdminSiteInfo(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmSiteInfo)

	for k, v := range map[string]string{
		"title":       fm.Title,
		"subTitle":    fm.SubTitle,
		"keywords":    fm.Keywords,
		"description": fm.Description,
		"copyright":   fm.Copyright,
	} {
		p.I18n.Set(lng, fmt.Sprintf("site.%s", k), v)
	}

	c.Redirect(http.StatusFound, "/admin/site/info")
	return nil
}

func (p *Engine) getAdminSiteAuthor(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "site.admin.author.title")
	fm := web.NewForm(c, "site-author", title, "/admin/site/author")
	for _, k := range []string{"email", "name"} {
		var v string
		if err := p.Settings.Get(fmt.Sprintf("site.author.%s", k), &v); err != nil {
			log.Error(err)
		}
		fm.AddFields(web.NewTextField(
			k,
			p.I18n.T(lng, fmt.Sprintf("site.attributes.author.%s", k)),
			v,
		))
	}

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/form", data)
}

type fmSiteAuthor struct {
	Email string `form:"email" binding:"email,max=255"`
	Name  string `form:"name" binding:"required,max=32"`
}

func (p *Engine) postAdminSiteAuthor(c *gin.Context, o interface{}) error {
	fm := o.(*fmSiteAuthor)

	for k, v := range map[string]string{
		"name":  fm.Name,
		"email": fm.Email,
	} {
		if err := p.Settings.Set(fmt.Sprintf("site.author.%s", k), v, false); err != nil {
			log.Error(err)
		}
	}

	c.Redirect(http.StatusFound, "/admin/site/author")
	return nil
}
