import { DynamicTable } from "./table.js";

class CPU extends DynamicTable {
    static tagName() {
        return 'cpu-table';
    }

    constructor() {
        super();
        this.id = "cpu";
    }

    addData(entry) {
        const newEntry = {...entry};
        ['C', 'D', 'DBR', 'K', 'PC', 'S', 'X', 'Y'].forEach(key => {
            newEntry[key] = `0x${entry[key].toString(16)} (${entry[key]})`;
        });
        this.addEntry(newEntry);
    }
}

customElements.define(CPU.tagName(), CPU, {extends: 'table'});

export function newCPU() {
    return document.createElement('table', {is: CPU.tagName()})
}
