package main

// Focus is a skill/domain focus tag used for filtering on the [I do] page.
type Focus string

const (
	FocusGameEngineDev    Focus = "game engine development"
	FocusGameDev          Focus = "game development"
	FocusFullStack        Focus = "full-stack development"
	FocusWebDev           Focus = "web development"
	FocusMobileAppDev     Focus = "mobile app development"
	FocusSystems          Focus = "systems programming"
	FocusProgrammingLangs Focus = "programming languages"
	FocusSystemSecurity   Focus = "systems security"
	FocusDevOps           Focus = "development operations"
	FocusHardwareTech     Focus = "hardware tech"
)

type TechTag string

const (
	TechTagPython     TechTag = "Python"
	TechTagNumPy      TechTag = "NumPy 2"
	TechTagMatplotlib TechTag = "Matplotlib"
	TechTagDatadog    TechTag = "Datadog"
	TechTagFlask      TechTag = "Flask"
	TechTagSphinx     TechTag = "Sphinx"
	TechTagTox        TechTag = "Tox"
	TechTagCoveragePy TechTag = "Coverage.py"
	TechTagPip        TechTag = "Pip"

	TechTagHTML               TechTag = "HTML"
	TechTagXML                TechTag = "XML"
	TechTagJavaScript         TechTag = "JavaScript"
	TechTagTypeScript         TechTag = "TypeScript"
	TechTagGraphQueryLanguage TechTag = "Graph Query Language"
	TechTagReact              TechTag = "React"
	TechTagNodeJS             TechTag = "Node.js"
	TechTagJSON               TechTag = "JSON"

	TechTagStorybook TechTag = "Storybook"
	TechTagWebpack   TechTag = "Webpack"
	TechTagVite      TechTag = "Vite"

	TechTagLinux             TechTag = "Linux"
	TechTagTrafficControl    TechTag = "Traffic Control (tc)"
	TechTagUnixSocketLibrary TechTag = "Unix Socket Library"

	TechTagBash TechTag = "Bash"

	TechTag2D TechTag = "2D"
	TechTag3D TechTag = "3D"

	TechTagAlgorithms          TechTag = "Algorithms"
	TechTagDataStructures      TechTag = "Data Structures"
	TechTagLogic               TechTag = "Logic"
	TechTagDesignPatterns      TechTag = "Design Patterns"
	TechTagArchitecture        TechTag = "Architecture"
	TechTagAsync               TechTag = "Async"
	TechTagSocketIO            TechTag = "SocketIO"
	TechTagComponentsLibrary   TechTag = "Components Library"
	TechTagServersideRendering TechTag = "Server-side Rendering"
	TechTagProfiling           TechTag = "Profiling"
	TechTagFFI                 TechTag = "FFI"
	TechTagETL                 TechTag = "ETL Architecture"
	TechTagTCP                 TechTag = "TCP"
	TechTagRPC                 TechTag = "RPC"
	TechTagMacros              TechTag = "Macros"
	TechTagTesting             TechTag = "Testing"
	TechTagPhysics             TechTag = "Physics"

	TechTagOpenGraph TechTag = "Open Graph"
	TechTagSiteMap   TechTag = "Sitemap"
	TechTagWebP      TechTag = "WebP"

	TechTagFlutter TechTag = "Flutter"
	TechTagDart    TechTag = "Dart"

	TechTagKotlin             TechTag = "Kotlin"
	TechTagAndroidDevelopment TechTag = "Android Development"
	TechTagJetpackCompose     TechTag = "Jetpack Compose"
	TechTagTimber             TechTag = "Timber"

	TechTagSwift TechTag = "Swift"

	TechTagC        TechTag = "C"
	TechTagCPP      TechTag = "C++"
	TechTagAssembly TechTag = "Assembly"

	TechTagOCaml TechTag = "OCaml"

	TechTagPuppeteer TechTag = "Puppeteer"
	TechTagIMGUI     TechTag = "IMGUI"
	TechTagUIUX      TechTag = "UI/UX"

	TechTagFirestore               TechTag = "Firestore"
	TechTagSlackAPI                TechTag = "Slack API"
	TechTagGoogleCloudRunFunctions TechTag = "Google Cloud Run Functions"
	TechTagGoogleCloudScheduler    TechTag = "Google Cloud Scheduler"
	TechTagGoogleCloudStorage      TechTag = "Google Cloud Storage"

	TechTagTypedRacket TechTag = "Typed Racket"
	TechTagRacket      TechTag = "Racket"

	TechTagCSharp TechTag = "C#"
	TechTagUnity  TechTag = "Unity"

	TechTagSTRIDE     TechTag = "STRIDE Threat Modelling"
	TechTagSoundMusic TechTag = "SFX/Music"
)

