package mux_test

import (
	"net/http"
	"testing"

	gorilla "github.com/gorilla/mux"
	"github.com/kapmahc/champak/web/mux"
)

func h(w http.ResponseWriter, r *http.Request) {

}

func TestGorilla(t *testing.T) {
	r := mux.Router{Router: gorilla.NewRouter()}
	r.Crud("cms.articles", "/cms/articles", h, h, h, h, h, h, h)
	r.Form("users.sign-in", "/users/sign-in", h, h)
	r.Walk(func(m, n, p string) error {
		t.Logf("%-5s %-16s %s", m, n, p)
		return nil
	})
	name := "cms.article.edit"
	t.Logf("%s => %s", name, r.URL(name, "id", 111))
}
