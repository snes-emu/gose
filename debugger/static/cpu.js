class CPU extends HTMLUListElement {
    static tagName() {
        return 'cpu-ul';
    }

    constructor() {
        super();
        this.id = "cpu";
    }

    addState(cpu) {
        const li = document.createElement("li");
        li.appendChild(document.createTextNode(Object.keys(cpu).map(key => `${key}: ${cpu[key]}`).join(',')));
        this.appendChild(li);
    }
}

customElements.define(CPU.tagName(), CPU, {extends: 'ul'});

export function newCPU() {
    return document.createElement('ul', {is: CPU.tagName()})
}
