// Переход на директорию ниже
function dirDown() {
    document.getElementById("showTable").innerHTML = null
    directoryPath = directoryPath + "/" + this.innerHTML
    get(directoryPath, sortFlag)
}