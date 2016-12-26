package site

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kapmahc/champak/web"
)

type fmLeaveWord struct {
	Body string
}

func (p *Engine) createLeaveWord(w http.ResponseWriter, r *http.Request, ps httprouter.Params, fm interface{}) (interface{}, error) {
	// TODO
	return web.H{}, nil
}

func (p *Engine) showLeaveWord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	// TODO
	return web.H{}, nil
}

func (p *Engine) destroyLeaveWord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	// TODO
	return web.H{}, nil
}

func (p *Engine) indexLeaveWords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (interface{}, error) {
	// TODO
	return web.H{}, nil
}
