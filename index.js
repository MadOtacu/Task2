document.addEventListener('DOMContentLoaded', async function() {
    let url = "http://localhost:9001/path?dst=.&sort=ASC"

    let response = await fetch(url)

    let commits = await response.json()

    createTableFromJson(commits)
}, false);

function createTableFromJson (response) {
    let arrResponse = []
    arrResponse = JSON.parse(response)

    var col = ["Тип", "Имя", "Размер"]

    var table = document.createElement("table")

    var tr = table.insertRow(-1)

    for (let headi = 0; i < col.length; i++) {
        var th = document.createElement("th");
        th.innerHTML = col[headi];
        tr.appendChild(th);
    }

    for (let rowi = 0; rowi < arrResponse.length; rowi++) {
        tr = table.insertRow(rowi)

        var tabType = tr.insertCell(0)
        var tabName = tr.insertCell(1)
        var tabSize = tr.insertCell(2)

        tabType.innerHTML = arrResponse[rowi].fileType
        tabName.innerHTML = arrResponse[rowi].name
        tabSize.innerHTML = arrResponse[rowi].convertedSize
    }

    var divContainer = document.getElementById("showTable");
    divContainer.innerHTML = "";
    divContainer.appendChild(table);
}