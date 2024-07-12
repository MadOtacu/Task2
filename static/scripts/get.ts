import { createTableFromJson } from "createTableFromJson"
import { directoryPathSetter } from "directoryPathSetter"

import { sort } from "./sort"
import { timeSetter } from "timeSetter";
import { showTable } from "showTable";
import { hideTable } from "hideTable";
import { showLoader } from "showLoader";
import { hideLoader } from "hideLoader";

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