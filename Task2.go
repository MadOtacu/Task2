package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
)

func dirSearcher(dst *string) {
	var size int64 = 0

	errPath := filepath.Walk(*dst, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			flag.PrintDefaults()
			os.Exit(1)
		}

		sizeToPaste := strconv.FormatInt(size, 10)
		if !info.IsDir() && *dst == path {
			size += info.Size()
			fmt.Println("Файл " + info.Name() + " размером " + sizeToPaste + " байт")
			size = 0
		} else if info.IsDir() && *dst == path {
			fmt.Println("Директория " + info.Name() + " размером " + sizeToPaste + " байт")
			fmt.Println(*dst)
			fmt.Println(path)
			size = 0
		} else if !info.IsDir() && *dst != path {
			size += info.Size()
			fmt.Println("gg Файл " + info.Name() + " размером " + sizeToPaste + " байт")
			fmt.Println(*dst)
			fmt.Println(path)
		}

		return nil
	})

	if errPath != nil {
		fmt.Println("Ошибка во время чтения файлов")
		os.Exit(1)
	}
}

func main() {
	var dst = flag.String("dst", "", "Путь к директории для рассчета размера")

	flag.Parse()

	dirSearcher(dst)
}
