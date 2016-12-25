package web

import (
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/unrolled/render"
)

// KEY http request context key type
type KEY string

const (
	// NOTICE notice-flash
	NOTICE = "notice"
	// ALERT alert-flash
	ALERT = "alert"
)

// H hash
type H map[string]interface{}

// Helper helper
type Helper struct {
	Render *render.Render `inject:""`
}

// HTML render html
func (p *Helper) HTML(name string, f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			ss := sessions.GetSession(r)
			ss.AddFlash(err.Error(), ALERT)
		}
		p.Render.HTML(w, http.StatusOK, name, r.Context().Value(DATA))
	}
}

// JSON json render
func (p *Helper) JSON(f func(w http.ResponseWriter, r *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if val, err := f(w, r); err == nil {
			p.Render.JSON(w, http.StatusOK, val)
		} else {
			p.Render.Text(w, http.StatusInternalServerError, err.Error())
		}
	}
}

// XML xml render
func (p *Helper) XML(f func(w http.ResponseWriter, r *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if val, err := f(w, r); err == nil {
			p.Render.XML(w, http.StatusOK, val)
		} else {
			p.Render.Text(w, http.StatusInternalServerError, err.Error())
		}
	}
}
