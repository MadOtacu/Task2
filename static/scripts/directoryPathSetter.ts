export let directoryPath = ""

// Установка значения пути директории
export function directoryPathSetter(dirPath: any) {
    directoryPath = dirPath
    document.getElementById("path").textContent = "Текущая директория: " + directoryPath;
}