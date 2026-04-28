package sitedata

// Tag is a skill/technology tag.
type Tag string

const (
	TagGameEngineDev    Tag = "game engine development"
	TagGameDev          Tag = "game development"
	TagFullStack        Tag = "full-stack development"
	TagSystems          Tag = "systems programming"
	TagProgrammingLangs Tag = "programming languages"
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
	WhatWentWell     string  `json:"whatWentWell"`
	WhatCouldBeBetter string `json:"whatCouldBeBetter"`
	WhatILearned     string  `json:"whatILearned"`
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
