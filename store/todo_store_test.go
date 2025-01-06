package store

import (
	"context"
	"just-do-it-2/todo"
	"reflect"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func TestSet(t *testing.T) {
	ctx := context.Background()

	t.Run("correctly set new todo", func(t *testing.T) {
		newTodo := todo.Todo{
			ID:          uuid.NewString(),
			Title:       "new-test-todo-1",
			IsCompleted: false,
		}
		id, err := mockStore.Set(ctx, newTodo)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := uuid.Parse(id); err != nil {
			t.Fatal(err)
		}

		t.Cleanup(func() {
			if _, err := mockStore.db.ExecContext(ctx, "DELETE FROM todos WHERE title = $1;", "new-test-todo-1"); err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Run("failed to set new todo, bad uuid", func(t *testing.T) {
		newTodo := todo.Todo{
			ID:          "bad-uuid-string",
			Title:       "new-test-todo-2",
			IsCompleted: false,
		}
		id, err := mockStore.Set(ctx, newTodo)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := uuid.Parse(id); err == nil {
			t.Error("expected bad uuid error occur, but didn't")
		}

		t.Cleanup(func() {
			if _, err := mockStore.db.ExecContext(ctx, "DELETE FROM todos WHERE uuid = $1;", "bad-uuid-string"); err != nil {
				t.Fatal(err)
			}
		})
	})

}

func TestFindByID(t *testing.T) {
	ctx := context.Background()

	t.Run("success found todo", func(t *testing.T) {
		id := "AF562BB4-6E5A-4A09-8BE5-A4F8C40061E5"
		got, err := mockStore.FindByID(ctx, id)
		if err != nil {
			t.Fatal(err)
		}

		if got.ID != id {
			t.Errorf("expected %s, got %s", id, got.ID)
		}
		if got.Title != "test-todo-1" {
			t.Errorf("expected %q, got %q", "test-todo-1", got.Title)
		}
		if got.IsCompleted != false {
			t.Errorf("expected %t, got %t", false, got.IsCompleted)
		}
	})

	t.Run("failed to get todo, not exist id", func(t *testing.T) {
		id := "not-found-id"
		_, err := mockStore.FindByID(ctx, id)
		if err == nil {
			t.Error("expected err occurs, but didn't")
		}
	})
}

var testTodoNum = 3

func TestFindAll(t *testing.T) {
	ctx := context.Background()

	t.Run("success find all todos", func(t *testing.T) {
		want := make([]todo.Todo, testTodoNum)
		want[0] = todo.Todo{ID: "AF562BB4-6E5A-4A09-8BE5-A4F8C40061E5", Title: "test-todo-1", IsCompleted: false}
		want[1] = todo.Todo{ID: "3E5B1DBB-19B5-4C7D-BF74-68D2F5DE1008", Title: "test-todo-2", IsCompleted: true}
		want[2] = todo.Todo{ID: "0BBF4EA6-CAF4-4AAB-B361-95A984D12412", Title: "test-todo-3", IsCompleted: false}

		got, err := mockStore.FindAll(ctx)
		if err != nil {
			t.Errorf("failed to get all: %q", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v, got %v", want, got)
		}
	})
}

func TestDeleteByID(t *testing.T) {
	ctx := context.Background()

	t.Run("success delete todo", func(t *testing.T) {
		id := "AF562BB4-6E5A-4A09-8BE5-A4F8C40061E5"
		err := mockStore.DeleteByID(ctx, id)
		if err != nil {
			t.Errorf("failed to delete: %q", err)
		}
	})

	t.Run("failed to delete todo, not exist id", func(t *testing.T) {
		id := "not-found-id"
		err := mockStore.DeleteByID(ctx, id)
		if err == nil {
			t.Error("expected err occurs, but didn't")
		}
	})
}
