package mux_test

import (
	"net/http"
	"testing"

	"github.com/kapmahc/champak/web/mux"
)

func h(w http.ResponseWriter, r *http.Request) {

}
func TestMux(t *testing.T) {
	mux.Crud("cms.articles", "/cms/articles", h, h, h, h, h, h, h)
	mux.Form("users.sign-in", "/users/sign-in", h, h)
	mux.Walk(func(r mux.Route) {
		t.Logf("%-5s %-16s %-16s %s", r.Method, r.Name, r.Path, r.Func)

	})

	name := "cms.article.edit"
	t.Logf("%s => %s", name, mux.URL(name, "id", 111))
}
