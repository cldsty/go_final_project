package main

import (
	"go-final-project/pkg/db"
	"go-final-project/pkg/server"
	"log"
	"os"
)

func main() {
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	defer db.Close()

	logger := log.New(os.Stdout, "server: ", log.LstdFlags)

	srv := server.NewServer(logger)

	logger.Println("Сервер запускается на http://localhost:7540 ...")
	if err := srv.HTTPServer.ListenAndServe(); err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
