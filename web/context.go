package web

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

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
	C *Cache              `inject:""`
}

// Rest wrap rest handles
func (p *Wrap) Rest(rt Router, path string, create, update, show, destroy, index httprouter.Handle) {
	if index != nil {
		rt.GET(path, index)
	}
	if create != nil {
		rt.POST(path, create)
	}
	path += "/:id"
	if show != nil {
		rt.GET(path, show)
	}
	if update != nil {
		rt.POST(path, update)
	}
	if destroy != nil {
		rt.DELETE(path, destroy)
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
func (p *Wrap) JSON(h Handle, c bool) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		const ct = "application/json; charset=UTF-8"
		key, err := p.cacheKey(w, r, ct)
		if err == nil {
			return
		}
		val, err := h(w, r, ps)
		if err == nil {
			if c {
				var buf bytes.Buffer
				enc := json.NewEncoder(&buf)
				err = enc.Encode(val)
				if err == nil {
					body := buf.Bytes()
					p.C.SetBytes(key, body, 24*time.Hour)
					p.write(w, body, ct)
					return
				}
			} else {
				p.R.JSON(w, http.StatusOK, val)
			}
		}
		p.R.Text(w, http.StatusInternalServerError, err.Error())

	}
}

// XML wrap handle
func (p *Wrap) XML(h Handle, c bool) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		const ct = "text/xml; charset=UTF-8"
		key, err := p.cacheKey(w, r, ct)
		if err == nil {
			return
		}
		val, err := h(w, r, ps)
		if err == nil {
			if c {
				var buf bytes.Buffer
				enc := xml.NewEncoder(&buf)
				err = enc.Encode(val)
				if err == nil {
					body := buf.Bytes()
					p.C.SetBytes(key, body, 24*time.Hour)
					p.write(w, body, ct)
					return
				}
			} else {
				p.R.XML(w, http.StatusOK, val)
				return
			}
		}
		p.R.Text(w, http.StatusInternalServerError, err.Error())
	}
}

func (p *Wrap) cacheKey(w http.ResponseWriter, r *http.Request, t string) (string, error) {
	key := fmt.Sprintf("pages%s", r.URL.RequestURI())
	buf, err := p.C.GetBytes(key)
	if err == nil {
		p.write(w, buf, t)
	}
	return key, err
}

func (p *Wrap) write(w http.ResponseWriter, b []byte, t string) {
	w.Header().Set("Content-Type", t)
	w.Write(b)
}
