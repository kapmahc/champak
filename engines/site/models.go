package site

import (
	"time"

	"github.com/kapmahc/champak/web"
)

// LeaveWord leave word
type LeaveWord struct {
	ID        uint      `json:"id"`
	Body      string    `json:"body"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName table name
func (LeaveWord) TableName() string {
	return "leave_words"
}

// Notice notice
type Notice struct {
	web.Model

	Body string `json:"body"`
	Type string `json:"type"`
}

// TableName table name
func (Notice) TableName() string {
	return "notices"
}

// Link link
type Link struct {
	web.Model

	Loc       string `json:"loc"`
	Label     string `json:"label"`
	Href      string `json:"href"`
	SortOrder int    `json:"sort_order"`
}

// TableName table name
func (Link) TableName() string {
	return "links"
}

// Card card
type Card struct {
	web.Model
	Loc       string `json:"loc"`
	Title     string `json:"title"`
	Sumamry   string `json:"summary"`
	Logo      string `json:"logo"`
	Href      string `json:"href"`
	SortOrder int    `json:"sort_order"`
}

// TableName table name
func (Card) TableName() string {
	return "cards"
}
