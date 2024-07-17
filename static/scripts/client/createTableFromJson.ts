import { dirDown } from "server/dirDown"
import fileImgLink from "Images/document.png"
import directoryImgLink from "Images/folder.png"

// Создание и генерация таблицы
export function createTableFromJson (response: any) {

    let table = document.createElement("table")

    for (let rowi = 0; rowi < response.length; rowi++) {
        let tr = table.insertRow(rowi)

        let tabType = tr.insertCell(0)
        let tabName = tr.insertCell(1)
        let tabSize = tr.insertCell(2)

        tabName.classList.add("tabWide")
        tabType.classList.add("tabShort")
        tabSize.classList.add("tabShort")

        tabName.textContent = response[rowi].name
        tabSize.textContent = response[rowi].convertedSize

        if (response[rowi].fileType == "Файл") {
            const fileImg = new Image()
            fileImg.src = fileImgLink
            tabType.appendChild(fileImg)
        }
        if (response[rowi].fileType == "Директория") {
            const directoryImg = new Image()
            directoryImg.src = directoryImgLink
            tabType.appendChild(directoryImg)
            tr.addEventListener("click", dirDown)
            tr.classList.add("buttonCell")
        }
    }

    let divContainer = document.getElementById("showTable");
    divContainer.innerHTML = "";
    divContainer.appendChild(table);
}