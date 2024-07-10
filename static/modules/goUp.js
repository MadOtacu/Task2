// Переход на директорию выше
document.getElementById("directoryUp").onclick = function goUp () {
    if (directoryPath != ".") {
        let directoryPathTemp = directoryPath.split("/")

        directoryPathTemp.splice(-1)

        directoryPath = directoryPathTemp.join("/")

        get(directoryPath, sortFlag)
    }
    else {
        alert("Вы на верху файловой системы!")
    }
}