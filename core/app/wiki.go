// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"wikipintia/database"
	"wikipintia/routes"

	"github.com/go-zoo/bone"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var validPath = regexp.MustCompile("^/(edit|save|view|obliterate)/([a-zA-Z0-9]+)$")
var templates = cacheTemplates()

func cacheTemplates() map[string]*template.Template {
	funcs := template.FuncMap{
		"html": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	files, _ := ioutil.ReadDir("templ")
	templates := make(map[string]*template.Template)
	for _, file := range files {
		if file.Name() == "master.html" {
			continue
		}

		templ, err := template.New(file.Name()).Funcs(funcs).ParseFiles("templ/master.html", "templ/"+file.Name())

		if err != nil {
			print("Error: " + err.Error())
		}

		templates[file.Name()] = templ
	}

	return templates
}

func main() {
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Wikipintia is running on production!")
	} else {
		log.Println("Wikipintia is running in dev mode!")
	}

	dbString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	db, _ := gorm.Open("mysql", dbString)
	db.LogMode(true)

	// if err != nil {
	// 	http.Error(response, err.Error(), http.StatusInternalServerError)
	// }

	database.Migrate(db)

	defer db.Close()

	mux := bone.New()

	fs := http.FileServer(http.Dir("static"))
	mux.Get("/static/", http.StripPrefix("/static/", fs))

	mux.Get("/", routes.Home(db, templates["home.html"]))
	mux.Get("/create", routes.Create(templates["create.html"]))
	mux.Post("/create", routes.New(db))
	mux.Get("/:id", routes.View(db, templates["view.html"]))
	mux.Get("/:id/edit", routes.Edit(db, templates["edit.html"]))
	mux.Post("/:id/edit", routes.Save(db))
	mux.Get("/:id/obliterate", routes.Obliterate(db))

	http.ListenAndServe(":5000", mux)
}
