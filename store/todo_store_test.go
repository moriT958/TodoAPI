package store

import (
	"context"
	"database/sql"
	"just-do-it-2/todo"
	"os"
	"reflect"
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

	todos := []todo.Todo{
		{ID: "test-uuid-1", Title: "test-todo-1", IsCompleted: false},
		{ID: "test-uuid-2", Title: "test-todo-2", IsCompleted: true},
		{ID: "test-uuid-3", Title: "test-todo-3", IsCompleted: false},
	}

	_, err = store.Set(ctx, todos[0])
	_, err = store.Set(ctx, todos[1])
	_, err = store.Set(ctx, todos[2])
	if err != nil {
		t.Fatal(err)
	}

	t.Run("success found todo", func(t *testing.T) {
		id := "test-uuid-1"
		got, err := store.FindByID(ctx, id)
		if err != nil {
			t.Fatal(err)
		}

		if got.ID != id {
			t.Errorf("expected %s, got %s", id, got.ID)
		}
		if got.Title != todos[0].Title {
			t.Errorf("expected %q, got %q", todos[0].Title, got.Title)
		}
		if got.IsCompleted != todos[0].IsCompleted {
			t.Errorf("expected %t, got %t", todos[0].IsCompleted, got.IsCompleted)
		}
	})

	t.Run("failed to get todo, not exist id", func(t *testing.T) {
		id := "not-found-id"
		_, err := store.FindByID(ctx, id)
		if err == nil {
			t.Error("expected err occurs, but didn't")
		}
	})
}

func TestFindAll(t *testing.T) {
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

	todos := []todo.Todo{
		{ID: "test-uuid-1", Title: "test-todo-1", IsCompleted: false},
		{ID: "test-uuid-2", Title: "test-todo-2", IsCompleted: true},
		{ID: "test-uuid-3", Title: "test-todo-3", IsCompleted: false},
	}

	_, err = store.Set(ctx, todos[0])
	_, err = store.Set(ctx, todos[1])
	_, err = store.Set(ctx, todos[2])
	if err != nil {
		t.Fatal(err)
	}

	t.Run("success find all todos", func(t *testing.T) {
		tl, err := store.FindAll(ctx)
		if err != nil {
			t.Errorf("failed to get all: %q", err)
		}

		if !reflect.DeepEqual(tl, todos) {
			t.Errorf("expected %v, got %v", todos, tl)
		}
	})
}

func TestDeleteByID(t *testing.T) {
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

	todos := []todo.Todo{
		{ID: "test-uuid-1", Title: "test-todo-1", IsCompleted: false},
		{ID: "test-uuid-2", Title: "test-todo-2", IsCompleted: true},
		{ID: "test-uuid-3", Title: "test-todo-3", IsCompleted: false},
	}

	_, err = store.Set(ctx, todos[0])
	_, err = store.Set(ctx, todos[1])
	_, err = store.Set(ctx, todos[2])
	if err != nil {
		t.Fatal(err)
	}

	t.Run("success delete todo", func(t *testing.T) {
		id := "test-uuid-1"
		err := store.DeleteByID(ctx, id)
		if err != nil {
			t.Errorf("failed to delete: %q", err)
		}
	})

	t.Run("failed to delete todo, not exist id", func(t *testing.T) {
		id := "not-found-id"
		err := store.DeleteByID(ctx, id)
		if err == nil {
			t.Error("expected err occurs, but didn't")
		}
	})
}
