package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	generate := flag.Bool("generate", false, "generate static site into docs/ and exit")
	flag.Parse()

	if *generate {
		runGenerate()
		return
	}

	runServer()
}

func runServer() {
	// Serve the generated static site from docs/ and static assets from static/
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", http.FileServer(http.Dir("docs")))

	addr := fmt.Sprintf(":%d", port)
	log.Println("Listening on http://localhost" + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
