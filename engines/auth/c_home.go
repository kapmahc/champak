package auth

import (
	"net/http"

	"github.com/kapmahc/champak/web"
)

func (p *Engine) getHome(wrt http.ResponseWriter, req *http.Request) {
	data := req.Context().Value(web.DATA).(web.H)
	p.Render.HTML(wrt, http.StatusOK, "home", data)
}
