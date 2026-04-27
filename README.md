# personal-site

A personal website with a Datastar-powered reactive project grid.
Words scroll in the hero; when one centres (or is clicked/locked), the
server streams fresh project cards via SSE.

## Stack

- **Go** standard library — no web framework needed
- **Datastar** (CDN) — signals + SSE merge-fragments
- Vanilla HTML/CSS/JS for the ticker animation

## Run locally

```bash
go run .
# → http://localhost:8080
```

No dependencies to install. `go.mod` uses only the standard library.

## Project structure

```
.
├── main.go                 # HTTP server, SSE endpoint, project data
├── templates/
│   ├── index.html          # Main page — Datastar store lives here
│   └── cards.html          # Fragment streamed into #card-grid via SSE
└── static/
    ├── style.css
    └── ticker.js           # Animation loop; writes $currentWord signal
```

## How it works

1. `ticker.js` runs a `requestAnimationFrame` loop scrolling the word list.
2. When the centred word changes, it calls `setSignal("currentWord", word)`,
   writing into the Datastar store on `#app`.
3. `data-on-signal__currentword` in `index.html` fires a `$$get` to
   `/projects?word=<word>` whenever the signal changes.
4. The Go handler looks up projects for that word, renders `cards.html`,
   and streams it back as a Datastar `merge-fragments` SSE event.
5. Datastar swaps the fragment into `#card-grid` with a CSS entry animation.

## Adding projects

Edit the `projects` map in `main.go`. Keys are the ticker words (lowercase).
Each value is a slice of `Project` structs with Title, Description, Tags,
and Link fields.

To add a new ticker word, also add it to the `WORDS` array in
`static/ticker.js`.
```
