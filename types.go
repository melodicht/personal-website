package main

// Tag is a skill/technology tag.
type Tag string

const (
	TagGameEngineDev    Tag = "game engine development"
	TagGameDev          Tag = "game development"
	TagFullStack        Tag = "full-stack development"
	TagWebDev           Tag = "web development"
	TagMobileAppDev     Tag = "Mobile App Development"
	TagSystems          Tag = "systems programming"
	TagProgrammingLangs Tag = "programming languages"
	TagSystemSecurity   Tag = "systems security"
	TagDevOps           Tag = "development operations"
	TagHardwareTech     Tag = "hardware tech"
)

type TechTag string

const (
	TechTagPython TechTag = "Python"
	TechTagNumPy TechTag = "NumPy"
	TechTagMatplotlib TechTag = "Matplotlib"
	TechTagDatadog TechTag = "Datadog"
	TechTagFlask TechTag = "Flask"
	TechTagSphinx TechTag = "Sphinx"
	TechTagTox TechTag = "Tox"
	TechTagCoveragePy TechTag = "Coverage.py"
	TechTagPip TechTag = "Pip"

	TechTagHTML TechTag = "HTML"
	TechTagXML TechTag = "XML"
	TechTagJavaScript TechTag = "JavaScript"
	TechTagTypeScript TechTag = "TypeScript"
	TechTagGraphQueryLanguage TechTag = "Graph Query Language (GQL)"
	TechTagReact TechTag = "React"

	TechTagStorybook TechTag = "Storybook"
	TechTagWebpack TechTag = "Webpack"
	TechTagVite TechTag = "Vite"

	TechTagLinux TechTag = "Linux"
	TechTagTrafficControl TechTag = "Traffic Control (tc)"
	
	TechTagBash TechTag = "Bash"
	
	TechTagAsync TechTag = "Async"
	TechTagSocketIO TechTag = "SocketIO"
	TechTagComponentsLibrary TechTag = "Components Library"
	TechTagServersideRendering TechTag = "Server-side Rendering"
	TechTagProfiling TechTag = "Profiling"
	TechTagFFI TechTag = "FFI"
	TechTagETL TechTag = "ETL Architecture"

	TechTagOpenGraph TechTag = "Open Graph"
	TechTagSiteMap TechTag = "Sitemap"
	TechTagWebP TechTag = "WebP"

	TechTagFlutter TechTag = "Flutter"
	TechTagDart TechTag = "Dart"

	TechTagKotlin TechTag = "Kotlin"
	TechTagAndroidDevelopment TechTag = "Android Development"
	TechTagJetpackCompose TechTag = "Jetpack Compose"
	TechTagTimber TechTag = "Timber"

	TechTagSwift TechTag = "Swift"

	TechTagC TechTag = "C"
	TechTagCPP TechTag = "C++"
	TechTagAssembly TechTag = "Assembly"

	TechTagPuppeteer TechTag = "Puppeteer"


	TechTagFirestore TechTag = "Firestore"
	TechTagSlackAPI TechTag = "Slack API"

	TechTagGoogleCloudRunFunctions TechTag = "Google Cloud Run Functions"
	TechTagGoogleCloudScheduler TechTag = "Google Cloud Scheduler"

	


	TechTagSTRIDE TechTag = "STRIDE Thread Modelling"
)

// ProjectCategory groups projects on the "I worked on" page.
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
	Reviews         []Review  `json:"reviews"`
}

// NonJobExperience holds reflection data for university and personal projects.
type NonJobExperience struct {
	Video            *Video  `json:"video,omitempty"`
	WhatWentWell     []string  `json:"whatWentWell"`
	WhatCouldBeBetter []string `json:"whatCouldBeBetter"`
	WhatILearned     []string  `json:"whatILearned"`
	SourceCodeLink   *string `json:"sourceCodeLink,omitempty"`
}

// ProjectTypeSpecifics is a union — exactly one field is non-nil.
type ProjectTypeSpecifics struct {
	Job    *JobExperience    `json:"job,omitempty"`
	NonJob *NonJobExperience `json:"nonJob,omitempty"`
}

// SubprojectInfo merges SubprojectType and SubprojectSpecific.
// If Video is nil, the subproject is Small; otherwise it is Big.
type SubprojectInfo struct {
	Video *Video `json:"video,omitempty"`
}

// Subproject is a unit of work within a Project.
type Subproject struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Tags        []Tag          `json:"tags"`
	TechTags    []TechTag      `json:"techTags"`
	Info        SubprojectInfo `json:"info"`
}

// Project is the top-level portfolio entry.
// Tags is omitted from the Go definition and computed by the generator.
type Project struct {
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Type        ProjectType          `json:"type"`
	Specifics   ProjectTypeSpecifics `json:"specifics"`
	Category    *ProjectCategory     `json:"category,omitempty"`
	Tags        []Tag                `json:"tags"` // computed by generator
	Subprojects []Subproject         `json:"subprojects"`
}

// SiteData is the root JSON structure embedded into the page.
type SiteData struct {
	Projects []Project `json:"projects"`
}
