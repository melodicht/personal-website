// ticker.js — tag ticker animation and grid swap for the [I do] page
(function () {
  const SPEED = window.TICKER_SPEED || 0.2;
  const EXT   = window.IMAGE_EXT   || "jpg";

  const ticker    = document.getElementById("tag-ticker");
  const container = document.getElementById("card-container");
  const store     = document.getElementById("template-store");

  if (!ticker || !container || !store) return;

  // ── Background crossfade ──────────────────────────────────────────
  let activeBg = "a";
  function setBackground(slug) {
    const url = "/static/images/tags/" + slug + "." + EXT;
    const inc = activeBg === "a" ? "bg-b" : "bg-a";
    const out = activeBg === "a" ? "bg-a" : "bg-b";
    document.getElementById(inc).style.backgroundImage = "url(" + url + ")";
    document.getElementById(inc).style.opacity = "1";
    document.getElementById(out).style.opacity = "0";
    activeBg = activeBg === "a" ? "b" : "a";
  }

  // ── Grid swap  ────────────────────────────────────────────────────
  let currentFocus = null;

  function swapGrid(slug) {
    const tmpl = store.querySelector("#grid-" + slug);
    if (!tmpl || slug === currentFocus) return;
    currentFocus = slug;
    const doSwap = function () {
      container.innerHTML = "";
      container.appendChild(tmpl.content.cloneNode(true));
    };
    // NOTE(marvin): Disabling view transition because of the lag it causes.
    // if (document.startViewTransition) {
      // document.startViewTransition(doSwap);
    // } else {
      doSwap();
    // }
  }

  // ── Ticker scroll ─────────────────────────────────────────────────
  function itemH() {
    const el = ticker.querySelector(".tick-item");
    return el ? el.offsetHeight : 40;
  }

  // Fixed logical wrap height so totalH() > wrapH() for looping
  function wrapH()  { return 480; }
  function totalH() { return ticker.querySelectorAll(".tick-item").length / 3 * itemH(); }

  function getItems() { return [...ticker.querySelectorAll(".tick-item")]; }

  let offset = 0, locked = false, lockedFocus = null, lastFocus = null, lastT = null;

  function focusAtCentre() {
    const IH = itemH(), centre = wrapH() / 2 + IH / 2;
    let best = null, bestDist = Infinity;
    getItems().forEach(function (el) {
      const dist = Math.abs(el.offsetTop - offset + IH / 2 - centre);
      if (dist < bestDist) { bestDist = dist; best = el.dataset.focus; }
    });
    return best;
  }

  function updateClasses() {
    const IH = itemH(), centre = wrapH() / 2 + IH / 2;
    getItems().forEach(function (el) {
      const dist = Math.abs(el.offsetTop - offset + IH / 2 - centre);
      el.classList.toggle("tick-item--centre", dist < IH * 0.55 && !locked);
      el.classList.toggle("tick-item--locked", locked && el.dataset.focus === lockedFocus);
    });
  }

  function normalise() {
    const TH = totalH();
    if (offset >= TH) offset -= TH;
    if (offset < 0)   offset += TH;
  }

  function snapTo(focus) {
    const IH = itemH(), centre = wrapH() / 2 - IH / 2, TH = totalH();
    let best = null, bestDist = Infinity;
    getItems().filter(function (el) { return el.dataset.focus === focus; }).forEach(function (el) {
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

  function onFocusChange(focus) {
    setBackground(focus);
    swapGrid(focus);
  }

  ticker.addEventListener("click", function (e) {
    const item = e.target.closest(".tick-item");
    if (!item) return;
    const focus = item.dataset.focus;
    if (locked && lockedFocus === focus) {
      locked = false; lockedFocus = null;
    } else {
      locked = true; lockedFocus = focus;
      lastFocus = focus;
      snapTo(focus);
      onFocusChange(focus);
    }
    updateClasses();
  });

  // Triple the items for seamless looping (items already in HTML once;
  // we clone them to make three copies)
  (function tripleItems() {
    const originals = [...ticker.querySelectorAll(".tick-item")];
    originals.forEach(function (el) {
      ticker.appendChild(el.cloneNode(true));
    });
    originals.forEach(function (el) {
      ticker.insertBefore(el.cloneNode(true), ticker.firstChild);
    });
  })();

  function step(ts) {
    if (!lastT) lastT = ts;
    const dt = Math.min(ts - lastT, 50);
    lastT = ts;
    if (!locked) {
      offset += (dt / 16) * SPEED;
      normalise();
      ticker.style.transform = "translateY(" + (-offset) + "px)";
      updateClasses();
      const f = focusAtCentre();
      if (f && f !== lastFocus) { lastFocus = f; onFocusChange(f); }
    }
    requestAnimationFrame(step);
  }

  // Init: start at middle third, trigger first focus
  offset = totalH();
  normalise();
  ticker.style.transform = "translateY(" + (-offset) + "px)";
  updateClasses();

  const firstItem = ticker.querySelector(".tick-item");
  if (firstItem) {
    const firstFocus = firstItem.dataset.focus;
    lastFocus = firstFocus;
    onFocusChange(firstFocus);
  }

  requestAnimationFrame(step);
})();
