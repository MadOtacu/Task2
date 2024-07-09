let sortFlag = true

let directoryPath = "."

document.addEventListener('DOMContentLoaded', get(directoryPath, sortFlag), false);

async function get (directoryPath, sortFlag) {
    document.getElementById("path").textContent="Путь к директории: " + directoryPath

    let url = "http://localhost:9001/path?dst=" + directoryPath + "&sort=" + sort(sortFlag)

    let response = await fetch(url)
        .then(e => e.json())
        .then(commits => {
            commits = commits.data
            createTableFromJson(commits)
        })
        .catch(e => {
            alert(e.errorText)
        })
}

function createTableFromJson (response) {

    let table = document.createElement("table")

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
    document.getElementById("showTable").innerHTML = null
    directoryPath = directoryPath + "/" + this.innerHTML
    get(directoryPath, sortFlag)
}

document.getElementById("directoryUp").onclick = function goUp () {
    if (directoryPath != ".") {
        let directoryPathTemp = directoryPath.split("/")

        directoryPathTemp.splice(-1)

        directoryPath = directoryPathTemp.join("/")

        get(directoryPath, sortFlag)
    }
    else {
        alert("Вы на верху файловой системы!")
    }
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