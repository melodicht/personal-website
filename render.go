package main

import (
	"fmt"
	"html"
	"regexp"
	"strings"
)

// ── Inline description parsing ────────────────────────────────────────────────
// Parses {tag} tokens in a description string and returns an HTML string where
// each token is rendered as an inline chip span.

var inlineTagRe = regexp.MustCompile(`\{([^}]+)\}`)

// ParsedSegment is either plain text or an inline chip.
type ParsedSegment struct {
	IsChip bool
	Value  string // canonical value (matched against known tech tags if possible)
}

func parseDescription(desc string, knownTechTags []TechTag) []ParsedSegment {
	var segments []ParsedSegment
	last := 0
	matches := inlineTagRe.FindAllStringIndex(desc, -1)
	for _, loc := range matches {
		if loc[0] > last {
			segments = append(segments, ParsedSegment{Value: desc[last:loc[0]]})
		}
		raw := desc[loc[0]+1 : loc[1]-1] // strip braces
		canonical := raw
		for _, tt := range knownTechTags {
			if strings.EqualFold(string(tt), raw) {
				canonical = string(tt)
				break
			}
		}
		segments = append(segments, ParsedSegment{IsChip: true, Value: canonical})
		last = loc[1]
	}
	if last < len(desc) {
		segments = append(segments, ParsedSegment{Value: desc[last:]})
	}
	return segments
}

// RenderDescription returns an HTML string with inline chips substituted.
// Safe to use with text/template since escaping is done manually.
func RenderDescription(desc string, knownTechTags []TechTag) string {
	var sb strings.Builder
	for _, seg := range parseDescription(desc, knownTechTags) {
		if seg.IsChip {
			sb.WriteString(`<span class="chip chip--tech chip--inline">`)
			sb.WriteString(html.EscapeString(seg.Value))
			sb.WriteString(`</span>`)
		} else {
			sb.WriteString(html.EscapeString(seg.Value))
		}
	}
	return sb.String()
}

// InlinedTechTags returns the lowercased set of tech tag strings inlined in a
// description via {tag} syntax, used to avoid duplicating them in the chip row.
func InlinedTechTags(desc string) map[string]bool {
	out := map[string]bool{}
	for _, seg := range parseDescription(desc, nil) {
		if seg.IsChip {
			out[strings.ToLower(seg.Value)] = true
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

// OwnTechTags returns only the subproject's own non-inherited, non-inlined tech tags.
func OwnTechTags(sp Subproject, inherited []TechTag) []TechTag {
	inlined := InlinedTechTags(sp.Description)
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
func EffectiveTechTags(sp Subproject, inherited []TechTag) []TechTag {
	merged := MergeTechTags(inherited, sp.TechTags)
	inlined := InlinedTechTags(sp.Description)
	var out []TechTag
	for _, t := range merged {
		if !inlined[strings.ToLower(string(t))] {
			out = append(out, t)
		}
	}
	return out
}

// ── Focus matching ────────────────────────────────────────────────────────────

// SubprojectEffectiveFocuses returns the full set of focuses for a subproject,
// merging project → subsection → subproject.
func SubprojectEffectiveFocuses(sp Subproject, secFocuses, projFocuses []Focus) []Focus {
	return MergeFocuses(projFocuses, MergeFocuses(secFocuses, sp.Focuses))
}

// HasFocus reports whether a subproject (with its own and project focuses) has the given focus.
func HasFocus(sp Subproject, spFocuses, projFocuses []Focus, f Focus) bool {
	effective := MergeFocuses(projFocuses, spFocuses)
	for _, ef := range effective {
		if ef == f {
			return true
		}
	}
	return false
}

// ── Flat subproject list for [I do] page ─────────────────────────────────────

// FlatSubproject is a subproject with its full inheritance context resolved.
type FlatSubproject struct {
	Subproject       Subproject
	Video            *Video
	ProjectTitle     string
	ProjectType      ProjectType
	InheritedTechTags []TechTag
	// Index into the global flat list, used as the selectedSubproject signal value
	Index int
}

// FlattenSubprojects returns all subprojects across all projects in order,
// with inherited tech tags resolved.
func FlattenSubprojects(projects []Project) []FlatSubproject {
	var result []FlatSubproject
	idx := 0
	for _, p := range projects {
		for _, sec := range p.Subsections {
			inherited := MergeTechTags(p.TechTags, sec.TechTags)
			for _, b := range sec.Bullets {
				result = append(result, FlatSubproject{
					Subproject:        b.Subproject,
					ProjectTitle:      p.Title,
					ProjectType:       p.Type,
					InheritedTechTags: inherited,
					Index:             idx,
				})
				idx++
			}
			for _, c := range sec.Cards {
				result = append(result, FlatSubproject{
					Subproject:        c.Subproject,
					ProjectTitle:      p.Title,
					ProjectType:       p.Type,
					InheritedTechTags: inherited,
					Index:             idx,
				})
				idx++
			}
			if sec.Major != nil {
				result = append(result, FlatSubproject{
					Subproject:        sec.Major.Subproject,
					Video:             sec.Major.Video,
					ProjectTitle:      p.Title,
					ProjectType:       p.Type,
					InheritedTechTags: inherited,
					Index:             idx,
				})
				idx++
			}
		}
	}
	return result
}

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
			allSps := func(sps []Subproject) {
				for _, sp := range sps {
					for _, f := range sp.Focuses {
						add(f)
					}
				}
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
			allSps(sps)
		}
	}
	return out
}

// FocusSlug converts a focus string to a URL/ID-safe slug.
func FocusSlug(f Focus) string {
	return strings.ReplaceAll(strings.ToLower(string(f)), " ", "-")
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
