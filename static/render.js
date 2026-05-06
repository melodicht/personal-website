// render.js — all content-column rendering functions
// Called by Datastar data-effect expressions.
(function () {
  const data   = window.SITE_DATA;
  const config = window.SITE_CONFIG;
  const jobs   = data.projects.filter(function (p) { return p.type === "Job"; });

  // Non-job projects in the same order as the "I worked on" grid
  const nonJobProjects = data.projects.filter(function (p) { return p.type !== "Job" && p.category; });

  function escHtml(s) {
    return String(s)
      .replace(/&/g, "&amp;").replace(/</g, "&lt;")
      .replace(/>/g, "&gt;").replace(/"/g, "&quot;");
  }

  function isExternal(url) {
    return typeof url === "string" && url.startsWith("http");
  }

  function patchSignal(id, value) {
    const el = document.getElementById(id);
    if (!el) return;
    el.value = typeof value === "number" ? String(value) : value;
    el.dispatchEvent(new Event("input", { bubbles: true }));
  }

  // ── Signal inputs for selectedProject / selectedSubproject ────────
  // We create hidden inputs here so Datastar can bind them.
  // They're created once and appended to #app.
  function ensureSignalInput(id, bindAttr, initialValue) {
    if (document.getElementById(id)) return;
    const input = document.createElement("input");
    input.id = id;
    input.type = "text";
    input.style.display = "none";
    input.value = String(initialValue);
    input.setAttribute("data-bind:" + bindAttr, "");
    document.getElementById("app").appendChild(input);
  }

  ensureSignalInput("selected-project-input",    "selected-project",    -1);
  ensureSignalInput("selected-subproject-input", "selected-subproject", -1);

  function setSelectedProject(idx) {
    patchSignal("selected-project-input", idx);
  }

  function setSelectedSubproject(idx) {
    patchSignal("selected-subproject-input", idx);
  }

  // ── Back button HTML ──────────────────────────────────────────────
  function backButtonHtml(label) {
    return "<button class='detail-back' id='detail-back-btn'>← " + escHtml(label) + "</button>";
  }

  // ── Tag chip HTML ─────────────────────────────────────────────────
  function tagsHtml(tags) {
    if (!tags || !tags.length) return "";
    return tags.map(function (t) {
      return "<span class='tag'>" + escHtml(t) + "</span>";
    }).join("");
  }

  function techTagsHtml(tags) {
    if (!tags || !tags.length) return "";
    return tags.map(function (t) {
      return "<span class='tag tag--tech'>" + escHtml(t) + "</span>";
    }).join("");
  }

  // ── Subproject card HTML (used in both detail views) ──────────────
  function subprojectCardHtml(sp) {
    var videoHtml = "";
    if (sp.info && sp.info.video) {
      videoHtml = "<video class='subproject-video' src='" + escHtml(sp.info.video) + "' controls></video>";
    }
    var techTags = techTagsHtml(sp.techTags);
    return "<div class='subproject-card" + (sp.info && sp.info.video ? " subproject-card--big" : "") + "'>" +
      (sp.title ? "<h4 class='subproject-card-title'>" + escHtml(sp.title) + "</h4>" : "") +
      "<p class='subproject-card-desc'>" + escHtml(sp.description) + "</p>" +
      videoHtml +
      (techTags ? "<div class='subproject-card-tags'>" + techTags + "</div>" : "") +
      "</div>";
  }

  // ── "I do" cards ──────────────────────────────────────────────────
  window.renderCards = function (tag) {
    const grid = document.getElementById("card-grid");
    const tmpl = document.getElementById("card-template");
    if (!grid || !tmpl) return;

    // Collect subprojects whose tags include the focused tag,
    // storing a global index so we can navigate to the right one.
    const matches = [];
    data.projects.forEach(function (p) {
      p.subprojects.forEach(function (sp, spIdx) {
        if (sp.tags.indexOf(tag) !== -1) {
          matches.push({ sp: sp, projectTitle: p.title });
        }
      });
    });

    // Build a flat subproject list index so clicking resolves correctly.
    // We store the flat index into window._allSubprojects.
    if (!window._allSubprojects) {
      window._allSubprojects = [];
      data.projects.forEach(function (p) {
        p.subprojects.forEach(function (sp) {
          window._allSubprojects.push({ sp: sp, projectTitle: p.title });
        });
      });
    }

    grid.innerHTML = "";
    matches.forEach(function (item) {
      // Find the flat index for this subproject
      var flatIdx = window._allSubprojects.findIndex(function (a) {
        return a.sp === item.sp;
      });

      const node   = tmpl.content.cloneNode(true);
      const card   = node.querySelector(".card");
      node.querySelector(".card-title").textContent = item.sp.title;
      node.querySelector(".card-desc").textContent  = item.sp.description;
      const footer = node.querySelector(".card-footer");
      item.sp.tags.forEach(function (t) {
        const span       = document.createElement("span");
        span.className   = "tag";
        span.textContent = t;
        footer.appendChild(span);
      });
      if (item.sp.info && item.sp.info.video) {
        const link       = document.createElement("a");
        link.className   = "card-link";
        link.href        = item.sp.info.video;
        link.textContent = "Watch →";
        footer.appendChild(link);
      }

      // Click to open subproject detail
      card.style.cursor = "pointer";
      card.addEventListener("click", function (e) {
        // Don't intercept link clicks
        if (e.target.tagName === "A") return;
        setSelectedSubproject(flatIdx);
      });

      grid.appendChild(node);
    });

    grid.classList.remove("cards-enter");
    void grid.offsetWidth;
    grid.classList.add("cards-enter");
  };

  // ── "I do" subproject detail ──────────────────────────────────────
  window.renderSubprojectDetail = function (flatIdx) {
    const container = document.getElementById("subproject-detail");
    if (!container) return;

    if (!window._allSubprojects) {
      window._allSubprojects = [];
      data.projects.forEach(function (p) {
        p.subprojects.forEach(function (sp) {
          window._allSubprojects.push({ sp: sp, projectTitle: p.title });
        });
      });
    }

    const item = window._allSubprojects[Number(flatIdx)];
    if (!item) return;
    const sp = item.sp;

    var videoHtml = "";
    if (sp.info && sp.info.video) {
      videoHtml = "<video class='detail-video' src='" + escHtml(sp.info.video) + "' controls></video>";
    }

    var techTags = techTagsHtml(sp.techTags);

    container.innerHTML =
      backButtonHtml("Back") +
      "<div class='detail-view'>" +
      "<div class='detail-header'>" +
      "<p class='detail-parent-label'>" + escHtml(item.projectTitle) + "</p>" +
      "<h2 class='detail-title'>" + escHtml(sp.title || "Untitled") + "</h2>" +
      (techTags ? "<div class='detail-tags'>" + techTags + "</div>" : "") +
      "</div>" +
      videoHtml +
      "<p class='detail-desc'>" + escHtml(sp.description) + "</p>" +
      "</div>";

    container.querySelector("#detail-back-btn").addEventListener("click", function () {
      setSelectedSubproject(-1);
    });
  };

  // ── "I worked at" job detail ──────────────────────────────────────
  window.renderJobDetail = function (idx) {
    const container = document.getElementById("job-detail");
    if (!container) return;
    const job = jobs[Number(idx)];
    if (!job) return;
    const exp   = job.specifics.job;
    const end   = exp.dateRange.end || "Present";
    const range = exp.dateRange.start + " – " + end;

    let portraitHtml = "";
    if (exp.portraitImage) {
      portraitHtml = "<img class='job-portrait' src='" + escHtml(exp.portraitImage) + "' alt='" + escHtml(exp.company) + "' />";
    }

    let reviewsHtml = "";
    if (exp.reviews && exp.reviews.length > 0) {
      reviewsHtml = "<h3 class='job-reviews-heading'>Recommendations</h3><div class='job-reviews'>" +
        exp.reviews.map(function (r) {
          return "<div class='review'>" +
            "<div class='review-header'>" +
            "<img class='review-avatar' src='" + escHtml(r.profilePicture) + "' alt='" + escHtml(r.name) + "' />" +
            "<div class='review-meta'>" +
            "<span class='review-name'>"  + escHtml(r.name) + "</span>" +
            "<span class='review-role'>"  + escHtml(r.role) + "</span>" +
            "</div></div>" +
            "<p class='review-text'>" + escHtml(r.text) + "</p>" +
            "</div>";
        }).join("") +
        "</div>";
    }

    let subprojectsHtml = "";
    if (job.subprojects && job.subprojects.length > 0) {
      subprojectsHtml = "<div class='detail-subprojects'>" +
        "<h3 class='detail-subprojects-heading'>Work done</h3>" +
        "<div class='subproject-list'>" +
        job.subprojects.map(subprojectCardHtml).join("") +
        "</div></div>";
    }

    container.innerHTML =
      "<div class='job-detail-hero' style='background-image:url(" + escHtml(exp.backgroundImage) + ")'>" +
      "<div class='job-detail-hero-overlay'></div>" +
      "<div class='job-detail-hero-content'>" +
      portraitHtml +
      "<div class='job-detail-titles'>" +
      "<h2 class='job-detail-company'>" + escHtml(exp.company) + "</h2>" +
      "<p class='job-detail-role'>"     + escHtml(exp.role)    + "</p>" +
      "<p class='job-detail-range'>"    + escHtml(range)       + "</p>" +
      "</div></div></div>" +
      "<div class='job-detail-body'>" +
      "<p class='job-detail-desc'>"     + escHtml(job.description) + "</p>" +
      reviewsHtml +
      subprojectsHtml +
      "</div>";
  };

  // ── "I worked on" grid by category ───────────────────────────────
  (function initWorkedOn() {
    const container = document.getElementById("worked-on-grid");
    if (!container) return;
    const tmpl = document.getElementById("card-template");

    // Group non-job projects by category, skip those without one
    const groups = {};
    const order  = [];
    data.projects.forEach(function (p) {
      if (p.type === "Job" || !p.category) return;
      if (!groups[p.category]) { groups[p.category] = []; order.push(p.category); }
      groups[p.category].push(p);
    });

    order.forEach(function (cat) {
      const heading       = document.createElement("h2");
      heading.className   = "category-heading";
      heading.textContent = cat;
      container.appendChild(heading);

      const grid       = document.createElement("div");
      grid.className   = "card-grid";
      container.appendChild(grid);

      groups[cat].forEach(function (p) {
        // Find the index in nonJobProjects for navigation
        var projIdx = nonJobProjects.indexOf(p);

        const node   = tmpl.content.cloneNode(true);
        const card   = node.querySelector(".card");
        node.querySelector(".card-title").textContent = p.title;
        node.querySelector(".card-desc").textContent  = p.description;
        const footer = node.querySelector(".card-footer");
        p.tags.forEach(function (t) {
          const span       = document.createElement("span");
          span.className   = "tag";
          span.textContent = t;
          footer.appendChild(span);
        });

        // Click to open project detail
        card.style.cursor = "pointer";
        card.addEventListener("click", function () {
          setSelectedProject(projIdx);
        });

        grid.appendChild(node);
      });
    });
  })();

  // ── "I worked on" project detail ─────────────────────────────────
  window.renderProjectDetail = function (idx) {
    const container = document.getElementById("project-detail");
    if (!container) return;
    const p = nonJobProjects[Number(idx)];
    if (!p) return;

    // Reflection section (NonJobExperience)
    var nj = p.specifics.nonJob;
    var reflectionHtml = "";
    if (nj) {
      function listHtml(heading, items) {
        if (!items || !items.length) return "";
        return "<div class='detail-reflection-block'>" +
          "<h4 class='detail-reflection-heading'>" + escHtml(heading) + "</h4>" +
          "<ul class='detail-reflection-list'>" +
          items.filter(function (s) { return s && s.trim(); }).map(function (s) {
            return "<li>" + escHtml(s) + "</li>";
          }).join("") +
          "</ul></div>";
      }

      var reflectionBlocks =
        listHtml("What went well", nj.whatWentWell) +
        listHtml("What could be better", nj.whatCouldBeBetter) +
        listHtml("What I learned", nj.whatILearned);

      if (reflectionBlocks) {
        reflectionHtml = "<div class='detail-reflection'>" + reflectionBlocks + "</div>";
      }

      if (nj.sourceCodeLink) {
        reflectionHtml += "<a class='detail-source-link' href='" + escHtml(nj.sourceCodeLink) + "' target='_blank' rel='noopener'>View source code →</a>";
      }
    }

    // Subprojects section
    var subprojectsHtml = "";
    if (p.subprojects && p.subprojects.length > 0) {
      subprojectsHtml = "<div class='detail-subprojects'>" +
        "<h3 class='detail-subprojects-heading'>Subprojects</h3>" +
        "<div class='subproject-list'>" +
        p.subprojects.map(subprojectCardHtml).join("") +
        "</div></div>";
    }

    // Type badge
    var typeBadgeHtml = "<span class='detail-type-badge detail-type-badge--" + p.type.toLowerCase() + "'>" + escHtml(p.type) + "</span>";

    container.innerHTML =
      backButtonHtml("Back to projects") +
      "<div class='detail-view'>" +
      "<div class='detail-header'>" +
      typeBadgeHtml +
      "<h2 class='detail-title'>" + escHtml(p.title) + "</h2>" +
      (p.tags && p.tags.length ? "<div class='detail-tags'>" + tagsHtml(p.tags) + "</div>" : "") +
      "</div>" +
      "<p class='detail-desc'>" + escHtml(p.description) + "</p>" +
      reflectionHtml +
      subprojectsHtml +
      "</div>";

    container.querySelector("#detail-back-btn").addEventListener("click", function () {
      setSelectedProject(-1);
    });

    // Scroll to top of content column when navigating into detail
    var col = container.closest(".col-content");
    if (col) col.scrollTop = 0;
  };

  // ── "I" bio ───────────────────────────────────────────────────────
  (function initBio() {
    const el = document.getElementById("bio-text");
    if (!el) return;
    const paragraphs = config.bioText.split("\n\n");
    el.innerHTML = paragraphs.map(function (p) {
      return "<p>" + escHtml(p) + "</p>";
    }).join("");
  })();

  // ── Reset detail state when mode changes ──────────────────────────
  // Watch for nav clicks so going to a different mode clears the detail.
  document.getElementById("nav-ticker").addEventListener("click", function () {
    setSelectedProject(-1);
    setSelectedSubproject(-1);
  });

})();
