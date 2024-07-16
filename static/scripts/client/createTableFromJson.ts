import { dirDown } from "server/dirDown"

// Создание и генерация таблицы
export function createTableFromJson (response: any) {

    let table = document.createElement("table")

    for (let rowi = 0; rowi < response.length; rowi++) {
        let tr = table.insertRow(rowi)

        let tabType = tr.insertCell(0)
        let tabName = tr.insertCell(1)
        let tabSize = tr.insertCell(2)

        tabType.textContent = response[rowi].fileType
        tabName.textContent = response[rowi].name
        tabSize.textContent = response[rowi].convertedSize

        if (response[rowi].fileType == "Директория") {
            tr.addEventListener("click", dirDown)
            tr.classList.add("buttonCell")
        }
    }

    let divContainer = document.getElementById("showTable");
    divContainer.innerHTML = "";
    divContainer.appendChild(table);
}