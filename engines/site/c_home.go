package site

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
)

func (p *Engine) getLocales(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	return p.I18n.Items(ps.ByName("lang")), nil
}

func (p *Engine) getSiteInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (interface{}, error) {
	lng := r.Context().Value(web.LOCALE).(string)
	rst := web.H{}
	for _, k := range []string{"title", "sub_title", "keywords", "description", "copyright"} {
		rst[k] = p.I18n.T(lng, fmt.Sprintf("site.%s", k))
	}
	author := web.H{}
	for _, k := range []string{"name", "email"} {
		author[k] = p.I18n.T(lng, fmt.Sprintf("site.author.%s", k))
	}

	rst[string(web.LOCALE)] = lng
	rst["author"] = author
	rst["languages"] = viper.GetStringSlice("languages")

	for _, k := range []string{"top", "bottom"} {
		var links []Link
		if err := p.Db.Select([]string{"href", "label"}).
			Where("loc = ?", k).
			Order("sort_order ASC").
			Find(&links).Error; err != nil {
			log.Error(err)
		}
		rst[k] = links
	}

	return rst, nil
}
