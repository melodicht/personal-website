package main

import (
	_ "embed"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed static/config.json
var configData []byte

//go:embed data.json
var embeddedSiteData []byte

var indexTmpl = template.Must(template.ParseFiles("templates/index.html"))

const port = 8080

type templateData struct {
	Config   template.JS
	SiteData template.JS
}

func main() {
	generate := flag.Bool("generate", false, "generate data.json and exit")
	flag.Parse()

	if *generate {
		runGenerate()
	} else {
		runServer()
	}
}

func runServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		td := templateData{
			Config:   template.JS(configData),
			SiteData: template.JS(embeddedSiteData),
		}
		if err := indexTmpl.Execute(w, td); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	// NOTE(marvin): setting Directory to static to prevent path
	// traversal exploits.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	addr := fmt.Sprintf(":%d", port)
	log.Println("Listening on http://localhost" + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
