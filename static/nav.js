// nav.js — the leftmost column nav ticker (no auto-scroll, click to snap)
(function () {
  const MODES = [
    { id: "about",      label: "Hi, I'm Marvin" },
    { id: "do",         label: "I do" },
    { id: "worked-at",  label: "I worked at" },
    { id: "worked-on",  label: "I worked on" },
    { id: "contact",    label: "Contact me at" },
  ];

  const container = document.getElementById("nav-ticker");
  let currentMode = "do";
  let datastarReady = false;

  function patchSignal(id, value) {
    const el = document.getElementById(id);
    if (!el) return;
    el.value = value;
    el.dispatchEvent(new Event("input", { bubbles: true }));
  }

  // Build nav items
  MODES.forEach(function (m) {
    const el = document.createElement("div");
    el.className = "nav-item";
    el.dataset.mode = m.id;
    el.textContent = m.label;
    el.addEventListener("click", function () { selectMode(m.id); });
    container.appendChild(el);
  });

  // We need a hidden input for $mode signal
  const modeInput = document.createElement("input");
  modeInput.id = "mode-input";
  modeInput.type = "text";
  modeInput.style.display = "none";
  modeInput.setAttribute("data-bind:mode", "");
  document.getElementById("app").appendChild(modeInput);

  function setMode(mode) {
    modeInput.value = mode;
    modeInput.dispatchEvent(new Event("input", { bubbles: true }));
  }

  function updateClasses(mode) {
    container.querySelectorAll(".nav-item").forEach(function (el) {
      el.classList.toggle("nav-item--active", el.dataset.mode === mode);
    });
  }

  function snapToMode(mode) {
    const el = container.querySelector("[data-mode='" + mode + "']");
    if (!el) return;
    const targetScroll = el.offsetTop - (container.clientHeight / 2) + (el.offsetHeight / 2);
    container.scrollTo({ top: Math.max(0, targetScroll), behavior: "smooth" });
  }

  function selectMode(mode) {
    currentMode = mode;
    setMode(mode);
    updateClasses(mode);
    snapToMode(mode);
  }

  function init() {
    setMode(currentMode);
    snapToMode(currentMode);
  }

  // Init
  updateClasses(currentMode);
  document.addEventListener("datastar-loaded", function () {
    datastarReady = true;
    init();
  });

  if (datastarReady) {
    init();
  }
})();
