package auth

import (
	"net"
	"net/http"

	"github.com/unrolled/render"
)

// KEY request context key type
type KEY string

// H hash
type H map[string]interface{}

const (
	// LOCALE locale key
	LOCALE = KEY("locale")
	// DATA data key
	DATA = KEY("data")
)

// New new context
func New(w http.ResponseWriter, r *http.Request, o render.Options) *Context {
	return &Context{wrt: w, req: r, rdr: render.New(o)}
}

// Context http context
type Context struct {
	wrt http.ResponseWriter
	req *http.Request
	rdr *render.Render
}

// ClientIP client ip
func (p *Context) ClientIP() string {
	if proxy := p.req.Header.Get("X-FORWARDED-FOR"); len(proxy) > 0 {
		return proxy
	}
	ip, _, _ := net.SplitHostPort(p.req.RemoteAddr)
	return ip
}

// XML render xml
func (p *Context) XML(v interface{}) {
	p.rdr.XML(p.wrt, http.StatusOK, v)
}

// JSON render json
func (p *Context) JSON(v interface{}) {
	p.rdr.JSON(p.wrt, http.StatusOK, v)
}
