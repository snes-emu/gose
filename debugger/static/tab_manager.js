class TabManager extends HTMLDivElement {
    static tagName() {
        return 'tab-manager';
    }

    constructor() {
        super();

        this.header = document.createElement('div');
        this.header.style.margin = "20px 0";
        this.header.style.borderBottom = "1px solid gray";
        this.appendChild(this.header);

        this.container = document.createElement('div');
        this.container.id = 'container_id';
        this.container.className = 'container'
        this.appendChild(this.container);
    }


    onInputChange(event) {
        Array.from(this.container.children).forEach(child => this.container.removeChild(child));
        this.container.appendChild(this.tabs[event.target.value])
    }


    setTabs(tabs) {
        this.tabs = {};
        tabs.forEach(({name, component}, index) => {
            const input = document.createElement('input');
            const id = `${name}_choice`
            input.type = 'radio';
            input.name = 'tab';
            input.value = name;
            input.id = id;
            input.onchange = this.onInputChange.bind(this);
            if (index == 0) {
                input.checked = true;
            }
            this.header.appendChild(input);

            const label = document.createElement('label');
            label.for = id;
            label.innerText = name;
            label.style.display = "initial";
            this.header.appendChild(label);

            this.tabs[name] = component;
        })

        this.container.appendChild(tabs[0].component);
    }
}

customElements.define(TabManager.tagName(), TabManager, {extends: 'div'});

export function newTabManager() {
    return document.createElement('div', {is: TabManager.tagName()})
}
