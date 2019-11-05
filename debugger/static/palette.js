class Palette extends HTMLDivElement {
    static tagName(){
        return 'palette-div';
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

customElements.define(Palette.tagName(), Palette, {extends: 'div'});

export function newPalette() {
    return document.createElement('div', {is: Palette.tagName()})
}
