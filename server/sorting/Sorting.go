package sorting

import "fmt"

// File - Структура данных тела ответа
type File struct {
	FileType      string `json:"fileType"`
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	ConvertedSize string `json:"convertedSize"`
}

// Sorting - Функция сортировки ответа
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
				// В противном случае возвращаем уведомление о неверном флаге
				fmt.Println("Неверное значение флага! Входными значениями могут быть только ASC или DESC.")
			}
		}
	}
}
