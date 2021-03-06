package models

import (
	"github.com/jinzhu/gorm"
)

type Entry struct {
	gorm.Model
	Title      string `gorm:"not null"`
	Slug       string `gorm:"not null"`
	Content    string `gorm:"type:text;not null"`
	Tags       []Tag  `gorm:"many2many:entry_tags"`
	CategoryID uint   `gorm:"default:NULL"`
	Category   Category
}

func MigrateEntry(db *gorm.DB) {
	if !db.HasTable(&Entry{}) {
		db.CreateTable(&Entry{})
		db.Model(&Entry{}).AddUniqueIndex("idx_entry_slug", "slug")
		db.Model(&Entry{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	}
}
