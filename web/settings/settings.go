package settings

import (
	"bytes"
	"encoding/gob"

	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/web/crypto"
)

var (
	db *gorm.DB
)

// Use use db
func Use(d *gorm.DB, auto bool) {
	db = d
	if auto {
		db.AutoMigrate(&Setting{})
	}
}

//Set save setting
func Set(k string, v interface{}, f bool) error {
	var m Setting
	null := db.Where("key = ?", k).First(&m).RecordNotFound()
	if null {
		m = Setting{Key: k}
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return err
	}
	if f {
		m.Val, err = crypto.Encrypt(buf.Bytes())
		if err != nil {
			return err
		}
	} else {
		m.Val = buf.Bytes()
	}
	m.Encode = f

	if null {
		err = db.Create(&m).Error
	} else {
		err = db.Model(&m).Updates(map[string]interface{}{
			"encode": f,
			"val":    buf,
		}).Error
	}
	return err
}

//Get get setting value by key
func Get(k string, v interface{}) error {
	var m Setting
	err := db.Where("key = ?", k).First(&m).Error
	if err != nil {
		return err
	}
	if m.Encode {
		if m.Val, err = crypto.Decrypt(m.Val); err != nil {
			return err
		}
	}

	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	buf.Write(m.Val)
	return dec.Decode(v)
}
