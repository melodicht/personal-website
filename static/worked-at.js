// worked-at.js — scroll the active job cell into the centre of the list on load
(function () {
  const list = document.getElementById("job-list");
  const active = list && list.querySelector(".job-cell--active");
  if (!list || !active) return;

  const listRect = list.getBoundingClientRect();
  const elRect = active.getBoundingClientRect();
  const lineHeight = 24;
  const targetMid = elRect.top - listRect.top + list.scrollTop - lineHeight / 2;
  const target = targetMid - list.clientHeight / 2;
  list.scrollTop = Math.max(0, target);
})();
