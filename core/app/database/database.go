package database

import (
	"wikipintia/database/models"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	models.MigrateCategory(db)
	models.MigrateTag(db)
	models.MigrateEntry(db)

	db.Table("entry_tags").AddForeignKey("entry_id", "entries(id)", "RESTRICT", "RESTRICT")
	db.Table("entry_tags").AddForeignKey("tag_id", "tags(id)", "RESTRICT", "RESTRICT")
}
