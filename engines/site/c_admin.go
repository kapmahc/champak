package site

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"github.com/kapmahc/champak/web"
)

func (p *Engine) getAdminSiteInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (interface{}, error) {
	lng := r.Context().Value(web.LOCALE).(string)
	info := web.H{}
	for _, k := range []string{"title", "subTitle", "keywords", "copyright", "description"} {
		info[k] = p.I18n.T(lng, fmt.Sprintf("site.%s", k))
	}
	return info, nil
}

func (p *Engine) postAdminSiteInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params, o interface{}) (interface{}, error) {
	lng := r.Context().Value(web.LOCALE).(string)
	fm := o.(*fmSiteInfo)
	p.I18n.Set(lng, "site.title", fm.Title)
	p.I18n.Set(lng, "site.subTitle", fm.SubTitle)
	p.I18n.Set(lng, "site.keywords", fm.Keywords)
	p.I18n.Set(lng, "site.description", fm.Description)
	p.I18n.Set(lng, "site.copyright", fm.Copyright)

	return web.H{}, nil
}

func (p *Engine) getAdminSiteAuthor(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (interface{}, error) {
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

func (p *Engine) postAdminSiteAuthor(w http.ResponseWriter, r *http.Request, _ httprouter.Params, o interface{}) (interface{}, error) {
	fm := o.(*fmSiteAuthor)
	p.Settings.Set("site.author.name", fm.Name, false)
	p.Settings.Set("site.author.email", fm.Email, false)
	return web.H{}, nil
}

func (p *Engine) getAdminSiteSeo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (interface{}, error) {
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

func (p *Engine) postAdminSiteSeo(w http.ResponseWriter, r *http.Request, _ httprouter.Params, o interface{}) (interface{}, error) {
	fm := o.(*fmSiteSeo)
	p.Settings.Set("site.seo.google_verify_id", fm.Google, false)
	p.Settings.Set("site.seo.baidu_verify_id", fm.Baidu, false)
	return web.H{}, nil
}

func (p *Engine) getAdminSiteSMTP(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (interface{}, error) {
	var val SMTP
	if err := p.Settings.Get("site.smtp", &val); err != nil {
		log.Error(err)
	}
	return val, nil
}

func (p *Engine) postAdminSiteSMTP(w http.ResponseWriter, r *http.Request, _ httprouter.Params, o interface{}) (interface{}, error) {
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
