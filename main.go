package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"just-do-it-2/server"
	"just-do-it-2/store"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := store.NewTodoStore(db)
	svr := server.NewTodoServer(store)

	svr.Run()
}
