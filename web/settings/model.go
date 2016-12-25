package settings

import "time"

// Setting setting
type Setting struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Key       string    `gorm:"not null;unique_index;type:VARCHAR(255)"`
	Val       []byte    `gorm:"not null"`
	Encode    bool      `gorm:"not null"`
	UpdatedAt time.Time `json:"created_at"`
	CreatedAt time.Time `json:"updated_at"`
}

// TableName table name
func (Setting) TableName() string {
	return "settings"
}
