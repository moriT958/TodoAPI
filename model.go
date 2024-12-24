package main

// todo model
type Todo struct {
	ID        string
	Title     string
	Completed bool
}

func (t *Todo) CompletedStr() string {
	if t.Completed {
		return "Done!"
	} else {
		return "Not yet."
	}
}

// store interface
type IInMemoryStore interface {
	save(Todo) (id string)
	getAll() []Todo
	getByID(id string) Todo
	deleteByID(id string)
}
