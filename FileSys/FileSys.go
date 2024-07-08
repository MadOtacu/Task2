package FileSys

import (
	"encoding/json"
	"flag"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type File struct {
	FileType      string `json:"fileType"`
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	ConvertedSize string `json:"convertedSize"`
}

// Функция поиска данных файлов и директорий
func DirSearcher(dst string, sort string) ([]byte, error) {
	var structFileArr []File
	var wg sync.WaitGroup

	// Считывание данных директории
	dirList, errRead := os.ReadDir(dst)
	if errRead != nil {
		flag.PrintDefaults()
	}

	// Обход директории
	for _, dirFile := range dirList {
		wg.Add(1)
		dirElement := dirFile
		structFile := new(File)
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
					flag.PrintDefaults()
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
	Sorting(structFileArr, sort)

	for i := 0; i < len(structFileArr); i++ {
		// Запись конвертированого значения
		structFileArr[i].ConvertedSize = UnitScaling(structFileArr[i].Size)
	}

	// Парсинг json-файла
	return json.MarshalIndent(structFileArr, "", "  ")
}

// Функция сортировки
func Sorting(arrToSort []File, sort string) {
	for i := 0; i < (len(arrToSort) - 1); i++ {
		for j := 0; j < ((len(arrToSort) - 1) - i); j++ {
			// Сортировка по возрастанию
			if sort == "ASC" {
				if arrToSort[j].Size > arrToSort[j+1].Size {
					temp := arrToSort[j]
					arrToSort[j] = arrToSort[j+1]
					arrToSort[j+1] = temp
				}
			} else if sort == "DESC" {
				// Сортировка по убыванию
				if arrToSort[j].Size < arrToSort[j+1].Size {
					temp := arrToSort[j]
					arrToSort[j] = arrToSort[j+1]
					arrToSort[j+1] = temp
				}
			} else {
				// В противном случае возвращаем значения флагов
				flag.PrintDefaults()
			}
		}
	}
}

// Функция скейлинга размера объектов
func UnitScaling(Size int64) string {
	var restOfSize int64
	var unitValue int
	var unit string
	const sizing int64 = 1000

	// Рассчет сокращенного размера объекта
	for i := 1; Size > sizing; i++ {
		restOfSize = Size % sizing
		Size = Size / sizing
		unitValue = i
	}

	// Запись единицы размера
	switch unitValue {
	case 0:
		unit = "байт"
	case 1:
		unit = "Кб"
	case 2:
		unit = "Мб"
	case 3:
		unit = "Гб"
	}

	// Вывод данных
	SizeToPaste := strconv.FormatInt(Size, 10)
	restOfSizeToPaste := strconv.FormatInt(restOfSize, 10)
	return string(SizeToPaste + "." + restOfSizeToPaste + " " + unit)
}
