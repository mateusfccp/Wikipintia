package models

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Name    string
	Entries []Entry
}

func MigrateTag(db *gorm.DB) {
	if !db.HasTable(&Tag{}) {
		db.CreateTable(&Tag{})
	}
}
