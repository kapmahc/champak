package auth

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"

	log "github.com/Sirupsen/logrus"
	sessions "github.com/goincremental/negroni-sessions"
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
const (
	// UID user's uid
	UID = "uid"
)

// CurrentUserMiddleware current user
type CurrentUserMiddleware struct {
	Dao *Dao `inject:""`
	Jwt *Jwt `inject:""`
}

func (p *CurrentUserMiddleware) ServeHTTP(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	uid, err := p.getByJwt(req)
	if err != nil {
		log.Debug(err)
		uid = p.getBySession(req)
	}

	if uid != nil {
		if user, err := p.getUser(uid.(string)); err == nil {
			req = req.WithContext(context.WithValue(req.Context(), CurrentUser, user))
		} else {
			log.Error(err)
		}
	}

	next(wrt, req)
}

func (p *CurrentUserMiddleware) getUser(uid string) (*User, error) {
	user, err := p.Dao.GetUserByUID(uid)
	if err != nil {
		return nil, err
	}
	if !user.IsConfirm() {
		return nil, fmt.Errorf("user %s is not confirm", user)
	}
	if user.IsLock() {
		return nil, fmt.Errorf("user %s is lock", user)
	}
	return user, nil
}

func (p *CurrentUserMiddleware) getBySession(req *http.Request) interface{} {
	ss := sessions.GetSession(req)
	return ss.Get(UID)
}

func (p *CurrentUserMiddleware) getByJwt(req *http.Request) (interface{}, error) {
	cm, err := p.Jwt.Parse(req)
	if err != nil {
		return nil, err
	}
	return cm.Get(UID), nil
}
