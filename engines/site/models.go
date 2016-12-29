package site

import (
	"time"

	"github.com/kapmahc/champak/web"
)

// SMTP smtp config
type SMTP struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"-"`
	Ssl      bool   `json:"ssl"`
}

// LeaveWord leave word
type LeaveWord struct {
	ID        uint      `json:"id"`
	Body      string    `json:"body"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
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
	SortOrder int    `json:"sortOrder"`
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
	SortOrder int    `json:"sortOrder"`
}

// TableName table name
func (Card) TableName() string {
	return "cards"
}
