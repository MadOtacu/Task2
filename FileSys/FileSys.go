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

func DirSearcher(dst string, sort string) ([]byte, error) {
	var structFileArr []File
	var wg sync.WaitGroup

	dirList, errRead := os.ReadDir(dst)
	if errRead != nil {
		flag.PrintDefaults()
	}
	for _, dirFile := range dirList {
		wg.Add(1)
		dirElement := dirFile
		structFile := new(File)
		go func() {
			defer wg.Done()
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

	Sorting(structFileArr, sort)

	for i := 0; i < len(structFileArr); i++ {
		structFileArr[i].ConvertedSize = UnitScaling(structFileArr[i].Size)
	}

	return json.MarshalIndent(structFileArr, "", "  ")
}

func Sorting(arrToSort []File, sort string) {
	for i := 0; i < (len(arrToSort) - 1); i++ {
		for j := 0; j < ((len(arrToSort) - 1) - i); j++ {
			if sort == "ASC" {
				if arrToSort[j].Size > arrToSort[j+1].Size {
					temp := arrToSort[j]
					arrToSort[j] = arrToSort[j+1]
					arrToSort[j+1] = temp
				}
			} else if sort == "DESC" {
				if arrToSort[j].Size < arrToSort[j+1].Size {
					temp := arrToSort[j]
					arrToSort[j] = arrToSort[j+1]
					arrToSort[j+1] = temp
				}
			} else {
				flag.PrintDefaults()
			}
		}
	}
}

func UnitScaling(Size int64) string {
	var restOfSize int64
	var unitValue int
	var unit string

	for i := 1; Size > 1000; i++ {
		restOfSize = Size % 1000
		Size = Size / 1000
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

	SizeToPaste := strconv.FormatInt(Size, 10)
	restOfSizeToPaste := strconv.FormatInt(restOfSize, 10)
	return string(SizeToPaste + "." + restOfSizeToPaste + " " + unit)
}
