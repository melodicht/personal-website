// jobs.js — job cell picker for "I worked at" mode
(function () {
  const data = window.SITE_DATA;
  const jobs = data.projects.filter(function (p) { return p.type === "Job"; });

  const list = document.getElementById("job-list");

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
    const listH  = list.clientHeight;
    const itemH  = el.offsetHeight;
    const target = el.offsetTop - listH / 2 + itemH / 2;
    list.scrollTo({ top: target, behavior: "smooth" });
  }

  function selectJob(idx) {
    updateCells(idx);
    snapToJob(idx);
    patchJobFocus(idx);
  }

  function escHtml(str) {
    return String(str)
      .replace(/&/g, "&amp;").replace(/</g, "&lt;")
      .replace(/>/g, "&gt;").replace(/"/g, "&quot;");
  }

  // Init first job selected
  setTimeout(function () { selectJob(0); }, 60);
})();
