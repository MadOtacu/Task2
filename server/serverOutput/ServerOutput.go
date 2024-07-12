package serverOutput

import (
	"encoding/json"
	"fmt"
	"net/http"

	Sorting "example.com/server/sorting"
)

// ServerResp - хранит ответ сервера
type ServerResp struct {
	Status    int            `json:"status"`
	ErrorText string         `json:"errorText"`
	Path      string         `json:"path"`
	StartPath string         `json:"startPath"`
	Data      []Sorting.File `json:"data"`
}

// ServerOutput - Функция создающая ответ для сервера
func ServerOutput(w http.ResponseWriter, status int, errorText string, path string, startPath string, data []Sorting.File) {
	resp := ServerResp{Status: status, ErrorText: errorText, Path: path, Data: data}
	resp.StartPath = startPath

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}

	w.Write(jsonResp)
}
