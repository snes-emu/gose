const paletteNum = 16
const paletteSize = 16
let palette = document.getElementById("palette");

for (let p = 0; p < paletteNum ; p++) {
    let row = document.createElement("div");
    row.appendChild(document.createTextNode(`${p}: `))
    palette.appendChild(row);
    for (let c = 0; c < paletteSize; c++) {
        palette.children[p].appendChild(document.createElement("div"))
    }
}

function pause() {
    fetch('/pause');
}

function step() {
    let count = document.getElementById("count");
    fetch('/step?count='+count.value)
        .then(resp => resp.json())
        .then(body => {
            let cpu = body.cpu;

            let lu = document.getElementById("cpu");
            let li = document.createElement("li");
            li.appendChild(document.createTextNode(Object.keys(cpu).map(key => `${key}: ${cpu[key]}`).join(',')));
            lu.appendChild(li);

            body.palette.forEach((color,i) => {
                let rgb = `rgb(${color.r},${color.g},${color.b})`;
                palette.children[Math.floor(i/paletteNum)].children[i%paletteSize].style = `background-color:${rgb}`;
            });
        })
}


function breakpoint() {
    const address = document.getElementById("breakpoint");
    fetch('/breakpoint?address='+address.value);
}