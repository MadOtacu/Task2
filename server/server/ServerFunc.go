package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/server/fileSys"
	"example.com/server/serverOutput"
	"gopkg.in/ini.v1"
)

// ServerFunc - Функция обработки запросов сервера
func ServerFunc(cfg *ini.File) {
	const StatusOK = 0
	const StatusBad = 1

	server := &http.Server{
		//Получение данных порта
		Addr: fmt.Sprintf(":%s", cfg.Section("server").Key("port").String()),
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))

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
		resp, err := fileSys.DirSearcher(dst, sort)
		if err != nil {
			serverOutput.ServerOutput(w, StatusBad, err.Error(), dst, startDst, start, nil)
			fmt.Println(err)
		}
		serverOutput.ServerOutput(w, StatusOK, "", dst, startDst, start, resp)
	})

	//Создание Горутины для запуска сервера
	go func() {
		fmt.Printf("Сервер запущен и слушает по порту %s", server.Addr)
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
