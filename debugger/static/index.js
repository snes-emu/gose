
function pause() {
    fetch('/pause');
}

function step() {
    count = document.getElementById("count");
    fetch('/step?count='+count.value)
        .then(resp => resp.json())
        .then(body => {
            cpu = document.getElementById("cpu");
            li = document.createElement("li");
            li.appendChild(document.createTextNode(Object.keys(body).map(key => `${key}: ${body[key]}`).join(',')));
            cpu.appendChild(li);
        })
}


function breakpoint() {
    address = document.getElementById("breakpoint");
    fetch('/breakpoint?address='+address.value);
}