package web

import "time"

const (
	// MARKDOWN markdown format
	MARKDOWN = "markdown"
	// HTML html format
	HTML = "html"
)

//Model base model
type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"updatedAt"`
	UpdatedAt time.Time `json:"createdAt"`
}