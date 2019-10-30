const paletteNum = 16
const paletteSize = 16
let palette = document.getElementById("palette");

for (let p = 0; p < paletteNum ; p++) {
    const row = document.createElement("div");
    row.appendChild(document.createTextNode(`${p}: `))
    palette.appendChild(row);
    for (let c = 0; c < paletteSize; c++) {
        palette.children[p].appendChild(document.createElement("div"))
    }
}

function pause() {
    fetch('/pause').then(resp => resp.json()).then(displayState);
}

function step() {
    const count = document.getElementById("count");
    fetch('/step?count='+count.value)
        .then(resp => resp.json())
        .then(displayState)
}

function displayState(body) {
    const cpu = body.cpu;

    const lu = document.getElementById("cpu");
    const li = document.createElement("li");
    li.appendChild(document.createTextNode(Object.keys(cpu).map(key => `${key}: ${cpu[key]}`).join(',')));
    lu.appendChild(li);

    body.palette.forEach((color,i) => {
        const rgb = `rgb(${color.r},${color.g},${color.b})`;
        palette.children[Math.floor(i/paletteNum)].children[i%paletteSize].style = `background-color:${rgb}`;
    });
}


function breakpoint() {
    const address = document.getElementById("breakpoint");
    fetch('/breakpoint?address='+address.value);
}