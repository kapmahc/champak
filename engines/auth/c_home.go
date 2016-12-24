package auth

import (
	"fmt"
	"net/http"

	"github.com/kapmahc/champak/web"
)

func (p *Engine) getHome(wrt http.ResponseWriter, req *http.Request) {

	data := req.Context().Value(web.DATA)
	fmt.Printf("### %+v \n", data)
	p.Render.HTML(wrt, http.StatusOK, "auth/home", data)
}
