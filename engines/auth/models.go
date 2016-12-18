package auth

import (
	"time"

	"github.com/kapmahc/champak/web"
)

const (
	// InputTypeMarkdown markdown format
	InputTypeMarkdown = "markdown"
	// InputTypeHTML html format
	InputTypeHTML = "html"
	// UserTypeEmail email user
	UserTypeEmail = "email"
)

// User user
type User struct {
	web.Model

	FullName        string
	Email           string
	UID             string
	Password        []byte
	ProviderID      string
	ProviderType    string
	Home            string
	Logo            string
	SignInCount     uint
	LastSignInAt    *time.Time
	LastSignInIP    string
	CurrentSignInAt *time.Time
	CurrentSignInIP string
	ConfirmedAt     *time.Time
	LockedAt        *time.Time
}

// TableName table name
func (User) TableName() string {
	return "users"
}

// Attachment attachment
type Attachment struct {
	ID           uint
	Title        string
	URL          string
	Length       uint
	MediaType    string
	ResourceType string
	ResourceID   uint
	SortOrder    int
	CreatedAt    time.Time

	UserID uint
	User   User
}

// TableName table name
func (Attachment) TableName() string {
	return "attachments"
}

// Card card
type Card struct {
	web.Model

	Title     string
	Summary   string
	Logo      string
	Href      string
	Loc       string
	SortOrder int
}

// TableName table name
func (Card) TableName() string {
	return "cards"
}

// Contact contact
type Contact struct {
	web.Model

	Key string
	Val string

	UserID uint
	User   User
}

// TableName table name
func (Contact) TableName() string {
	return "contacts"
}

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

// Link link
type Link struct {
	web.Model

	Label     string
	Href      string
	Loc       string
	SortOrder int
}

// TableName table name
func (Link) TableName() string {
	return "links"
}

// Log log
type Log struct {
	ID        uint
	Message   string
	CreatedAt time.Time

	UserID uint
	User   User
}

// TableName table name
func (Log) TableName() string {
	return "logs"
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

// Policy policy
type Policy struct {
	web.Model

	StartUp  time.Time
	ShutDown time.Time

	UserID uint
	User   User
	RoleID uint
	Role   Role
}

// TableName table name
func (Policy) TableName() string {
	return "policies"
}

// Role role
type Role struct {
	web.Model

	Name         string
	ResourceID   uint
	ResourceType string
}

// TableName table name
func (Role) TableName() string {
	return "roles"
}

// Setting setting
type Setting struct {
	web.Model

	Key  string
	Val  []byte
	Flag bool

	UserID uint
	User   *User
}

// TableName table name
func (Setting) TableName() string {
	return "settings"
}
