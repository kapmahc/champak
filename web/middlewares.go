package web

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/csrf"
	"golang.org/x/text/language"
)

const (
	// DATA data key
	DATA = KEY("data")
	// LOCALE locale key
	LOCALE = KEY("locale")
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
	const key = string(LOCALE)
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

	ctx := context.WithValue(req.Context(), LOCALE, tag.String())
	ctx = context.WithValue(ctx, DATA, H{
		"locale":    tag.String(),
		"languages": p.languages,
	})
	next(wrt, req.WithContext(ctx))
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
