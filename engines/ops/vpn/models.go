package vpn

import (
	"time"

	"github.com/kapmahc/champak/web"
)

// User user
type User struct {
	web.Model

	FullName string
	Email    string
	Password string
	Details  string
	Online   bool
	Enable   bool
	StartUp  time.Time
	ShutDown time.Time
}

// TableName table name
func (User) TableName() string {
	return "vpn_users"
}

// Log log
type Log struct {
	ID          uint
	TrustedIP   string
	TrustedPort uint
	RemoteIP    string
	RemotePort  uint
	StartUp     *time.Time
	ShutDown    *time.Time
	Received    float64
	Send        float64

	UserID uint
	User   User
}

// TableName table name
func (Log) TableName() string {
	return "vpn_users"
}
