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

// IsAvailable is available?
func (p *User) IsAvailable() bool {
	return p.IsConfirm() && !p.IsLock()
}

// IsConfirm is confirm?
func (p *User) IsConfirm() bool {
	return p.ConfirmedAt != nil
}

// IsLock is lock?
func (p *User) IsLock() bool {
	return p.LockedAt != nil
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
