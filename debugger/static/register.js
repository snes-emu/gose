class Register extends HTMLUListElement {
    static tagName() {
        return 'register-ul';
    }

    constructor() {
        super();
        this.id = "register";
    }

    addState(reg) {
        const li = document.createElement("li");
        li.appendChild(document.createTextNode(
            `name: ${reg.name}, type: ${reg.type}, data: 0x${reg.data.toString(16)}`
        ));
        this.appendChild(li);
    }
}

customElements.define(Register.tagName(), Register, {extends: 'ul'});

export function newRegister() {
    return document.createElement('ul', {is: Register.tagName()})
}
