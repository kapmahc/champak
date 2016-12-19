package auth

import "github.com/jinzhu/gorm"

// Dao auth dao
type Dao struct {
	Db *gorm.DB `inject:""`
}

// Authority get roles
func (p *Dao) Authority(user uint, rty string, rit uint) []string {
	//TODO
	return []string{}
}
