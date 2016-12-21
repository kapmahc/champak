package auth

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/web"
)

const ()

// Dao auth dao
type Dao struct {
	Db       *gorm.DB      `inject:""`
	Security *web.Security `inject:""`
}

// signIn set sign-in info
func (p *Dao) signIn(user uint, ip string) {
	var u User
	if err := p.Db.Where("id = ?", user).First(&u).Error; err != nil {
		log.Error(err)
		return
	}
	if err := p.Db.Model(&u).Updates(map[string]interface{}{
		"sign_in_count":      u.SignInCount + 1,
		"last_sign_in_at":    u.CurrentSignInAt,
		"last_sign_in_ip":    u.CurrentSignInIP,
		"current_sign_in_ip": ip,
		"current_sign_in_at": time.Now(),
	}).Error; err != nil {
		log.Error(err)
	}
}

// GetUserByUID get user by uid
func (p *Dao) GetUserByUID(uid string) (*User, error) {
	var u User
	err := p.Db.Where("uid = ?", uid).First(&u).Error
	return &u, err
}

// GetByEmail get user by email
func (p *Dao) GetByEmail(email string) (*User, error) {
	var user User
	err := p.Db.
		Where("resource_type = ? AND resource_id = ?", UserTypeEmail, email).
		First(&user).Error
	return &user, err
}

// Log add log
func (p *Dao) Log(user uint, message string) {
	if err := p.Db.Create(&Log{UserID: user, Message: message}).Error; err != nil {
		log.Error(err)
	}
}

// AddEmailUser add email user
func (p *Dao) AddEmailUser(fullName, email, password string) (*User, error) {

	user := User{
		Email:           email,
		Password:        p.Security.Sum([]byte(password)),
		FullName:        fullName,
		ProviderType:    UserTypeEmail,
		ProviderID:      email,
		Home:            "/users",
		LastSignInIP:    "0.0.0.0",
		CurrentSignInIP: "0.0.0.0",
	}
	user.SetUID()
	user.SetGravatarLogo()
	user.Home = fmt.Sprintf("/users/%s", user.UID)

	err := p.Db.Create(&user).Error
	return &user, err
}

// Authority get roles
func (p *Dao) Authority(user uint, rty string, rid uint) []string {
	var items []Role
	if err := p.Db.
		Where("resource_type = ? AND resource_id = ?", rty, rid).
		Find(&items).Error; err != nil {
		log.Error(err)
	}
	var roles []string
	for _, r := range items {
		var pm Policy
		if err := p.Db.
			Where("role_id = ? AND user_id = ? ", r.ID, user).
			First(&pm).Error; err != nil {
			log.Error(err)
			continue
		}
		if pm.Enable() {
			roles = append(roles, r.Name)
		}
	}
	return roles
}

//Is is role ?
func (p *Dao) Is(user uint, name string) bool {
	return p.Can(user, name, "-", 0)
}

//Can can?
func (p *Dao) Can(user uint, name string, rty string, rid uint) bool {
	var r Role
	if p.Db.
		Where("name = ? AND resource_type = ? AND resource_id = ?", name, rty, rid).
		First(&r).
		RecordNotFound() {
		return false
	}
	var pm Policy
	if p.Db.
		Where("user_id = ? AND role_id = ?", user, r.ID).
		First(&pm).
		RecordNotFound() {
		return false
	}

	return pm.Enable()
}

//Role check role exist
func (p *Dao) Role(name string, rty string, rid uint) (*Role, error) {
	var e error
	r := Role{}
	db := p.Db
	if db.
		Where("name = ? AND resource_type = ? AND resource_id = ?", name, rty, rid).
		First(&r).
		RecordNotFound() {
		r = Role{
			Name:         name,
			ResourceType: rty,
			ResourceID:   rid,
		}
		e = db.Create(&r).Error

	}
	return &r, e
}

//Deny deny permission
func (p *Dao) Deny(role uint, user uint) error {
	return p.Db.
		Where("role_id = ? AND user_id = ?", role, user).
		Delete(Policy{}).Error
}

//Allow allow permission
func (p *Dao) Allow(role uint, user uint, years, months, days int) error {
	begin := time.Now()
	end := begin.AddDate(years, months, days)
	var count int
	p.Db.
		Model(&Policy{}).
		Where("role_id = ? AND user_id = ?", role, user).
		Count(&count)
	if count == 0 {
		return p.Db.Create(&Policy{
			UserID:   user,
			RoleID:   role,
			StartUp:  begin,
			ShutDown: end,
		}).Error
	}
	return p.Db.
		Model(&Policy{}).
		Where("role_id = ? AND user_id = ?", role, user).
		UpdateColumns(map[string]interface{}{"begin": begin, "end": end}).Error

}
