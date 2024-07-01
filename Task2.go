package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	var dst = flag.String("dst", "", "Путь к директории для рассчета размера")

	flag.Parse()

	var size int64 = 0

	errPath := filepath.Walk(*dst, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			flag.PrintDefaults()
			os.Exit(1)
		}

		if !info.IsDir() {
			size += info.Size()
		}

		return nil
	})

	if errPath != nil {
		fmt.Println("Ошибка во время чтения файлов")
		os.Exit(1)
	}

	sizeToPaste := strconv.FormatInt(size, 10)
	fmt.Println("Размер вложенных директорий составляет " + sizeToPaste + " байт")
}
