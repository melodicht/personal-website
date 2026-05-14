// worked-at.js — snap job list to the focused job
// All signal logic (jobFocus updates, active class) handled by Datastar in the HTML.
window.snapToJob = function (idx) {
  const list = document.getElementById("job-list");
  if (!list) return;
  const el = list.querySelector("[data-on-click*='jobFocus = " + idx + "']");
  if (!el) return;
  const listRect = list.getBoundingClientRect();
  const elRect   = el.getBoundingClientRect();
  const lineHeight = 24;
  const targetMid  = elRect.top - listRect.top + list.scrollTop - lineHeight / 2;
  const target     = targetMid - list.clientHeight / 2;
  list.scrollTo({ top: target, behavior: "smooth" });
};
