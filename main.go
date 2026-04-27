package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

//go:embed static/config.json
var configData []byte

type Project struct {
	Title       string
	Description string
	Tags        []string
	Link        string
}

var projects = map[string][]Project{
	"websites": {
		{Title: "Portfolio v2", Description: "A brutalist personal site built with plain HTML and a lot of opinions.", Tags: []string{"HTML", "CSS"}, Link: "#"},
		{Title: "Studio Nori", Description: "Branding site for a small ceramics studio. Scroll-driven animations, zero JS frameworks.", Tags: []string{"Astro", "CSS"}, Link: "#"},
		{Title: "Docs Redesign", Description: "Redesigned the docs for an open-source CLI tool used by 12k developers.", Tags: []string{"VitePress", "Figma"}, Link: "#"},
	},
	"games": {
		{Title: "Dungeon Typist", Description: "A typing-speed roguelike where every keystroke is an action. Made in 72 hours for Ludum Dare.", Tags: []string{"Go", "Ebitengine"}, Link: "#"},
		{Title: "Gravity Flip", Description: "Puzzle platformer with one mechanic taken to its logical extreme.", Tags: []string{"Unity", "C#"}, Link: "#"},
		{Title: "Terminal Quest", Description: "Text adventure that runs inside your actual shell. ls to explore, rm to fight.", Tags: []string{"Go", "Bubble Tea"}, Link: "#"},
	},
	"music": {
		{Title: "Slow Burn EP", Description: "Four tracks of ambient electronic music composed during a winter residency.", Tags: []string{"Ableton", "Max/MSP"}, Link: "#"},
		{Title: "Generative Drones", Description: "Infinite ambient music engine — no two listens are the same.", Tags: []string{"SuperCollider"}, Link: "#"},
	},
	"tools": {
		{Title: "Grip", Description: "A fast grep-like CLI for structured logs. Parses JSON lines and filters by field.", Tags: []string{"Go", "CLI"}, Link: "#"},
		{Title: "Recap", Description: "Weekly digest generator that pulls from GitHub, Linear, and your calendar.", Tags: []string{"Go", "API"}, Link: "#"},
		{Title: "Typecheck Action", Description: "GitHub Action that runs tsc, mypy, and go vet in parallel and posts a unified report.", Tags: []string{"Go", "GitHub Actions"}, Link: "#"},
	},
	"art": {
		{Title: "Grid Studies", Description: "A series of 100 generative pieces exploring grid distortion. Each one is unique.", Tags: []string{"p5.js", "Generative"}, Link: "#"},
		{Title: "Ink & Code", Description: "Physical prints produced by a pen plotter driven by Go programs.", Tags: []string{"Go", "SVG", "Plotter"}, Link: "#"},
	},
	"apps": {
		{Title: "Moodboard", Description: "Offline-first image board for designers. Drag, drop, done. No account needed.", Tags: []string{"Tauri", "Rust"}, Link: "#"},
		{Title: "Flashpack", Description: "Spaced-repetition flashcards that live in your terminal.", Tags: []string{"Go", "Bubble Tea"}, Link: "#"},
		{Title: "Daylog", Description: "A one-line-a-day journal that emails you your entry from exactly one year ago.", Tags: []string{"Go", "SQLite"}, Link: "#"},
	},
}

var fallbackProjects = []Project{
	{Title: "Coming soon", Description: "Projects in this category are on the way.", Tags: []string{}, Link: "#"},
}

var indexTmpl = template.Must(template.ParseFiles("templates/index.html"))
var cardsTmpl = template.Must(template.ParseFiles("templates/cards.html"))

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/projects", handleProjects)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Pass config as a template.JS value so it is injected unescaped into <script>.
	if err := indexTmpl.Execute(w, template.JS(configData)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleProjects(w http.ResponseWriter, r *http.Request) {
	word := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("word")))

	cards, ok := projects[word]
	if !ok || len(cards) == 0 {
		cards = fallbackProjects
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("X-Accel-Buffering", "no")

	var sb strings.Builder
	if err := cardsTmpl.Execute(&sb, cards); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "event: datastar-patch-elements\n")
	for _, line := range strings.Split(strings.TrimSpace(sb.String()), "\n") {
		fmt.Fprintf(w, "data: elements %s\n", line)
	}
	fmt.Fprintf(w, "\n")

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}
