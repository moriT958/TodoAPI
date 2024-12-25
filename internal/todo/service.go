package todo

import (
	"errors"
	"just-do-it-2/internal/model"

	"github.com/google/uuid"
)

type TodoService struct {
	store model.ITodoStore
}

func NewTodoService(s model.ITodoStore) *TodoService {
	return &TodoService{store: s}
}

func (ts *TodoService) CreateTodo(title string) (model.Todo, error) {
	newTodo := model.Todo{
		ID:        uuid.NewString(),
		Title:     title,
		Completed: false,
	}
	ts.store.Save(newTodo)
	return newTodo, nil
}

func (ts *TodoService) GetAllTodo(offset, limit int) ([]model.Todo, error) {
	todos := ts.store.GetAll(offset, limit)
	return todos, nil
}

func (ts *TodoService) CompleteTodo(todoID string) (string, error) {
	todo := ts.store.FindByID(todoID)
	if todo == (model.Todo{}) {
		return "", errors.New("todo not exists")
	}
	todo.Completed = true
	ts.store.Save(todo)
	return todo.ID, nil
}

func (ts *TodoService) DeleteTodo(todoID string) error {
	if ok := ts.store.DeleteByID(todoID); ok {
		return nil
	}
	return errors.New("failed to delete data")
}
