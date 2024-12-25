package server

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

var (
	errParseResponse = errors.New("handler: failed to parse response")
	errParseRequest  = errors.New("handler: failed to parse request")
	errDeleteData    = errors.New("handler: failed to delete data")
)

const jsonContentType = "application/json; charset=utf-8"

func (srv *TodoServer) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todoRequest TodoTitleRequest
	if err := json.NewDecoder(r.Body).Decode(&todoRequest); err != nil {
		http.Error(w, errParseRequest.Error(), http.StatusInternalServerError)
		slog.Error("parse json request error", slog.String("err", err.Error()))
		return
	}

	newTodo := srv.todoService.CreateTodo(todoRequest.Title)

	res := TodoIDResponse{
		ID: newTodo.ID,
	}

	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, errParseResponse.Error(), http.StatusInternalServerError)
		slog.Error("parse json response error", slog.String("err", err.Error()))
		return
	}
}

const maxTodoNum = 3

func (srv *TodoServer) getAllTodoHandler(w http.ResponseWriter, r *http.Request) {
	todos := srv.todoService.GetAllTodo(maxTodoNum)

	res := TodoListResponse{Todos: make([]TodoResponse, len(todos))}
	for i := 0; i < len(todos); i++ {
		res.Todos[i] = TodoResponse{
			ID:        todos[i].ID,
			Title:     todos[i].Title,
			Completed: todos[i].CompletedStr(),
		}
	}

	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, errParseResponse.Error(), http.StatusInternalServerError)
		slog.Error("parse json response error", slog.String("err", err.Error()))
		return
	}
}

func (srv *TodoServer) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := r.PathValue(PathValueID)

	id := srv.todoService.CompleteTodo(todoID)

	res := TodoIDResponse{
		ID: id,
	}
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, errParseResponse.Error(), http.StatusInternalServerError)
		slog.Error("parse json response error", slog.String("err", err.Error()))
		return
	}
}

func (srv *TodoServer) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := r.PathValue(PathValueID)

	ok := srv.todoService.DeleteTodo(todoID)
	if !ok {
		http.Error(w, errParseResponse.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
}
