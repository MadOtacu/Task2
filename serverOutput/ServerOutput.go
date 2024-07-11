package serverOutput

import (
	"encoding/json"
	"fmt"
	"net/http"

	Sorting "example.com/sorting"
)

type ServerResp struct {
	Status    int            `json:"status"`
	ErrorText string         `json:"errorText"`
	Path      string         `json:"path"`
	Data      []Sorting.File `json:"data"`
}

func ServerOutput(w http.ResponseWriter, status int, errorText string, path string, data []Sorting.File) {
	resp := ServerResp{Status: status, ErrorText: errorText, Path: path, Data: data}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(status)
	w.Write(jsonResp)
}
