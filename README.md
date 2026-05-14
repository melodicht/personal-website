# personal-site

A statically-generated personal portfolio site using Go templates and Datastar.

## Stack

- **Go** standard library — static site generation and dev server
- **Datastar** (CDN) — reactive signals for in-page interactions
- **Vanilla JS** — ticker animation and grid swap only
- **Atomic design** — components organized as atoms → molecules → organisms → pages

## Run locally

```bash
# Generate the static site (writes dist/ and static/style.css)
go run . -generate

# Serve it
go run .
```

OR
```
go build -o site.exe .
.\site.exe -generate

go build -o site.exe .
.\site.exe
```

Open http://localhost:8080

## Project structure

```
.
├── types.go                    # Data types (Focus, TechTag, Project, etc.)
├── render.go                   # Go-side rendering helpers (RenderDescription, etc.)
├── generate.go                 # Static site generator (template rendering, CSS concat)
├── generate_data.go            # Project data definitions
├── main.go                     # Dev server (serves dist/ and static/)
│
├── templates/
│   ├── layout/
│   │   ├── base.html           # Shared HTML shell (head, body, nav slot)
│   │   └── base.css            # Global reset, layout variables, mobile styles
│   ├── components/
│   │   ├── atoms/
│   │   │   └── chip/           # Single tech tag or focus chip
│   │   ├── molecules/
│   │   │   ├── subproject-card/
│   │   │   ├── bullet-point/
│   │   │   ├── major-subproject/
│   │   │   └── source-code-link/
│   │   └── organisms/
│   │       ├── nav/            # Left-hand navigation column
│   │       ├── subsection/     # A titled group of bullets/cards/major within a project
│   │       ├── job-detail/     # Full job detail view
│   │       ├── project-detail/ # Full non-job project detail view
│   │       └── tag-grid/       # Card grid for one focus tag ([I do] page)
│   └── pages/
│       ├── index.html          # [I do] — ticker + card grids
│       ├── worked-at.html      # [I worked at]
│       ├── worked-on.html      # [I worked on]
│       ├── about.html          # [Hi, I'm Marvin]
│       └── contact.html        # [Contact me at]
│
├── static/
│   ├── style.css               # AUTO-GENERATED — do not edit by hand.
│   │                           # Concatenation of all component CSS files.
│   ├── ticker.js               # Ticker animation and grid swap (index page only)
│   └── worked-at.js            # snapToJob helper (worked-at page only)
│
└── dist/                       # AUTO-GENERATED — output of go run . -generate
```

## CSS scope

**All CSS defined in component files is global.** There is no Shadow DOM or
scoped CSS. The concatenated `static/style.css` contains every component's
styles in one file. Class names follow a BEM-like convention (e.g.
`.subproject-card`, `.subproject-card-title`) to avoid collisions.

`static/style.css` is auto-generated — never edit it directly. Edit the
individual `.css` files in `templates/components/` or `templates/layout/`.

## Adding a project

Edit `generate_data.go`. Run `go run . -generate` to rebuild the site.

## Component conventions

- A **component** is a folder containing a `.html` Go template partial and a `.css` file.
- Components in a given layer may only use components from the **same or lower** layers:
  - atoms → no dependencies
  - molecules → atoms only
  - organisms → atoms + molecules
  - pages → any component
- Pages (`templates/pages/`) are full HTML documents that extend `base.html`.
