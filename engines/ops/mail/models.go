package mail

import "github.com/kapmahc/champak/web"

// Domain domain
type Domain struct {
	web.Model

	Name string
}

// TableName table name
func (Domain) TableName() string {
	return "mail_domains"
}

// User user
type User struct {
	web.Model

	FullName string
	Email    string
	Password string

	DomainID uint
	Domain   Domain
}

// TableName table name
func (User) TableName() string {
	return "mail_users"
}

// Alias alias
type Alias struct {
	web.Model

	Source      string
	Destination string

	DomainID uint
	Domain   Domain
}

// TableName table name
func (Alias) TableName() string {
	return "mail_aliases"
}
