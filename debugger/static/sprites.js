class Sprites extends HTMLDivElement {
    static tagName() {
        return 'sprites-div';
    }

    constructor() {
        super();

        this.id = "sprites";
    }

    clear() {
        while (this.firstChild) {
            this.removeChild(this.firstChild);
        }
    }

    updateSprites(sprites) {
        this.clear();

        sprites.forEach(sprite => {
            const img = document.createElement("img");
            img.setAttribute("src", `data:image/png;base64, ${sprite}`);
            img.style.margin = "3px";
            img.style.border = "thin grey dotted";
            img.style.imageRendering = "crisp-edges";
            img.style.height = "64px";
            img.style.width = "64px";
            this.appendChild(img);
        });
    }
}

customElements.define(Sprites.tagName(), Sprites, { extends: 'div' });

export function newSprites() {
    return document.createElement('div', { is: Sprites.tagName() })
}
