package gorm

import (
	_gorm "github.com/jinzhu/gorm"
)

// New new gorm store
func New(db *_gorm.DB, auto bool) *Store {
	if auto {
		db.AutoMigrate(&Locale{})
		db.Model(&Locale{}).AddUniqueIndex("idx_locales_lang_code", "lang", "code")
	}
	return &Store{
		db: db,
	}
}

// Store gorm store
type Store struct {
	db *_gorm.DB
}

// Set set record
func (p *Store) Set(lang, code, message string, override bool) error {
	var l Locale
	var err error
	if p.db.Where("lang = ? AND code = ?", lang, code).First(&l).RecordNotFound() {
		l.Lang = lang
		l.Code = code
		l.Message = message
		err = p.db.Create(&l).Error
	} else {
		if override {
			l.Message = message
			err = p.db.Save(&l).Error
		}
	}
	return err
}

// All get all items
func (p *Store) All(lang string) (map[string]string, error) {
	var items []Locale
	err := p.db.
		Select([]string{"code", "message"}).
		Where("lang = ?", lang).Find(&items).Error
	if err != nil {
		return nil, err
	}
	rst := make(map[string]string)
	for _, i := range items {
		rst[i.Code] = i.Message
	}
	return rst, nil
}

// Get get message
func (p *Store) Get(lang, code string) (string, error) {

	var l Locale
	if err := p.db.
		Select("message").
		Where("lang = ? AND code = ?", lang, code).
		First(&l).Error; err != nil {
		return "", err
	}
	return l.Message, nil
}

// Del delete
func (p *Store) Del(lang, code string) error {
	return p.db.Where("lang = ? AND code = ?", lang, code).Delete(Locale{}).Error
}
