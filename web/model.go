package web

import "time"

const (
	// TypeMARKDOWN markdown format
	TypeMARKDOWN = "markdown"
	// TypeHTML html format
	TypeHTML = "html"
)

//Model base model
type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"updatedAt"`
	UpdatedAt time.Time `json:"createdAt"`
}

// Dropdown dropdown
type Dropdown struct {
	Label string
	Href  string
	Links []*Link
}

// Link link
type Link struct {
	Label string
	Href  string
}
