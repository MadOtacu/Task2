import { dirDown } from "server/dirDown"
import fileImgLink from "../../static/images/document.png"
import directoryImgLink from "../../static/images/folder.png"

// Создание и генерация таблицы
export function createTableFromJson (response: any) {

    let table = document.createElement("table")

    const fileImg = new Image()
    const directoryImg = new Image()

    for (let rowi = 0; rowi < response.length; rowi++) {
        let tr = table.insertRow(rowi)

        let tabType = tr.insertCell(0)
        let tabName = tr.insertCell(1)
        let tabSize = tr.insertCell(2)

        tabName.textContent = response[rowi].name
        tabSize.textContent = response[rowi].convertedSize

        if (response[rowi].fileType == "Файл") {
            tabType.appendChild(fileImg)
        }
        if (response[rowi].fileType == "Директория") {
            tabType.appendChild(directoryImg)
            tr.addEventListener("click", dirDown)
            tr.classList.add("buttonCell")
        }
    }

    let divContainer = document.getElementById("showTable");
    divContainer.innerHTML = "";
    divContainer.appendChild(table);
}