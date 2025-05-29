package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/server/file_sys"
	"example.com/server/server_output"
	"gopkg.in/ini.v1"
)

type PHPData struct {
	Root       string `json:"root"`
	Size       int    `json:"size"`
	ParsedTime string `json:"parsedTime"`
}

// ServerFunc - Функция обработки запросов сервера
func ServerFunc(cfg *ini.File) {
	const StatusOK = 0
	const StatusBad = 1

	server := &http.Server{
		//Получение данных порта
		Addr: fmt.Sprintf(":%s", cfg.Section("server").Key("port").String()),
	}

	http.Handle("/", http.FileServer(http.Dir("./static/dist")))

	http.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		//Получение данных из строки браузера (dst начинается с пути, в котором находится консоль)
		dst := r.URL.Query().Get("dst")
		if dst == "" {
			dst = cfg.Section("servAtr").Key("dst").String()
		}
		sort := r.URL.Query().Get("sort")
		if sort == "" {
			sort = cfg.Section("servAtr").Key("sort").String()
		}

		startDst := cfg.Section("servAtr").Key("dst").String()
		//Вызов функции из пакета FileSys
		resp, err := file_sys.DirSearcher(dst, sort)
		if err != nil {
			server_output.ServerOutput(w, StatusBad, err.Error(), dst, startDst, start, nil)
			fmt.Println(err)
		}
		fmt.Println("AAAAAAAAAAA")
		server_output.ServerOutput(w, StatusOK, "", dst, startDst, start, resp)

		var dirSize int
		for _, obj := range resp {
			dirSize += int(obj.Size)
		}

		var elapsed = float32(time.Since(start)) / float32(time.Second)

		phpdata := PHPData{}

		phpdata.Root = dst
		phpdata.ParsedTime = fmt.Sprintf("%.3f", elapsed)
		phpdata.Size = dirSize

		phpResp, err := json.Marshal(phpdata)
		if err != nil {
			fmt.Println(err)
		}

		phpBuffer := bytes.NewBuffer(phpResp)

		url := "http://localhost:80/writer.php"
		reqAA, _ := http.NewRequest("POST", url, phpBuffer)
		reqAA.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		respAA, err := client.Do(reqAA)
		if err != nil {
			panic(err)
		}
		defer respAA.Body.Close()

		fmt.Println("response Status:", respAA.Status)
		fmt.Println("response Headers:", respAA.Header)
		body, _ := ioutil.ReadAll(respAA.Body)
		fmt.Println("response Body:", string(body))
	})

	// Получение контекста сервера
	shutdownCtx, shutdownRelease := context.WithCancel(context.Background())
	defer shutdownRelease()

	//Создание Горутины для запуска сервера
	go func() {
		fmt.Printf("Сервер запускается по порту %s\n", server.Addr)
		errServ := server.ListenAndServe()
		if errServ != nil {
			fmt.Println(errServ)
			return
		}
		fmt.Println("Остановка приема подключений к серверу.")
	}()

	// Реализация Graceful Shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Завершение работы сервера
	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("HTTP ошибка остановки: %v\n", err)
	}
	fmt.Println("Завершение Graceful shutdown.")
}
