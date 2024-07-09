let sortFlag = true

let directoryPath = "."

document.addEventListener('DOMContentLoaded', get(directoryPath, sortFlag), false);

async function get (directoryPath, sortFlag) {
    let url = "http://localhost:9001/path?dst=" + directoryPath + "&sort=" + sort(sortFlag)

    let response = await fetch(url)

    let commits = await response.json()

    createTableFromJson(commits)
}

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

        if (response[rowi].fileType == "Директория") {
            tabName.addEventListener("click", dirDown)
        }
    }

    let divContainer = document.getElementById("showTable");
    divContainer.innerHTML = "";
    divContainer.appendChild(table);
}

function dirDown() {
    directoryPath = directoryPath + this.innerHTML
    get(directoryPath, sortFlag)
}

document.getElementById("sort").onclick = function switchSort () {
    if (sortFlag == true) {
        sortFlag = false

        get(directoryPath, sortFlag)
    }
    else {
        sortFlag = true

        get(directoryPath, sortFlag)
    }
}

function sort (flag) {
    if (flag == true) {
        return "ASC"
    }
    else {
        return "DESC"
    }
}