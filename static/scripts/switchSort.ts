import { get } from "./get"

import { directoryPath } from "./directoryPathSetter"

import { sortFlag, sortFlagSetter } from "./sortFlagSetter"

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