package main

// store
type InMemoryStore struct {
	storage string
	mem     []Todo
	count   int
}

// constructor
func NewInMemoryStore(filename string) *InMemoryStore {
	return &InMemoryStore{
		storage: filename,
		mem:     make([]Todo, 0),
	}
}

func (s *InMemoryStore) save(newTodo Todo) string {
	for _, todo := range s.mem {
		if todo.ID == newTodo.ID {
			todo = newTodo
			return newTodo.ID
		}
	}

	s.mem = append(s.mem, newTodo)
	s.count++
	return newTodo.ID
}

func (s *InMemoryStore) getAll() []Todo {
	return s.mem
}

func (s *InMemoryStore) getByID(id string) Todo {
	for i := range s.mem {
		if id == s.mem[i].ID {
			return s.mem[i]
		}
	}
	return Todo{}
}

func (s *InMemoryStore) deleteByID(id string) {
	for i := range s.mem {
		if id == s.mem[i].ID {
			s.mem = append(s.mem[:i], s.mem[i+1:]...)
		}
	}
}

//func (s *InMemoryStore) dump() error
