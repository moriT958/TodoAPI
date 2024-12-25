package todo

import "just-do-it-2/internal/model"

type TodoService struct {
	store model.ITodoStore
}

func NewTodoService(s model.ITodoStore) *TodoService {
	return &TodoService{store: s}
}
