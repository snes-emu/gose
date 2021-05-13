import { DynamicTable } from "./table.js";

class CPU extends DynamicTable {
    static tagName() {
        return 'cpu-table';
    }

    constructor() {
        super();
        this.id = "cpu";
    }
}

customElements.define(CPU.tagName(), CPU, { extends: 'table' });

export function newCPU() {
    return document.createElement('table', { is: CPU.tagName() })
}
