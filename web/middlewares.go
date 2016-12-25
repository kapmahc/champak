package web

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/csrf"

	"golang.org/x/text/language"
)

const (
	// DATA data-key
	DATA = KEY("data")
)

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
	const key = "locale"
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

	next(wrt, req.WithContext(
		context.WithValue(
			req.Context(),
			DATA,
			H{
				"locale":    tag.String(),
				"languages": p.languages,
			},
		),
	))
}

// CsrfMiddleware csrf
type CsrfMiddleware struct {
}

func (p *CsrfMiddleware) ServeHTTP(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	data := req.Context().Value(DATA).(H)
	tkn := csrf.Token(req)
	wrt.Header().Set("X-CSRF-Token", tkn)
	data["csrf"] = tkn
	next(wrt, req.WithContext(context.WithValue(req.Context(), DATA, data)))
}

// ClientIPMiddleware client ip
type ClientIPMiddleware struct {
}

func (p *ClientIPMiddleware) ServeHTTP(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	ip := req.Header.Get("X-FORWARDED-FOR")
	if ip == "" {
		ip, _, _ = net.SplitHostPort(req.RemoteAddr)
	}

	data := req.Context().Value(DATA).(H)
	data["client-ip"] = ip
	next(wrt, req.WithContext(context.WithValue(req.Context(), DATA, data)))
}
