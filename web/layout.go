package web

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

// Layout layout
type Layout struct {
	Db *gorm.DB `inject:""`
}

// Links links
func (p *Layout) Links(loc string) []Link {
	var items []Link
	if err := p.Db.
		Select([]string{"label", "href"}).
		Where("loc = ?", loc).
		Order("sort_order ASC").
		Find(&items).Error; err != nil {
		log.Error(err)
	}
	return items
}

// Cards cards
func (p *Layout) Cards(loc string) []Card {
	var items []Card
	if err := p.Db.
		Select([]string{"title", "summary", "logo", "href"}).
		Where("loc = ?", loc).
		Order("sort_order ASC").
		Find(&items).Error; err != nil {
		log.Error(err)
	}
	return items
}

// Link link
type Link struct {
	Model
	Loc       string
	Label     string
	Href      string
	SortOrder int
}

// Card card
type Card struct {
	Model
	Loc       string
	Title     string
	Sumamry   string
	Logo      string
	Href      string
	SortOrder int
}
