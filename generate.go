package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var projects = []Project{

	// ── JOB: Eyebot ───────────────────────────────────────────────
	{
		Title:       "Eyebot",
		Description: "A startup making vision healthcare more accessible through an automated kiosk that performs on-site eye tests and routes results to a certified doctor for a same-day prescription.",
		Type:        ProjectTypeJob,
		Category:    catPtr(CategorySystems),
		Specifics: ProjectTypeSpecifics{
			Job: &JobExperience{
				Company:         "Eyebot",
				Role:            "Software Engineer Intern",
				BackgroundImage: "static/images/jobs/eyebot-bg.jpg",
				PortraitImage:   imgPtr("static/images/jobs/eyebot-portrait.jpg"),
				DateRange:       DateRange{Start: "May 2025", End: strPtr("Dec 2025")},
				Reviews: []Review{
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
		Subprojects: []Subproject{
			{
				Title: "Bland-Altman Analysis with NumPy 2 and Matplotlib",
				Description: "Analyzed kiosk performance data using NumPy 2 and Matplotlib, implementing Bland-Altman statistical techniques to quantify compliance with operational standards.",
				Tags: []Tag{TagFullStack},
				TechTags: []TechTag{TechTagPython, TechTagNumPy, TechTagMatplotlib},
			},
			{
				Title: "Production and Lifetime-testing Telemetry",
				Description: "Integrated Datadog for production and lifetime-testing telemetry, recording anomalous behaviour and incidents, with data visualization",
				Tags: []Tag{TagFullStack},
				TechTags: []TechTag{TechTagPython, TechTagDatadog},
			},
			{
				Title: "Lifetime-testing Framework",
				Description: "Developed a Python framework for lifetime testing that enabled non-specialist engineers to create robust test scripts through a simplified async interface with built-in error handling, data collection and logging, reducing development time.",
				Tags: []Tag{TagFullStack},
				TechTags: []TechTag{TechTagPython, TechTagAsync},
			},
			{
				Title: "Bash-scripted Deployment Automation",
				Description: "Managed deployment lifecycle for kiosk software across development environments through Bash-scripted automation, reducing development time, allowing quick iteration and consistent releases.",
				Tags: []Tag{TagFullStack},
				TechTags: []TechTag{TechTagBash, TechTagLinux},
			},
			{
				Title: "Android App for Calibrating Kiosks",
				Description: "Developed Android calibration app using Jetpack Compose, Timber logging, and SocketIO networking to enable non-technical staff to configure and maintain kiosks without engineering support.",
				Tags: []Tag{TagFullStack},
				TechTags: []TechTag{TechTagKotlin, TechTagAndroidDevelopment, TechTagJetpackCompose, TechTagTimber},
			},
			{
				Title: "Improve Flask Web App with Server-side Rendering & Components Library",
				Description: "Enhanced Flask web portal performance by implementing server-side rendering and a reusable component library based on atomic design principles, reducing page load times and ensuring UI consistency.",
				Tags: []Tag{TagFullStack},
				TechTags: []TechTag{TechTagPython, TechTagFlask, TechTagServersideRendering, TechTagComponentsLibrary},
			},
			{
				Title: "Pupil Labeling Feature on Web App",
				Description: "Developed full-stack pupil labeling feature integrating JavaScript/HTML frontend with Firestore backend for persistent data management.",
				Tags: []Tag{TagFullStack},
				TechTags: []TechTag{TechTagPython, TechTagFlask, TechTagJavaScript, TechTagHTML, TechTagFirestore},
			},
			{
				Title: "Kiosk Performance Report through Slack",
				Description: "Automated kiosk performance monitoring by deploying a Slack bot using Google Cloud Services, Cloud Scheduler, and Slack Python API to deliver real-time technical and commercial metrics to stakeholders.",
				Tags: []Tag{TagFullStack},
				TechTags: []TechTag{TechTagPython, TechTagSlackAPI, TechTagGoogleCloudRunFunctions, TechTagGoogleCloudScheduler},
			},
			{
				Title: "Internal Python Library",
				Description: "Established internal Python library repository with automated documentation generation (Sphinx), testing pipeline (tox), code coverage reporting (Coverage.py), and SSH-authenticated pip installation, eliminating code duplication across projects and streamlining source access.",
				Tags: []Tag{TagFullStack},
				TechTags: []TechTag{TechTagPython, TechTagSphinx, TechTagTox, TechTagCoveragePy, TechTagPip},
			},
			{
				Title: "Penetration Testing the Kiosk",
				Description: "Conducted penetration testing on kiosk infrastructure using STRIDE threat modeling, identifying 4 critical vulnerabilities and delivering remediation strategies that eliminated or mitigated security risks.",
				Tags: []Tag{TagSystemSecurity},
				TechTags: []TechTag{TechTagSTRIDE},
			},
			{
				Title: "Kiosk Hardware Assembly",
				Description: "Assembled kiosk hardware following ESD-safe procedures, performing precision soldering, optical component alignment, and torque-controlled fastening while documenting standard operating procedures.",
				Tags: []Tag{TagHardwareTech},
				TechTags: []TechTag{},
			},
			{
				Title: "",
				Description: "Developed SOP for simulating degraded network conditions using Linux traffic control to validate kiosk performance across varying connectivity scenarios prior to deployment.",
				Tags: []Tag{TagDevOps},
				TechTags: []TechTag{TechTagLinux, TechTagTrafficControl},
			},
		},
	},

	// ── JOB: AirAsia ──────────────────────────────────────────────
	{
		Title:       "IKHLAS (subsidiary of AirAsia)",
		Description: "Provides Muslim communities around the world access to faith-based practices, much of it through digital technology.",
		Type:        ProjectTypeJob,
		Category:    catPtr(CategorySystems),
		Specifics: ProjectTypeSpecifics{
			Job: &JobExperience{
				Company:         "AirAsia",
				Role:            "Software Engineer Apprentice",
				BackgroundImage: "static/images/jobs/airasia-bg.jpg",
				DateRange:       DateRange{Start: "Nov 2023", End: strPtr("Aug 2024")},
				Reviews: []Review{
					{
						ProfilePicture: "static/images/reviews/airasia-manager.jpg",
						Name:           "Wei Liang Tan",
						Role:           "Tech Lead, AirAsia Platform",
						Text:           "Marvin was one of the strongest apprentices we've had on the platform team. He picked up our internal tooling stack quickly and delivered the weekly digest feature with minimal guidance. His code quality and attention to edge cases were well above what we'd typically expect at this stage.",
					},
				},
			},
		},
		Subprojects: []Subproject{
			{
				Title: "Reducing Flutter App Startup Time",
				Description: "Improved Flutter app user experience by developing and open-sourcing dartprofiler, an instrumental profiler for Dart programming language that can tailor to specific device chipsets by using Dart FFI into C++ and inlined assembly, identifying and eliminating 40-80% of startup time.",
				Tags: []Tag{TagFullStack, TagSystems},
				TechTags: []TechTag{TechTagFlutter, TechTagDart, TechTagProfiling, TechTagFFI, TechTagCPP, TechTagAssembly},
			},
			{
				Title: "Automated Scraper & Ingestion Pipeline",
				Description: "Developed and deployed a scalable prayer time scraper and ingester using ETL architecture, Graph query language, TypeScript, Cloud Functions for Firebase, and Puppeteer, providing 4800 daily active users with accurate, government-approved prayer times 24/7.",
				Tags: []Tag{TagFullStack},
				TechTags: []TechTag{TechTagTypeScript, TechTagETL, TechTagGraphQueryLanguage, TechTagGoogleCloudRunFunctions, TechTagFirestore, TechTagPuppeteer},
			},
			{
				Title: "Flutter Native Widget Integration",
				Description: "Implemented a new prayer times widget for our Flutter app natively in Swift (iOS) and Kotlin (Android), ensuring seamless communication between native and Flutter sides, resulting in a 100% increase in daily active users.",
				Tags: []Tag{TagFullStack, TagMobileAppDev},
				TechTags: []TechTag{TechTagDart, TechTagSwift, TechTagKotlin, TechTagFlutter},
			},
			{
				Title: "Optimized Web Blog's SEO",
				Description: "By implementing Open Graph meta tags, generating a sitemap.xml for 400+ blogs, and using WebP images for faster load times, resulting in an 11% increase in active users.",
				Tags: []Tag{TagWebDev},
				TechTags: []TechTag{TechTagHTML, TechTagXML, TechTagOpenGraph, TechTagSiteMap, TechTagWebP},
			},
			{
				Title: "Internal React Components Library",
				Description: "Spearheaded the development of an internal React components library; utilized Storybook and webpack for deploying the components workshop for documentation and visual testing, and Vite for library deployment; formalized the development pipeline from UI/UX design to implementation, improving documentation and enforcing a unifying design system.",
				Tags: []Tag{TagWebDev, TagFullStack, TagDevOps},
				TechTags: []TechTag{TechTagReact, TechTagJavaScript, TechTagTypeScript, TechTagStorybook, TechTagWebpack, TechTagVite},
			},
		},
	},

	// ── JOB: Khoury College ───────────────────────────────────────
	{
		Title:       "Khoury College of Computer Sciences",
		Description: "I was teaching assistant for undergraduate computer science courses, supporting students and professors through office hours, grading, exam review sessions and exam proctoring.",
		Type:        ProjectTypeJob,
		Category:    catPtr(CategorySystems),
		Specifics: ProjectTypeSpecifics{
			Job: &JobExperience{
				Company:         "Khoury College",
				Role:            "Teaching Assistant",
				BackgroundImage: "static/images/jobs/khoury-bg.jpg",
				DateRange:       DateRange{Start: "Sep 2022", End: nil},
				Reviews:         []Review{},
			},
		},
		Subprojects: []Subproject{
			{
				Title:       "Typecheck Action",
				Description: "GitHub Action that runs tsc, mypy, and go vet in parallel and posts a unified type-error report as a PR comment.",
				Tags:        []Tag{TagFullStack},
				Info:        SubprojectInfo{},
			},
		},
	},

	// ── UNIVERSITY: Ocaml Compiler  ───────────────────────────────
	{
		Title: "Designing a Compiler",
		Description: "(Insert explanation of what a compiler is and what kind of compiler we are building.) Written in OCaml, that generates my language into X86-64 assembly with C as the run-time.",
		Type: ProjectTypeUniversity,
		Category: catPtr(CategoryProgrammingLanguages),
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell: []string{},
				WhatCouldBeBetter: []string{},
				WhatILearned: []string{},
			},
		},
		Subprojects: []Subproject{
			{
				Title: "Garbage Collection",
				Description: "A stop-the-world garbage collector, by applying graph theory to implement Cheney’s semispace collector.",
				Tags: []Tag{TagSystems, TagProgrammingLangs},
				TechTags: []TechTag{TechTagC},
			},
			// Other things to include: ANF, graph colouring register allocation, built-in testing, error control flow
		},
	},

	// ── UNIVERSITY: Typed-untyped interactions  ───────────────────
	{
		Title: "Typed-untyped Interactions Through Machine",
		Description: "There are many way to implement a programming languages. One of the is through an abstract machine. On top of have all the basic features of a programming lanaguage (arithmetic, conditionals...), this programming language features the co-existence of typed and untyped modules. There are two distinct versions. One where the types are checked to the extent at possible, and then thrown away when the program actually runs. Another one where the types are checked, and become run-time checks during the program run.",
		Type: ProjectTypeUniversity,
		Category: catPtr(CategoryProgrammingLanguages),
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell: []string{"Macros for tests, nice organization"},
				WhatCouldBeBetter: []string{"The macros could be better"},
				WhatILearned: []string{"Making a language via a machine, logic, "},
			},
		},

		// Subprojects: CESK machine, Classes, Modules, Statements, Tail Calls, Type Stripping, Run-time Checks
		Subprojects: []Subproject{},
	},

	// ── UNIVERSITY: A Networked Card Game  ───────────────────
	{
		Title: "A Networked Card Game",
		Description: "For the Software Development course at Northeastern University, also known as \"hell\", we implemented a networked card game in about 12 weeks.",
		Type: ProjectTypeUniversity,
		Category: catPtr(CategorySystems),
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell: []string{"Fairly good organization, network/logic interactions with that design pattern"},
				WhatCouldBeBetter: []string{"Brute force algorithm should have used streams instead of whole lists."},
				WhatILearned: []string{"Software comes in layers.", "Got to be really meticulous when reading the specifications, take notes while reading."},
			},
		},

		// Subprojects: Greedy algorithm, fault-tolerant networking layer,
		Subprojects: []Subproject{},
	},

	// ── UNIVERSITY: RAFT  ───────────────────
	{
		Title: "RAFT: Consensus Algorithm For Distributed Systems",
		Description: "Even when we are working on a single device, the digital services we use normally have several computers working together behind the scenes. But how could these computers work together when communication between them could be faulty and go down at any time? Consensus algorithms are a key feature to make communication possible, by allowing many computers to agree on a value. RAFT is one such algorithm for this, which I have implemented for my final project for the Distributed Systems course at Northeastern University. It is important to note that I was under a very strict time constrain.",
		Type: ProjectTypeUniversity,
		Category: catPtr(CategorySystems),
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell: []string{"The custom allocator built for the log structure makes it memory efficient. A lot of the memory are contiguous. There is no fragmentation.", "As long as one node lives, the log structure will continue to live."},
				WhatCouldBeBetter: []string{"Leader will block for response after sending RPC for a period of time. If the response misses the window, it will be dropped.", "The logs are not very clean. When a server is shut down, all the other servers will flood the logs saying that they cannot connect to the server.", "There isn’t a clear clean separation between the logical RAFT layer and the server layer. This makes no difference to the user, but isn’t good for maintenance.", "The printing node data looks ugly. Invalid log index is long string of digits instead of saying invalid log index."},
				WhatILearned: []string{"Take advantage of leeways in the specifications."},
			},
		},

		// Subprojects: Leader Election, Log Replication, Fault Tolerance
		Subprojects: []Subproject{},
	},

	// ── PERSONAL: Skyline Engine  ─────────────────────────────────
	{
		Title: "A Custom Game Engine",
		Description: "C++ 20 (even though the code looks like C),- SDL3 for the platform layer, - Vulkan for hardware-accelerated graphics,- Jolt for physics (the physics engine developed for the second Horizon Zero Dawn game), - Dear Imgui for UI for internal tooling (though we will probably that for the actual game's UI),- CMake for the build system, - We require GCC or Clang for the C++ compiler, with compiler extensions enabled",
		Type: ProjectTypePersonal,
		Category: catPtr(CategorySystems),
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell: []string{""},
				WhatCouldBeBetter: []string{""},
				WhatILearned: []string{""},
			},
		},

		// Subprojects: Hot-reloading, Looped-live Playback, Scene Editor, Engine Architecture
		Subprojects: []Subproject{},
	},

	// ── PERSONAL: My Personal Portfolio  ──────────────────────────
	{
		Title:       "My Personal Portfolio (This Website)",
		Description: "The website is designed to be able to be read by hiring managers efficiently from different fields, game development, game engine development, full-stack development and so on. The [I do] mode allows hiring managers to focus on seeing what content that's relevant to them only. The other modes provide all data unfiltered.",
		Type:        ProjectTypePersonal,
		Category:    catPtr(CategorySystems),
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell:      []string{},
				WhatCouldBeBetter: []string{},
				WhatILearned:      []string{"Using Datastar for the first time"},
			},
		},

		// Subprojects: Static-site generation, Hyper-media driven
		Subprojects: []Subproject{},
	},

	// ── UNIVERSITY: Dreams of Celestial Pull  ──────────────────────
	{
		Title:       "Dreams of Celestial Pull: A Physics-based First-person Shooter Platformer",
		Description: "For Game Capstone, the final games course at Northeastern University where you spend an entire semester developing a game, I single-handedly developed Dreams of Celestial Pull. The game is made with the custom game engine, Skyline Engine. If you are into game design, I recommend giving the game a try first before reading the below, as there will be game design spoilers.",
		Type:        ProjectTypeUniversity,
		Category:    catPtr(CategoryGames),
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell:      []string{"Lots of content from simple mechanics, gravcity ball allows for lots of different consequences."},
				WhatCouldBeBetter: []string{"Due to limitations of the graphics-portion of the engine, a lot of visual feedback that really needed to be there weren't. It would have gave players intuition about the nature of gravity balls.", "There is also no sound."},
				WhatILearned:      []string{"Before you can walk, you must crawl. I thought the game was going to have enemies that could shoot bullets, but I decided to focus on the mechanics even without enemies. With what I know now, I know what kind of enemies to design for that would optimize for how gravity balls work."},
			},
		},
		Subprojects: []Subproject{
			{
				Title:       "Moodboard",
				Description: "Offline-first image board for designers. Drag, drop, done. No account needed.",
				Tags:        []Tag{TagFullStack, TagSystems},
				Info:        SubprojectInfo{Video: vidPtr("static/videos/moodboard.mp4")},
			},
		},

		// Subprojects: FPS movement and RK4, game and level design (Jon Blow inspired)
	},

	// ── UNIVERSITY: Boids with Goals  ──────────────────────
	{
		Title: "Boids with Goals: A Game AI Project",
		Description: "XXX studied how a flock of birds move in the air and found a simple algorithm that seem to mimic how they actually do it in real-life. However, the algorithm itself doesn't consider boids interacting with obstructions, and also boids flying towards a goal. So, this projects takes the idea one step further to create boids that do.",
		Type: ProjectTypeUniversity,
		Category: catPtr(CategoryGames),
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell:      []string{},
				WhatCouldBeBetter: []string{"Boids like to fly towards a wall, and at the very last minute make a sharp turn. This isn't very realistic.", "Boids ocassionally fly in circles far from the goals. likely caused by counter-balancing forces between the the structure of the walls and the goals.", "There is no way to level-edit while the program is running. The placement of walls are hard-coded.", "Boids sometimes move in a staircase pattern, likely caused by the fact that the path-finding only considers cardinal directios and doesn't smooth out the paths. "},
				WhatILearned:      []string{"The sum of forces is a powerful idea."},
			},
		},

		// Subprojects: Collision, Path-finding behaviour
		Subprojects: []Subproject{},
	},

	// ── PERSONAL: Toxic Texting  ──────────────────────
	{
		Title: "Toxic Texting: A Chill Texting Game Where You Only Respond with Yes or No",
		Description: "Made initially for the Summer 2021 NEU Game Development Club 48-hour game jam, Toxic Texting is a fun, short and sweet texting game where you respond with either yes or no. It is made in Unity 2D, and I was one of the two programmers, the composer, sounnd designer and writer.",
		Type: ProjectTypePersonal,
		Category: catPtr(CategoryGames),
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell:      []string{"The user experience is on point."},
				WhatCouldBeBetter: []string{"The dialogue system is a pain."},
				WhatILearned:      []string{"Take advantage of Unity's UI system as much as possible, so that there's less work on my part."},
			},
		},

		// Subprojects: Dialogue system, the dialogue themselves, sound and music, gameplay system.
		Subprojects: []Subproject{},
	},

	// ── PERSONAL: Tower Takeover  ──────────────────────
	{
		Title: "Tower Takeover: A vanilla JavaScript Tower Defense Game That Runs on your Web Browser",
		Description: "This is a game that I led for formerly Hometeam Game Dev (now DevPods). Hometeam Game Dev is a community of game developers that makes games without pay. I was responsible for structuring tasks on a kanban board, running playtests, and making sure that the game ships.",
		Type: ProjectTypePersonal,
		Category: catPtr(CategoryGames),
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell:      []string{"Gameplay system is easy to work with."},
				WhatCouldBeBetter: []string{"The game is too hard, and without written instructions, it's hard to figure out how to play the game."},
				WhatILearned:      []string{"This game has a lot of information and actions. I should have designed the game around that, instead of the raw systems. Start with what the user experience, and not the underlying physics of the game."},
			},
		},
		Subprojects: []Subproject{},
	},
}

func strPtr(s string) *string                                      { return &s }
func imgPtr(i Image) *Image                      { return &i }
func vidPtr(v Video) *Video                      { return &v }
func catPtr(c ProjectCategory) *ProjectCategory  { return &c }

func runGenerate() {
	// Compute Tags on each Project as the union of its subprojects' tags.
	for i := range projects {
		seen := map[Tag]bool{}
		for _, sp := range projects[i].Subprojects {
			for _, t := range sp.Tags {
				seen[t] = true
			}
		}
		tags := make([]Tag, 0, len(seen))
		for t := range seen {
			tags = append(tags, t)
		}
		projects[i].Tags = tags
	}

	data := SiteData{Projects: projects}

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
