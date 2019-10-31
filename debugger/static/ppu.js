class PPU extends HTMLDivElement {
    static tagName(){
        return 'ppu-div';
    }

    constructor() {
        super();

        this.paletteNum = 16;
        this.paletteSize = 16;
        this.id = "palette";

        for (let p = 0; p < this.paletteNum ; p++) {
            const row = document.createElement("div");
            row.appendChild(document.createTextNode(`${p}: `))
            this.appendChild(row);
            for (let c = 0; c < this.paletteSize; c++) {
                this.children[p].appendChild(document.createElement("div"))
            }
        }
    }

    updatePalette(palette) {
        palette.forEach((color,i) => {
            const rgb = `rgb(${color.r},${color.g},${color.b})`;
            this.children[Math.floor(i/this.paletteNum)].children[i%this.paletteSize].style = `background-color:${rgb}`;
        });
    }
}

customElements.define(PPU.tagName(), PPU, {extends: 'div'});

export function newPPU() {
    return document.createElement('div', {is: PPU.tagName()})
}
