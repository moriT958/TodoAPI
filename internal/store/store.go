package store

import (
	"errors"
	"just-do-it-2/internal/model"
	"log"
	"sync"
)

var _ model.ITodoStore = (*TodoStore)(nil)

type TodoStore struct {
	sync.RWMutex
	count int
	mem   []model.Todo
}

func New() *TodoStore {
	mem := make([]model.Todo, 0)
	return &TodoStore{mem: mem}
}

func (s *TodoStore) Save(todo model.Todo) {
	s.Lock()
	defer s.Unlock()

	for i := 0; i < s.count; i++ {
		if s.mem[i].ID == todo.ID {
			s.mem[i] = todo
			return
		}
	}
	s.mem = append(s.mem, todo)
	s.count++
}

func (s *TodoStore) GetAll(offset, limit int) []model.Todo {
	s.RLock()
	defer s.RUnlock()

	if err := s.checkCount(offset, limit); err != nil {
		log.Println(err)
		return nil
	}

	// size is a number of data that is able to be get
	size := s.count - offset

	var todos []model.Todo

	if limit > size {
		todos = make([]model.Todo, size)
	} else {
		todos = make([]model.Todo, limit)
	}

	idx := 0
	for i := offset; i < size; i++ {
		todos[idx] = s.mem[i]
		idx++
	}
	return todos
}

func (s *TodoStore) checkCount(offset, limit int) error {
	if offset > s.count || offset < 0 {
		return errors.New("invalid offset value")
	}

	if limit < 0 {
		return errors.New("invalid limit value")
	}

	return nil
}

func (s *TodoStore) FindByID(id string) model.Todo {
	s.RLock()
	defer s.RUnlock()

	for i := 0; i < s.count; i++ {
		if s.mem[i].ID == id {
			return s.mem[i]
		}
	}
	return model.Todo{}
}
func (s *TodoStore) DeleteByID(id string) bool {
	s.Lock()
	defer s.Unlock()

	for i := 0; i < s.count; i++ {
		if s.mem[i].ID == id {
			s.mem = append(s.mem[:i], s.mem[i+1:]...)
			s.count--
			return true
		}
	}

	return false
}
