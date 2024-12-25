package main

import (
	"just-do-it-2/config"
	"just-do-it-2/internal/server"
	"just-do-it-2/internal/store"
	"just-do-it-2/internal/todo"
	"log"
	"log/slog"
	"os"
)

func init() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func main() {
	if err := config.Load("config.json"); err != nil {
		log.Fatal(err)
	}

	store := store.New()
	todoSevice := todo.NewTodoService(store)
	server := server.New(todoSevice)

	server.Run()
}
