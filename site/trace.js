const steps = [
  { state: "pending", tone: "", line: "Waiting for the first attempt.", fill: 8, note: "Row written for endpoint lab." },
  { state: "attempt", tone: "", line: "POST is in flight.", fill: 36, note: "Claimed for 30s so a second drainer skips it." },
  { state: "retry", tone: "", line: "Endpoint timed out. Backoff scheduled.", fill: 58, note: "Still pending. Attempt 1 of 5." },
  { state: "delivered", tone: "ok", line: "The endpoint accepted the body.", fill: 100, note: "Signed header checked against the raw body." },
  { state: "failed", tone: "bad", line: "Five attempts. This delivery stops.", fill: 100, note: "A later event starts a new row." }
];

const legend = document.getElementById("legend");
const stateEl = document.getElementById("trace-state");
const lineEl = document.getElementById("trace-line");
const dot = document.getElementById("trace-dot");
const fill = document.getElementById("rail-fill");
const attempts = document.getElementById("attempts");

steps.forEach((step) => {
  const li = document.createElement("li");
  li.dataset.state = step.state;
  li.innerHTML = `<b>${step.state}</b><span>${step.note}</span>`;
  legend.appendChild(li);
});

let i = 0;
function paint() {
  const step = steps[i];
  stateEl.textContent = step.state;
  lineEl.textContent = step.line;
  fill.style.width = step.fill + "%";
  fill.style.background = step.tone === "bad" ? "var(--fail)" : "var(--signal)";
  dot.className = "dot " + step.tone;
  legend.querySelectorAll("li").forEach((li) => {
    li.classList.toggle("on", li.dataset.state === step.state);
  });
  const item = document.createElement("li");
  item.textContent = `${String(i + 1).padStart(2, "0")}  ${step.state} — ${step.note}`;
  attempts.appendChild(item);
  while (attempts.children.length > 4) attempts.removeChild(attempts.firstChild);
  i = (i + 1) % steps.length;
  if (i === 0) attempts.replaceChildren();
}

paint();
const reduce = window.matchMedia("(prefers-reduced-motion: reduce)").matches;
if (!reduce) setInterval(paint, 2200);

async function copyText(text) {
  if (navigator.clipboard && window.isSecureContext) {
    try {
      await navigator.clipboard.writeText(text);
      return true;
    } catch {
      /* fall through to the selection copy */
    }
  }
  const area = document.createElement("textarea");
  area.value = text;
  area.setAttribute("readonly", "");
  area.style.position = "fixed";
  area.style.left = "-9999px";
  document.body.appendChild(area);
  area.select();
  let ok = false;
  try {
    ok = document.execCommand("copy");
  } catch {
    ok = false;
  }
  area.remove();
  return ok;
}

document.querySelectorAll("[data-copy]").forEach((button) => {
  button.addEventListener("click", async () => {
    const code = document.getElementById(button.dataset.copy);
    const ok = await copyText(code.textContent);
    button.textContent = ok ? "Copied" : "Select the block";
    button.classList.toggle("done", ok);
    setTimeout(() => {
      button.textContent = "Copy";
      button.classList.remove("done");
    }, 3000);
  });
});
