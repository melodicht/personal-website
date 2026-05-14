package main


var projects = []Project{

	// ── JOB: Eyebot ───────────────────────────────────────────────
	{
		Title:       "Eyebot",
		Description: "A startup making vision healthcare more accessible through an automated kiosk that performs on-site eye tests and routes results to a certified doctor for a same-day prescription.",
		Type:        ProjectTypeJob,
		Category:    catPtr(CategorySystems),
		Focuses:        []Focus{},
		TechTags:    []TechTag{},
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
		Subsections: []Subsection{
			{
				Title:    "Full-stack Development",
				Focuses:     []Focus{},
				TechTags: []TechTag{},
				Bullets: []BulletPoint{
					{Subproject: Subproject{
						Title:       "Bland-Altman Analysis with NumPy 2 and Matplotlib",
						Description: "Analyzed kiosk performance data using {NumPy 2} and {Matplotlib}, implementing Bland-Altman statistical techniques to quantify compliance with operational standards, and generating graphs and PDF reports.",
						Focuses:        []Focus{FocusFullStack},
						TechTags:    []TechTag{TechTagPython, TechTagNumPy, TechTagMatplotlib, TechTagAlgorithms},
					}},
					{Subproject: Subproject{
						Title:       "Lifetime-testing Framework",
						Description: "Developed a {Python} framework for lifetime testing that enabled non-specialist engineers to create robust test scripts through a simplified {async} interface with built-in error handling, data collection with {Firestore} and {Google Cloud Storage}, and logging, reducing development time.",
						Focuses:        []Focus{FocusFullStack},
						TechTags:    []TechTag{TechTagPython, TechTagAsync, TechTagFirestore, TechTagGoogleCloudStorage},
					}},
					{Subproject: Subproject{
						Title:       "Android App for Calibrating Kiosks",
						Description: "Developed {Android} calibration app using {Jetpack Compose}, {Timber} logging, {Firestore}, and {SocketIO} networking to enable non-technical staff to configure and maintain kiosks without engineering support.",
						Focuses:        []Focus{FocusMobileAppDev, FocusFullStack},
						TechTags:    []TechTag{TechTagKotlin, TechTagAndroidDevelopment, TechTagJetpackCompose, TechTagTimber, TechTagFirestore, TechTagSocketIO},
					}},
					{Subproject: Subproject{
						Title:       "Improve Flask Web App with Server-side Rendering & Components Library",
						Description: "Enhanced {Python} {Flask} web portal performance by implementing {server-side rendering} and a reusable {components library} based on atomic design principles, reducing page load times and ensuring UI consistency.",
						Focuses:        []Focus{FocusWebDev, FocusFullStack},
						TechTags:    []TechTag{TechTagPython, TechTagFlask, TechTagServersideRendering, TechTagComponentsLibrary},
					}},
					{Subproject: Subproject{
						Title:       "Pupil Labeling Feature on Web App",
						Description: "Developed full-stack pupil labeling feature for {Python} {Flask} app integrating {JavaScript}/{HTML} frontend with {Firestore} backend for persistent data management.",
						Focuses:        []Focus{FocusWebDev, FocusFullStack},
						TechTags:    []TechTag{TechTagPython, TechTagFlask, TechTagJavaScript, TechTagHTML, TechTagFirestore},
					}},
					{Subproject: Subproject{
						Title:       "Kiosk Performance Report through Slack",
						Description: "Automated kiosk performance monitoring by deploying a Slack bot using {Python}, {Google Cloud Run Functions}, {Google Cloud Scheduler}, and {Slack API} to deliver real-time technical and commercial metrics to stakeholders.",
						Focuses:        []Focus{FocusFullStack},
						TechTags:    []TechTag{TechTagPython, TechTagSlackAPI, TechTagGoogleCloudRunFunctions, TechTagGoogleCloudScheduler},
					}},
					{Subproject: Subproject{
						Title:       "Internal Python Library",
						Description: "Established internal {Python} library repository with automated documentation generation ({Sphinx}), testing pipeline ({tox}), code coverage reporting ({Coverage.py}), and SSH-authenticated {pip} installation, eliminating code duplication across projects and streamlining souprce access.",
						Focuses:        []Focus{FocusFullStack, FocusDevOps},
						TechTags:    []TechTag{TechTagPython, TechTagSphinx, TechTagTox, TechTagCoveragePy, TechTagPip},
					}},
				},
			},
			{
				Title:    "Development Operations & Test Engineering",
				Focuses:     []Focus{FocusDevOps},
				TechTags: []TechTag{TechTagLinux},
				Bullets: []BulletPoint{
					{Subproject: Subproject{
						Title:       "Production and Lifetime-testing Telemetry",
						Description: "Integrated {Datadog} for production and lifetime-testing telemetry on kiosks running {Linux}, recording anomalous behaviour and incidents, with data visualization",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagDatadog},
					}},
					{Subproject: Subproject{
						Title:       "Bash-scripted Deployment Automation",
						Description: "Managed deployment lifecycle for kiosk software across {Linux} development environments through {Bash}-scripted automation, reducing development time, allowing quick iteration and consistent releases.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagBash},
					}},
					{Subproject: Subproject{
						Title:       "Simulating Degraded Network Conditions",
						Description: "Developed SOP for simulating degraded network conditions using {Linux} {traffic control (tc)} to validate kiosk performance across varying connectivity scenarios prior to deployment.",
						Focuses:        []Focus{FocusDevOps},
						TechTags:    []TechTag{TechTagTrafficControl},
					}},
				},
			},
			{
				Title:    "Misc",
				Focuses:     []Focus{},
				TechTags: []TechTag{},
				Bullets: []BulletPoint{
					{Subproject: Subproject{
						Title:       "Penetration Testing the Kiosk",
						Description: "Conducted penetration testing on kiosk infrastructure using {STRIDE threat modelling}, identifying 4 critical vulnerabilities and delivering remediation strategies that eliminated or mitigated security risks.",
						Focuses:        []Focus{FocusSystemSecurity},
						TechTags:    []TechTag{TechTagSTRIDE},
					}},
					{Subproject: Subproject{
						Title:       "Kiosk Hardware Assembly",
						Description: "Assembled kiosk hardware following ESD-safe procedures, performing precision soldering, optical component alignment, and torque-controlled fastening while documenting standard operating procedures.",
						Focuses:        []Focus{FocusHardwareTech},
						TechTags:    []TechTag{},
					}},
				},
			},
			
		},
	},

	// ── JOB: AirAsia ──────────────────────────────────────────────
	{
		Title:       "IKHLAS (subsidiary of AirAsia)",
		Description: "Provides Muslim communities around the world access to faith-based practices, much of it through digital technology.",
		Type:        ProjectTypeJob,
		Category:    catPtr(CategorySystems),
		Focuses:        []Focus{},
		TechTags:    []TechTag{},
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
		Subsections: []Subsection{
			{
				Title:    "Mobile App Development",
				Focuses:     []Focus{FocusMobileAppDev},
				TechTags: []TechTag{},
				Bullets: []BulletPoint{
					{Subproject: Subproject{
						Title:       "Reducing Flutter App Startup Time",
						Description: "Improved {Flutter} app user experience by developing and open-sourcing dartprofiler, an instrumental profiler for {Dart} programming language that can tailor to specific device chipsets by using Dart {FFI} into {C++} and inlined {assembly}, identifying and eliminating 40-80% of startup time.",
						Focuses:        []Focus{FocusSystems, FocusMobileAppDev},
						TechTags:    []TechTag{TechTagFlutter, TechTagDart, TechTagProfiling, TechTagFFI, TechTagCPP, TechTagAssembly},
					}},
					{Subproject: Subproject{
						Title:       "Flutter Native Widget Integration",
						Description: "Implemented a new prayer times widget for our {Flutter} app natively in {Swift} (iOS) and {Kotlin} (Android), with appropriate interfacing {Dart} code, ensuring seamless communication between native and Flutter sides, resulting in a 100% increase in daily active users.",
						Focuses:        []Focus{FocusMobileAppDev, FocusFullStack},
						TechTags:    []TechTag{TechTagDart, TechTagSwift, TechTagKotlin, TechTagFlutter},
					}},
				},
			},
			{
				Title:    "Web Development",
				Focuses:     []Focus{FocusWebDev},
				TechTags: []TechTag{},
				Bullets: []BulletPoint{
					{Subproject: Subproject{
						Title:       "Optimized Web Blog's SEO",
						Description: "Implemented {Open Graph} meta tags, generated a sitemap.xml for 400+ blogs, and used {WebP} images for faster load times, resulting in an 11% increase in active users.",
						Focuses:        []Focus{FocusWebDev},
						TechTags:    []TechTag{TechTagHTML, TechTagXML, TechTagOpenGraph, TechTagSiteMap, TechTagWebP},
					}},
					{Subproject: Subproject{
						Title:       "Internal React Components Library",
						Description: "Spearheaded the development of an internal {React} components library compatible with both {JavaScript} and {TypeScript} projects; utilized {Storybook} and {webpack} for deploying the components workshop for documentation and visual testing, and {Vite} for library deployment; formalized the development pipeline from UI/UX design to implementation, improving documentation and enforcing a unifying design system.",
						Focuses:        []Focus{FocusWebDev, FocusDevOps},
						TechTags:    []TechTag{TechTagJavaScript, TechTagTypeScript, TechTagReact, TechTagStorybook, TechTagWebpack, TechTagVite},
					}},
				},
			},
			{
				Title:    "Scraper & Data Pipeline",
				Focuses:     []Focus{},
				TechTags: []TechTag{},
				Bullets: []BulletPoint{
					
					{Subproject: Subproject{
						Title:       "Automated Scraper & Ingestion Pipeline",
						Description: "Developed and deployed a scalable prayer time scraper and ingester using {ETL architecture}, {Graph query language}, {Node.js}, {TypeScript}, {Google Cloud Run Functions}, and {Puppeteer}, providing 4800 daily active users with accurate, government-approved prayer times 24/7.",
						Focuses:        []Focus{FocusFullStack, FocusDevOps},
						TechTags:    []TechTag{TechTagTypeScript, TechTagNodeJS, TechTagETL, TechTagGraphQueryLanguage, TechTagGoogleCloudRunFunctions, TechTagFirestore, TechTagPuppeteer},
					}},
					
				},
			},
		},
	},

	// ── JOB: Khoury College ───────────────────────────────────────
	{
		Title:       "Khoury College of Computer Sciences",
		Description: "I was teaching assistant for undergraduate computer science courses, supporting students and professors through office hours, grading, exam review sessions and exam proctoring.",
		Type:        ProjectTypeJob,
		Category:    catPtr(CategorySystems),
		Focuses:        []Focus{},
		TechTags:    []TechTag{},
		Specifics: ProjectTypeSpecifics{
			Job: &JobExperience{
				Company:         "Khoury College",
				Role:            "Teaching Assistant",
				BackgroundImage: "static/images/jobs/khoury-bg.jpg",
				DateRange:       DateRange{Start: "Sep 2022", End: nil},
				Reviews:         []Review{},
			},
		},
		Subsections: []Subsection{},
	},

	// ── UNIVERSITY: Ocaml Compiler  ───────────────────────────────
	{
		Title:       "Designing a Compiler",
		Description: "(Insert explanation of what a compiler is and what kind of compiler we are building.) Written in OCaml, that generates my language into X86-64 assembly with C as the run-time.",
		Type:        ProjectTypeUniversity,
		Category:    catPtr(CategoryProgrammingLanguages),
		Focuses:        []Focus{FocusProgrammingLangs},
		TechTags:    []TechTag{TechTagOCaml},
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell:      []string{},
				WhatCouldBeBetter: []string{},
				WhatILearned:      []string{},
			},
		},
		Subsections: []Subsection{
			{
				Title:    "Subprojects",
				Focuses:     []Focus{},
				TechTags: []TechTag{},
				Cards: []Card{
					{Subproject: Subproject{
						Title:       "Garbage Collection",
						Description: "A stop-the-world garbage collector, by applying graph theory to implement Cheney's semispace collector, implemented with {C}.",
						Focuses:        []Focus{FocusSystems},
						TechTags:    []TechTag{TechTagC},
					}},
					{Subproject: Subproject{
						Title:       "A-normal Form (ANF)",
						Description: "There is a compiler phase that turns the abstract syntax tree to ANF, an intermediary representation that makes every intermediate computation named. This makes it easier to compile down to Assembly.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagDataStructures},
					}},
					{Subproject: Subproject{
						Title:       "Graph Colouring Register Allocation",
						Description: "Accessing a value in a register is faster than accessing it on the stack, so we want to keep values in registers for as long as possible. To do this, we track the liveness of each value — the range of the program over which it is needed — and use Chaitin's graph colouring algorithm to assign values to registers optimally",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagAlgorithms, TechTagDataStructures},
					}},
					{Subproject: Subproject{
						Title:       "Exceptions",
						Description: "This programming language allows users to throw and catch exceptions.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{},
					}},
					{Subproject: Subproject{
						Title:       "Built-in Testing Functionality",
						Description: "Normally, to run tests in a programming language, one has to use a library. This makes test organization awkward, and the library would have to use reflection magic to test for errors. In this programming language, tests are first-class citizens and can be woven into the code, and the code can be executed without the tests. The {C} run-time accumulates the test information and prints them out at the end.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagC},
					}},
				},
			},
		},
	},

	// ── UNIVERSITY: Typed-untyped interactions  ───────────────────
	{
		Title:       "Typed-untyped Interactions Through Machine",
		Description: "There are many way to implement a programming languages. One of the is through an abstract machine. On top of have all the basic features of a programming lanaguage (arithmetic, conditionals...), this programming language features the co-existence of typed and untyped modules. There are two distinct versions. One where the types are checked to the extent at possible, and then thrown away when the program actually runs. Another one where the types are checked, and become run-time checks during the program run.",
		Type:        ProjectTypeUniversity,
		Category:    catPtr(CategoryProgrammingLanguages),
		Focuses:        []Focus{FocusProgrammingLangs},
		TechTags:    []TechTag{TechTagTypedRacket},
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell:      []string{"Macros for tests, nice organization"},
				WhatCouldBeBetter: []string{"The macros could be better"},
				WhatILearned:      []string{"Making a language via a machine, logic, "},
			},
		},
		Subsections: []Subsection{
			{
				Title:    "CESK Machine",
				Focuses:     []Focus{},
				TechTags: []TechTag{TechTagLogic},
				Major: &MajorSubproject{
					Subproject: Subproject{
						Title:       "CESK Machine",
						Description: "Implemented a programming language via a CESK abstract machine — a formal model of computation where program execution is expressed as a sequence of discrete state transitions. Each state is a tuple of four components: the Control (the expression currently being evaluated), the Environment (bindings of variables to values), the Store (mutable memory), and the Kontinuation (the rest of the computation). Execution begins at a well-defined initial state and steps through intermediate states until it reaches a final one, giving the language a rigorous, mathematically grounded operational semantics.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{},
					},
				},
			},
			{
				Title:    "Notable Language Features",
				Focuses:     []Focus{},
				TechTags: []TechTag{},
				Cards: []Card{
					{Subproject: Subproject{
						Title:       "Classes & Modules",
						Description: "This language features classes, methods, and objects. On top of that, the language also features typed and untyped modules, and both can import each other. A module can only contain one class.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{},
					}},
					{Subproject: Subproject{
						Title:       "Tail Calls",
						Description: "This language has tail calls. That is, when a function call is the last thing that is being computed, a new stack will not be allocated.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{},
					}},
				},
			},
			{
				Title:    "Typed-untyped Interactions",
				Focuses:     []Focus{},
				TechTags: []TechTag{},
				Cards: []Card{
					{Subproject: Subproject{
						Title:       "Typed-untyped Interactions (JavaScript Style)",
						Description: "The {JavaScript} way of handling typed-untyped interactions is by running all type checks statically, and then strip away all the types prior to execution.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagJavaScript},
					}},
					{Subproject: Subproject{
						Title:       "Typed-untyped Interactions (Racket Style)",
						Description: "The Racket way of handling typed-untyped interactions is by running all type checks statically, and wrapping values that would be passed from typed modules to untyped modules in run-time checks for actual execution.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{},
					}},
					
				},
			},
			{
				Title:    "Unit and Integration Tests",
				Focuses:     []Focus{},
				TechTags: []TechTag{TechTagTesting},
				Bullets: []BulletPoint{
					{Subproject: Subproject{
						Title:       "Unit and Integration Testing, Featuring Macros",
						Description: "On top of unit tests for specific functions, there exist integration tests for specific phases of the language implementation. There are integration tests that targets the core logic of the CESK machine specifically, and through the use of {macros}, allows me to write the intermediate states of the abstract machine in a sequence, and the test runner would check that the state of machine goes through the sequence correctly.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagMacros},
					}},
					{Subproject: Subproject{
						Title:       "Programming-language-level Test Harness",
						Description: "An overarching test harness, a script written in Typed Racket, takes in entire programs written in the programming language, as well as a file that just contains the expected output, and checks if the actual output as produced by the implementation of the programming language equals to the expected output.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{},
					}},
				},
			},
		},
	},

	// ── UNIVERSITY: A Networked Card Game  ───────────────────
	{
		Title:       "A Networked Card Game",
		Description: "For the Software Development course at Northeastern University, also known as \"hell\", we implemented a networked card game in about 12 weeks.",
		Type:        ProjectTypeUniversity,
		Category:    catPtr(CategorySystems),
		Focuses:        []Focus{FocusFullStack},
		TechTags:    []TechTag{TechTagRacket},
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell:      []string{"Fairly good organization, network/logic interactions with that design pattern"},
				WhatCouldBeBetter: []string{"Brute force algorithm should have used streams instead of whole lists."},
				WhatILearned:      []string{"Software comes in layers.", "Got to be really meticulous when reading the specifications, take notes while reading."},
			},
		},
		Subsections: []Subsection{
			{
				Title:    "Notable Features",
				Focuses:     []Focus{},
				TechTags: []TechTag{},
				Cards: []Card{
					{Subproject: Subproject{
						Title:       "Client-server Interactions with Fault-tolerant Networking",
						Description: "Upon starting the server, it waits until the maximum number of clients connects or until the timer runs out, before starting the server. If the minimum number of players needed to play the game is not reached, the server terminates. Players can connect to the server. Once the game starts, the server is responsible for running all the game logic, listening to player actions, and informing the players of the game state and results through the network. Players crashing does not bring down the server.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagTCP},
					}},
					{Subproject: Subproject{
						Title:       "AI Players",
						Description: "For testing purposes, we have AI players that behave deterministically given some game state. The AI player uses a greedy algorithm to maximize some value in the short-term. The strategy pattern is used. This AI player is one such strategy, and in theory any other algorithms could be subbed in. For the released game, a strategy that communicates with the actual player and waits for user input would be used.",
						Focuses:        []Focus{FocusGameDev},
						TechTags:    []TechTag{TechTagAlgorithms, TechTagDesignPatterns},
					}},
					{Subproject: Subproject{
						Title:       "Remote Proxy Pattern",
						Description: "When interacting with players, the code on the host server goes through a logical interface. An implementation interface could make {RPC}s. However, for testing, it is useful for no RPCs to be made, and for the host server to run the logic immediately.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagRPC, TechTagDesignPatterns},
					}},
					{Subproject: Subproject{
						Title:       "JSON Messages De/serialization",
						Description: "The server and connecting players communicate with {JSON}, and so de/serialization code is needed. The would-be duplicate and boilerplate de/serialization code responsible are handled by macros, and also makes the code more present to look at.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagJSON, TechTagMacros},
					}},
				},
			},
			{
				Title:    "Testing",
				Focuses:     []Focus{},
				TechTags: []TechTag{TechTagTesting},
				Bullets: []BulletPoint{
					{Subproject: Subproject{
						Title:       "Unit Testing Game Logic",
						Description: "There are unit tests written for game logic.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{},
					}},
					{Subproject: Subproject{
						Title:       "Integration Testing Without Networks",
						Description: "There are integration tests that could simulate the game and (AI) players without involving the network, by using the aforementioned remote proxy design pattern.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{},
					}},
					{Subproject: Subproject{
						Title:       "Integration Testing With Networks",
						Description: "There are {Bash} scripts that can launch servers and clients, and the clients use the AI players, to test the game with networking without having to manually provide input.",
						Focuses:        []Focus{FocusDevOps},
						TechTags:    []TechTag{TechTagBash},
					}},
				},
			},
		},
	},

	// ── UNIVERSITY: RAFT  ───────────────────
	{
		Title:       "RAFT: Consensus Algorithm For Distributed Systems",
		Description: "Even when we are working on a single device, the digital services we use normally have several computers working together behind the scenes. But how could these computers work together when communication between them could be faulty and go down at any time? Consensus algorithms are a key feature to make communication possible, by allowing many computers to agree on a value. RAFT is one such algorithm for this, which I have implemented for my final project for the Distributed Systems course at Northeastern University. It is important to note that I was under a very strict time constrain.",
		Type:        ProjectTypeUniversity,
		Category:    catPtr(CategorySystems),
		Focuses:        []Focus{FocusSystems, FocusFullStack},
		TechTags:    []TechTag{TechTagCPP},
		SourceCode: &SourceCode{OnRequest: true},
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell:      []string{"The custom allocator built for the log structure makes it memory efficient. A lot of the memory are contiguous. There is no fragmentation.", "As long as one node lives, the log structure will continue to live."},
				WhatCouldBeBetter: []string{"Leader will block for response after sending RPC for a period of time. If the response misses the window, it will be dropped.", "The logs are not very clean. When a server is shut down, all the other servers will flood the logs saying that they cannot connect to the server.", "There isn't a clear clean separation between the logical RAFT layer and the server layer. This makes no difference to the user, but isn't good for maintenance.", "The printing node data looks ugly. Invalid log index is long string of digits instead of saying invalid log index."},
				WhatILearned:      []string{"Take advantage of leeways in the specifications."},
			},
		},
		Subsections: []Subsection{
			{
				Title:    "Subprojects",
				Focuses:     []Focus{},
				TechTags: []TechTag{},
				Cards: []Card{
					{Subproject: Subproject{
						Title:       "Client-server Interactions with Fault-tolerance",
						Description: "The servers are nodes of the RAFT algorithm. Clients are those who wish to submit commands for the RAFT algorithm to persist. The networking interactions are done through the {Unix socket library}. Any of the servers can crash (and that is assumed to be the only mode of failure), without crashing the other servers. As long as majority of the nodes are alive, the system can continue to make progress. If majority of the nodes are no longer alive, the data will persist as one node is alive, but the system can no longer agree on new values.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagLinux, TechTagTCP, TechTagUnixSocketLibrary},
					}},
					{Subproject: Subproject{
						Title:       "Leader Election",
						Description: "A key element of the algorithm is leader election. A server by itself is just a hunk of metal capable of computation, but RAFT consensus algorithm associates each server node with a role. The two essentials roles are leader and follower. There can only at most be one leader. Followers are able to be promoted to leader, which happens in the case where the system just started and there are no leaders to begin with, or the leader crashes. How this happens is defined in the RAFT specifications, specifically the \"Vote Request\" {RPC}, which I implemented. A follower that wants to be a leader sends the request to all the other nodes, and if it receives yes from the majority, it becomes the leader.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagAlgorithms, TechTagRPC},
					}},
					{Subproject: Subproject{
						Title:       "Log Replication With Custom Memory Slab Allocator",
						Description: "Another key element is that each node keeps track of the full history of commands, in a log data structure. What it means for nodes to be in sync is for their logs to be the same. The leader's log is replicated to the followers. My implementation of the log data structure uses my own memory allocator that allocates in memory-aligned chunks, that are a multiple of page sizes. The allocator holds a linked list of the chunks, and a chunk holds tightly-packed log entries.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagDataStructures},
					}},
				},
			},
		},
	},

	// ── PERSONAL: Skyline Engine  ─────────────────────────────────
	{
		Title:       "A Custom Game Engine",
		Description: "C++ 20 (even though the code looks like C),- SDL3 for the platform layer, - Vulkan for hardware-accelerated graphics,- Jolt for physics (the physics engine developed for the second Horizon Zero Dawn game), - Dear Imgui for UI for internal tooling (though we will probably that for the actual game's UI),- CMake for the build system, - We require GCC or Clang for the C++ compiler, with compiler extensions enabled",
		Type:        ProjectTypePersonal,
		Category:    catPtr(CategorySystems),
		Focuses:        []Focus{FocusGameEngineDev},
		TechTags:    []TechTag{TechTagCPP},
		SourceCode: &SourceCode{Link: strPtr("https://github.com/melodicht/skyline-engine")},
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell:      []string{""},
				WhatCouldBeBetter: []string{""},
				WhatILearned:      []string{""},
			},
		},
		Subsections: []Subsection{
			{
				Title:    "Subprojects",
				Focuses:     []Focus{},
				TechTags: []TechTag{},
				Cards: []Card{
					{Subproject: Subproject{
						Title:       "Hot Reloading",
						Description: "The user is able to keep the game running, make a change to the game source code, recompile the game, and immediately see the new changes take place in the running game.",
						Focuses:        []Focus{FocusSystems},
						TechTags:    []TechTag{},
					}},
					{Subproject: Subproject{
						Title:       "Looped Live Playback & Input Streaming",
						Description: "The user is able record a segment of gameplay and loop it for as long as they want. Any inputs recorded in the loop will be played back. This goes well with hot reloading. The user can make a change in the game source, and immediately see the changes when the loop restarts.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{},
					}},
					{Subproject: Subproject{
						Title:       "Scene Editor",
						Description: "I worked on using the Dear {IMGUI} library to display the ECS menu, and the UI to change the fields of an entity, and various buttons to do things like add/destroy entities.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagIMGUI},
					}},
					{Subproject: Subproject{
						Title:       "Engine Architecture",
						Description: "I was inspired by Handmade Hero when establishing much of the engine's {architecture}. There are four major software components: renderer, platform, engine and game. The renderer encapsulates all the messiness of displaying things on the screen and interactions with the GPU. The platform encapsulates the messiness of operating system. There is a strong relationship between the platform and the renderer. The engine is the entry point to the game-side of things, and calls into the game code. It holds code that all the game uses. The game-side of things is for code that is specific to one game. Note that the engine and game modules do not interact with the renderer directly.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagArchitecture},
					}},
				},
			},
		},
	},

	// ── PERSONAL: My Personal Portfolio  ──────────────────────────
	{
		Title:       "My Personal Portfolio (This Website)",
		Description: "The website is designed to be able to be read by hiring managers efficiently from different fields, game development, game engine development, full-stack development and so on. The [I do] mode allows hiring managers to focus on seeing what content that's relevant to them only. The other modes provide all data unfiltered.",
		Type:        ProjectTypePersonal,
		Category:    catPtr(CategorySystems),
		Focuses:        []Focus{FocusFullStack},
		TechTags:    []TechTag{},
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell:      []string{},
				WhatCouldBeBetter: []string{},
				WhatILearned:      []string{"Using Datastar for the first time"},
			},
		},
		Subsections: []Subsection{},
	},

	// ── UNIVERSITY: Dreams of Celestial Pull  ──────────────────────
	{
		Title:       "Dreams of Celestial Pull: A Physics-based First-person Shooter Platformer",
		Description: "For Game Capstone, the final games course at Northeastern University where you spend an entire semester developing a game, I single-handedly developed Dreams of Celestial Pull. The game is made with the custom game engine, Skyline Engine. If you are into game design, I recommend giving the game a try first before reading the below, as there will be game design spoilers.",
		Type:        ProjectTypeUniversity,
		Category:    catPtr(CategoryGames),
		Focuses:        []Focus{FocusGameDev},
		TechTags:    []TechTag{TechTagCPP},
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell:      []string{"Lots of content from simple mechanics, gravity ball allows for lots of different consequences."},
				WhatCouldBeBetter: []string{"Due to limitations of the graphics-portion of the engine, a lot of visual feedback that really needed to be there weren't. It would have gave players intuition about the nature of gravity balls.", "There is also no sound."},
				WhatILearned:      []string{"Before you can walk, you must crawl. I thought the game was going to have enemies that could shoot bullets, but I decided to focus on the mechanics even without enemies. With what I know now, I know what kind of enemies to design for that would optimize for how gravity balls work."},
			},
		},
		Subsections: []Subsection{
			{
				Title:    "FPS Movement & RK4",
				Focuses:     []Focus{},
				TechTags: []TechTag{},
				Major: &MajorSubproject{
					Subproject: Subproject{
						Title:       "FPS Movement & RK4",
						Description: "This game features an interesting marriage between two contrasting notions of {physics} simulation. FPS movement includes on-ground movement and jumping. On-ground movement has acceleration with max velocity. As the game is a platformer, the player can cut a jump short by letting go of jump fast rather than holding down the button. There is also air control, where acceleration in the air is lower. The other notion of physics simulation is RK4, or fourth-order Runge-kutta method. Given a position and velocity of an object, and a function that computes the object's acceleration, the method can produce the next object's position and velocity. The method uses calculus effectively in order to get good approximations with minimal additional computations. This is not the standard approach to movement physics. I used it in my game anyway because the movements physics is tremendously better and the game still runs smoothly. These two notions of movement physics don't naturally work together. The FPS-specific physics has a lot of edge cases and handles player input, whereas RK4 is a general purpose function and does not consider player input. Still, I made it work.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagPhysics, TechTagAlgorithms, TechTag3D},
					},
				},
			},
		},
	},

	// ── UNIVERSITY: Boids with Goals  ──────────────────────
	{
		Title:       "Boids with Goals: A Game AI Project",
		Description: "XXX studied how a flock of birds move in the air and found a simple algorithm that seem to mimic how they actually do it in real-life. However, the algorithm itself doesn't consider boids interacting with obstructions, and also boids flying towards a goal. So, this projects takes the idea one step further to create boids that do.",
		Type:        ProjectTypeUniversity,
		Category:    catPtr(CategoryGames),
		Focuses:        []Focus{FocusGameDev},
		TechTags:    []TechTag{TechTagJavaScript, TechTagPhysics, TechTagAlgorithms, TechTag2D},
		SourceCode: &SourceCode{Link: strPtr("https://github.com/melodicht/boids")},
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell:      []string{},
				WhatCouldBeBetter: []string{"Boids like to fly towards a wall, and at the very last minute make a sharp turn. This isn't very realistic.", "Boids ocassionally fly in circles far from the goals. likely caused by counter-balancing forces between the the structure of the walls and the goals.", "There is no way to level-edit while the program is running. The placement of walls are hard-coded.", "Boids sometimes move in a staircase pattern, likely caused by the fact that the path-finding only considers cardinal directios and doesn't smooth out the paths. "},
				WhatILearned:      []string{"The sum of forces is a powerful idea."},
			},
		},
		Subsections: []Subsection{
			{
				Title:    "Boids Movement with Multi-layered Dijkstra Path-finding",
				Focuses:     []Focus{},
				TechTags: []TechTag{},
				Major: &MajorSubproject{
					Subproject: Subproject{
						Title:       "Boids Movement with Multi-layered Dijkstra Path-finding",
						Description: "A boid's movement is based on local conditions close to its proximity, like distance and alignment to neighbouring boids, as well as obstacles to not fly into. However, goal-oriented behaviour requires the boid to navigate to a point which may exist beyond said proximity. With obstacles, there needs to exist some intelligence to not fly blindly into a dead end. While the twin goals of navigation by local and global conditions conflict, my approach uses multi-layered Dijkstra's algorithm to create a field of forces, generating local conditions from global ones. Where a boid moves is the sum of its regular boid's movement and the vector at the position of the map as produced by Dikstra's algorithm.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{},
					},
				},
			},
			{
				Title:    "Collision Detection",
				Focuses:     []Focus{},
				TechTags: []TechTag{},
				Major: &MajorSubproject{
					Subproject: Subproject{
						Title:       "Collision Detection",
						Description: "Each boid determines whether it's about to fly into an obstacle by shooting a ray in front of it and seeing if it intersects with an obstacle. All obstacles are rectangles.",
						Focuses:        []Focus{FocusGameEngineDev},
						TechTags:    []TechTag{},
					},
				},
			},
		},
	},

	// ── PERSONAL: Toxic Texting  ──────────────────────
	{
		Title:       "Toxic Texting: A Chill Texting Game Where You Only Respond with Yes or No",
		Description: "Made initially for the Summer 2021 NEU Game Development Club 48-hour game jam, Toxic Texting is a fun, short and sweet texting game where you respond with either yes or no. It is made in Unity 2D, and I was one of the two programmers, the composer, sounnd designer and writer.",
		Type:        ProjectTypePersonal,
		Category:    catPtr(CategoryGames),
		Focuses:        []Focus{FocusGameDev},
		TechTags:    []TechTag{TechTagCSharp, TechTagUnity, TechTag2D},
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell:      []string{"The user experience is on point."},
				WhatCouldBeBetter: []string{"The dialogue system is a pain."},
				WhatILearned:      []string{"Take advantage of Unity's UI system as much as possible, so that there's less work on my part."},
			},
		},
		Subsections: []Subsection{
			{
				Title:    "Subprojects",
				Focuses:     []Focus{},
				TechTags: []TechTag{},
				Cards: []Card{
					{Subproject: Subproject{
						Title:       "Dialogue & Game Systems",
						Description: "The writer writes all the dialogue in the spreadsheet, which is then imported into the game in CSV format. The game parses the CSV file and constructs the dialogue tree. A level can be understood as the navigation of the tree, where the user's yes/no moves the player from one dialogue to the next. A writer can also write game events into the dialogue, like music change or screen shake, and the dialogue system will invoke the corresponding methods in the corresponding systems.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagDataStructures},
					}},
					{Subproject: Subproject{
						Title:       "UI/UX",
						Description: "Using Unity's UI system and with the help of the artist on our team, I implemented and partially designed the phone UI and animations.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagUIUX},
					}},
					{Subproject: Subproject{
						Title:       "Sound & Music",
						Description: "On top of making all the sound effects and music in Logic Pro X (except for one song that is creditted in the game), I also hooked them into the game. A lot of it is through Unity's event system. The writer can also make sound cueues in the dialogue. Upon seeing a sound queue, the dialogue system tells the sound system to play sounds. When the music switches, there is a crossfade.",
						Focuses:        []Focus{},
						TechTags:    []TechTag{TechTagSoundMusic},
					}},
				},
			},
		},
	},

	// ── PERSONAL: Tower Takeover  ──────────────────────
	{
		Title:       "Tower Takeover: A vanilla JavaScript Tower Defense Game That Runs on your Web Browser",
		Description: "This is a game that I led for formerly Hometeam Game Dev (now DevPods). Hometeam Game Dev is a community of game developers that makes games without pay. I was responsible for structuring tasks on a kanban board, running playtests, and making sure that the game ships.",
		Type:        ProjectTypePersonal,
		Category:    catPtr(CategoryGames),
		Focuses:        []Focus{FocusGameDev, FocusWebDev},
		TechTags:    []TechTag{TechTagJavaScript},
		Specifics: ProjectTypeSpecifics{
			NonJob: &NonJobExperience{
				WhatWentWell:      []string{"Gameplay system is easy to work with."},
				WhatCouldBeBetter: []string{"The game is too hard, and without written instructions, it's hard to figure out how to play the game."},
				WhatILearned:      []string{"This game has a lot of information and actions. I should have designed the game around that, instead of the raw systems. Start with what the user experience, and not the underlying physics of the game."},
			},
		},
		Subsections: []Subsection{},
	},
}

func strPtr(s string) *string                   { return &s }
func imgPtr(i Image) *Image                     { return &i }
func vidPtr(v Video) *Video                     { return &v }
func catPtr(c ProjectCategory) *ProjectCategory { return &c }
