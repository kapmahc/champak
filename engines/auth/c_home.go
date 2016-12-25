package auth

import (
	"net/http"

	"github.com/kapmahc/champak/web"
)

func (p *Engine) getHome(wrt http.ResponseWriter, req *http.Request) (web.H, error) {
	data := req.Context().Value(web.DATA).(web.H)
	return data, nil
}
