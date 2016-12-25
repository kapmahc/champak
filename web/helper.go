package web

import (
	"net"
	"net/http"

	"github.com/go-playground/form"
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

// ClientIP get client ip
func ClientIP(r *http.Request) string {
	ip := r.Header.Get("X-FORWARDED-FOR")
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}

// Bind bind request to form
func Bind(r *http.Request, v interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	dec := form.NewDecoder()
	return dec.Decode(v, r.Form)
}
