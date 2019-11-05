class Palette extends HTMLTableElement {
    static tagName(){
        return 'palette-table';
    }

    constructor() {
        super();

        this.paletteNum = 16;
        this.paletteSize = 16;
        this.id = "palette";

        this.body = document.createElement("tbody");
        this.appendChild(this.body);

        for (let p = 0; p < this.paletteNum ; p++) {
            const row = document.createElement("tr");
            this.body.appendChild(row);
            for (let c = 0; c < this.paletteSize; c++) {
                row.appendChild(document.createElement("td"));
            }
        }
    }

    updatePalette(palette) {
        palette.forEach((color,i) => {
            const rgb = `rgb(${color.r >> 8},${color.g >> 8},${color.b >> 8})`;
            this.body.children[Math.floor(i/this.paletteNum)].children[i%this.paletteSize].style = `background-color:${rgb}`;
        });
    }
}

customElements.define(Palette.tagName(), Palette, {extends: 'table'});

export function newPalette() {
    return document.createElement('table', {is: Palette.tagName()})
}
