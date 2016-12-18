package auth

import "github.com/jinzhu/gorm"

// Dao auth dao
type Dao struct {
	Db *gorm.DB `inject:""`
}
