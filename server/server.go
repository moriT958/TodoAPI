package server

import (
	"fmt"
	"just-do-it-2/todo"
	"log"
	"net/http"
	"os"
	"os/signal"
)

const Address = "0.0.0.0:8080"

type TodoServer struct {
	http.Server
	store todo.ITodoStore
}

func NewTodoServer(s todo.ITodoStore) *TodoServer {
	svr := new(TodoServer)
	mux := http.NewServeMux()
	mux.Handle("GET /todos", http.HandlerFunc(svr.GetAllTodos))
	mux.Handle("PATCH /todos/{id}", http.HandlerFunc(svr.CompleteTodo))
	mux.Handle("DELETE /todos/{id}", http.HandlerFunc(svr.DeleteTodo))
	mux.Handle("POST /todos", http.HandlerFunc(svr.CreateTodo))

	svr.Addr = Address
	svr.Handler = mux

	svr.store = s
	return svr
}

func (s *TodoServer) Run() {
	fmt.Printf("JustDoIt Server starting on %s...\n", Address)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	<-shutdown
	if err := s.Close(); err != nil {
		log.Fatalf("server failed to shutdown gracefully: %v", err)
	}
}
