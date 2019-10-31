const paletteNum = 16
const paletteSize = 16


//tab management
let container = document.getElementById('tab_container');

let cpuTab = document.createElement("ul");
cpuTab.id = 'cpu';

let ppuTab = document.createElement('div');
ppuTab.id = 'palette';
for (let p = 0; p < paletteNum ; p++) {
    const row = document.createElement("div");
    row.appendChild(document.createTextNode(`${p}: `))
    ppuTab.appendChild(row);
    for (let c = 0; c < paletteSize; c++) {
        ppuTab.children[p].appendChild(document.createElement("div"))
    }
}


let tabs = document.querySelectorAll('input[name=tab]');
tabs.forEach(tab => tab.onchange= event => {
    Array.from(container.children).forEach(child => container.removeChild(child));

    switch (event.target.value) {
        case 'cpu':
            container.appendChild(cpuTab);
            break;
        case 'ppu':
            container.appendChild(ppuTab);
            break;
    }

});

//mount initial tab
container.appendChild(cpuTab);

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

    const li = document.createElement("li");
    li.appendChild(document.createTextNode(Object.keys(cpu).map(key => `${key}: ${cpu[key]}`).join(',')));
    cpuTab.appendChild(li);

    body.palette.forEach((color,i) => {
        const rgb = `rgb(${color.r},${color.g},${color.b})`;
        ppuTab.children[Math.floor(i/paletteNum)].children[i%paletteSize].style = `background-color:${rgb}`;
    });
}


function breakpoint() {
    const address = document.getElementById("breakpoint");
    fetch('/breakpoint?address='+address.value);
}