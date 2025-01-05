package server

import (
	"just-do-it-2/store"
	"net/http"
)

type TodoServer struct {
	http.Server
	store store.TodoStore
}

func NewTodoServer(s *store.TodoStore) *TodoServer {
	svr := new(TodoServer)
	mux := http.NewServeMux()
}
