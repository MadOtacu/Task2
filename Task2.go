package main

import (
	"fmt"

	Server "example.com/server"
	"gopkg.in/ini.v1"
)

func main() {
	// Загрузка файла конфигурации
	cfg, errLoad := ini.Load("conf.ini")
	if errLoad != nil {
		fmt.Printf("Ошибка чтения файла: %v", errLoad)
	}
	// Вызов функции запуска сервера
	Server.ServerFunc(cfg)
}
