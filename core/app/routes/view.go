package routes

import (
	"html/template"
	"log"
	"net/http"

	"wikipintia/database/models"

	"github.com/chaseadamsio/goorgeous"

	"github.com/go-zoo/bone"
	"github.com/jinzhu/gorm"
)

func View(db *gorm.DB, templ *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var entry models.Entry
		err := db.Where(&models.Entry{Title: bone.GetValue(r, "id")}).First(&entry)
		if err != nil {
			log.Println(err)
		}
		entry.Content = string(goorgeous.OrgCommon([]byte(entry.Content)))
		templ.ExecuteTemplate(w, "master", entry)
	})
}
