package main

import (
	"log"

	"github.com/srg77global/home_work_basic/todo-list"
	"github.com/srg77global/home_work_basic/todo-list/internal/handler"
)

func main() {
	// инициализация хендлеров
	handler := new(handler.Handler)

	// инициализация и запуск сервера
	server := todo.Server{}
	if err := server.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error running server: %v", err.Error())
	}
}
