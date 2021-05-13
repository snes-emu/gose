import { DynamicTable } from "./table.js";

class Register extends DynamicTable {
    static tagName() {
        return 'register-table';
    }

    constructor() {
        super();
        this.id = "register";
    }

    addData(entry) {
        const newEntry = { ...entry };
        newEntry.data = `0x${entry.data.toString(16)}`;
        this.addEntry(newEntry);
    }
}

customElements.define(Register.tagName(), Register, { extends: 'table' });

export function newRegister() {
    return document.createElement('table', { is: Register.tagName() })
}
