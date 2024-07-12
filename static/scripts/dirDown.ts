import { directoryPath, directoryPathSetter } from "./directoryPathSetter"

import { sortFlag } from "./sortFlagSetter"

import { get } from "./get"

// Переход на директорию ниже
export function dirDown() {
    document.getElementById("showTable").innerHTML = null
    let directoryPathSet = `${directoryPath}/${this.innerHTML}`
    directoryPathSetter(directoryPathSet)
    get(directoryPathSet, sortFlag)

}