package testdata

import (
	"just-do-it-2/internal/model"
	"just-do-it-2/internal/store"
)

var Todos = []model.Todo{
	{ID: "test-uuid-1", Title: "test-task-1", Completed: false},
	{ID: "test-uuid-2", Title: "test-task-2", Completed: false},
	{ID: "test-uuid-3", Title: "test-task-3", Completed: false},
}

func InitTestStore() *store.TodoStore {
	store := store.New()
	for _, todo := range Todos {
		store.Save(todo)
	}
	return store
}
