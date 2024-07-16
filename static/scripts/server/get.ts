import { createTableFromJson } from "client/createTableFromJson"
import { directoryPathSetter } from "server/directoryPathSetter"

import { sort } from "server/sort"
import { timeSetter } from "server/timeSetter";
import { showTable } from "client/showTable";
import { hideTable } from "client/hideTable";
import { showLoader } from "client/showLoader";
import { hideLoader } from "client/hideLoader";

export let startPath: any

// Получение данных с сервера
export function get (directoryPath: any, sortFlag: any) {
    hideTable()

    showLoader()

    let url = `/path?dst=${directoryPath}&sort=${sort(sortFlag)}`

    let response = fetch(url)
        .then(e => e.json())
        .then(commits => {
            if (commits.status == 0) {
                timeSetter(commits.elapsedTime)
                startPath = commits.startPath
                directoryPathSetter(commits.path)
                let commitsData = commits.data
                hideLoader()
                createTableFromJson(commitsData)
                showTable()
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