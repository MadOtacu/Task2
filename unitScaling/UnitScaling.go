package unitScaling

import "strconv"

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
