// ticker.js — tag ticker for "I do" mode
(function () {
  const config  = window.SITE_CONFIG;
  const data    = window.SITE_DATA;
  const SPEED   = config.tickerSpeed;
  const EXT     = config.imageExt || "jpg";

  // Derive unique tags from all subprojects
  const tagSet = new Set();
  data.projects.forEach(function (p) {
    p.subprojects.forEach(function (sp) {
      sp.tags.forEach(function (t) { tagSet.add(t); });
    });
  });
  const TAGS = Array.from(tagSet);

  function patchSignal(id, value) {
    const el = document.getElementById(id);
    if (!el) return;
    el.value = value;
    el.dispatchEvent(new Event("input", { bubbles: true }));
  }

  // Background crossfade
  let activeBg = "a";
  function setBackground(tag) {
    const slug = tag.replace(/\s+/g, "-");
    const url  = "/static/images/tags/" + slug + "." + EXT;
    const inc  = activeBg === "a" ? "bg-b" : "bg-a";
    const out  = activeBg === "a" ? "bg-a" : "bg-b";
    document.getElementById(inc).style.backgroundImage = "url(" + url + ")";
    document.getElementById(inc).style.opacity = "1";
    document.getElementById(out).style.opacity = "0";
    activeBg = activeBg === "a" ? "b" : "a";
  }

  const ticker = document.getElementById("tag-ticker");
  const LINE_HEIGHT_REM = 1.35;

  function itemH() {
    const ref = document.querySelector(".i-do-label");
    if (!ref) return 60;
    return parseFloat(getComputedStyle(ref).fontSize) * LINE_HEIGHT_REM;
  }

  function buildItems() {
    ticker.innerHTML = "";
    [...TAGS, ...TAGS, ...TAGS].forEach(function (t) {
      const el = document.createElement("div");
      el.className   = "tick-item";
      el.textContent = t;
      el.dataset.tag = t;
      el.addEventListener("click", function () { onItemClick(t); });
      ticker.appendChild(el);
    });
  }

  let offset = 0, locked = false, lockedTag = null, lastTag = null, lastT = null;

  function totalH()  { return TAGS.length * itemH(); }
  function wrapH()   { return document.querySelector(".ticker-wrap") ? document.querySelector(".ticker-wrap").clientHeight : 400; }
  function getItems(){ return [...ticker.querySelectorAll(".tick-item")]; }

  function tagAtCentre() {
    const IH = itemH(), centre = wrapH() / 2;
    let best = null, bestDist = Infinity;
    getItems().forEach(function (el) {
      const dist = Math.abs(el.offsetTop - offset + IH / 2 - centre);
      if (dist < bestDist) { bestDist = dist; best = el.dataset.tag; }
    });
    return best;
  }

  function updateClasses() {
    const IH = itemH(), centre = wrapH() / 2;
    getItems().forEach(function (el) {
      const dist = Math.abs(el.offsetTop - offset + IH / 2 - centre);
      el.classList.toggle("tick-item--centre", dist < IH * 0.55 && !locked);
      el.classList.toggle("tick-item--locked",  locked && el.dataset.tag === lockedTag);
    });
  }

  function snapTo(tag) {
    const IH = itemH(), centre = wrapH() / 2 - IH / 2, TH = totalH();
    let best = null, bestDist = Infinity;
    getItems().filter(function (el) { return el.dataset.tag === tag; }).forEach(function (el) {
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

  function onTagChange(tag) {
    setBackground(tag);
    patchSignal("focus-input", tag);
  }

  function onItemClick(tag) {
    if (locked && lockedTag === tag) {
      locked    = false;
      lockedTag = null;
      patchSignal("locked-input", "false");
    } else {
      locked    = true;
      lockedTag = tag;
      lastTag   = tag;
      snapTo(tag);
      onTagChange(tag);
      patchSignal("locked-input", "true");
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
      const t = tagAtCentre();
      if (t && t !== lastTag) { lastTag = t; onTagChange(t); }
    }
    requestAnimationFrame(step);
  }

  // locked hidden input
  const lockedInput = document.createElement("input");
  lockedInput.id = "locked-input";
  lockedInput.type = "checkbox";
  lockedInput.style.display = "none";
  lockedInput.setAttribute("data-bind:locked", "");
  document.getElementById("app").appendChild(lockedInput);

  buildItems();
  const IH = itemH();
  offset = TAGS.length * IH - (wrapH() / 2 - IH / 2);
  normalise();
  ticker.style.transform = "translateY(" + (-offset) + "px)";
  updateClasses();

  // Init first tag
  const firstTag = TAGS[0];
  document.getElementById("bg-a").style.backgroundImage =
    "url(/static/images/tags/" + firstTag.replace(/\s+/g, "-") + "." + EXT + ")";
  document.getElementById("bg-a").style.opacity = "1";
  setTimeout(function () { patchSignal("focus-input", firstTag); lastTag = firstTag; }, 50);

  requestAnimationFrame(step);
})();
