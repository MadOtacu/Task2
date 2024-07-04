package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

type File struct {
	FileType string `json:"fileType"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
}

func dirSearcher(dst string, sort string) ([]byte, error) {
	var structFileArr []File
	var wg sync.WaitGroup

	dirList, errRead := os.ReadDir(dst)
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
				errWalking := filepath.Walk(dst+"/"+dirElement.Name(), func(path string, info fs.FileInfo, err error) error {
					if !info.IsDir() {
						structFile.Size += info.Size()
					}
					return nil
				})
				if errWalking != nil {
					flag.PrintDefaults()
					os.Exit(1)
				}

				structFile.FileType = "Директория"
				structFile.Name = dirElement.Name()

				structFileArr = append(structFileArr, *structFile)
			} else {
				infoFile, errGettingInfo := dirElement.Info()
				if errGettingInfo != nil {
					panic(errGettingInfo)
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

	return json.MarshalIndent(structFileArr, "", "  ")

	//for _, dirPrint := range structFileArr {
	//	Size, restOfSize, unit := UnitScaling(dirPrint.Size)
	//	SizeToPaste := strconv.FormatInt(Size, 10)
	//	restOfSizeToPaste := strconv.FormatInt(restOfSize, 10)
	//	fmt.Println(dirPrint.FileType + " " + dirPrint.Name + " размером " + SizeToPaste + "." + restOfSizeToPaste + " " + unit)
	//}
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
				os.Exit(1)
			}
		}
	}
}

/*
func UnitScaling(Size int64) (int64, int64, string) {
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
	return Size, restOfSize, unit
}
*/

func main() {
	server := &http.Server{
		Addr: ":9001",
	}

	http.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {

		dst := r.URL.Query().Get("dst")
		sort := r.URL.Query().Get("sort")
		resp, err := dirSearcher(dst, sort)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Fprint(w, string(resp))
	})

	go func() {
		errServ := server.ListenAndServe()
		if errServ != nil {
			log.Println(errServ)
			os.Exit(1)
		}
		log.Println("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")
}
