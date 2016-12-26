package site

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kapmahc/champak/web"
)

type fmNotice struct {
	Body string
}

func (p *Engine) createNotice(w http.ResponseWriter, r *http.Request, ps httprouter.Params, fm interface{}) (interface{}, error) {
	// TODO
	return web.H{}, nil
}

func (p *Engine) updateNotice(w http.ResponseWriter, r *http.Request, ps httprouter.Params, fm interface{}) (interface{}, error) {
	// TODO
	return web.H{}, nil
}

func (p *Engine) showNotice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	// TODO
	return web.H{}, nil
}

func (p *Engine) destroyNotice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	// TODO
	return web.H{}, nil
}

func (p *Engine) indexNotices(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (interface{}, error) {
	// TODO
	return web.H{}, nil
}
