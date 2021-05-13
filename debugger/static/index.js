import { newPalette } from "./palette.js";
import { newCPU } from "./cpu.js";
import { newSprites } from "./sprites.js";
import { newRegister } from "./register.js";
import { newTabManager } from "./tab_manager.js";
import { newMemory } from "./memory.js";


//tab management
const cpuTab = newCPU();
const paletteTab = newPalette();
const spritesTab = newSprites();
const registerTab = newRegister();
const memoryTab = newMemory();

const tabManager = newTabManager();
tabManager.setTabs([
    {
        "name": "CPU",
        "component": cpuTab,
    },
    {
        "name": "Palette",
        "component": paletteTab,
    },
    {
        "name": "Sprites",
        "component": spritesTab,
    },
    {
        "name": "Registers",
        "component": registerTab,
    }, {
        "name": "Memory",
        "component": memoryTab,
    }
]);

const root = document.getElementById("root");
root.appendChild(tabManager);

const resumeButton = document.getElementById("resume_button");
resumeButton.onclick = function () {
    fetch('/resume').then(resp => resp.json()).then(displayState);
}

const pauseButton = document.getElementById("pause_button");
pauseButton.onclick = function () {
    fetch('/pause').then(resp => resp.json()).then(displayState);
}

const stepButton = document.getElementById("step_button");
stepButton.onclick = function () {
    const count = document.getElementById("count");
    fetch('/step?count=' + count.value)
        .then(resp => resp.json())
        .then(displayState)
}

const breakpointButton = document.getElementById("breakpoint_button");
breakpointButton.onclick = function () {
    const address = document.getElementById("breakpoint");
    fetch('/breakpoint?address=' + address.value);
}
const clearBreakpointButton = document.getElementById("clear_breakpoint_button");
clearBreakpointButton.onclick = function () {
    const address = document.getElementById("breakpoint");
    address.value = "";
    fetch('/breakpoint?clear=address');
}

const registerBreakpointButton = document.getElementById("register_breakpoint_button");
registerBreakpointButton.onclick = function () {
    const register = document.getElementById("register_breakpoint");
    fetch('/breakpoint?registers=' + register.value);
}
const clearRegisterBreakpointButton = document.getElementById("clear_register_breakpoint_button");
clearRegisterBreakpointButton.onclick = function () {
    const register = document.getElementById("register_breakpoint");
    register.value = ""
    fetch('/breakpoint?clear=registers');
}

function displayState(body) {
    cpuTab.addEntry(body.cpu);
    paletteTab.updatePalette(body.palette);
    spritesTab.updateSprites(body.sprites);
    memoryTab.updateRam(body.ram);
    memoryTab.updateVRam(body.vram);
    if (body.register) {
        registerTab.addData(body.register);
    }
}
