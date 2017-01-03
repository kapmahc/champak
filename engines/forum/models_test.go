package forum_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kapmahc/champak/engines/forum"
)

func OpenDatabase() (*gorm.DB, error) {
	// db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open("postgres", "user=postgres dbname=champak_test sslmode=disable")
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return db, nil
}

func TestManyToMany(t *testing.T) {
	db, err := OpenDatabase()
	if err != nil {
		t.Fatal(err)
	}
	// db.AutoMigrate(&forum.Article{}, &forum.Tag{}, &forum.Comment{})
	var c int
	db.Model(&forum.Tag{}).Count(&c)
	if c == 0 {
		for _, n := range []string{"aaa", "bbb", "ccc"} {
			db.Create(&forum.Tag{Name: n})
		}
	}
	db.Model(&forum.Article{}).Count(&c)
	if c == 0 {
		for i := 1; i <= 10; i++ {
			db.Create(&forum.Article{
				Title:   fmt.Sprintf("ttt %d", i),
				Summary: fmt.Sprintf("sss %d", i),
				Body:    fmt.Sprintf("bbb %d", i),
				UserID:  1,
			})
		}
	}
	var a forum.Article
	db.First(&a)
	log.Printf("article: %+v", a)
	var tags []forum.Tag
	db.Limit(2).Find(&tags)
	db.Model(&a).Association("Tags").Replace(tags)
}
