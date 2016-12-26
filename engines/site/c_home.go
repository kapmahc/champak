package site

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
)

func (p *Engine) getSiteInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	p.R.JSON(w, http.StatusOK, rst)
}
