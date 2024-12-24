package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"just-do-it/config"
	"just-do-it/logger"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/google/uuid"
)

// todo server
type TodoServer struct {
	http.Server
	store IInMemoryStore
}

// constructor
func NewTodoServer(store IInMemoryStore) *TodoServer {
	srv := new(TodoServer)

	mux := http.NewServeMux()
	mux.Handle("GET /todos", http.HandlerFunc(srv.getTodoHandler))
	mux.Handle("POST /todos", http.HandlerFunc(srv.postTodoHandler))
	mux.Handle("GET /todos/{id}", http.HandlerFunc(srv.completeTodoHandler))
	mux.Handle("DELETE /todos/{id}", http.HandlerFunc(srv.deleteTodoHandler))

	srv.Addr = config.Cfg.Address
	srv.Handler = mux
	srv.store = store

	return srv
}

func (s *TodoServer) run() {
	fmt.Printf("App version %s:\nServer starting at %s\n", config.Cfg.Version, config.Cfg.Address)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	<-shutdown
	log.Println("Shutting down server...")
	if err := s.Close(); err != nil {
		log.Fatalf("Server failed to shutdonw: %v", err)
	}
}

var (
	errParseRequest    = errors.New("failed to parse request json")
	errParseResponse   = errors.New("failed to parse response json")
	errSaveTodo        = errors.New("failed to save todo data")
	errParsePathParams = errors.New("failed to parse path parameters")
)

const (
	contentType = "application/json; charset=utf-8"
)

type todoResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed string `json:"completed"`
}

func (s *TodoServer) getTodoHandler(w http.ResponseWriter, r *http.Request) {
	todos := s.store.getAll()
	res := make([]todoResponse, len(todos))

	for i, t := range todos {
		res[i] = todoResponse{
			ID:        t.ID,
			Title:     t.Title,
			Completed: t.CompletedStr(),
		}
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, errParseResponse.Error(), http.StatusInternalServerError)
		logger.Log.Error(err.Error())
		return
	}
}

type postTodoReq struct {
	Title string `json:"title"`
}

func (s *TodoServer) postTodoHandler(w http.ResponseWriter, r *http.Request) {
	var postReq postTodoReq
	if err := json.NewDecoder(r.Body).Decode(&postReq); err != nil {
		http.Error(w, errParseRequest.Error(), http.StatusBadRequest)
		logger.Log.Error(err.Error(), slog.String("result", "failed"))
		return
	}

	newTodo := Todo{
		ID:        uuid.NewString(),
		Title:     postReq.Title,
		Completed: false,
	}
	s.store.save(newTodo)

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusCreated)
	logger.Log.Info("POST Create todo request",
		slog.String("result", "success"))
}

const idPathKey = "id"

func (s *TodoServer) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue(idPathKey)
	todo := s.store.getByID(id)
	todo.Completed = true
	s.store.save(todo)

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)

	res := todoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.CompletedStr(),
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		logger.Log.Error(err.Error())
		http.Error(w, errParseResponse.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *TodoServer) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue(idPathKey)
	s.store.deleteByID(id)

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
}
