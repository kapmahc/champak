package mux_test

import (
	"net/http"
	"testing"

	_mux "github.com/gorilla/mux"
	"github.com/kapmahc/champak/web/mux"
)

func h(w http.ResponseWriter, r *http.Request) {

}

func TestGorilla(t *testing.T) {
	mux.Use(_mux.NewRouter())
	mux.Crud("cms.articles", "/cms/articles", h, h, h, h, h, h, h)
	mux.Form("users.sign-in", "/users/sign-in", h, h)
	mux.Walk(func(m, n, p string) error {
		t.Logf("%-5s %-16s %s", m, n, p)
		return nil
	})
	name := "cms.article.edit"
	t.Logf("%s => %s", name, mux.URL(name, "id", 111))
}
