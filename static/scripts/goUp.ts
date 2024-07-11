import { directoryPath, directoryPathSetter } from "./directoryPathSetter"

import { sortFlag } from "./sortFlagSetter"

import { get } from "./get"

// Переход на директорию выше
export function goUp () {
    if (directoryPath != "/home") {
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