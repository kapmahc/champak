package mux_test

import (
	"net/http"
	"testing"

	"github.com/kapmahc/champak/web/mux"
	"github.com/kapmahc/champak/web/mux/gorilla"
	"github.com/kapmahc/champak/web/mux/logging"
)

func h(w http.ResponseWriter, r *http.Request) {

}

func TestGorills(t *testing.T) {
	testMux(t, gorilla.New())
}

func TestLogging(t *testing.T) {
	r := logging.New()
	testMux(t, r)
	r.Walk(func(r logging.Route) {
		t.Logf("%-5s %-16s %-16s %s", r.Method, r.Name, r.Path, r.Func)
	})
}

func testMux(t *testing.T, r mux.Router) {
	mux.Use(r)
	mux.Crud("cms.articles", "/cms/articles", h, h, h, h, h, h, h)
	mux.Form("users.sign-in", "/users/sign-in", h, h)

	name := "cms.article.edit"
	t.Logf("%s => %s", name, mux.URL(name, "id", 111))
}
