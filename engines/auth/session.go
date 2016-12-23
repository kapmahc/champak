package auth

import "github.com/kapmahc/champak/web"

// Session session
type Session struct {
	Dao  *Dao      `inject:""`
	I18n *web.I18n `inject:""`
}
