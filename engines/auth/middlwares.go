package auth

import (
	"context"
	"io"
	"net/http"
	"regexp"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/csrf"
	"github.com/kapmahc/champak/web"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"
	"golang.org/x/text/language"
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

// NewLocaleMiddleware create a locale middleware
func NewLocaleMiddleware(languages ...string) (*LocaleMiddleware, error) {
	var tags []language.Tag
	for _, l := range languages {
		if lng, err := language.Parse(l); err == nil {
			tags = append(tags, lng)
		} else {
			return nil, err
		}
	}
	return &LocaleMiddleware{
		matcher:   language.NewMatcher(tags),
		languages: languages,
	}, nil
}

// LocaleMiddleware detect locale from http header
type LocaleMiddleware struct {
	matcher   language.Matcher
	languages []string
}

func (p *LocaleMiddleware) ServeHTTP(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	const key = string(web.LOCALE)
	write := false

	// 1. Check URL arguments.
	lng := req.URL.Query().Get(key)

	// 2. Get language information from cookies.
	if len(lng) == 0 {
		if ck, er := req.Cookie(key); er == nil {
			lng = ck.Value
		} else {
			write = true
		}
	} else {
		write = true
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lng) == 0 {
		write = true
		al := req.Header.Get("Accept-Language")
		if len(al) > 4 {
			lng = al[:5]
		}
	}

	tag, _, _ := p.matcher.Match(language.Make(lng))

	// Write cookie
	if write {
		http.SetCookie(wrt, &http.Cookie{
			Name:    key,
			Value:   tag.String(),
			Expires: time.Now().AddDate(10, 0, 0),
			Path:    "/",
		})
	}
	ctx := context.WithValue(req.Context(), web.LOCALE, tag.String())
	ctx = context.WithValue(ctx, web.DATA, web.H{
		"locale":    tag.String(),
		"languages": p.languages,
	})
	next(wrt, req.WithContext(ctx))
}

// CsrfMiddleware csrf
type CsrfMiddleware struct {
}

func (p *CsrfMiddleware) ServeHTTP(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	data := req.Context().Value(web.DATA).(web.H)
	tkn := csrf.Token(req)
	wrt.Header().Set("X-CSRF-Token", tkn)
	data["csrf"] = tkn
	next(wrt, req.WithContext(context.WithValue(req.Context(), web.DATA, data)))
}
