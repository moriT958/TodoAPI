package server

import (
	"encoding/json"
	"errors"
	"log"
	"log/slog"
	"net/http"
)

var (
	errParseResponse = errors.New("handler: failed to parse response")
	errParseRequest  = errors.New("handler: failed to parse request")
	errDeleteData    = errors.New("handler: failed to delete data")
	errCreateData    = errors.New("handler: failed to create data")
	errGetData       = errors.New("handler: failed to get data")
	errUpdateData    = errors.New("handler: failed to update data")
)

const jsonContentType = "application/json; charset=utf-8"

func (srv *TodoServer) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todoRequest TodoTitleRequest
	if err := json.NewDecoder(r.Body).Decode(&todoRequest); err != nil {
		http.Error(w, errParseRequest.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	newTodo, err := srv.todoService.CreateTodo(todoRequest.Title)
	if err != nil {
		http.Error(w, errCreateData.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

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

const todoOffsetNum = 0
const todoLimitNum = 3

func (srv *TodoServer) GetAllTodoHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := srv.todoService.GetAllTodo(todoOffsetNum, todoLimitNum)
	if err != nil {
		http.Error(w, errGetData.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

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

func (srv *TodoServer) CompleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := r.PathValue(PathValueID)

	id, err := srv.todoService.CompleteTodo(todoID)
	if err != nil {
		http.Error(w, errUpdateData.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

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

func (srv *TodoServer) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := r.PathValue(PathValueID)

	if err := srv.todoService.DeleteTodo(todoID); err != nil {
		http.Error(w, errDeleteData.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
}
