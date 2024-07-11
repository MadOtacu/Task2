import { createTableFromJson } from "./createTableFromJson"

import { sort } from "./sort"

// Получение данных с сервера
export function get (directoryPath: any, sortFlag: any) {
    document.getElementById("path").textContent="Путь к директории: " + directoryPath

    let url = "/path?dst=" + directoryPath + "&sort=" + sort(sortFlag)

    let response = fetch(url)
        .then(e => e.json())
        .then(commits => {
            if (commits.status == 200) {
                let commitsData = commits.data
                createTableFromJson(commitsData)
            }
            else {
                alert(commits.errorText)
            }
        })
        .catch(e => {
            alert(e.errorText)
        })
}