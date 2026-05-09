// jobs.js — job cell picker for "I worked at" mode
(function () {
  const data = window.SITE_DATA;
  const jobs = data.projects.filter(function (p) { return p.type === "Job"; });

  const list = document.getElementById("job-list");

  let datastarReady = false;
  let currentIdx = 0;

  // Build job cells
  jobs.forEach(function (job, i) {
    const exp   = job.specifics.job;
    const end   = exp.dateRange.end || "Present";
    const range = exp.dateRange.start + " – " + end;

    const cell = document.createElement("div");
    cell.className   = "job-cell";
    cell.dataset.idx = i;
    cell.innerHTML =
      "<span class='job-cell-company'>" + escHtml(exp.company) + "</span>" +
      "<span class='job-cell-role'>"    + escHtml(exp.role)    + "</span>" +
      "<span class='job-cell-range'>"   + escHtml(range)       + "</span>";

    cell.addEventListener("click", function () { selectJob(i); });
    list.appendChild(cell);
  });

  // Hidden input for $jobFocus signal
  const jobInput = document.createElement("input");
  jobInput.id    = "job-focus-input";
  jobInput.type  = "text";
  jobInput.style.display = "none";
  jobInput.setAttribute("data-bind:job-focus", "");
  document.getElementById("app").appendChild(jobInput);

  function patchJobFocus(idx) {
    jobInput.value = String(idx);
    jobInput.dispatchEvent(new Event("input", { bubbles: true }));
  }

  function updateCells(idx) {
    list.querySelectorAll(".job-cell").forEach(function (el) {
      el.classList.toggle("job-cell--active", Number(el.dataset.idx) === idx);
    });
  }

  function snapToJob(idx) {
    const el = list.querySelector("[data-idx='" + idx + "']");
    if (!el) return;
    const listRect  = list.getBoundingClientRect();
    const elRect    = el.getBoundingClientRect();
    // Target the top of the cell (where the company name starts) plus half a line height
    const lineHeight = 24; // ~1rem at 16px base, matching .job-cell-company font-size
    const targetMid  = elRect.top - listRect.top + list.scrollTop - lineHeight / 2;
    const target     = targetMid - list.clientHeight / 2;
    list.scrollTo({ top: target, behavior: "smooth" });
  }

  // Expose so renderJobDetail can snap correctly on first mode switch
  window._snapToJob = function (idx) { snapToJob(idx); };

  function selectJob(idx) {
    currentIdx = idx;
    updateCells(idx);
    snapToJob(idx);
    patchJobFocus(idx);
  }

  function escHtml(str) {
    return String(str)
      .replace(/&/g, "&amp;").replace(/</g, "&lt;")
      .replace(/>/g, "&gt;").replace(/"/g, "&quot;");
  }

  function init() {
    selectJob(0);
  }

  document.addEventListener("datastar-loaded", function () {
    datastarReady = true;
    init();
  });

  if (datastarReady) {
    init();
  }
})();
