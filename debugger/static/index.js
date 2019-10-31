import { newPPU } from "./ppu.js";
import { newCPU } from "./cpu.js";
import { newTabManager } from "./tab_manager.js";


//tab management
const cpuTab = newCPU();

const ppuTab = newPPU();

const tabManager = newTabManager();
tabManager.setTabs([
    {
        "name": "CPU",
        "component": cpuTab,
    },
    {
        "name": "PPU",
        "component": ppuTab,
    }
]);

const root = document.getElementById("root");
root.appendChild(tabManager);

const pauseButton = document.getElementById("pause_button");
pauseButton.onclick = function() {
    fetch('/pause').then(resp => resp.json()).then(displayState);
}

const stepButton = document.getElementById("step_button");
stepButton.onclick = function() {
    const count = document.getElementById("count");
    fetch('/step?count='+count.value)
        .then(resp => resp.json())
        .then(displayState)
}

const breakpointButton = document.getElementById("breakpoint_button");
breakpointButton.onclick = function() {
    const address = document.getElementById("breakpoint");
    fetch('/breakpoint?address='+address.value);
}


function displayState(body) {
    cpuTab.addState(body.cpu);
    ppuTab.updatePalette(body.palette);
}
