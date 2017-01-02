package site

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kapmahc/champak/web"
)

func (p *Engine) createNotice(w http.ResponseWriter, r *http.Request, _ httprouter.Params, o interface{}) (interface{}, error) {
	fm := o.(*fmBody)
	if err := p.Db.Create(&Notice{Body: fm.Body, Type: web.MARKDOWN}).Error; err != nil {
		return nil, err
	}
	return fm, nil
}

func (p *Engine) updateNotice(w http.ResponseWriter, r *http.Request, ps httprouter.Params, o interface{}) (interface{}, error) {
	fm := o.(*fmBody)
	if err := p.Db.Model(&Notice{}).Where("id = ?", ps.ByName("id")).Update("body", fm.Body).Error; err != nil {
		return nil, err
	}
	return web.H{}, nil
}

func (p *Engine) destroyNotice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	if err := p.Db.Where("id = ?", ps.ByName("id")).Delete(&Notice{}).Error; err != nil {
		return nil, err
	}
	return web.H{}, nil
}

func (p *Engine) indexNotices(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (interface{}, error) {
	var items []Notice
	if err := p.Db.Model(&Notice{}).Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
