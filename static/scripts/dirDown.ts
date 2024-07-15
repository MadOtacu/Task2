import { directoryPath, directoryPathSetter } from "./directoryPathSetter"

import { sortFlag } from "./sortFlagSetter"

import { get } from "./get"

// Переход на директорию ниже
export function dirDown(e: any) {
    document.getElementById("showTable").innerHTML = null
    let directoryPathSet = `${directoryPath}/${e.currentTarget.cells[1].textContent}`
    directoryPathSetter(directoryPathSet)
    get(directoryPathSet, sortFlag)

}