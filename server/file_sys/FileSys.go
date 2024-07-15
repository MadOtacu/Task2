package file_sys

import (
	"os"
	"sync"

	"example.com/server/sorting"
	"example.com/server/unit_scaling"
	"example.com/server/walker"
)

// DirSearcher - Функция поиска данных файлов и директорий
func DirSearcher(dst string, sort string) ([]sorting.File, error) {
	var wg sync.WaitGroup

	// Считывание данных директории
	dirList, errRead := os.ReadDir(dst)
	if errRead != nil {
		return nil, errRead
	}

	structFileArr := make([]sorting.File, len(dirList))

	// Обход директории
	for i, dirFile := range dirList {
		wg.Add(1)
		structFile := new(sorting.File)
		go walker.Walker(i, dst, &wg, dirFile, structFile, structFileArr)
	}

	wg.Wait()

	// Функция сортировки массива
	sorting.Sorting(structFileArr, sort)

	for i := 0; i < len(structFileArr); i++ {
		// Запись конвертированого значения
		structFileArr[i].ConvertedSize = unit_scaling.UnitScaling(structFileArr[i].Size)
	}

	// Парсинг json-файла
	return structFileArr, nil
}
