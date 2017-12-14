package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Entries  []Entry
	Parent   *Category
	ParentID uint
}

func MigrateCategory(db *gorm.DB) {
	if !db.HasTable(&Category{}) {
		db.CreateTable(&Category{})
		db.Model(&Category{}).AddForeignKey("parent_id", "categories(id)", "RESTRICT", "RESTRICT")
	}
}
