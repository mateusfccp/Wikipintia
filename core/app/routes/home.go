package routes

import (
	"html/template"
	"log"
	"net/http"

	"wikipintia/database/models"

	"github.com/jinzhu/gorm"
)

func Home(db *gorm.DB, templ *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var entries []models.Entry
		err := db.Find(&entries)
		if err != nil {
			log.Println(err)
		}
		templ.ExecuteTemplate(w, "master", entries)
	})
}
