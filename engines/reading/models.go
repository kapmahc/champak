package reading

import (
	"github.com/kapmahc/champak/engines/auth"
	"github.com/kapmahc/champak/web"
)

// Book book
type Book struct {
	web.Model

	Author      string
	Publisher   string
	Title       string
	Type        string
	Lang        string
	File        string
	Subject     string
	Description string
	PublishedAt string
	Cover       string
}

// TableName table name
func (Book) TableName() string {
	return "reading_books"
}

// Note note
type Note struct {
	web.Model

	Body string
	Type string

	UserID uint
	User   auth.User
	BookID uint
	Book   Book
}

// TableName table name
func (Note) TableName() string {
	return "reading_notes"
}
