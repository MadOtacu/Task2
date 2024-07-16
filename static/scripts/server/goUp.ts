import { directoryPath, directoryPathSetter } from "server/directoryPathSetter"

import { sortFlag } from "server/sortFlagSetter"

import { get, startPath } from "server/get"

// Переход на директорию выше
export function goUp () {
    if (directoryPath != startPath) {
        let directoryPathTemp = directoryPath.split("/")

        directoryPathTemp.splice(-1)

        let directoryPathTotal = directoryPathTemp.join("/")

        directoryPathSetter(directoryPathTotal)

        get(directoryPathTotal, sortFlag)
    }
    else {
        alert("Вы на верху файловой системы!")
    }
}