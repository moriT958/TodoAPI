package store

import (
	"just-do-it-2/internal/model"
	"reflect"
	"testing"
)

func TestSave(t *testing.T) {
	t.Run("Save a new Todo", func(t *testing.T) {
		store := New()
		todo := model.Todo{ID: "test-uuid", Title: "test-task", Completed: false}

		store.Save(todo)

		if len(store.mem) != 1 {
			t.Errorf("expected mem size to be 1, got %d", len(store.mem))
		}

		if len(store.mem) != store.count {
			t.Errorf("expected mem size is same to count.")
		}

		if store.mem[0].ID != todo.ID || store.mem[0].Title != todo.Title || store.mem[0].Completed != todo.Completed {
			t.Errorf("expected todo to be %v, got %v", todo, store.mem[0])
		}
	})

	t.Run("Update an existing Todo", func(t *testing.T) {
		store := New()
		todo := model.Todo{ID: "test-uuid", Title: "test-task", Completed: false}
		store.Save(todo)

		todoUpdated := model.Todo{ID: "test-uuid", Title: "test-task-updated", Completed: false}
		store.Save(todoUpdated)

		if len(store.mem) != 1 {
			t.Errorf("expected mem size to be 1 after update, got %d", len(store.mem))
		}

		if store.mem[0].Title != todoUpdated.Title {
			t.Errorf("expected updated title to be %q, got %q", todoUpdated.Title, store.mem[0].Title)
		}
	})

	t.Run("Add a second Todo", func(t *testing.T) {
		store := New()
		todo1 := model.Todo{ID: "test-uuid-1", Title: "test-task-1", Completed: false}
		todo2 := model.Todo{ID: "test-uuid-2", Title: "test-task-2", Completed: false}
		store.Save(todo1)
		store.Save(todo2)

		if len(store.mem) != 2 {
			t.Errorf("expected mem size to be 2, got %d", len(store.mem))
		}

		if store.mem[1].ID != todo2.ID || store.mem[1].Title != todo2.Title || store.mem[1].Completed != todo2.Completed {
			t.Errorf("expected todo to be %v, got %v", todo2, store.mem[1])
		}
	})

	t.Run("Avoid duplicate Todo with same ID", func(t *testing.T) {
		store := New()
		todo1 := model.Todo{ID: "test-uuid-1", Title: "test-task-1", Completed: false}
		store.Save(todo1)

		todo2 := model.Todo{ID: "test-uuid-1", Title: "test-task-2", Completed: false}
		store.Save(todo2)

		if len(store.mem) != 1 {
			t.Errorf("expected mem size to remain 1 after duplicate save, got %d", len(store.mem))
		}

		if store.mem[0].Title != todo2.Title {
			t.Errorf("expected updated title to be %q, got %q", todo2.Title, store.mem[0].Title)
		}
	})
}

func TestGetAll(t *testing.T) {
	store := New()
	todos := []model.Todo{
		{ID: "test-uuid-1", Title: "test-task-1", Completed: false},
		{ID: "test-uuid-2", Title: "test-task-2", Completed: false},
		{ID: "test-uuid-3", Title: "test-task-3", Completed: false},
	}
	store.Save(todos[0])
	store.Save(todos[1])
	store.Save(todos[2])

	tests := []struct {
		name   string
		offset int
		limit  int
		want   []model.Todo
	}{
		{name: "valid offset and limit should return all todos", offset: 0, limit: 3, want: todos},
		{name: "invalid offset should return nil", offset: -1, limit: 3, want: nil},
		{name: "invalid limit should return nil", offset: 0, limit: -1, want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := store.GetAll(tt.offset, tt.limit)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected %v, but got %v", tt.want, got)
			}
		})
	}
}

func TestFindByID(t *testing.T) {
	store := New()
	todo := model.Todo{ID: "test-uuid-1", Title: "test-task-1", Completed: false}
	store.Save(todo)

	testcases := []struct {
		name string
		id   string
		want model.Todo
	}{
		{name: "should return correct todo", id: "test-uuid-1", want: todo},
		{name: "not exit id should return nothin", id: "wrong-uuid-1", want: model.Todo{}},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			gotTodo := store.FindByID(tt.id)
			if gotTodo != tt.want {
				t.Errorf("expected %v, got %v", tt.want, gotTodo)
			}
		})
	}
}

func TestDeleteByID(t *testing.T) {
	store := New()
	todo1 := model.Todo{ID: "test-uuid-1", Title: "test-task-1", Completed: false}
	store.Save(todo1)

	tests := []struct {
		name string
		id   string
		want bool
	}{
		{name: "valid id should return true", id: "test-uuid-1", want: true},
		{name: "invalid id should return false", id: "wrong-uuid-1", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := store.DeleteByID(tt.id); got != tt.want {
				t.Errorf("expected %t, got %t", tt.want, got)
			}
		})
	}

}
