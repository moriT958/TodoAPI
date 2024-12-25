package todo

import (
	"fmt"
	"just-do-it-2/internal/model"
)

type TodoService struct {
	store model.ITodoStore
}

func NewTodoService(s model.ITodoStore) *TodoService {
	return &TodoService{store: s}
}

func (ts *TodoService) CreateTodo(title string) model.Todo {
	newTodo := model.Todo{
		ID:        "test-uuid-1",
		Title:     title,
		Completed: false,
	}
	ts.store.Save(newTodo)
	return newTodo
}

func (ts *TodoService) GetAllTodo(num int) []model.Todo {
	todos := make([]model.Todo, num)
	todos = ts.store.GetAll(num)

	for i := 0; i < num; i++ {
		todos[i] = model.Todo{
			ID:        fmt.Sprintf("test-id-%d", i),
			Title:     fmt.Sprintf("test-title-%d", i),
			Completed: false,
		}
	}
	return todos
}

func (ts *TodoService) CompleteTodo(todoID string) string {
	todo := ts.store.FindByID(todoID)
	todo.Completed = true
	ts.store.Save(todo)
	return todo.ID
}

func (ts *TodoService) DeleteTodo(todoID string) bool {
	if ok := ts.store.DeleteByID(todoID); ok {
		return true
	}
	return false
}
