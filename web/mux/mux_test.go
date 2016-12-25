package mux_test

import (
	"net/http"
	"testing"

	_mux "github.com/gorilla/mux"
	"github.com/kapmahc/champak/web/mux"
	"github.com/kapmahc/champak/web/mux/gorilla"
	"github.com/kapmahc/champak/web/mux/logging"
)

func h(w http.ResponseWriter, r *http.Request) {

}

func TestGorills(t *testing.T) {
	r := gorilla.New(_mux.NewRouter())
	testMux(t, r)
	name := "cms.article.edit"
	t.Logf("%s => %s", name, r.URL(name, "id", 111))
}

func TestLogging(t *testing.T) {
	r := logging.New()
	testMux(t, r)
	r.Walk(func(r logging.Route) {
		t.Logf("%-5s %-16s %-16s %s", r.Method, r.Name, r.Path, r.Func)
	})
}

func testMux(t *testing.T, r mux.Router) {
	m := mux.New(r)
	m.Crud("cms.articles", "/cms/articles", h, h, h, h, h, h, h)
	m.Form("users.sign-in", "/users/sign-in", h, h)
}
