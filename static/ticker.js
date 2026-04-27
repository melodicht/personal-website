(function () {
  const config = window.SITE_CONFIG;
  const WORDS = config.words;
  const SPEED = config.tickerSpeed;
  const EXT   = config.imageExt || "jpg";

  // ── Background crossfade ─────────────────────────────────────────
  // We keep two layers (bg-a, bg-b). The "active" layer is fully opaque;
  // the "inactive" one is transparent. To transition:
  //   1. Load the new image onto the inactive layer.
  //   2. Fade the inactive layer to opaque.
  //   3. Fade the active layer to transparent.
  //   4. Swap which layer is considered active.
  let activeBg = "a";

  function setBackground(word) {
    const url = "/static/images/" + word + "." + EXT;
    const incoming = activeBg === "a" ? "bg-b" : "bg-a";
    const outgoing = activeBg === "a" ? "bg-a" : "bg-b";

    const inEl  = document.getElementById(incoming);
    const outEl = document.getElementById(outgoing);

    inEl.style.backgroundImage = "url(" + url + ")";
    inEl.style.opacity = "1";
    outEl.style.opacity = "0";

    activeBg = activeBg === "a" ? "b" : "a";
  }

  // ── Signal bridge ────────────────────────────────────────────────
  function patchSignal(id, value) {
    const el = document.getElementById(id);
    if (!el) return;
    el.value = value;
    el.dispatchEvent(new Event("input", { bubbles: true }));
  }

  // ── Ticker setup ─────────────────────────────────────────────────
  const ticker = document.getElementById("ticker");
  const LINE_HEIGHT_REM = 1.35;

  function itemH() {
    return parseFloat(getComputedStyle(document.querySelector(".i-make")).fontSize) * LINE_HEIGHT_REM;
  }

  function buildItems() {
    ticker.innerHTML = "";
    [...WORDS, ...WORDS, ...WORDS].forEach((w) => {
      const el = document.createElement("div");
      el.className = "tick-item";
      el.textContent = w;
      el.dataset.word = w;
      el.addEventListener("click", () => onItemClick(w));
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
    getItems().forEach((el) => {
      const dist = Math.abs(el.offsetTop - offset + IH / 2 - centre);
      if (dist < bestDist) { bestDist = dist; best = el.dataset.word; }
    });
    return best;
  }

  function updateClasses() {
    const IH = itemH(), centre = wrapH() / 2;
    getItems().forEach((el) => {
      const dist = Math.abs(el.offsetTop - offset + IH / 2 - centre);
      el.classList.toggle("tick-item--centre", dist < IH * 0.55 && !locked);
      el.classList.toggle("tick-item--locked", locked && el.dataset.word === lockedWord);
    });
  }

  function snapTo(word) {
    const IH = itemH(), centre = wrapH() / 2 - IH / 2, TH = totalH();
    let best = null, bestDist = Infinity;
    getItems().filter(el => el.dataset.word === word).forEach((el) => {
      const dist = Math.abs(el.offsetTop - offset - centre);
      if (dist < bestDist) { bestDist = dist; best = el; }
    });
    if (!best) return;
    let delta = best.offsetTop - centre - offset;
    if (delta > TH / 2) delta -= TH;
    if (delta < -TH / 2) delta += TH;
    offset += delta;
    normalise();
    ticker.style.transform = `translateY(${-offset}px)`;
  }

  function normalise() {
    const TH = totalH();
    if (offset >= TH) offset -= TH;
    if (offset < 0) offset += TH;
  }

  function onWordChange(word) {
    patchSignal("word-input", word);
    setBackground(word);
  }

  function onItemClick(word) {
    if (locked && lockedWord === word) {
      locked = false;
      lockedWord = null;
      patchSignal("locked-input", "false");
    } else {
      locked = true;
      lockedWord = word;
      lastWord = word;
      snapTo(word);
      patchSignal("locked-input", "true");
      onWordChange(word);
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
      ticker.style.transform = `translateY(${-offset}px)`;
      updateClasses();

      const w = wordAtCentre();
      if (w && w !== lastWord) {
        lastWord = w;
        onWordChange(w);
      }
    }
    requestAnimationFrame(step);
  }

  // ── Init ─────────────────────────────────────────────────────────
  buildItems();
  const IH = itemH();
  offset = WORDS.length * IH - (wrapH() / 2 - IH / 2);
  normalise();
  ticker.style.transform = `translateY(${-offset}px)`;
  updateClasses();

  // Set initial background immediately (no transition on first load)
  const firstWord = WORDS[0];
  document.getElementById("bg-a").style.backgroundImage =
    "url(/static/images/" + firstWord + "." + EXT + ")";
  document.getElementById("bg-a").style.opacity = "1";

  requestAnimationFrame(step);
})();
