// Получение данных с сервера
function get (directoryPath, sortFlag) {
    document.getElementById("path").textContent="Путь к директории: " + directoryPath

    let url = "http://localhost:9001/path?dst=" + directoryPath + "&sort=" + sort(sortFlag)

    let response = fetch(url)
        .then(e => e.json())
        .then(commits => commits.data)
        .then(data => createTableFromJson(data))
        .catch(e => {
            alert(e.errorText)
        })
}