package ops

import (
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	og := rt.Group("/ops", p.Session.MustSignInHandler(), p.Session.MustAdminHandler())

	og.GET("/site/info", p.getSiteInfo)
	og.POST(
		"/site/info",
		web.PostFormHandler("/ops/site/info", &fmSiteInfo{}, p.postSiteInfo),
	)
	og.GET("/site/author", p.getSiteAuthor)
	og.POST(
		"/site/author",
		web.PostFormHandler("/ops/site/author", &fmSiteAuthor{}, p.postSiteAuthor),
	)
	og.GET("/site/seo", p.getSiteSeo)
	og.POST(
		"/site/seo",
		web.PostFormHandler("/ops/site/seo", &fmSiteSeo{}, p.postSiteSeo),
	)
	og.GET("/site/status", p.getSiteStatus)

	og.GET("/leave-words", p.indexLeaveWords)
	og.DELETE("/leave-words/:id", web.FlashHandler("/ops/leave-words", p.destoryLeaveWord))

	og.GET("/notices", p.indexNotices)
	og.GET("/notices/new", p.newNotice)
	og.POST("/notices", web.PostFormHandler("/ops/notices", &fmNotice{}, p.createNotice))
	og.POST("/notices/:id", web.PostFormHandler("/ops/notices", &fmNotice{}, p.updateNotice))
	og.DELETE("/notices/:id", web.FlashHandler("/ops/notices", p.destoryNotice))
	og.GET("/notices/:id/edit", web.FlashHandler("/ops/notices", p.editNotice))

}
