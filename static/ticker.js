(function () {
  const WORDS = ["websites", "games", "music", "art", "tools", "apps"];

  // Write a value into a hidden input and fire "input" so data-bind picks it up.
  function patchSignal(id, value) {
    const el = document.getElementById(id);
    if (!el) return;
    el.value = value;
    el.dispatchEvent(new Event("input", { bubbles: true }));
  }

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
      patchSignal("word-input", word);
    }
    updateClasses();
  }

  function step(ts) {
    if (!lastT) lastT = ts;
    const dt = Math.min(ts - lastT, 50);
    lastT = ts;

    if (!locked) {
      offset += (dt / 16) * 0.2;
      normalise();
      ticker.style.transform = `translateY(${-offset}px)`;
      updateClasses();

      const w = wordAtCentre();
      if (w && w !== lastWord) {
        lastWord = w;
        patchSignal("word-input", w);
      }
    }
    requestAnimationFrame(step);
  }

  buildItems();
  const IH = itemH();
  offset = WORDS.length * IH - (wrapH() / 2 - IH / 2);
  normalise();
  ticker.style.transform = `translateY(${-offset}px)`;
  updateClasses();
  requestAnimationFrame(step);
})();
