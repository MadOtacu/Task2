package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type File struct {
	fileType string
	name     string
	size     int64
}

func dirSearcher(dst *string, sort *string) {
	var structFileArr []File
	var wg sync.WaitGroup

	dirList, errRead := os.ReadDir(*dst)
	if errRead != nil {
		flag.PrintDefaults()
		os.Exit(1)
	}
	for _, dirFile := range dirList {
		wg.Add(1)
		dirElement := dirFile
		structFile := new(File)
		go func() {
			defer wg.Done()
			if dirElement.IsDir() {
				errWalking := filepath.Walk(*dst+"/"+dirElement.Name(), func(path string, info fs.FileInfo, err error) error {
					if !info.IsDir() {
						structFile.size += info.Size()
					}
					return nil
				})
				if errWalking != nil {
					flag.PrintDefaults()
					os.Exit(1)
				}

				structFile.fileType = "Директория"
				structFile.name = dirElement.Name()

				structFileArr = append(structFileArr, *structFile)
			} else {
				infoFile, errGettingInfo := dirElement.Info()
				if errGettingInfo != nil {
					panic(errGettingInfo)
				}

				structFile.fileType = "Файл"
				structFile.name = dirElement.Name()
				structFile.size = infoFile.Size()

				structFileArr = append(structFileArr, *structFile)
			}
		}()
	}

	wg.Wait()

	Sorting(structFileArr, sort)

	for _, dirPrint := range structFileArr {
		size, restOfSize, unit := UnitScaling(dirPrint.size)
		sizeToPaste := strconv.FormatInt(size, 10)
		restOfSizeToPaste := strconv.FormatInt(restOfSize, 10)
		fmt.Println(dirPrint.fileType + " " + dirPrint.name + " размером " + sizeToPaste + "." + restOfSizeToPaste + " " + unit)
	}
}

func Sorting(arrToSort []File, sort *string) {
	for i := 0; i < (len(arrToSort) - 1); i++ {
		for j := 0; j < ((len(arrToSort) - 1) - i); j++ {
			if *sort == "ASC" {
				if arrToSort[j].size > arrToSort[j+1].size {
					temp := arrToSort[j]
					arrToSort[j] = arrToSort[j+1]
					arrToSort[j+1] = temp
				}
			} else if *sort == "DESC" {
				if arrToSort[j].size < arrToSort[j+1].size {
					temp := arrToSort[j]
					arrToSort[j] = arrToSort[j+1]
					arrToSort[j+1] = temp
				}
			} else {
				flag.PrintDefaults()
				os.Exit(1)
			}
		}
	}
}

func UnitScaling(size int64) (int64, int64, string) {
	var restOfSize int64
	var unitValue int
	var unit string

	for i := 1; size > 1000; i++ {
		restOfSize = size % 1000
		size = size / 1000
		unitValue = i
	}

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
	return size, restOfSize, unit
}

func main() {
	var dst = flag.String("dst", "", "Путь к директории для вывода данных")
	var sort = flag.String("sort", "ASC", "Сортировка по размеру (Входные значения ASC или DESC). По умолчанию применяется ASC")

	flag.Parse()

	dirSearcher(dst, sort)
}
