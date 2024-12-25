package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"just-do-it-2/internal/model"
	"log/slog"
	"net/http"
)

var (
	errParseResponse = errors.New("handler: failed to parse response")
	errParseRequest  = errors.New("handler: failed to parse request")
)

const jsonContentType = "application/json; charset=utf-8"

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todoRequest TodoTitleRequest
	if err := json.NewDecoder(r.Body).Decode(&todoRequest); err != nil {
		http.Error(w, errParseRequest.Error(), http.StatusInternalServerError)
		slog.Error("parse json request error", slog.String("err", err.Error()))
		return
	}

	// TODO:
	// write create todo logic here.
	newTodo := model.Todo{
		ID: "test-uuid-1",
		//Title:     todoRequest.Title,
		//Completed: false,
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

func getAllTodoHandler(w http.ResponseWriter, r *http.Request) {
	res := TodoListResponse{Todos: make([]TodoResponse, 3)}
	for i := 0; i < 3; i++ {
		res.Todos[i] = TodoResponse{
			ID:        fmt.Sprintf("test-id-%d", i),
			Title:     fmt.Sprintf("test-title-%d", i),
			Completed: fmt.Sprintf("test-completed-%d", i),
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

func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := r.PathValue(PathValueID)

	// TODO:
	// add complete todo logic here.

	res := TodoIDResponse{
		ID: todoID,
	}
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, errParseResponse.Error(), http.StatusInternalServerError)
		slog.Error("parse json response error", slog.String("err", err.Error()))
		return
	}
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	//todoID := r.PathValue(PathValueID)

	// TODO:
	// write delete todo logic here.

	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
}
