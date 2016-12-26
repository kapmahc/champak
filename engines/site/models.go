package site

import (
	"time"

	"github.com/kapmahc/champak/web"
)

// LeaveWord leave word
type LeaveWord struct {
	ID        uint
	Body      string
	Type      string
	CreatedAt time.Time
}

// TableName table name
func (LeaveWord) TableName() string {
	return "leave_words"
}

// Notice notice
type Notice struct {
	web.Model

	Body string
	Type string
}

// TableName table name
func (Notice) TableName() string {
	return "notices"
}
