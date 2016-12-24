package mux

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
)

// URL get url by name
func URL(name string, args ...interface{}) string {
	var pairs []string
	for _, v := range args {
		switch v.(type) {
		case string:
			pairs = append(pairs, v.(string))
		default:
			pairs = append(pairs, fmt.Sprintf("%v", v))
		}
	}
	if r := root.Get(name); r != nil {
		u, e := r.URL(pairs...)
		if e == nil {
			return u.String()
		}
		log.Error(e)
	}
	return "not-found"
}

// Walk walk routes
func Walk(fn func(Route)) {
	for _, r := range routes {
		fn(r)
	}
}
