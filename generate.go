package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// NavMode describes a navigation item.
type NavMode struct {
	ID    string
	Label string
	URL   string
}

var navModes = []NavMode{
	{ID: "about",      Label: "Hi, I'm Marvin", URL: "/about.html"},
	{ID: "do",         Label: "I do",            URL: "/index.html"},
	{ID: "worked-at",  Label: "I worked at",     URL: "/worked-at.html"},
	{ID: "worked-on",  Label: "I worked on",     URL: "/worked-on.html"},
	{ID: "contact",    Label: "Contact me at",   URL: "/contact.html"},
}

// IndexedProject wraps a non-job project with its index in the nonJobProjects slice.
type IndexedProject struct {
	Project
	GlobalIndex int
}

func runGenerate() {
	// ── Load and parse all templates ─────────────────────────────────
	// execTemplate needs the tmpl set but must be in the FuncMap before parsing.
	// We use a pointer-to-template trick to break the circular dependency.
	var tmplPtr *template.Template
	funcs := templateFuncs()
	funcs["execTemplate"] = func(name string, data interface{}) (string, error) {
		var buf bytes.Buffer
		err := tmplPtr.ExecuteTemplate(&buf, name, data)
		if err != nil {
			return "", fmt.Errorf("execTemplate %q: %w", name, err)
		}
		return buf.String(), nil
	}

	tmpl := template.New("").Funcs(funcs)

	var allTemplateContent strings.Builder
	err := filepath.WalkDir("templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".html") {
			return err
		}
		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		fmt.Printf("adding template file: %s (%d bytes)\n", path, len(data))
		// Strip carriage returns for Windows CRLF compatibility
		cleaned := strings.ReplaceAll(string(data), "\r\n", "\n")
		allTemplateContent.WriteString(cleaned)
		allTemplateContent.WriteString("\n")
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "template walk error: %v\n", err)
		os.Exit(1)
	}

	tmpl, err = tmpl.Parse(allTemplateContent.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "template parse error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Registered templates:")
	for _, t := range tmpl.Templates() {
		fmt.Println(" -", t.Name())
	}

	// Point the pointer at the fully-parsed template set
	tmplPtr = tmpl

	// ── Concatenate all CSS files ─────────────────────────────────────
	var cssBuilder strings.Builder
	err = filepath.WalkDir("templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".css") {
			return err
		}
		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		cssBuilder.WriteString("/* " + path + " */\n")
		cssBuilder.Write(data)
		cssBuilder.WriteString("\n")
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "css concat error: %v\n", err)
		os.Exit(1)
	}
	if err := os.MkdirAll("static", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "mkdir error: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile("static/style.css", []byte(cssBuilder.String()), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "write style.css error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("wrote static/style.css")

	// ── Prepare data ─────────────────────────────────────────────────
	jobs := filterJobs(projects)
	nonJobProjects := filterNonJobs(projects)
	allFocuses := AllFocuses(projects)
	flatSps := FlattenSubprojects(projects)
	_ = flatSps // available for future use

	// ── Render pages ─────────────────────────────────────────────────
	type pageSpec struct {
		name     string
		outFile  string
		activeID string
		signals  string
		effect   string
		data     map[string]interface{}
	}

	pages := []pageSpec{
		{
			name:     "index.html",
			outFile:  "dist/index.html",
			activeID: "do",
			signals:  `{}`,
			data: map[string]interface{}{
				"Projects":   projects,
				"AllFocuses": allFocuses,
			},
		},
		{
			name:     "worked-at.html",
			outFile:  "dist/worked-at.html",
			activeID: "worked-at",
			signals:  `{"jobFocus": 0}`,
			effect:   `console.log('jobFocus:', $jobFocus)`,
			data: map[string]interface{}{
				"Jobs": jobs,
			},
		},
		{
			name:     "worked-on.html",
			outFile:  "dist/worked-on.html",
			activeID: "worked-on",
			signals:  `{"selectedProject": -1}`,
			data: map[string]interface{}{
				"NonJobProjects":     indexedNonJobs(nonJobProjects),
				"ProjectsByCategory": groupByCategory(nonJobProjects),
			},
		},
		{
			name:     "about.html",
			outFile:  "dist/about.html",
			activeID: "about",
			signals:  `{}`,
			data:     map[string]interface{}{},
		},
		{
			name:     "contact.html",
			outFile:  "dist/contact.html",
			activeID: "contact",
			signals:  `{}`,
			data:     map[string]interface{}{},
		},
	}

	if err := os.MkdirAll("dist", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "mkdir dist error: %v\n", err)
		os.Exit(1)
	}

	for _, p := range pages {
		pageData := p.data
		pageData["Nav"] = map[string]interface{}{
			"ActiveMode": p.activeID,
			"Modes":      navModes,
		}
		pageData["Signals"] = p.signals
		pageData["Effect"] = p.effect

		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, p.name, pageData); err != nil {
			fmt.Fprintf(os.Stderr, "render %s error: %v\n", p.name, err)
			os.Exit(1)
		}
		if err := os.WriteFile(p.outFile, buf.Bytes(), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "write %s error: %v\n", p.outFile, err)
			os.Exit(1)
		}
		fmt.Println("wrote", p.outFile)
	}
}

// ── Helpers ───────────────────────────────────────────────────────────────────

func filterJobs(ps []Project) []Project {
	var out []Project
	for _, p := range ps {
		if p.Type == ProjectTypeJob {
			out = append(out, p)
		}
	}
	return out
}

func filterNonJobs(ps []Project) []Project {
	var out []Project
	for _, p := range ps {
		if p.Type != ProjectTypeJob && p.Category != nil {
			out = append(out, p)
		}
	}
	return out
}

func indexedNonJobs(ps []Project) []IndexedProject {
	out := make([]IndexedProject, len(ps))
	for i, p := range ps {
		out[i] = IndexedProject{Project: p, GlobalIndex: i}
	}
	return out
}

func groupByCategory(ps []Project) map[string][]IndexedProject {
	out := map[string][]IndexedProject{}
	for i, p := range ps {
		cat := string(*p.Category)
		out[cat] = append(out[cat], IndexedProject{Project: p, GlobalIndex: i})
	}
	return out
}

// templateFuncs returns the Go template function map.
func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"FocusSlug":         FocusSlug,
		"RenderDescription": RenderDescription,
		"MergeTechTags":     MergeTechTags,
		"OwnTechTags":       OwnTechTags,
		"EffectiveTechTags": EffectiveTechTags,
		"CardAnimationDelay": CardAnimationDelay,
		"HasFocus":          HasFocus,
		"lower": func(v interface{}) string {
			return strings.ToLower(fmt.Sprintf("%v", v))
		},
		"printf": fmt.Sprintf,
		"deref":             func(s *string) string { if s == nil { return "" }; return *s },
		"derefImg":          func(i *Image) string { if i == nil { return "" }; return string(*i) },
		"add":               func(a, b int) int { return a + b },
		// dict builds a map from alternating key/value pairs, used in template calls
		"dict": func(values ...interface{}) map[string]interface{} {
			m := map[string]interface{}{}
			for i := 0; i+1 < len(values); i += 2 {
				m[values[i].(string)] = values[i+1]
			}
			return m
		},
	}
}
