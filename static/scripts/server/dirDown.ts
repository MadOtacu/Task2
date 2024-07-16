import { directoryPath, directoryPathSetter } from "server/directoryPathSetter"

import { sortFlag } from "server/sortFlagSetter"

import { get } from "server/get"

// Переход на директорию ниже
export function dirDown(e: any) {
    document.getElementById("showTable").innerHTML = null
    let directoryPathSet = `${directoryPath}/${e.currentTarget.cells[1].textContent}`
    directoryPathSetter(directoryPathSet)
    get(directoryPathSet, sortFlag)

}