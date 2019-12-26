class Memory extends HTMLDivElement {
    static tagName() {
        return 'memory-div';
    }

    constructor() {
        super();

        this.id = "memory";
        this.content = document.createElement("pre");
        this.appendChild(this.content);

        this.content.style.height = "500px";
        this.content.style.overflow_y = "scroll";
    }

    clear() {
        this.content.innerText = "";
    }

    updateMemory(memory) {
        this.clear();
        console.log(memory);
        let buffer = Uint8Array.from(atob(memory).split('').map(c => c.charCodeAt()));
        let text = "";
        for (let i = 0; i < buffer.length / 0x10; i++) {
            text += `${toHex(i * 0x10, 6)} `;
            for (let j = 0; j < 0x10; j++) {
                text += ` ${toHex(buffer[i * 0x10 + j], 2)}`;
            }
            text += "\n";
        }
        this.content.innerText = text;
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
