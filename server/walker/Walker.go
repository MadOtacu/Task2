package walker

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"sync"

	"example.com/server/sorting"
)

// Walker - Функция обхода дериктории
func Walker(i int, dst string, wg *sync.WaitGroup, dirElement fs.DirEntry, structFile *sorting.File, structFileArr []sorting.File) {
	defer wg.Done()
	// Если элемент является директорией проходимся по ее структуре и записываем ее размер
	if dirElement.IsDir() {
		errWalking := filepath.Walk(fmt.Sprintf("%s/%s", dst, dirElement.Name()), func(path string, info fs.FileInfo, err error) error {
			if !info.IsDir() {
				structFile.Size += info.Size()
			}
			return nil
		})
		if errWalking != nil {
			fmt.Printf("Ошибка обхода директории: %s", errWalking.Error())
		}

		structFile.FileType = "Директория"
		structFile.Name = dirElement.Name()

		structFileArr[i] = *structFile
	} else {
		// Если элемент является файлом считываем его данные
		infoFile, errGettingInfo := dirElement.Info()
		if errGettingInfo != nil {
			log.Println(errGettingInfo)
		}

		structFile.FileType = "Файл"
		structFile.Name = dirElement.Name()
		structFile.Size = infoFile.Size()

		structFileArr[i] = *structFile
	}
}
