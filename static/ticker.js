(function () {
  const config   = window.SITE_CONFIG;
  const WORDS    = config.words;
  const SPEED    = config.tickerSpeed;
  const EXT      = config.imageExt || "jpg";
  const PROJECTS = config.projects;

  // ── Card rendering ───────────────────────────────────────────────
  // Called by Datastar's data-effect whenever $currentWord changes.
  // Props down via argument; no global state touched.
  window.renderCards = function(word) {
    const grid  = document.getElementById("card-grid");
    const tmpl  = document.getElementById("card-template");
    const cards = PROJECTS[word] || [];

    grid.innerHTML = "";
    cards.forEach(function(p) {
      const node = tmpl.content.cloneNode(true);
      node.querySelector(".card-title").textContent = p.title;
      node.querySelector(".card-desc").textContent  = p.description;

      const footer = node.querySelector(".card-footer");
      p.tags.forEach(function(t) {
        const tag = document.createElement("span");
        tag.className   = "tag";
        tag.textContent = t;
        footer.appendChild(tag);
      });

      const link = document.createElement("a");
      link.className   = "card-link";
      link.href        = p.link;
      link.textContent = "View →";
      footer.appendChild(link);

      grid.appendChild(node);
    });

    // Re-trigger entry animation
    grid.classList.remove("cards-enter");
    void grid.offsetWidth;
    grid.classList.add("cards-enter");
  };

  // ── Background crossfade ─────────────────────────────────────────
  let activeBg = "a";

  function setBackground(word) {
    const url      = "/static/images/" + word + "." + EXT;
    const incoming = activeBg === "a" ? "bg-b" : "bg-a";
    const outgoing = activeBg === "a" ? "bg-a" : "bg-b";
    document.getElementById(incoming).style.backgroundImage = "url(" + url + ")";
    document.getElementById(incoming).style.opacity = "1";
    document.getElementById(outgoing).style.opacity = "0";
    activeBg = activeBg === "a" ? "b" : "a";
  }

  // ── Signal bridge ────────────────────────────────────────────────
  function patchSignal(id, value) {
    const el = document.getElementById(id);
    if (!el) return;
    el.value = value;
    el.dispatchEvent(new Event("input", { bubbles: true }));
  }

  // ── Ticker ───────────────────────────────────────────────────────
  const ticker = document.getElementById("ticker");
  const LINE_HEIGHT_REM = 1.35;

  function itemH() {
    return parseFloat(getComputedStyle(document.querySelector(".i-make")).fontSize) * LINE_HEIGHT_REM;
  }

  function buildItems() {
    ticker.innerHTML = "";
    [...WORDS, ...WORDS, ...WORDS].forEach(function(w) {
      const el = document.createElement("div");
      el.className    = "tick-item";
      el.textContent  = w;
      el.dataset.word = w;
      el.addEventListener("click", function() { onItemClick(w); });
      ticker.appendChild(el);
    });
  }

  let offset = 0, locked = false, lockedWord = null, lastWord = null, lastT = null;

  function totalH() { return WORDS.length * itemH(); }
  function wrapH()  { return document.querySelector(".ticker-wrap").clientHeight; }
  function getItems() { return [...ticker.querySelectorAll(".tick-item")]; }

  function wordAtCentre() {
    const IH = itemH(), centre = wrapH() / 2;
    let best = null, bestDist = Infinity;
    getItems().forEach(function(el) {
      const dist = Math.abs(el.offsetTop - offset + IH / 2 - centre);
      if (dist < bestDist) { bestDist = dist; best = el.dataset.word; }
    });
    return best;
  }

  function updateClasses() {
    const IH = itemH(), centre = wrapH() / 2;
    getItems().forEach(function(el) {
      const dist = Math.abs(el.offsetTop - offset + IH / 2 - centre);
      el.classList.toggle("tick-item--centre", dist < IH * 0.55 && !locked);
      el.classList.toggle("tick-item--locked", locked && el.dataset.word === lockedWord);
    });
  }

  function snapTo(word) {
    const IH = itemH(), centre = wrapH() / 2 - IH / 2, TH = totalH();
    let best = null, bestDist = Infinity;
    getItems().filter(function(el) { return el.dataset.word === word; }).forEach(function(el) {
      const dist = Math.abs(el.offsetTop - offset - centre);
      if (dist < bestDist) { bestDist = dist; best = el; }
    });
    if (!best) return;
    let delta = best.offsetTop - centre - offset;
    if (delta > TH / 2) delta -= TH;
    if (delta < -TH / 2) delta += TH;
    offset += delta;
    normalise();
    ticker.style.transform = "translateY(" + (-offset) + "px)";
  }

  function normalise() {
    const TH = totalH();
    if (offset >= TH) offset -= TH;
    if (offset < 0)   offset += TH;
  }

  function onWordChange(word) {
    setBackground(word);
    patchSignal("word-input", word);
  }

  function patchLocked(val) {
    const el = document.getElementById("locked-input");
    if (!el) return;
    el.checked = val;
    el.dispatchEvent(new Event("change", { bubbles: true }));
  }

  function onItemClick(word) {
    if (locked && lockedWord === word) {
      locked     = false;
      lockedWord = null;
      patchLocked(false);
    } else {
      locked     = true;
      lockedWord = word;
      lastWord   = word;
      snapTo(word);
      onWordChange(word);
      patchLocked(true);
    }
    updateClasses();
  }

  function step(ts) {
    if (!lastT) lastT = ts;
    const dt = Math.min(ts - lastT, 50);
    lastT = ts;
    if (!locked) {
      offset += (dt / 16) * SPEED;
      normalise();
      ticker.style.transform = "translateY(" + (-offset) + "px)";
      updateClasses();
      const w = wordAtCentre();
      if (w && w !== lastWord) { lastWord = w; onWordChange(w); }
    }
    requestAnimationFrame(step);
  }

  // ── Init ─────────────────────────────────────────────────────────
  buildItems();
  const IH = itemH();
  offset = WORDS.length * IH - (wrapH() / 2 - IH / 2);
  normalise();
  ticker.style.transform = "translateY(" + (-offset) + "px)";
  updateClasses();

  const firstWord = WORDS[0];
  document.getElementById("bg-a").style.backgroundImage = "url(/static/images/" + firstWord + "." + EXT + ")";
  document.getElementById("bg-a").style.opacity = "1";

  // Trigger initial word signal so data-effect fires on load
  setTimeout(function() { patchSignal("word-input", firstWord); }, 50);
  lastWord = firstWord;

  requestAnimationFrame(step);
})();
