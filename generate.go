package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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

func runGenerate() {
	// ── Load and parse all templates ─────────────────────────────────
	// tmplPtr starts nil; templateFuncs captures &tmpl so execTemplate
	// dereferences the correct set at call time, after parsing completes.
	var tmpl *template.Template
	tmpl = template.New("").Funcs(templateFuncs(&tmpl))

	var allTemplateContent strings.Builder
	err := filepath.WalkDir("templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".html") {
			return err
		}
		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}
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
	rendered := RenderProjects(projects)
	jobs := filterJobs(rendered)
	nonJobProjects := filterNonJobs(rendered)
	allFocuses := AllFocuses(projects)

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
				"Projects":   rendered,
				"AllFocuses": allFocuses,
			},
		},
		{
			name:     "worked-at.html",
			outFile:  "dist/worked-at.html",
			activeID: "worked-at",
			signals:  `{}`,
			data: map[string]interface{}{
				"Jobs": jobs,
				"Job":  jobs[0],
			},
		},
		{
			name:     "worked-on.html",
			outFile:  "dist/worked-on.html",
			activeID: "worked-on",
			signals:  `{}`,
			data: map[string]interface{}{
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
		pageData["Signals"] = template.HTMLAttr(`data-signals='` + p.signals + `'`)
		pageData["Effect"] = template.HTMLAttr(p.effect)

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

	// ── Per-job pages ─────────────────────────────────────────────────
	if err := os.MkdirAll("dist/i-worked-at", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "mkdir error: %v\n", err)
		os.Exit(1)
	}
	for _, job := range jobs {
		dir := "dist/i-worked-at/" + job.URLSlug
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "mkdir error: %v\n", err)
			os.Exit(1)
		}
		pageData := map[string]interface{}{
			"Nav":     map[string]interface{}{"ActiveMode": "worked-at", "Modes": navModes},
			"Signals": template.HTMLAttr(`data-signals='{}'`),
			"Effect":  template.HTMLAttr(""),
			"Job":     job,
			"Jobs":    jobs,
		}
		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, "job.html", pageData); err != nil {
			fmt.Fprintf(os.Stderr, "render job %s error: %v\n", job.URLSlug, err)
			os.Exit(1)
		}
		outFile := dir + "/index.html"
		if err := os.WriteFile(outFile, buf.Bytes(), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "write %s error: %v\n", outFile, err)
			os.Exit(1)
		}
		fmt.Println("wrote", outFile)
	}

	// ── Per-project pages ─────────────────────────────────────────────
	if err := os.MkdirAll("dist/i-worked-on", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "mkdir error: %v\n", err)
		os.Exit(1)
	}
	for _, p := range nonJobProjects {
		dir := "dist/i-worked-on/" + p.URLSlug
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "mkdir error: %v\n", err)
			os.Exit(1)
		}
		pageData := map[string]interface{}{
			"Nav":     map[string]interface{}{"ActiveMode": "worked-on", "Modes": navModes},
			"Signals": template.HTMLAttr(`data-signals='{}'`),
			"Effect":  template.HTMLAttr(""),
			"Project": p,
		}
		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, "project.html", pageData); err != nil {
			fmt.Fprintf(os.Stderr, "render project %s error: %v\n", p.URLSlug, err)
			os.Exit(1)
		}
		outFile := dir + "/index.html"
		if err := os.WriteFile(outFile, buf.Bytes(), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "write %s error: %v\n", outFile, err)
			os.Exit(1)
		}
		fmt.Println("wrote", outFile)
	}
}

// ── Helpers ───────────────────────────────────────────────────────────────────

func filterJobs(ps []RenderedProject) []RenderedProject {
	var out []RenderedProject
	for _, p := range ps {
		if p.Type == ProjectTypeJob {
			out = append(out, p)
		}
	}
	return out
}

func filterNonJobs(ps []RenderedProject) []RenderedProject {
	var out []RenderedProject
	for _, p := range ps {
		if p.Type != ProjectTypeJob && p.Category != nil {
			out = append(out, p)
		}
	}
	return out
}

func groupByCategory(ps []RenderedProject) map[string][]RenderedProject {
	out := map[string][]RenderedProject{}
	for _, p := range ps {
		cat := string(*p.Category)
		out[cat] = append(out[cat], p)
	}
	return out
}

// templateFuncs returns the Go template function map.
// tmplPtr is a pointer to the template set, filled in after parsing.
// execTemplate dereferences it at call time, so it is safe to pass
// a pointer that is nil during parsing and assigned afterwards.
func templateFuncs(tmplPtr **template.Template) template.FuncMap {
	return template.FuncMap{
		"execTemplate": func(name string, data interface{}) (template.HTML, error) {
			var buf bytes.Buffer
			err := (*tmplPtr).ExecuteTemplate(&buf, name, data)
			if err != nil {
				return "", fmt.Errorf("execTemplate %q: %w", name, err)
			}
			return template.HTML(buf.String()), nil
		},
		"FocusSlug":          FocusSlug,
		"TitleSlug":          TitleSlug,
		"MergeTechTags":      MergeTechTags,
		"OwnTechTags":        OwnTechTags,
		"EffectiveTechTags":  EffectiveTechTags,
		"CardAnimationDelay": CardAnimationDelay,
		"HasFocus":           HasFocus,
		"lower": func(v interface{}) string {
			return strings.ToLower(fmt.Sprintf("%v", v))
		},
		"deref":    func(s *string) string { if s == nil { return "" }; return *s },
		"derefImg": func(i *Image) string { if i == nil { return "" }; return string(*i) },
		"add":      func(a, b int) int { return a + b },
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
