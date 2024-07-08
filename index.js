document.addEventListener('DOMContentLoaded', async function() {
    let url = "http://localhost:9001/path?dst=.&sort=ASC"

    let response = await fetch(url)

    let commits = await response.json()

    createTableFromJson(commits)
}, false);

function createTableFromJson (response) {

    let col = ["Тип", "Имя", "Размер"]

    let table = document.createElement("table")

    let tr = table.insertRow(-1)

    for (let rowi = 0; rowi < response.length; rowi++) {
        tr = table.insertRow(rowi)

        let tabType = tr.insertCell(0)
        let tabName = tr.insertCell(1)
        let tabSize = tr.insertCell(2)

        tabType.innerHTML = response[rowi].fileType
        tabName.innerHTML = response[rowi].name
        tabSize.innerHTML = response[rowi].convertedSize
    }

    let divContainer = document.getElementById("showTable");
    divContainer.innerHTML = "";
    divContainer.appendChild(table);
}