package site

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	"github.com/julienschmidt/httprouter"
	"github.com/kapmahc/champak/web"
)

func (p *Engine) getAdminSiteInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	lng := r.Context().Value(web.LOCALE).(string)
	info := web.H{}
	for _, k := range []string{"title", "subTitle", "keywords", "copyright", "description"} {
		info[k] = p.I18n.T(lng, fmt.Sprintf("site.%s", k))
	}
	return info, nil
}

func (p *Engine) postAdminSiteInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params, o interface{}) (interface{}, error) {
	lng := r.Context().Value(web.LOCALE).(string)
	fm := o.(*fmSiteInfo)
	p.I18n.Set(lng, "site.title", fm.Title)
	p.I18n.Set(lng, "site.subTitle", fm.SubTitle)
	p.I18n.Set(lng, "site.keywords", fm.Keywords)
	p.I18n.Set(lng, "site.description", fm.Description)
	p.I18n.Set(lng, "site.copyright", fm.Copyright)

	return web.H{}, nil
}

func (p *Engine) getAdminSiteAuthor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	author := web.H{}
	for _, k := range []string{"name", "email"} {
		var v string
		if err := p.Settings.Get(fmt.Sprintf("site.author.%s", k), &v); err != nil {
			log.Error(err)
		}
		author[k] = v
	}
	return author, nil
}

func (p *Engine) postAdminSiteAuthor(w http.ResponseWriter, r *http.Request, ps httprouter.Params, o interface{}) (interface{}, error) {
	fm := o.(*fmSiteAuthor)
	p.Settings.Set("site.author.name", fm.Name, false)
	p.Settings.Set("site.author.email", fm.Email, false)
	return web.H{}, nil
}

func (p *Engine) getAdminDbStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	type Status struct {
		Version string
	}
	var sts Status
	if err := p.Db.Raw("SELECT VERSION() AS version").Scan(&sts).Error; err != nil {
		return nil, err
	}
	return web.H{
		"version": sts.Version,
	}, nil
}

func (p *Engine) getAdminRedisStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	c := p.Redis.Get()
	defer c.Close()

	sts, err := redis.String(c.Do("INFO"))
	if err != nil {
		return nil, err
	}
	return string(sts), nil
}

func (p *Engine) getAdminRuntimeStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return web.H{
		"GO":           runtime.Version(),
		"OS":           runtime.GOOS,
		"ARCH":         runtime.GOARCH,
		"CPUS":         runtime.NumCPU(),
		"MEMORY USAGE": fmt.Sprintf("%d/%d MB", mem.Alloc/(1024*1024), mem.Sys/(1024*1024)),
		"LAST GC":      time.Unix(0, int64(mem.LastGC)),
	}, nil
}

func (p *Engine) getAdminSiteSeo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	seo := web.H{}
	for _, k := range []string{"google", "baidu"} {
		var v string
		if err := p.Settings.Get(fmt.Sprintf("site.seo.%s_verify_id", k), &v); err != nil {
			log.Error(err)
		}
		seo[k] = v
	}
	return seo, nil
}

func (p *Engine) postAdminSiteSeo(w http.ResponseWriter, r *http.Request, ps httprouter.Params, o interface{}) (interface{}, error) {
	fm := o.(*fmSiteSeo)
	p.Settings.Set("site.seo.google_verify_id", fm.Google, false)
	p.Settings.Set("site.seo.baidu_verify_id", fm.Baidu, false)
	return web.H{}, nil
}

func (p *Engine) getAdminSiteSMTP(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	var val SMTP
	if err := p.Settings.Get("site.smtp", &val); err != nil {
		log.Error(err)
	}
	return val, nil
}

func (p *Engine) postAdminSiteSMTP(w http.ResponseWriter, r *http.Request, ps httprouter.Params, o interface{}) (interface{}, error) {
	fm := o.(*fmSiteSMTP)
	if err := p.Settings.Set("site.smtp", &SMTP{
		Host:     fm.Host,
		Port:     fm.Port,
		User:     fm.User,
		Password: fm.Password,
		Ssl:      fm.Ssl,
	}, true); err != nil {
		return nil, err
	}
	return web.H{}, nil
}
