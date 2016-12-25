package auth

import (
	"io"
	"net/http"
	"regexp"

	log "github.com/Sirupsen/logrus"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"
)

// NewMinifyMiddleware new minify middleware
func NewMinifyMiddleware() *MinifyMiddleware {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/javascript", js.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
	return &MinifyMiddleware{m: m}
}

// MinifyResponseWriter minify response writer
type MinifyResponseWriter struct {
	http.ResponseWriter
	io.WriteCloser
}

func (m MinifyResponseWriter) Write(b []byte) (int, error) {
	return m.WriteCloser.Write(b)
}

// MinifyMiddleware minify html output
type MinifyMiddleware struct {
	m *minify.M
}

func (p *MinifyMiddleware) ServeHTTP(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	w := MinifyResponseWriter{wrt, p.m.Writer("text/html", wrt)}
	next(w, req)

	if err := w.Close(); err != nil {
		log.Error(err)
	}
}

// -----------------------------------------------------------------------------
