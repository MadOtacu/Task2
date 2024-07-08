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

func serverFunc(cfg *ini.File) {
	server := &http.Server{
		Addr: cfg.Section("server").Key("port").String(),
	}

	http.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {

		dst := r.URL.Query().Get("dst")
		sort := r.URL.Query().Get("sort")
		resp, err := FileSys.DirSearcher(dst, sort)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Fprint(w, string(resp))
	})

	go func() {
		errServ := server.ListenAndServe()
		if errServ != nil {
			fmt.Println(errServ)
		}
		fmt.Println("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("HTTP shutdown error: %v", err)
	}
	fmt.Println("Graceful shutdown complete.")
}

func main() {
	cfg, errLoad := ini.Load("conf.ini")
	if errLoad != nil {
		fmt.Printf("Fail to read file: %v", errLoad)
		os.Exit(1)
	}
	serverFunc(cfg)

}
