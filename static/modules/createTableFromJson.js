// Создание и генерация таблицы
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