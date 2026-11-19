const steps = [
  { state: "molecule", tone: "", line: "Three amino acids plus copper.", fill: 25, note: "GHK-Cu. On the label the name is Copper Tripeptide-1." },
  { state: "percent", tone: "", line: "The concentration is printed.", fill: 50, note: "{percent} on the label. The figure, in plain type." },
  { state: "serum", tone: "ok", line: "A few drops, morning and evening.", fill: 75, note: "For the look of fine lines and the feel of moisturized skin." },
  { state: "lot", tone: "ok", line: "The certificate matches the glass.", fill: 100, note: "Lot {lot}. Identity by {method}. Purity {purity}." }
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