// ProjectCategory groups projects on the [I worked on] page.
type ProjectCategory string

const (
	CategoryProgrammingLanguages ProjectCategory = "Programming Languages"
	CategoryGames                ProjectCategory = "Games"
	CategorySystems              ProjectCategory = "Low-level/Distributed/Full-stack Systems"
)

// ProjectType distinguishes job, university, and personal projects.
type ProjectType string

const (
	ProjectTypeJob        ProjectType = "Job"
	ProjectTypeUniversity ProjectType = "University"
	ProjectTypePersonal   ProjectType = "Personal"
)

// Image and Video are paths to static assets or external URLs.
type Image string
type Video string

// DateRange represents a start date and an optional end date (nil = present).
type DateRange struct {
	Start string  `json:"start"`
	End   *string `json:"end"`
}

// Review is a LinkedIn-style recommendation.
type Review struct {
	ProfilePicture Image  `json:"profilePicture"`
	Name           string `json:"name"`
	Role           string `json:"role"`
	Text           string `json:"text"`
}

// JobExperience holds data specific to a job project.
type JobExperience struct {
	Company         string    `json:"company"`
	Role            string    `json:"role"`
	BackgroundImage Image     `json:"backgroundImage"`
	PortraitImage   *Image    `json:"portraitImage,omitempty"`
	DateRange       DateRange `json:"dateRange"`
	Website         *string   `json:"website"`
	Reviews         []Review  `json:"reviews"`
}

// NonJobExperience holds reflection data for university and personal projects.
type NonJobExperience struct {
}

// SourceCode represents the availability of source code.
// Exactly one of Link or OnRequest should be set.
type SourceCode struct {
	Link      *string `json:"link,omitempty"`
	OnRequest bool    `json:"onRequest,omitempty"`
}

// ProjectTypeSpecifics is a union — exactly one field is non-nil.
type ProjectTypeSpecifics struct {
	Job    *JobExperience    `json:"job,omitempty"`
	NonJob *NonJobExperience `json:"nonJob,omitempty"`
}

// Subproject is the shared content unit used by BulletPoint, Card, and MajorSubproject.
type Subproject struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Focuses     []Focus     `json:"focuses"`
	TechTags    []TechTag   `json:"techTags"`
	SourceCode  *SourceCode `json:"sourceCode,omitempty"`
}

// BulletPoint renders its subproject as a prose bullet line.
type BulletPoint struct {
	Subproject Subproject `json:"subproject"`
}

// Card renders its subproject as a card.
type Card struct {
	Subproject Subproject `json:"subproject"`
}

// MajorSubproject renders its subproject full-width with optional video.
type MajorSubproject struct {
	Subproject Subproject `json:"subproject"`
	Video      *Video     `json:"video,omitempty"`
}

// Subsection groups related items under a heading within a project.
// Exactly one of Bullets, Cards, or Major should be populated.
type Subsection struct {
	Title      string           `json:"title"`
	Focuses    []Focus          `json:"focuses"`
	TechTags   []TechTag        `json:"techTags"`
	SourceCode *SourceCode      `json:"sourceCode,omitempty"`
	Bullets    []BulletPoint    `json:"bullets,omitempty"`
	Cards      []Card           `json:"cards,omitempty"`
	Major      *MajorSubproject `json:"major,omitempty"`
}

// Project is the top-level portfolio entry.
// Focuses and TechTags are inherited by all subsections and their subprojects.
type Project struct {
	Title       string               `json:"title"`
	URLSlug     string               `json:"urlSlug,omitempty"`
	Description string               `json:"description"`
	Type        ProjectType          `json:"type"`
	Specifics   ProjectTypeSpecifics `json:"specifics"`
	Category    *ProjectCategory     `json:"category,omitempty"`
	Focuses     []Focus              `json:"focuses"`
	TechTags    []TechTag            `json:"techTags"`
	SourceCode  *SourceCode          `json:"sourceCode,omitempty"`
	Subsections []Subsection         `json:"subsections"`
}

// SiteData is the root structure passed to all templates.
type SiteData struct {
	Projects []Project `json:"projects"`
}
