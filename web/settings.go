package web

import (
	"bytes"
	"encoding/gob"

	"github.com/jinzhu/gorm"
)

// Setting setting
type Setting struct {
	Model

	Key    string `gorm:"not null;unique_index;type:VARCHAR(255)"`
	Val    []byte `gorm:"not null"`
	Encode bool   `gorm:"not null"`
}

// TableName table name
func (Setting) TableName() string {
	return "settings"
}

// Settings setting helper
type Settings struct {
	Security *Security `inject:""`
	Db       *gorm.DB  `inject:""`
}

//Set save setting
func (p *Settings) Set(k string, v interface{}, f bool) error {
	var m Setting
	null := p.Db.Where("key = ?", k).First(&m).RecordNotFound()
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
		m.Val, err = p.Security.Encrypt(buf.Bytes())
		if err != nil {
			return err
		}
	} else {
		m.Val = buf.Bytes()
	}
	m.Encode = f

	if null {
		err = p.Db.Create(&m).Error
	} else {
		err = p.Db.Model(&m).Updates(map[string]interface{}{
			"encode": f,
			"val":    buf,
		}).Error
	}
	return err
}

//Get get setting value by key
func (p *Settings) Get(k string, v interface{}) error {
	var m Setting
	err := p.Db.Where("key = ?", k).First(&m).Error
	if err != nil {
		return err
	}
	if m.Encode {
		if m.Val, err = p.Security.Decrypt(m.Val); err != nil {
			return err
		}
	}

	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	buf.Write(m.Val)
	return dec.Decode(v)
}
