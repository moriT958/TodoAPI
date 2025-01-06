package server

import (
	"encoding/json"
	"just-do-it-2/store"
	"net/http"
)

const Address = "127.0.0.1:8080"

const jsonContentType = "application/json; charset=utf-8"

type TodoServer struct {
	http.Server
	store store.TodoStore
}

func NewTodoServer(s *store.TodoStore) *TodoServer {
	svr := new(TodoServer)
	mux := http.NewServeMux()
	mux.Handle("GET /todos", http.HandlerFunc(svr.GetAllTodos))
	mux.Handle("PATCH /todos/{id}", http.HandlerFunc(svr.CompleteTodo))
	mux.Handle("DELETE /todos/{id}", http.HandlerFunc(svr.DeleteTodo))
	mux.Handle("POST /todos", http.HandlerFunc(svr.CreateTodo))

	svr.Addr = Address
	svr.Handler = mux

	svr.store = *s
	return svr
}

func (s *TodoServer) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todos, err := s.store.FindAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := RespGetAllTodos{make([]RespTodo, len(todos))}
	for i, todo := range todos {
		res.Todos[i] = RespTodo{
			ID:          todo.ID,
			Title:       todo.Title,
			IsCompleted: todo.IsCompletedString(),
		}
	}

	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *TodoServer) CompleteTodo(w http.ResponseWriter, r *http.Request) {

}

func (s *TodoServer) DeleteTodo(w http.ResponseWriter, r *http.Request) {

}

func (s *TodoServer) CreateTodo(w http.ResponseWriter, r *http.Request) {

}
