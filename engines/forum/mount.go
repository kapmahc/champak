package forum

import (
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	gt := rt.Group("/forum")
	// -----------
	gt.GET("/tags", p.indexTags)
	gt.GET("/tags/new",
		p.Session.MustSignInHandler(), p.Session.MustAdminHandler(),
		p.newTag,
	)
	gt.POST(
		"/tags",
		p.Session.MustSignInHandler(), p.Session.MustAdminHandler(),
		web.PostFormHandler(&fmTag{}, p.createTag),
	)
	gt.POST(
		"/tags/:id",
		p.Session.MustSignInHandler(), p.Session.MustAdminHandler(),
		web.PostFormHandler(&fmTag{}, p.updateTag),
	)
	gt.DELETE(
		"/tags/:id",
		p.Session.MustSignInHandler(), p.Session.MustAdminHandler(),
		web.JSON(p.destoryTag),
	)
	gt.GET(
		"/tags/edit/:id",
		p.Session.MustSignInHandler(), p.Session.MustAdminHandler(),
		web.HTML(p.editTag),
	)
	gt.GET(
		"/tags/show/:id",
		web.HTML(p.showTag),
	)
	// ---------------
	gt.GET("/articles/latest", p.latestArticles)
	gt.GET("/articles", p.Session.MustSignInHandler(), p.indexArticles)
	gt.GET("/articles/new",
		p.Session.MustSignInHandler(),
		p.newArticle,
	)
	gt.POST(
		"/articles",
		p.Session.MustSignInHandler(),
		web.PostFormHandler(&fmArticle{}, p.createArticle),
	)
	gt.POST(
		"/articles/:id",
		p.Session.MustSignInHandler(),
		web.PostFormHandler(&fmArticle{}, p.updateArticle),
	)
	gt.DELETE(
		"/articles/:id",
		p.Session.MustSignInHandler(),
		web.JSON(p.destoryArticle),
	)
	gt.GET(
		"/articles/edit/:id",
		p.Session.MustSignInHandler(),
		web.HTML(p.editArticle),
	)
	gt.GET(
		"/articles/show/:id",
		web.HTML(p.showArticle),
	)
	// ---------------
	gt.GET("/comments/latest", p.latestComments)
	gt.GET("/comments", p.Session.MustSignInHandler(), p.indexComments)
	gt.GET("/comments/new",
		p.Session.MustSignInHandler(),
		p.newComment,
	)
	gt.POST(
		"/comments",
		p.Session.MustSignInHandler(),
		web.PostFormHandler(&fmComment{}, p.createComment),
	)
	gt.POST(
		"/comments/:id",
		p.Session.MustSignInHandler(),
		web.PostFormHandler(&fmComment{}, p.updateComment),
	)
	gt.DELETE(
		"/comments/:id",
		p.Session.MustSignInHandler(),
		web.JSON(p.destoryComment),
	)
	gt.GET(
		"/comments/edit/:id",
		p.Session.MustSignInHandler(),
		web.HTML(p.editComment),
	)
}
