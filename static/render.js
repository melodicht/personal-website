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

  function patchSignal(id, value) {
    const el = document.getElementById(id);
    if (!el) return;
    el.value = typeof value === "number" ? String(value) : value;
    el.dispatchEvent(new Event("input", { bubbles: true }));
  }

  // ── Signal inputs ─────────────────────────────────────────────────
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

  function setSelectedProject(idx)    { patchSignal("selected-project-input",    idx); }
  function setSelectedSubproject(idx) { patchSignal("selected-subproject-input", idx); }

  // ── Back button ───────────────────────────────────────────────────
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

  // ── Inline description parsing ────────────────────────────────────
  function parseDescription(str) {
    var segments = [];
    var re = /\{([^}]+)\}/g;
    var last = 0, match;
    while ((match = re.exec(str)) !== null) {
      if (match.index > last) {
        segments.push({ type: "text", value: str.slice(last, match.index) });
      }
      segments.push({ type: "tag", value: match[1] });
      last = re.lastIndex;
    }
    if (last < str.length) {
      segments.push({ type: "text", value: str.slice(last) });
    }
    return segments;
  }

  function renderDescription(str, allTechTags) {
    var known = allTechTags || [];
    return parseDescription(str).map(function (seg) {
      if (seg.type === "tag") {
        var lower = seg.value.toLowerCase();
        var canonical = known.find(function (t) { return t.toLowerCase() === lower; }) || seg.value;
        return "<span class='tag tag--tech tag--inline'>" + escHtml(canonical) + "</span>";
      }
      return escHtml(seg.value);
    }).join("");
  }

  function inlinedTags(str) {
    return parseDescription(str)
      .filter(function (seg) { return seg.type === "tag"; })
      .map(function (seg) { return seg.value.toLowerCase(); });
  }

  // ── Tag inheritance helpers ───────────────────────────────────────
  // Merges arrays deduped, with earlier arrays taking priority.
  function mergeTags(a, b) {
    return a.concat((b || []).filter(function (t) { return a.indexOf(t) === -1; }));
  }

  // Returns the effective tech tags for a subproject given its inheritance chain.
  // Filters out inlined tags from the description.
  function effectiveTechTags(sp, inheritedTechTags) {
    var inherited = inheritedTechTags || [];
    var inlined   = inlinedTags(sp.description);
    return mergeTags(inherited, sp.techTags || []).filter(function (t) {
      return inlined.indexOf(t.toLowerCase()) === -1;
    });
  }

  // Returns only the subproject's own non-inherited, non-inlined tech tags
  // (for display on cards where inherited ones are already implied by context).
  function ownTechTags(sp, inheritedTechTags) {
    var inherited = inheritedTechTags || [];
    var inlined   = inlinedTags(sp.description);
    return (sp.techTags || []).filter(function (t) {
      return inherited.indexOf(t) === -1 && inlined.indexOf(t.toLowerCase()) === -1;
    });
  }

  // ── Flatten subprojects from subsections ─────────────────────────
  // Returns a flat array of { sp, subsection, project } for all subprojects
  // in a project, regardless of whether they are bullets, cards, or major.
  function flattenSubprojects(project) {
    var result = [];
    (project.subsections || []).forEach(function (sec) {
      (sec.bullets || []).forEach(function (b) {
        result.push({ sp: b.subproject, subsection: sec, project: project });
      });
      (sec.cards || []).forEach(function (c) {
        result.push({ sp: c.subproject, subsection: sec, project: project });
      });
      if (sec.major) {
        result.push({ sp: sec.major.subproject, video: sec.major.video, subsection: sec, project: project });
      }
    });
    return result;
  }

  // ── Card renderer (for subproject cards) ─────────────────────────
  function subprojectCardHtml(sp, inheritedTechTags) {
    var own      = ownTechTags(sp, inheritedTechTags);
    var techTags = techTagsHtml(own);
    return "<div class='subproject-card'>" +
      (sp.title ? "<h4 class='subproject-card-title'>" + escHtml(sp.title) + "</h4>" : "") +
      (techTags ? "<div class='subproject-card-tags'>" + techTags + "</div>" : "") +
      "<p class='subproject-card-desc'>" + renderDescription(sp.description, mergeTags(inheritedTechTags || [], sp.techTags || [])) + "</p>" +
      "</div>";
  }

  // ── Bullet point renderer ─────────────────────────────────────────
  function bulletPointHtml(sp, inheritedTechTags) {
    var own        = ownTechTags(sp, inheritedTechTags);
    var allKnown   = mergeTags(inheritedTechTags || [], sp.techTags || []);
    var techChips  = own.length ? " " + own.map(function (t) {
      return "<span class='tag tag--tech tag--inline'>" + escHtml(t) + "</span>";
    }).join(" ") : "";
    return "<li class='subsection-bullet'>" +
      "<p class='subsection-bullet-text'>" +
      renderDescription(sp.description, allKnown) +
      techChips +
      "</p>" +
      "</li>";
  }

  // ── Major subproject renderer ─────────────────────────────────────
  function majorSubprojectHtml(sp, video, inheritedTechTags) {
    var eff      = effectiveTechTags(sp, inheritedTechTags);
    var techTags = techTagsHtml(eff);
    var allKnown = mergeTags(inheritedTechTags || [], sp.techTags || []);
    var videoHtml = video
      ? "<video class='major-subproject-video' src='" + escHtml(video) + "' controls></video>"
      : "";
    return "<div class='major-subproject'>" +
      (sp.title ? "<h4 class='major-subproject-title'>" + escHtml(sp.title) + "</h4>" : "") +
      (techTags ? "<div class='major-subproject-tags'>" + techTags + "</div>" : "") +
      "<p class='major-subproject-desc'>" + renderDescription(sp.description, allKnown) + "</p>" +
      videoHtml +
      "</div>";
  }

  // ── Subsection renderer ───────────────────────────────────────────
  function subsectionHtml(sec, projectTechTags) {
    var inheritedTechTags = mergeTags(projectTechTags || [], sec.techTags || []);

    var contentHtml = "";
    if (sec.bullets && sec.bullets.length) {
      contentHtml = "<ul class='subsection-bullet-list'>" +
        sec.bullets.map(function (b) {
          return bulletPointHtml(b.subproject, inheritedTechTags);
        }).join("") +
        "</ul>";
    } else if (sec.cards && sec.cards.length) {
      contentHtml = "<div class='subproject-list'>" +
        sec.cards.map(function (c) {
          return subprojectCardHtml(c.subproject, inheritedTechTags);
        }).join("") +
        "</div>";
    } else if (sec.major) {
      contentHtml = majorSubprojectHtml(sec.major.subproject, sec.major.video, inheritedTechTags);
    }

    var secTechTagsHtml = techTagsHtml(sec.techTags);
    return "<div class='detail-subsection'>" +
      "<div class='detail-subprojects-heading'>" +
      escHtml(sec.title) +
      (secTechTagsHtml ? " <span class='subsection-heading-tags'>" + secTechTagsHtml + "</span>" : "") +
      "</div>" +
      contentHtml +
      "</div>";
  }

  // ── Build subsections HTML for a project ─────────────────────────
  function projectSubsectionsHtml(p) {
    if (!p.subsections || !p.subsections.length) return "";
    return p.subsections.map(function (sec) {
      return subsectionHtml(sec, p.techTags || []);
    }).join("");
  }

  // ── "I do" cards ──────────────────────────────────────────────────
  window.renderCards = function (tag) {
    const grid = document.getElementById("card-grid");
    const tmpl = document.getElementById("card-template");
    if (!grid || !tmpl) return;

    // Build _allSubprojects flat list if not yet done
    if (!window._allSubprojects) {
      window._allSubprojects = [];
      data.projects.forEach(function (p) {
        flattenSubprojects(p).forEach(function (item) {
          window._allSubprojects.push(item);
        });
      });
    }

    // Group matching subprojects by project
    const groups = [];
    const groupIndex = {};
    data.projects.forEach(function (p) {
      var projTags     = p.tags || [];
      var projTechTags = p.techTags || [];
      flattenSubprojects(p).forEach(function (item) {
        var secTags = item.subsection.tags || [];
        var spTags  = item.sp.tags || [];
        var effectiveTags = mergeTags(projTags, mergeTags(secTags, spTags));
        if (effectiveTags.indexOf(tag) === -1) return;
        if (groupIndex[p.title] === undefined) {
          groupIndex[p.title] = groups.length;
          groups.push({ project: p, items: [] });
        }
        groups[groupIndex[p.title]].items.push(item);
      });
    });

    grid.innerHTML = "";

    groups.forEach(function (group) {
      var p = group.project;
      var isJob  = p.type === "Job";
      var navIdx = isJob ? jobs.indexOf(p) : nonJobProjects.indexOf(p);

      // ── Group header ──
      const header = document.createElement("div");
      header.className = "card-group-header";
      header.innerHTML =
        "<span class='card-group-title'>" + escHtml(p.title) + "</span>" +
        "<span class='card-group-type detail-type-badge detail-type-badge--" + p.type.toLowerCase() + "'>" + escHtml(p.type) + "</span>" +
        (p.techTags && p.techTags.length ? techTagsHtml(p.techTags) : "") +
        "<span class='card-group-link'>View project →</span>";

      header.style.cursor = "pointer";
      header.addEventListener("click", function () {
        if (isJob) {
          patchSignal("mode-input", "worked-at");
          patchSignal("job-focus-input", navIdx);
        } else {
          patchSignal("mode-input", "worked-on");
          patchSignal("selected-project-input", navIdx);
        }
        document.querySelectorAll(".nav-item").forEach(function (el) {
          el.classList.toggle("nav-item--active", el.dataset.mode === (isJob ? "worked-at" : "worked-on"));
        });
      });

      grid.appendChild(header);

      // ── Cards row ──
      const row = document.createElement("div");
      row.className = "card-grid card-group-grid";
      grid.appendChild(row);

      group.items.forEach(function (item) {
        var flatIdx = window._allSubprojects.findIndex(function (a) {
          return a.sp === item.sp;
        });

        var inheritedTechTags = mergeTags(p.techTags || [], item.subsection.techTags || []);
        var own     = ownTechTags(item.sp, inheritedTechTags);
        var inlined = inlinedTags(item.sp.description);
        var allKnown = mergeTags(inheritedTechTags, item.sp.techTags || []);

        const node = tmpl.content.cloneNode(true);
        const card = node.querySelector(".card");
        node.querySelector(".card-title").textContent = item.sp.title;
        node.querySelector(".card-desc").innerHTML = renderDescription(item.sp.description, allKnown);

        if (own.length) {
          const tagsRow     = document.createElement("div");
          tagsRow.className = "card-tech-tags";
          own.forEach(function (t) {
            const span       = document.createElement("span");
            span.className   = "tag tag--tech";
            span.textContent = t;
            tagsRow.appendChild(span);
          });
          const body = node.querySelector(".card-body");
          const desc = node.querySelector(".card-desc");
          body.insertBefore(tagsRow, desc);
        }

        card.style.cursor = "pointer";
        card.addEventListener("click", function (e) {
          if (e.target.tagName === "A") return;
          setSelectedSubproject(flatIdx);
        });

        row.appendChild(node);
      });
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
        flattenSubprojects(p).forEach(function (item) {
          window._allSubprojects.push(item);
        });
      });
    }

    const item = window._allSubprojects[Number(flatIdx)];
    if (!item) return;
    const sp = item.sp;

    var inheritedTechTags = mergeTags(item.project.techTags || [], item.subsection.techTags || []);
    var eff      = effectiveTechTags(sp, inheritedTechTags);
    var techTags = techTagsHtml(eff);
    var allKnown = mergeTags(inheritedTechTags, sp.techTags || []);

    var videoHtml = item.video
      ? "<video class='detail-video' src='" + escHtml(item.video) + "' controls></video>"
      : "";

    container.innerHTML =
      backButtonHtml("Back") +
      "<div class='detail-view'>" +
      "<div class='detail-header'>" +
      "<p class='detail-parent-label'>" + escHtml(item.project.title) + "</p>" +
      "<h2 class='detail-title'>" + escHtml(sp.title || "Untitled") + "</h2>" +
      (techTags ? "<div class='detail-tags'>" + techTags + "</div>" : "") +
      "</div>" +
      videoHtml +
      "<p class='detail-desc'>" + renderDescription(sp.description, allKnown) + "</p>" +
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

    document.querySelectorAll(".job-cell").forEach(function (el) {
      el.classList.toggle("job-cell--active", Number(el.dataset.idx) === Number(idx));
    });
    if (window._snapToJob) window._snapToJob(Number(idx));

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
      ((job.tags && job.tags.length) || (job.techTags && job.techTags.length) ?
        "<div class='detail-tags'>" + tagsHtml(job.tags) + techTagsHtml(job.techTags) + "</div>" : "") +
      projectSubsectionsHtml(job) +
      reviewsHtml +
      "</div>";
  };

  // ── "I worked on" grid by category ───────────────────────────────
  (function initWorkedOn() {
    const container = document.getElementById("worked-on-grid");
    if (!container) return;
    const tmpl = document.getElementById("card-template");

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

    var typeBadgeHtml = "<span class='detail-type-badge detail-type-badge--" + p.type.toLowerCase() + "'>" + escHtml(p.type) + "</span>";

    container.innerHTML =
      backButtonHtml("Back to projects") +
      "<div class='detail-view'>" +
      "<div class='detail-header'>" +
      typeBadgeHtml +
      "<h2 class='detail-title'>" + escHtml(p.title) + "</h2>" +
      (p.tags && p.tags.length ? "<div class='detail-tags'>" + tagsHtml(p.tags) + "</div>" : "") +
      (p.techTags && p.techTags.length ? "<div class='detail-tags'>" + techTagsHtml(p.techTags) + "</div>" : "") +
      "</div>" +
      "<p class='detail-desc'>" + escHtml(p.description) + "</p>" +
      projectSubsectionsHtml(p) +
      reflectionHtml +
      "</div>";

    container.querySelector("#detail-back-btn").addEventListener("click", function () {
      setSelectedProject(-1);
    });

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

  // ── Reset detail state on mode change ────────────────────────────
  document.getElementById("nav-ticker").addEventListener("click", function () {
    setSelectedProject(-1);
    setSelectedSubproject(-1);
  });

})();
