package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	FileSys "example.com/scripts/fileSys"
	ServerOutput "example.com/scripts/serverOutput"
	"gopkg.in/ini.v1"
)

// serverFunc - Функция обработки запросов сервера
func ServerFunc(cfg *ini.File) {
	server := &http.Server{
		//Получение данных порта
		Addr: cfg.Section("server").Key("port").String(),
	}

	http.Handle("/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		//Получение данных из строки браузера (dst начинается с пути, в котором находится консоль)
		dst := r.URL.Query().Get("dst")
		if dst == "" {
			dst = cfg.Section("servAtr").Key("dst").String()
		}
		sort := r.URL.Query().Get("sort")
		if sort == "" {
			sort = cfg.Section("servAtr").Key("sort").String()
		}

		//Вызов функции из пакета FileSys
		resp, err := FileSys.DirSearcher(dst, sort)
		if err != nil {
			ServerOutput.ServerOutput(w, http.StatusBadRequest, err.Error(), dst, nil)
			fmt.Println(err)
		}
		ServerOutput.ServerOutput(w, http.StatusOK, "", dst, resp)
	})

	//Создание Горутины для запуска сервера
	go func() {
		fmt.Println("Сервер запущен и слушает.")
		errServ := server.ListenAndServe()
		if errServ != nil {
			fmt.Println(errServ)
		}
		fmt.Println("Остановка приема подключений к серверу.")
	}()

	// Реализация Graceful Shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Получение контекста сервера
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	// Завершение работы сервера
	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("HTTP ошибка остановки: %v", err)
	}
	fmt.Println("Завершение Graceful shutdown.")
}
