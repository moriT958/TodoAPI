package store_test

import (
	"just-do-it-2/internal/model"
	"just-do-it-2/internal/store/testdata"
	"log"
	"reflect"
	"testing"
)

func TestSave(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want int
	}{
		{name: "save correctly", id: "correct-uuid", want: 4},
		{name: "add an existing todo", id: "test-uuid-2", want: 3},
	}

	t.Run("correctly save init data", func(t *testing.T) {
		store := testdata.InitTestStore()
		if len(store.Mem()) != store.Count() {
			t.Errorf("not match data num(%d) and count(%d)", len(store.Mem()), store.Count())
		}
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := testdata.InitTestStore()
			store.Save(model.Todo{ID: tt.id})
			if len(store.Mem()) != tt.want {
				log.Println(store.Mem())
				t.Errorf("data num expected %d, but got %d", tt.want, len(store.Mem()))
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	tests := []struct {
		name   string
		offset int
		limit  int
		want   []model.Todo
	}{
		{name: "valid offset and limit should return all todos", offset: 0, limit: 3, want: testdata.Todos},
		{name: "invalid offset should return nil", offset: -1, limit: 3, want: nil},
		{name: "invalid limit should return nil", offset: 0, limit: -1, want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := testdata.InitTestStore()
			got := store.GetAll(tt.offset, tt.limit)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected %v, but got %v", tt.want, got)
			}
		})
	}
}

func TestFindByID(t *testing.T) {
	testcases := []struct {
		name string
		id   string
		want model.Todo
	}{
		{name: "should return correct todo", id: "test-uuid-1", want: testdata.Todos[0]},
		{name: "not exit id should return nothin", id: "wrong-uuid-1", want: model.Todo{}},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			store := testdata.InitTestStore()
			gotTodo := store.FindByID(tt.id)
			if gotTodo != tt.want {
				t.Errorf("expected %v, got %v", tt.want, gotTodo)
			}
		})
	}
}

func TestDeleteByID(t *testing.T) {
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
			store := testdata.InitTestStore()
			if got := store.DeleteByID(tt.id); got != tt.want {
				t.Errorf("expected %t, got %t", tt.want, got)
			}
		})
	}

}
