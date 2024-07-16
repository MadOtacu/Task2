import { get } from "server/get"

import { directoryPath } from "server/directoryPathSetter"

import { sortFlag, sortFlagSetter } from "server/sortFlagSetter"

// Функция сортировки
export function switchSort () {
    if (sortFlag == true) {
        let sortFlagTemp = false

        sortFlagSetter(sortFlagTemp)

        get(directoryPath, sortFlagTemp)
    }
    else {
        let sortFlagTemp = true

        sortFlagSetter(sortFlagTemp)

        get(directoryPath, sortFlagTemp)
    }
}