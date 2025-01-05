package store

import (
	"context"
	"just-do-it-2/todo"
)

type TodoStore struct {
	db DBTX
}

var _ todo.ITodoStore = (*TodoStore)(nil)

func NewTodoStore(dbtx DBTX) *TodoStore {
	return &TodoStore{
		db: dbtx,
	}
}

func (s *TodoStore) Set(ctx context.Context, todo todo.Todo) (string, error) {
	var isExist bool
	if err := s.db.QueryRowContext(ctx, "SELECT EXISTS (SELECT * FROM todos WHERE uuid = $1);", todo.ID).Scan(&isExist); err != nil {
		return "", err
	}

	if isExist {
		if _, err := s.db.ExecContext(ctx, "UPDATE todos SET (title, is_completed) = ($1, $2);", todo.Title, todo.IsCompleted); err != nil {
			return "", err
		}
		return todo.ID, nil
	}

	if _, err := s.db.ExecContext(ctx, "INSERT INTO todos (uuid, title) VALUES ($1, $2);", todo.ID, todo.Title); err != nil {
		return "", err
	}

	return todo.ID, nil
}

func (s *TodoStore) FindByID(ctx context.Context, id string) (todo.Todo, error) {
	return todo.Todo{}, nil
}

func (s *TodoStore) DeleteByID(ctx context.Context, id string) error {
	return nil
}
