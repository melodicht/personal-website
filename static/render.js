// render.js — all content-column rendering functions
// Called by Datastar data-effect expressions.
(function () {
  const data   = window.SITE_DATA;
  const config = window.SITE_CONFIG;
  const jobs   = data.projects.filter(function (p) { return p.type === "Job"; });

  function escHtml(s) {
    return String(s)
      .replace(/&/g, "&amp;").replace(/</g, "&lt;")
      .replace(/>/g, "&gt;").replace(/"/g, "&quot;");
  }

  function isExternal(url) {
    return typeof url === "string" && url.startsWith("http");
  }

  // ── "I do" cards ──────────────────────────────────────────────────
  window.renderCards = function (tag) {
    const grid = document.getElementById("card-grid");
    const tmpl = document.getElementById("card-template");
    if (!grid || !tmpl) return;

    // Collect subprojects whose tags include the focused tag
    const matches = [];
    data.projects.forEach(function (p) {
      p.subprojects.forEach(function (sp) {
        if (sp.tags.indexOf(tag) !== -1) matches.push(sp);
      });
    });

    grid.innerHTML = "";
    matches.forEach(function (sp) {
      const node   = tmpl.content.cloneNode(true);
      node.querySelector(".card-title").textContent = sp.title;
      node.querySelector(".card-desc").textContent  = sp.description;
      const footer = node.querySelector(".card-footer");
      sp.tags.forEach(function (t) {
        const span       = document.createElement("span");
        span.className   = "tag";
        span.textContent = t;
        footer.appendChild(span);
      });
      if (sp.info && sp.info.video) {
        const link       = document.createElement("a");
        link.className   = "card-link";
        link.href        = sp.info.video;
        link.textContent = "Watch →";
        footer.appendChild(link);
      }
      grid.appendChild(node);
    });

    grid.classList.remove("cards-enter");
    void grid.offsetWidth;
    grid.classList.add("cards-enter");
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
        const node   = tmpl.content.cloneNode(true);
        node.querySelector(".card-title").textContent = p.title;
        node.querySelector(".card-desc").textContent  = p.description;
        const footer = node.querySelector(".card-footer");
        p.tags.forEach(function (t) {
          const span       = document.createElement("span");
          span.className   = "tag";
          span.textContent = t;
          footer.appendChild(span);
        });
        grid.appendChild(node);
      });
    });
  })();

  // ── "I" bio ───────────────────────────────────────────────────────
  (function initBio() {
    const el = document.getElementById("bio-text");
    if (!el) return;
    const paragraphs = config.bioText.split("\n\n");
    el.innerHTML = paragraphs.map(function (p) {
      return "<p>" + escHtml(p) + "</p>";
    }).join("");
  })();

})();
