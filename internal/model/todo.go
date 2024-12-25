package model

type Todo struct {
	ID        string
	Title     string
	Completed bool
}

func (t *Todo) CompletedStr() string {
	if t.Completed {
		return "Done!"
	} else {
		return "Not Yet."
	}
}

type ITodoStore interface {
	Save(Todo)
	GetAll(num int) []Todo
	FindByID(id string) Todo
	DeleteByID(id string) bool
}
