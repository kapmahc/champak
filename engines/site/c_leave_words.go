package site

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kapmahc/champak/web"
)

func (p *Engine) createLeaveWord(w http.ResponseWriter, r *http.Request, _ httprouter.Params, o interface{}) (interface{}, error) {
	fm := o.(*fmBody)
	if err := p.Db.Create(&LeaveWord{
		Body: fm.Body,
		Type: web.MARKDOWN,
	}).Error; err != nil {
		return nil, err
	}
	return web.H{}, nil
}

func (p *Engine) destroyLeaveWord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	if err := p.Db.Where("id = ?", ps.ByName("id")).Delete(&LeaveWord{}).Error; err != nil {
		return nil, err
	}
	return web.H{}, nil
}

func (p *Engine) indexLeaveWords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (interface{}, error) {
	var items []LeaveWord
	if err := p.Db.Order("created_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
