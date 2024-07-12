package fileSys

import (
	"os"
	"sync"

	Sorting "example.com/scripts/sorting"
	UnitScaling "example.com/scripts/unitScaling"
	Walker "example.com/scripts/walker"
)

// Функция поиска данных файлов и директорий
func DirSearcher(dst string, sort string) ([]Sorting.File, error) {
	var wg sync.WaitGroup

	// Считывание данных директории
	dirList, errRead := os.ReadDir(dst)
	if errRead != nil {
		return nil, errRead
	}

	structFileArr := make([]Sorting.File, len(dirList))

	// Обход директории
	for i, dirFile := range dirList {
		wg.Add(1)
		structFile := new(Sorting.File)
		go Walker.Walker(i, dst, &wg, dirFile, structFile, structFileArr)
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
