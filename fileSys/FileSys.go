package fileSys

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"

	Sorting "example.com/sorting"
	UnitScaling "example.com/unitScaling"
)

// Функция поиска данных файлов и директорий
func DirSearcher(dst string, sort string) ([]Sorting.File, error) {
	structFileArr := make([]Sorting.File, 0)
	var wg sync.WaitGroup

	// Считывание данных директории
	dirList, errRead := os.ReadDir(dst)
	if errRead != nil {
		return nil, errRead
	}

	// Обход директории
	for _, dirFile := range dirList {
		wg.Add(1)
		dirElement := dirFile
		structFile := new(Sorting.File)
		go func() {
			defer wg.Done()
			// Если элемент является директорией проходимся по ее структуре и записываем ее размер
			if dirElement.IsDir() {
				errWalking := filepath.Walk(dst+"/"+dirElement.Name(), func(path string, info fs.FileInfo, err error) error {
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

				structFileArr = append(structFileArr, *structFile)
			} else {
				// Если элемент является файлом считываем его данные
				infoFile, errGettingInfo := dirElement.Info()
				if errGettingInfo != nil {
					log.Println(errGettingInfo)
				}

				structFile.FileType = "Файл"
				structFile.Name = dirElement.Name()
				structFile.Size = infoFile.Size()

				structFileArr = append(structFileArr, *structFile)
			}
		}()
	}

	wg.Wait()

	// Функция сортировки массива
	Sorting.Sorting(structFileArr, sort)

	for i := 0; i < len(structFileArr); i++ {
		// Запись конвертированого значения
		structFileArr[i].ConvertedSize = UnitScaling.UnitScaling(structFileArr[i].Size)
	}

	// Парсинг json-файла
	return structFileArr, nil
}
