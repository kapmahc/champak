package gorm

import "time"

//Locale locale model
type Locale struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Lang      string    `gorm:"not null;type:varchar(8);index"`
	Code      string    `gorm:"not null;index;type:VARCHAR(255)"`
	Message   string    `gorm:"not null;type:varchar(800)"`
	UpdatedAt time.Time `json:"created_at"`
	CreatedAt time.Time `json:"updated_at"`
}

// TableName table name
func (Locale) TableName() string {
	return "locales"
}
