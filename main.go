package main

import (
	"database/sql"
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

	//store := store.NewTodoStore(db)
}
