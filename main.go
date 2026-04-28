package main

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed static/config.json
var configData []byte

var indexTmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := indexTmpl.Execute(w, template.JS(configData)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
