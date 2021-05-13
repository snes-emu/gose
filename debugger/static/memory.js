class Memory extends HTMLDivElement {
    static tagName() {
        return 'memory-div';
    }

    constructor() {
        super();

        this.id = "memory";

        this.style="display: flex";

        const ramContainer = document.createElement("div");
        ramContainer.innerHTML = "<h>RAM</h>";
        this.appendChild(ramContainer);

        this.ram = document.createElement("pre");
        ramContainer.appendChild(this.ram);
        this.ram.style.height = "500px";
        this.ram.style.overflow_y = "scroll";

        const vramContainer = document.createElement("div");
        vramContainer.innerHTML = "<h>VRAM</h>"
        this.appendChild(vramContainer);

        this.vram = document.createElement("pre");
        vramContainer.appendChild(this.vram);
        this.vram.style.height = "500px";
        this.vram.style.overflow_y = "scroll";
    }

    updateRam(memory) {
        let buffer = Uint8Array.from(atob(memory).split('').map(c => c.charCodeAt()));
        let text = "";
        for (let i = 0; i < buffer.length / 0x10; i++) {
            text += `${toHex(i * 0x10, 6)} `;
            for (let j = 0; j < 0x10; j++) {
                text += ` ${toHex(buffer[i * 0x10 + j], 2)}`;
            }
            text += "\n";
        }
        this.ram.innerText = text;
    }

    updateVRam(vram) {
        let buffer = Uint8Array.from(atob(vram).split('').map(c => c.charCodeAt()));
        let text = "";
        for (let i = 0; i < buffer.length / 0x10; i++) {
            text += `${toHex(i * 0x10, 4)} `;
            for (let j = 0; j < 0x10; j++) {
                text += ` ${toHex(buffer[i * 0x10 + j], 2)}`;
            }
            text += "\n";
        }
        this.vram.innerText = text;
    }
}


function toHex(n, digitNumber) {
    let padding = "";
    for (let i = 0; i < digitNumber; i++) {
        padding += "0";
    }
    return (padding + n.toString(16)).substr(-digitNumber)
}

customElements.define(Memory.tagName(), Memory, { extends: 'div' });

export function newMemory() {
    return document.createElement('div', { is: Memory.tagName() })
}
