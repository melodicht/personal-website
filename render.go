package main

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

// ── Description parsing ───────────────────────────────────────────────────────

var inlineTagRe = regexp.MustCompile(`\{([^}]+)\}`)

// DescriptionSegment is either a plain text run or an inline tech tag chip.
type DescriptionSegment struct {
	IsChip bool
	Value  template.HTML // already HTML-escaped, safe to output directly
}

func parseDescription(desc string, knownTechTags []TechTag) []DescriptionSegment {
	var segments []DescriptionSegment
	last := 0
	matches := inlineTagRe.FindAllStringIndex(desc, -1)
	for _, loc := range matches {
		if loc[0] > last {
			segments = append(segments, DescriptionSegment{
				Value: template.HTML(template.HTMLEscapeString(desc[last:loc[0]])),
			})
		}
		raw := desc[loc[0]+1 : loc[1]-1] // strip braces
		canonical := raw
		for _, tt := range knownTechTags {
			if strings.EqualFold(string(tt), raw) {
				canonical = string(tt)
				break
			}
		}
		segments = append(segments, DescriptionSegment{
			IsChip: true,
			Value:  template.HTML(template.HTMLEscapeString(canonical)),
		})
		last = loc[1]
	}
	if last < len(desc) {
		segments = append(segments, DescriptionSegment{
			Value: template.HTML(template.HTMLEscapeString(desc[last:])),
		})
	}
	return segments
}

// ── Rendered types ────────────────────────────────────────────────────────────

// RenderedSubproject is the template-ready version of Subproject.
// Description has been parsed into DescriptionSegments; all other fields are copied.
type RenderedSubproject struct {
	Title             string               `json:"title"`
	ParsedDescription []DescriptionSegment `json:"parsedDescription"`
	Focuses           []Focus              `json:"focuses"`
	TechTags          []TechTag            `json:"techTags"`
	SourceCode        *SourceCode          `json:"sourceCode,omitempty"`
}

// RenderSubproject transforms an authoring Subproject into a RenderedSubproject,
// parsing the description against the known tech tags for canonical chip values.
func RenderSubproject(sp Subproject, knownTechTags []TechTag) RenderedSubproject {
	return RenderedSubproject{
		Title:             sp.Title,
		ParsedDescription: parseDescription(sp.Description, knownTechTags),
		Focuses:           sp.Focuses,
		TechTags:          sp.TechTags,
		SourceCode:        sp.SourceCode,
	}
}

// ── Inlined tag helpers ───────────────────────────────────────────────────────

// inlinedTechTagsFromSegments returns the lowercased set of chip values from
// an already-parsed description, used to avoid duplicating them in the chip row.
func inlinedTechTagsFromSegments(segs []DescriptionSegment) map[string]bool {
	out := map[string]bool{}
	for _, seg := range segs {
		if seg.IsChip {
			out[strings.ToLower(string(seg.Value))] = true
		}
	}
	return out
}

// ── Tag inheritance helpers ───────────────────────────────────────────────────

// MergeTechTags merges two slices, deduped, with a taking priority.
func MergeTechTags(a, b []TechTag) []TechTag {
	seen := map[TechTag]bool{}
	var out []TechTag
	for _, t := range a {
		if !seen[t] {
			seen[t] = true
			out = append(out, t)
		}
	}
	for _, t := range b {
		if !seen[t] {
			seen[t] = true
			out = append(out, t)
		}
	}
	return out
}

// MergeFocuses merges two Focus slices, deduped.
func MergeFocuses(a, b []Focus) []Focus {
	seen := map[Focus]bool{}
	var out []Focus
	for _, f := range a {
		if !seen[f] {
			seen[f] = true
			out = append(out, f)
		}
	}
	for _, f := range b {
		if !seen[f] {
			seen[f] = true
			out = append(out, f)
		}
	}
	return out
}

// OwnTechTags returns the subproject's own non-inherited, non-inlined tech tags.
func OwnTechTags(sp RenderedSubproject, inherited []TechTag) []TechTag {
	inlined := inlinedTechTagsFromSegments(sp.ParsedDescription)
	inheritedSet := map[TechTag]bool{}
	for _, t := range inherited {
		inheritedSet[t] = true
	}
	var out []TechTag
	for _, t := range sp.TechTags {
		if !inheritedSet[t] && !inlined[strings.ToLower(string(t))] {
			out = append(out, t)
		}
	}
	return out
}

// EffectiveTechTags returns the full merged tech tag list for a subproject,
// excluding any that are inlined in the description.
func EffectiveTechTags(sp RenderedSubproject, inherited []TechTag) []TechTag {
	merged := MergeTechTags(inherited, sp.TechTags)
	inlined := inlinedTechTagsFromSegments(sp.ParsedDescription)
	var out []TechTag
	for _, t := range merged {
		if !inlined[strings.ToLower(string(t))] {
			out = append(out, t)
		}
	}
	return out
}

// ── Rendered container types ──────────────────────────────────────────────────
// These mirror the authoring types but with Subproject replaced by RenderedSubproject.

type RenderedBullet struct {
	Subproject RenderedSubproject
}

type RenderedCard struct {
	Subproject RenderedSubproject
}

type RenderedMajor struct {
	Subproject RenderedSubproject
	Video      *Video
}

