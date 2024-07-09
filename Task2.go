package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/example/FileSys"
	"gopkg.in/ini.v1"
)

// Функция обработки запросов сервера
func serverFunc(cfg *ini.File) {
	server := &http.Server{
		//Получение данных порта
		Addr: cfg.Section("server").Key("port").String(),
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))

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
			fmt.Println(err)
		}
		fmt.Fprint(w, string(resp))
	})

	//Создание Горутины для запуска сервера
	go func() {
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

func main() {
	// Загрузка файла конфигурации
	cfg, errLoad := ini.Load("conf.ini")
	if errLoad != nil {
		fmt.Printf("Ошибка чтения файла: %v", errLoad)
	}
	// Вызов функции запуска сервера
	serverFunc(cfg)
}
