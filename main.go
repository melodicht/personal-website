package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

const PROD_BASE_PATH = "/personal-website"

func main() {
	local := flag.Bool("local", false, "generating for hosting locally")
	generate := flag.Bool("generate", false, "generate static site into docs/ and exit")
	flag.Parse()

	basePath := ""
	if !*local {
		basePath = PROD_BASE_PATH
	}
	
	if *generate {
		runGenerate(basePath)
		return
	}

	runServer(basePath)
}

func runServer(basePath string) {
	// Serve the generated static site from docs/ and static assets from static/
	http.Handle(basePath+"/static/", http.StripPrefix(basePath+"/static/", http.FileServer(http.Dir("docs/static"))))
	http.Handle(basePath+"/", http.StripPrefix(basePath+"/", http.FileServer(http.Dir("docs"))))

	addr := fmt.Sprintf(":%d", port)
	log.Println("Listening on http://localhost" + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
