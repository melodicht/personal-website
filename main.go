package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"ds_demo/sitedata"
)

//go:embed static/config.json
var configData []byte

//go:embed data.json
var embeddedSiteData []byte

var indexTmpl = template.Must(template.ParseFiles("templates/index.html"))

type templateData struct {
	Config   template.JS
	SiteData template.JS
}

func main() {
	generate := flag.Bool("generate", false, "generate data.json and exit")
	flag.Parse()

	if *generate {
		runGenerate()
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		td := templateData{
			Config:   template.JS(configData),
			SiteData: template.JS(embeddedSiteData),
		}
		if err := indexTmpl.Execute(w, td); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func strPtr(s string) *string                              { return &s }
func imgPtr(i sitedata.Image) *sitedata.Image              { return &i }
func vidPtr(v sitedata.Video) *sitedata.Video              { return &v }
func catPtr(c sitedata.ProjectCategory) *sitedata.ProjectCategory { return &c }

func runGenerate() {
	projects := []sitedata.Project{

		// ── JOB: Eyebot ───────────────────────────────────────────────
		{
			Title:       "Eyebot",
			Description: "Software engineering internship building real-time computer vision pipelines and tooling for an early-stage robotics startup.",
			Type:        sitedata.ProjectTypeJob,
			Category:    catPtr(sitedata.CategorySystems),
			Specifics: sitedata.ProjectTypeSpecifics{
				Job: &sitedata.JobExperience{
					Company:         "Eyebot",
					Role:            "Software Engineer Intern",
					BackgroundImage: "static/images/jobs/eyebot-bg.jpg",
					PortraitImage:   imgPtr("static/images/jobs/eyebot-portrait.jpg"),
					DateRange:       sitedata.DateRange{Start: "May 2025", End: strPtr("Dec 2025")},
					Reviews: []sitedata.Review{
						{
							ProfilePicture: "static/images/reviews/eyebot-manager.jpg",
							Name:           "Sarah Chen",
							Role:           "Engineering Manager, Eyebot",
							Text:           "Marvin joined us mid-sprint and was shipping production code within the first week. His ability to reason about low-latency systems and communicate trade-offs clearly made him an asset to the team. I'd hire him full-time without hesitation.",
						},
						{
							ProfilePicture: "static/images/reviews/eyebot-peer1.jpg",
							Name:           "James Okafor",
							Role:           "Senior Software Engineer, Eyebot",
							Text:           "Working alongside Marvin was a genuine pleasure. He took ownership of the game engine integration end-to-end, asked sharp questions, and consistently delivered clean, well-documented code. A rare intern who makes the people around him better.",
						},
						{
							ProfilePicture: "static/images/reviews/eyebot-peer2.jpg",
							Name:           "Priya Nair",
							Role:           "Software Engineer, Eyebot",
							Text:           "Marvin has an impressive range — he could jump between writing a performant Go service in the morning and debugging a gnarly rendering issue in the afternoon. He's also just a great person to have on a team.",
						},
					},
				},
			},
			Subprojects: []sitedata.Subproject{
				{
					Title:       "Dungeon Typist",
					Description: "A typing-speed roguelike where every keystroke is an action. Built the core game loop and entity system using Ebitengine.",
					Tags:        []sitedata.Tag{sitedata.TagGameDev, sitedata.TagGameEngineDev},
					Info:        sitedata.SubprojectInfo{Video: vidPtr("static/videos/dungeon-typist.mp4")},
				},
				{
					Title:       "Grip",
					Description: "A fast grep-like CLI for structured logs. Parses JSON lines and filters by field with sub-millisecond latency.",
					Tags:        []sitedata.Tag{sitedata.TagSystems},
					Info:        sitedata.SubprojectInfo{},
				},
			},
		},

		// ── JOB: AirAsia ──────────────────────────────────────────────
		{
			Title:       "AirAsia",
			Description: "Software engineering apprenticeship on the platform team, building internal tooling and customer-facing web features.",
			Type:        sitedata.ProjectTypeJob,
			Category:    catPtr(sitedata.CategorySystems),
			Specifics: sitedata.ProjectTypeSpecifics{
				Job: &sitedata.JobExperience{
					Company:         "AirAsia",
					Role:            "Software Engineer Apprentice",
					BackgroundImage: "static/images/jobs/airasia-bg.jpg",
					DateRange:       sitedata.DateRange{Start: "Nov 2023", End: strPtr("Aug 2024")},
					Reviews: []sitedata.Review{
						{
							ProfilePicture: "static/images/reviews/airasia-manager.jpg",
							Name:           "Wei Liang Tan",
							Role:           "Tech Lead, AirAsia Platform",
							Text:           "Marvin was one of the strongest apprentices we've had on the platform team. He picked up our internal tooling stack quickly and delivered the weekly digest feature with minimal guidance. His code quality and attention to edge cases were well above what we'd typically expect at this stage.",
						},
					},
				},
			},
			Subprojects: []sitedata.Subproject{
				{
					Title:       "Recap",
					Description: "Weekly digest generator that pulls from GitHub, Linear, and the team calendar and emails a unified summary every Monday.",
					Tags:        []sitedata.Tag{sitedata.TagFullStack},
					Info:        sitedata.SubprojectInfo{},
				},
				{
					Title:       "Studio Nori",
					Description: "Branding site for an internal ceramics studio initiative. Scroll-driven animations, zero JS frameworks.",
					Tags:        []sitedata.Tag{sitedata.TagFullStack},
					Info:        sitedata.SubprojectInfo{},
				},
			},
		},

		// ── JOB: Khoury College ───────────────────────────────────────
		{
			Title:       "Khoury College",
			Description: "Teaching assistant for undergraduate computer science courses, supporting students through office hours, grading, and curriculum design.",
			Type:        sitedata.ProjectTypeJob,
			Category:    catPtr(sitedata.CategorySystems),
			Specifics: sitedata.ProjectTypeSpecifics{
				Job: &sitedata.JobExperience{
					Company:         "Khoury College",
					Role:            "Teaching Assistant",
					BackgroundImage: "static/images/jobs/khoury-bg.jpg",
					DateRange:       sitedata.DateRange{Start: "Sep 2022", End: nil},
					Reviews:         []sitedata.Review{},
				},
			},
			Subprojects: []sitedata.Subproject{
				{
					Title:       "Typecheck Action",
					Description: "GitHub Action that runs tsc, mypy, and go vet in parallel and posts a unified type-error report as a PR comment.",
					Tags:        []sitedata.Tag{sitedata.TagFullStack},
					Info:        sitedata.SubprojectInfo{},
				},
			},
		},

		// ── UNIVERSITY: Systems & Networks ────────────────────────────
		{
			Title:       "Systems & Networks",
			Description: "A pair of university projects exploring low-level systems programming and networked applications.",
			Type:        sitedata.ProjectTypeUniversity,
			Category:    catPtr(sitedata.CategorySystems),
			Specifics: sitedata.ProjectTypeSpecifics{
				NonJob: &sitedata.NonJobExperience{
					WhatWentWell:      "Both projects shipped on time with clean architectures. Terminal Quest's shell integration was particularly well-received.",
					WhatCouldBeBetter: "Test coverage was thinner than I'd have liked, especially around network edge cases in the Flashpack sync layer.",
					WhatILearned:      "How to reason carefully about concurrency and how small protocol decisions early on create large maintenance costs later.",
					SourceCodeLink:    strPtr("https://github.com/placeholder/systems-networks"),
				},
			},
			Subprojects: []sitedata.Subproject{
				{
					Title:       "Terminal Quest",
					Description: "Text adventure that runs inside your actual shell. ls to explore, rm to fight.",
					Tags:        []sitedata.Tag{sitedata.TagGameDev, sitedata.TagProgrammingLangs},
					Info:        sitedata.SubprojectInfo{},
				},
				{
					Title:       "Flashpack",
					Description: "Spaced-repetition flashcards that live in your terminal, with an optional peer-sync protocol over UDP.",
					Tags:        []sitedata.Tag{sitedata.TagSystems},
					Info:        sitedata.SubprojectInfo{},
				},
			},
		},

		// ── UNIVERSITY: Web & Tools Practicum ─────────────────────────
		{
			Title:       "Web & Tools Practicum",
			Description: "University practicum covering full-stack web development and developer tooling, culminating in two shipped projects.",
			Type:        sitedata.ProjectTypeUniversity,
			Category:    catPtr(sitedata.CategorySystems),
			Specifics: sitedata.ProjectTypeSpecifics{
				NonJob: &sitedata.NonJobExperience{
					WhatWentWell:      "Docs Redesign was adopted by the open-source project maintainers within a week of submission. Daylog has been in personal daily use ever since.",
					WhatCouldBeBetter: "I underestimated the time needed for cross-browser testing on Docs Redesign and had to rush the Safari fixes.",
					WhatILearned:      "The importance of writing for an audience you can't talk to, and how much a good email cadence matters for retention in consumer tools.",
					SourceCodeLink:    nil,
				},
			},
			Subprojects: []sitedata.Subproject{
				{
					Title:       "Docs Redesign",
					Description: "Redesigned the docs for an open-source CLI tool used by 12k developers. New information architecture and search.",
					Tags:        []sitedata.Tag{sitedata.TagFullStack},
					Info:        sitedata.SubprojectInfo{},
				},
				{
					Title:       "Daylog",
					Description: "A one-line-a-day journal that emails you your entry from exactly one year ago.",
					Tags:        []sitedata.Tag{sitedata.TagFullStack},
					Info:        sitedata.SubprojectInfo{},
				},
			},
		},

		// ── PERSONAL: Generative Art Studio ───────────────────────────
		{
			Title:       "Generative Art Studio",
			Description: "An ongoing personal project exploring procedural image generation and physical output via pen plotter.",
			Type:        sitedata.ProjectTypePersonal,
			Category:    catPtr(sitedata.CategorySystems),
			Specifics: sitedata.ProjectTypeSpecifics{
				NonJob: &sitedata.NonJobExperience{
					WhatWentWell:      "The plotter output exceeded expectations — the physical prints have a quality that screen rendering can't replicate.",
					WhatCouldBeBetter: "The SuperCollider drone engine grew into a monolith. Should have modularised the signal chain earlier.",
					WhatILearned:      "How to bridge creative and technical thinking, and that constraints (pen speed, paper texture) produce better art than infinite freedom.",
					SourceCodeLink:    strPtr("https://github.com/placeholder/generative-art"),
				},
			},
			Subprojects: []sitedata.Subproject{
				{
					Title:       "Grid Studies",
					Description: "A series of 100 generative pieces exploring grid distortion. Each one is unique, rendered with p5.js.",
					Tags:        []sitedata.Tag{sitedata.TagSystems},
					Info:        sitedata.SubprojectInfo{Video: vidPtr("static/videos/grid-studies.mp4")},
				},
				{
					Title:       "Ink & Code",
					Description: "Physical prints produced by a pen plotter driven by Go programs that generate SVG paths.",
					Tags:        []sitedata.Tag{sitedata.TagSystems},
					Info:        sitedata.SubprojectInfo{},
				},
				{
					Title:       "Generative Drones",
					Description: "Infinite ambient music engine written in SuperCollider. No two listens are the same.",
					Tags:        []sitedata.Tag{sitedata.TagSystems},
					Info:        sitedata.SubprojectInfo{},
				},
			},
		},

		// ── PERSONAL: Moodboard ───────────────────────────────────────
		{
			Title:       "Moodboard",
			Description: "Offline-first image board for designers. Drag, drop, done. No account needed. Built with Tauri and Rust.",
			Type:        sitedata.ProjectTypePersonal,
			Category:    catPtr(sitedata.CategorySystems),
			Specifics: sitedata.ProjectTypeSpecifics{
				NonJob: &sitedata.NonJobExperience{
					WhatWentWell:      "The offline-first architecture worked flawlessly. Users appreciated the zero-signup experience.",
					WhatCouldBeBetter: "Tauri's file system API had sharp edges that cost me two weeks. I'd reach for a different IPC approach next time.",
					WhatILearned:      "How to design for data ownership, and that Rust's borrow checker is genuinely your friend once you stop fighting it.",
					SourceCodeLink:    strPtr("https://github.com/placeholder/moodboard"),
				},
			},
			Subprojects: []sitedata.Subproject{
				{
					Title:       "Moodboard",
					Description: "Offline-first image board for designers. Drag, drop, done. No account needed.",
					Tags:        []sitedata.Tag{sitedata.TagFullStack, sitedata.TagSystems},
					Info:        sitedata.SubprojectInfo{Video: vidPtr("static/videos/moodboard.mp4")},
				},
			},
		},

		// ── PERSONAL: Gravity Flip ────────────────────────────────────
		{
			Title:       "Gravity Flip",
			Description: "Puzzle platformer with one mechanic — gravity inversion — taken to its logical extreme across 30 handcrafted levels.",
			Type:        sitedata.ProjectTypePersonal,
			Category:    catPtr(sitedata.CategoryGames),
			Specifics: sitedata.ProjectTypeSpecifics{
				NonJob: &sitedata.NonJobExperience{
					WhatWentWell:      "Level design came together quickly once I committed to the single-mechanic constraint. Less is more.",
					WhatCouldBeBetter: "The physics engine had floating-point precision issues in edge cases that I never fully resolved.",
					WhatILearned:      "How to scope a game to actually finish it, and that Unity's physics system rewards experimentation over documentation-reading.",
					SourceCodeLink:    nil,
				},
			},
			Subprojects: []sitedata.Subproject{
				{
					Title:       "Gravity Flip",
					Description: "Puzzle platformer with one mechanic taken to its logical extreme.",
					Tags:        []sitedata.Tag{sitedata.TagGameDev},
					Info:        sitedata.SubprojectInfo{Video: vidPtr("static/videos/gravity-flip.mp4")},
				},
			},
		},

		// ── PERSONAL: Portfolio v2 ────────────────────────────────────
		{
			Title:       "Portfolio v2",
			Description: "This website. Built with Go, Datastar, and a lot of CSS.",
			Type:        sitedata.ProjectTypePersonal,
			Category:    catPtr(sitedata.CategorySystems),
			Specifics: sitedata.ProjectTypeSpecifics{
				NonJob: &sitedata.NonJobExperience{
					WhatWentWell:      "The Datastar + Go stack turned out to be a genuinely pleasant pairing for a content-heavy static site.",
					WhatCouldBeBetter: "The initial signal architecture needed several iterations to get right.",
					WhatILearned:      "That hypermedia-first frameworks reward thinking carefully about data flow before writing a single line of HTML.",
					SourceCodeLink:    strPtr("https://github.com/placeholder/portfolio"),
				},
			},
			Subprojects: []sitedata.Subproject{
				{
					Title:       "Portfolio v2",
					Description: "A personal site with a brutalist aesthetic and a lot of opinions about typography.",
					Tags:        []sitedata.Tag{sitedata.TagFullStack},
					Info:        sitedata.SubprojectInfo{},
				},
			},
		},
	}

	// Compute Tags on each Project as the union of its subprojects' tags.
	for i := range projects {
		seen := map[sitedata.Tag]bool{}
		for _, sp := range projects[i].Subprojects {
			for _, t := range sp.Tags {
				seen[t] = true
			}
		}
		tags := make([]sitedata.Tag, 0, len(seen))
		for t := range seen {
			tags = append(tags, t)
		}
		projects[i].Tags = tags
	}

	data := sitedata.SiteData{Projects: projects}

	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "marshal error: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile("data.json", b, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "write error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("wrote data.json")
}
