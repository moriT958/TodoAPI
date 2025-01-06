package server

import (
	"encoding/json"
	"errors"
	"just-do-it-2/todo"
	"net/http"

	"github.com/google/uuid"
)

const jsonContentType = "application/json; charset=utf-8"

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
	ctx := r.Context()

	id := r.PathValue(PathValueID)
	t, err := s.store.FindByID(ctx, id)
	if (err != nil) || (t == todo.Todo{}) {
		http.Error(w, errors.New("todo doesn't exist").Error(), http.StatusBadRequest)
		return
	}

	t.IsCompleted = true
	if _, err := s.store.Set(ctx, t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
}

func (s *TodoServer) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue(PathValueID)
	if err := s.store.DeleteByID(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
}

func (s *TodoServer) CreateTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var reqSchema ReqCreateTodo
	if err := json.NewDecoder(r.Body).Decode(&reqSchema); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newTodo := todo.Todo{
		ID:          uuid.NewString(),
		Title:       reqSchema.Title,
		IsCompleted: false,
	}

	id, err := s.store.Set(ctx, newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusCreated)

	res := RespCreateTodo{
		ID: id,
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
