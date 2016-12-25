package web

import (
	"net"
	"net/http"
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
