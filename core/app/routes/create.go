package routes

import (
	"encoding/json"
	"html/template"
	"net/http"
	"net/url"
	"wikipintia/database/models"

	"github.com/jinzhu/gorm"
)

func PrettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	println(string(b))
}

func Create(templ *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.ExecuteTemplate(w, "master", nil)
	})
}

func New(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")

		entry := models.Entry{
			Title:   title,
			Slug:    url.PathEscape(title),
			Content: r.FormValue("content"),
		}

		dbc := db.Create(&entry)
		if dbc.Error != nil {
			http.Error(w, dbc.Error.Error(), http.StatusExpectationFailed)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
	})
}
