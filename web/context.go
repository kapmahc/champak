package web

import (
	"net/http"

	validator "gopkg.in/go-playground/validator.v8"

	"github.com/go-playground/form"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
)

// KEY context key type
type KEY string

// H hash
type H map[string]interface{}

// Handle http handle
type Handle func(http.ResponseWriter, *http.Request, httprouter.Params) (interface{}, error)

// FormHandle http form handle
type FormHandle func(http.ResponseWriter, *http.Request, httprouter.Params, interface{}) (interface{}, error)

// Wrap wrap
type Wrap struct {
	R *render.Render      `inject:""`
	V *validator.Validate `inject:""`
}

// Rest wrap rest handles
func (p *Wrap) Rest(rt Router, path string, create, update httprouter.Handle, show, destroy, index Handle) {
	if index != nil {
		rt.GET(path, p.JSON(index))
	}
	if create != nil {
		rt.POST(path, create)
	}
	path += "/:id"
	if show != nil {
		rt.GET(path, p.JSON(show))
	}
	if update != nil {
		rt.POST(path, update)
	}
	if destroy != nil {
		rt.DELETE(path, p.JSON(destroy))
	}
}

// Form wrap form handle
func (p *Wrap) Form(f interface{}, h FormHandle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		dec := form.NewDecoder()
		err := r.ParseForm()
		var val interface{}

		if err == nil {
			err = dec.Decode(f, r.Form)
		}
		if err == nil {
			val, err = h(w, r, ps, f)
		}
		if err == nil {
			err = p.V.Struct(f)
		}
		if err == nil {
			p.R.JSON(w, http.StatusOK, val)
		} else {
			p.R.Text(w, http.StatusInternalServerError, err.Error())
		}
	}
}

// JSON wrap handle
func (p *Wrap) JSON(h Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if v, e := h(w, r, ps); e == nil {
			p.R.JSON(w, http.StatusOK, v)
		} else {
			p.R.Text(w, http.StatusInternalServerError, e.Error())
		}
	}
}

// XML wrap handle
func (p *Wrap) XML(h Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if v, e := h(w, r, ps); e == nil {
			p.R.XML(w, http.StatusOK, v)
		} else {
			p.R.Text(w, http.StatusInternalServerError, e.Error())
		}
	}
}