type RenderedSubsection struct {
	Title      string
	Focuses    []Focus
	TechTags   []TechTag
	SourceCode *SourceCode
	Bullets    []RenderedBullet
	Cards      []RenderedCard
	Major      *RenderedMajor
}

type RenderedProject struct {
	Title       string
	URLSlug     string
	Description string
	Type        ProjectType
	Specifics   ProjectTypeSpecifics
	Category    *ProjectCategory
	Focuses     []Focus
	TechTags    []TechTag
	SourceCode  *SourceCode
	Subsections []RenderedSubsection
}

// RenderProject transforms an authoring Project into a RenderedProject,
// pre-parsing all subproject descriptions against the inherited tech tag context.
func RenderProject(p Project) RenderedProject {
	var secs []RenderedSubsection
	for _, sec := range p.Subsections {
		inherited := MergeTechTags(p.TechTags, sec.TechTags)
		var bullets []RenderedBullet
		for _, b := range sec.Bullets {
			bullets = append(bullets, RenderedBullet{
				Subproject: RenderSubproject(b.Subproject, MergeTechTags(inherited, b.Subproject.TechTags)),
			})
		}
		var cards []RenderedCard
		for _, c := range sec.Cards {
			cards = append(cards, RenderedCard{
				Subproject: RenderSubproject(c.Subproject, MergeTechTags(inherited, c.Subproject.TechTags)),
			})
		}
		var major *RenderedMajor
		if sec.Major != nil {
			rs := RenderSubproject(sec.Major.Subproject, MergeTechTags(inherited, sec.Major.Subproject.TechTags))
			major = &RenderedMajor{Subproject: rs, Video: sec.Major.Video}
		}
		secs = append(secs, RenderedSubsection{
			Title:      sec.Title,
			Focuses:    sec.Focuses,
			TechTags:   sec.TechTags,
			SourceCode: sec.SourceCode,
			Bullets:    bullets,
			Cards:      cards,
			Major:      major,
		})
	}
	return RenderedProject{
		Title:       p.Title,
		URLSlug:     ProjectSlug(p),
		Description: p.Description,
		Type:        p.Type,
		Specifics:   p.Specifics,
		Category:    p.Category,
		Focuses:     p.Focuses,
		TechTags:    p.TechTags,
		SourceCode:  p.SourceCode,
		Subsections: secs,
	}
}

// RenderProjects transforms a slice of Projects into RenderedProjects.
func RenderProjects(ps []Project) []RenderedProject {
	out := make([]RenderedProject, len(ps))
	for i, p := range ps {
		out[i] = RenderProject(p)
	}
	return out
}

// ── Focus matching ────────────────────────────────────────────────────────────

// HasFocus reports whether a subproject has the given focus.
func HasFocus(sp RenderedSubproject, spFocuses, projFocuses []Focus, f Focus) bool {
	effective := MergeFocuses(projFocuses, spFocuses)
	for _, ef := range effective {
		if ef == f {
			return true
		}
	}
	return false
}

// ── Flat subproject list ──────────────────────────────────────────────────────

// AllFocuses returns the deduplicated set of all focuses used across all projects.
func AllFocuses(projects []Project) []Focus {
	seen := map[Focus]bool{}
	var out []Focus
	add := func(f Focus) {
		if !seen[f] {
			seen[f] = true
			out = append(out, f)
		}
	}
	for _, p := range projects {
		for _, f := range p.Focuses {
			add(f)
		}
		for _, sec := range p.Subsections {
			for _, f := range sec.Focuses {
				add(f)
			}
			var sps []Subproject
			for _, b := range sec.Bullets {
				sps = append(sps, b.Subproject)
			}
			for _, c := range sec.Cards {
				sps = append(sps, c.Subproject)
			}
			if sec.Major != nil {
				sps = append(sps, sec.Major.Subproject)
			}
			for _, sp := range sps {
				for _, f := range sp.Focuses {
					add(f)
				}
			}
		}
	}
	return out
}

// FocusSlug converts a focus string to a URL/ID-safe slug.
func FocusSlug(f Focus) string {
	return strings.ReplaceAll(strings.ToLower(string(f)), " ", "-")
}

// TitleSlug converts a title string to a URL/ID-safe slug.
func TitleSlug(s string) string {
	// lowercase, replace spaces with hyphens, strip non-alphanumeric-or-hyphen chars
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			b.WriteRune(r)
		}
	}
	// collapse multiple hyphens
	result := b.String()
	for strings.Contains(result, "--") {
		result = strings.ReplaceAll(result, "--", "-")
	}
	return strings.Trim(result, "-")
}

// ProjectSlug returns the URL slug for a project — URLSlug if set, otherwise
// derived from the title via TitleSlug.
func ProjectSlug(p Project) string {
	if p.URLSlug != "" {
		return p.URLSlug
	}
	return TitleSlug(p.Title)
}

// CardAnimationDelay returns the CSS animation-delay for a card at position i (0-based).
func CardAnimationDelay(i int) string {
	return fmt.Sprintf("%.2fs", float64(i)*0.05)
}

// SourceCodeURL returns the URL for a source code link, or empty string.
func SourceCodeURL(sc *SourceCode) string {
	if sc == nil || sc.Link == nil {
		return ""
	}
	return *sc.Link
}
