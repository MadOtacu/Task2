package server_output

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"example.com/server/sorting"
)

// ServerResp - хранит ответ сервера
type ServerResp struct {
	Status      int            `json:"status"`      // Статус ответа
	ErrorText   string         `json:"errorText"`   // Текст ошибки
	Path        string         `json:"path"`        // Текущий путь
	StartPath   string         `json:"startPath"`   // Стартовый путь
	ElapsedTime string         `json:"elapsedTime"` // Время выполнения
	Data        []sorting.File `json:"data"`        // Данные для отрисовки таблицы
}

// ServerOutput - Функция создающая ответ для сервера
func ServerOutput(w http.ResponseWriter, status int, errorText string, path string, startPath string, start time.Time, data []sorting.File) {
	resp := ServerResp{Status: status, ErrorText: errorText, Path: path, Data: data, StartPath: startPath}
	var elapsed = float32(time.Since(start)) / float32(time.Second)
	resp.ElapsedTime = fmt.Sprintf("%.3f", elapsed)

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}

	w.Write(jsonResp)
}
