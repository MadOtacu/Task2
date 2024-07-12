import { createTableFromJson } from "./createTableFromJson"
import { directoryPathSetter } from "./directoryPathSetter"

import { sort } from "./sort"

export let startPath: any

// Получение данных с сервера
export function get (directoryPath: any, sortFlag: any) {
    document.getElementById("showTable").style.display = "none";

    document.getElementById("loader").style.display = "block";

    let url = "/path?dst=" + directoryPath + "&sort=" + sort(sortFlag)

    let response = fetch(url)
        .then(e => e.json())
        .then(commits => {
            if (commits.status == 0) {
                startPath = commits.startPath
                directoryPathSetter(commits.path)
                let commitsData = commits.data
                document.getElementById("loader").style.display = "none";
                createTableFromJson(commitsData)
                document.getElementById("showTable").style.display = "block";
            }
            else {
                document.getElementById("loader").style.display = "none";
                alert(commits.errorText)
            }
        })
        .catch(e => {
            alert(e.errorText)
        })
}