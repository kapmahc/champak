package cms

import (
	"github.com/kapmahc/champak/engines/auth"
	"github.com/kapmahc/champak/web"
)

// Article article
type Article struct {
	web.Model

	Title   string
	Summary string
	Body    string
	Type    string

	UserID uint
	User   auth.User
	Tags   []Tag
}

// TableName table name
func (Article) TableName() string {
	return "cms_articles"
}

// Tag tag
type Tag struct {
	web.Model

	Name string

	Articles []Article
}

// TableName table name
func (Tag) TableName() string {
	return "cms_tags"
}

//Comment comment
type Comment struct {
	web.Model

	Body string
	Type string

	UserID    uint
	User      auth.User
	ArticleID uint
	Article   Article
}

// TableName table name
func (Comment) TableName() string {
	return "cms_comments"
}
