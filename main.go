package main

import (
	"just-do-it/config"
	"just-do-it/logger"
	"log"
)

func main() {
	logger.InitLogger()

	if err := config.InitConfig("config.json"); err != nil {
		log.Fatal(err)
	}

	store := NewInMemoryStore("todo.db.json")
	server := NewTodoServer(store)

	server.run()
}
