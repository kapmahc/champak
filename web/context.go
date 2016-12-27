package web

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/go-playground/form"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	validator "gopkg.in/go-playground/validator.v9"
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

// ClientIP get client ip
func (p *Wrap) ClientIP(r *http.Request) string {
	ip := r.Header.Get("X-FORWARDED-FOR")
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}

//Redirect redirect
func (p *Wrap) Redirect(to string, h Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var act, msg string
		if val, err := h(w, r, ps); err == nil {
			act = "notice"
			msg = val.(string)
		} else {
			act = "alert"
			msg = err.Error()
		}
		http.Redirect(w, r, fmt.Sprintf("%s?%s=%s", to, act, msg), http.StatusFound)
	}
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
		r.ParseMultipartForm(32 << 10)
		log.Debugf("bind form %+v", r.Form)
		if err == nil {
			err = dec.Decode(f, r.Form)
		}
		if err == nil {
			err = p.V.Struct(f)
		}
		var val interface{}
		if err == nil {
			val, err = h(w, r, ps, f)
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
		val, err := h(w, r, ps)
		if err == nil {
			p.R.JSON(w, http.StatusOK, val)
			return
		}
		p.R.Text(w, http.StatusInternalServerError, err.Error())
	}
}

// XML wrap handle
func (p *Wrap) XML(h Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		val, err := h(w, r, ps)
		if err == nil {
			p.R.XML(w, http.StatusOK, val)
			return
		}
		p.R.Text(w, http.StatusInternalServerError, err.Error())
	}
}

// JSONC wrap handle
func (p *Wrap) JSONC(h Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		const ct = "application/json; charset=UTF-8"
		key, err := p.cacheKey(w, r, ct)
		if err == nil {
			return
		}
		val, err := h(w, r, ps)
		if err == nil {
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			err = enc.Encode(val)
			if err == nil {
				body := buf.Bytes()
				p.C.SetBytes(key, body, 24*time.Hour)
				p.write(w, body, ct)
				return
			}
		}

		p.R.Text(w, http.StatusInternalServerError, err.Error())
	}
}

// XMLC wrap cache handle
func (p *Wrap) XMLC(h Handle, c bool) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		const ct = "text/xml; charset=UTF-8"
		key, err := p.cacheKey(w, r, ct)
		if err == nil {
			return
		}
		val, err := h(w, r, ps)
		if err == nil {
			var buf bytes.Buffer
			enc := xml.NewEncoder(&buf)
			err = enc.Encode(val)
			if err == nil {
				body := buf.Bytes()
				p.C.SetBytes(key, body, 24*time.Hour)
				p.write(w, body, ct)
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
