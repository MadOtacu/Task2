// Функция сортировки
document.getElementById("sort").onclick = function switchSort () {
    if (sortFlag == true) {
        sortFlag = false

        get(directoryPath, sortFlag)
    }
    else {
        sortFlag = true

        get(directoryPath, sortFlag)
    }
}