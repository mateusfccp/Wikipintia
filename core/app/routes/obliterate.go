package routes

import (
	"log"
	"net/http"

	"wikipintia/database/models"

	"github.com/go-zoo/bone"
	"github.com/jinzhu/gorm"
)

func Obliterate(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var entry models.Entry

		err := db.Model(&entry).Where(&models.Entry{Slug: bone.GetValue(r, "slug")}).Delete(&entry)
		if err != nil {
			log.Println(err)
		}

		log.Println("Redirecting...")
		http.Redirect(w, r, "/", http.StatusFound)
	})
}
