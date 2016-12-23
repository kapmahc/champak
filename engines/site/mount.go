package site

import "github.com/gorilla/mux"

// Mount mount
func (p *Engine) Mount(rt *mux.Router) {
	// rt.GET("/", p.getHome)
	// rt.GET("/rss.atom", p.getRss)
	//
	// rt.GET("/leave-words/new", p.newLeaveWord)
	// rt.POST(
	// 	"/leave-words",
	// 	web.PostFormHandler("/leave-words/new", &fmLeaveWord{}, p.createLeaveWord),
	// )
	//
	// rt.GET("/dashboard", p.Session.MustSignInHandler(), p.getDashboard)
}
