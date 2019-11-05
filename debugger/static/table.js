export class DynamicTable extends HTMLTableElement {
    constructor() {
        super();
        this.body = document.createElement("tbody");
        this.appendChild(this.body);
    }

    addEntry(entry) {
        // If we haven't set the table header yet, do it
        if (!this.header) {
            this.header = document.createElement("thead");
            this.fields = [];
            const hdrRow = document.createElement("tr");
            Object.keys(entry).forEach(key => {
                this.fields.push(key);

                const h = document.createElement("th");
                h.appendChild(document.createTextNode(key));
                hdrRow.appendChild(h);
            });

            this.header.appendChild(hdrRow);
            this.appendChild(this.header);
        }

        const tr = document.createElement("tr");
        this.fields.forEach(key => {
            const td = document.createElement("td");
            td.appendChild(document.createTextNode(entry[key]));
            tr.appendChild(td);
        })
        this.body.appendChild(tr);
    }
}
