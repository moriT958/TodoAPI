package store

import "just-do-it-2/internal/model"

var _ model.ITodoStore = (*TodoStore)(nil)

type TodoStore struct {
}

func New() *TodoStore {
	return &TodoStore{}
}

func (s *TodoStore) Save(todo model.Todo) {}

func (s *TodoStore) GetAll(num int) []model.Todo {
	return nil
}
func (s *TodoStore) FindByID(id string) model.Todo {
	return model.Todo{}
}
func (s TodoStore) DeleteByID(id string) bool {
	return false
}
