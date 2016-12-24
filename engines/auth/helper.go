package auth

import (
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/kapmahc/champak/web"
	"github.com/unrolled/render"
)

// Helper controller wrapper
type Helper struct {
	Render *Render `inject:""`
}

// JSON render json
func (p *Helper) JSON(fn func(http.ResponseWriter, *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(wrt http.ResponseWriter, req *http.Request) {
		if val, err := fn(wrt, req); err == nil {
			p.Render.JSON(wrt, http.StatusOK, val)
		} else {
			p.Render.Text(wrt, http.StatusInternalServerError, err.Error())
		}
	}
}

// HTML render html
func (p *Helper) HTML(view, layout, ego string, fn func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(wrt http.ResponseWriter, req *http.Request) {
		if err := fn(wrt, req); err == nil {
			p.Render.HTML(
				wrt,
				http.StatusOK,
				view,
				req.Context().Value(web.DATA),
				render.HTMLOptions{Layout: layout},
			)
		} else {
			ss := sessions.GetSession(req)
			ss.AddFlash(err.Error(), web.ALERT)
			http.Redirect(wrt, req, ego, http.StatusFound)
		}
	}
}
