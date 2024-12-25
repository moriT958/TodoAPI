package server

import (
	"fmt"
	"just-do-it-2/config"
	"just-do-it-2/internal/todo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type TodoServer struct {
	http.Server
	todoService todo.TodoService
}

func New(ts *todo.TodoService) *TodoServer {
	srv := new(TodoServer)

	mux := http.NewServeMux()
	mux.Handle("GET /hello", http.HandlerFunc(srv.helloHandler))
	mux.Handle("POST /todos", http.HandlerFunc(srv.createTodoHandler))
	mux.Handle("GET /todos", http.HandlerFunc(srv.getAllTodoHandler))
	mux.Handle("PATCH /todos/{id}", http.HandlerFunc(srv.completeTodoHandler))
	mux.Handle("DELETE /todos/{id}", http.HandlerFunc(srv.deleteTodoHandler))
	srv.Handler = mux

	srv.Addr = config.Address()
	srv.ReadTimeout = time.Duration(config.ReadTimeout() * int64(time.Second))
	srv.WriteTimeout = time.Duration(config.WriteTimeout() * int64(time.Second))

	srv.todoService = *ts

	return srv
}

func (srv *TodoServer) Run() {
	fmt.Printf("JustDoIt Version %s:\nServer starting at %s...\n", config.Version(), config.Address())
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	<-shutdown
	fmt.Println("Sever shutdowning...")
	if err := srv.Close(); err != nil {
		log.Fatalf("Server failed to shutdown gracefully: %v", err)
	}
}

func (srv *TodoServer) helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
