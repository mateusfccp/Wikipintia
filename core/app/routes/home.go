package routes

import (
	"html/template"
	"net/http"

	"wikipintia/database/models"

	"github.com/jinzhu/gorm"
)

func Home(db *gorm.DB, templ *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var entries []models.Entry

		dbc := db.Find(&entries)
		if dbc.Error != nil {
			http.Error(w, dbc.Error.Error(), http.StatusExpectationFailed)
		}

		templ.ExecuteTemplate(w, "master", entries)
	})
}
