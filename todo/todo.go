package todo

import "context"

type Todo struct {
	ID          string
	Title       string
	IsCompleted bool
}

func (t *Todo) IsCompletedString() string {
	if t.IsCompleted {
		return "Done!"
	}

	return "not Done."
}

type ITodoStore interface {
	Set(ctx context.Context, todo Todo) (id string, err error)
	FindAll(ctx context.Context) ([]Todo, error)
	FindByID(ctx context.Context, id string) (Todo, error)
	DeleteByID(ctx context.Context, id string) error
}
