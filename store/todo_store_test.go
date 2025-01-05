package store

import (
	"context"
	"database/sql"
	"just-do-it-2/todo"
	"os"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var testDsn = os.Getenv("DATABASE_URL")

func TestSet(t *testing.T) {
	ctx := context.Background()

	db, err := sql.Open("postgres", testDsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	store := NewTodoStore(tx)

	t.Run("correctly set new todo", func(t *testing.T) {
		newTodo := todo.Todo{
			ID:          uuid.NewString(),
			Title:       "test-todo-1",
			IsCompleted: false,
		}
		id, err := store.Set(ctx, newTodo)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := uuid.Parse(id); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("failed to set new todo, bad uuid", func(t *testing.T) {
		newTodo := todo.Todo{
			ID:          "bad-uuid-string",
			Title:       "test-todo-2",
			IsCompleted: false,
		}
		id, err := store.Set(ctx, newTodo)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := uuid.Parse(id); err == nil {
			t.Error("expected bad uuid error occur, but didn't")
		}
	})
}

func TestFindByID(t *testing.T) {
}

func TestDeleteByID(t *testing.T) {
}
