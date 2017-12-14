package routes

import (
	"html/template"
	"log"
	"net/http"

	"wikipintia/database/models"

	"github.com/go-zoo/bone"
	"github.com/jinzhu/gorm"
)

func Edit(db *gorm.DB, templ *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var entry models.Entry
		err := db.Where(&models.Entry{Title: bone.GetValue(r, "id")}).First(&entry)
		if err != nil {
			log.Println(err)
		}
		templ.ExecuteTemplate(w, "master", entry)
	})
}

func Save(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		title := bone.GetValue(r, "id")
		var entry models.Entry
		err := db.Model(&entry).Where(&models.Entry{Title: title}).Update("content", r.FormValue("content"))
		if err != nil {
			log.Println(err)
		}

		log.Println("Redirecting...")
		http.Redirect(w, r, "/"+title, http.StatusFound)
	})
}
