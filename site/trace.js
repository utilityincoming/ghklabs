const steps = [
  { state: "molecule", tone: "", line: "Three amino acids plus copper.", fill: 12, note: "Copper Tripeptide-1 on a cosmetic label. GHK-Cu when identity is the point." },
  { state: "cells", tone: "", line: "Maquart et al., 1988 — more collagen synthesis.", fill: 40, note: "Measured in cultured fibroblasts. A cell result, not a result on a face." },
  { state: "people", tone: "", line: "A later human study, small and mixed.", fill: 64, note: "Set after laser resurfacing. Satisfaction moved; several blinded measures did not." },
  { state: "this serum", tone: "ok", line: "What this bottle claims.", fill: 100, note: "The look of fine lines, the feel of moisturized skin. Nothing more." },
  { state: "not a drug", tone: "bad", line: "Where the serum stops.", fill: 100, note: "Not a treatment, a wound product, or a dose. Those stay on the paper, labeled as papers." }
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
