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
