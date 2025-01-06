package server

import (
	"context"
	"encoding/json"
	"errors"
	"just-do-it-2/todo"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type stubTodoStore struct{}

var _ todo.ITodoStore = (*stubTodoStore)(nil)

func (st *stubTodoStore) Set(ctx context.Context, todo todo.Todo) (string, error) {
	return "FB6514BC-E78C-4570-B5A0-120DFFD9A56E", nil
}

var testTodoNum = 3

func (st *stubTodoStore) FindAll(ctx context.Context) ([]todo.Todo, error) {
	todos := make([]todo.Todo, testTodoNum)
	todos[0] = todo.Todo{
		ID:          "00C867A8-2202-44E1-9FB9-3023549B5A11",
		Title:       "test-todo-0",
		IsCompleted: false,
	}
	todos[1] = todo.Todo{
		ID:          "092C0E4F-301B-4614-8232-6BA0CFD9DF22",
		Title:       "test-todo-1",
		IsCompleted: true,
	}
	todos[2] = todo.Todo{
		ID:          "1D0FA1F2-851C-462B-8CAC-DFC112EC9F3D",
		Title:       "test-todo-2",
		IsCompleted: false,
	}
	return todos, nil
}

func (st *stubTodoStore) FindByID(ctx context.Context, id string) (todo.Todo, error) {
	if id != "00C867A8-2202-44E1-9FB9-3023549B5A11" {
		return todo.Todo{}, errors.New("input id should be 00C867A8-2202-44E1-9FB9-3023549B5A11")
	}
	todo := todo.Todo{
		ID:          "00C867A8-2202-44E1-9FB9-3023549B5A11",
		Title:       "test-todo-0",
		IsCompleted: false,
	}
	return todo, nil
}

func (st *stubTodoStore) DeleteByID(ctx context.Context, id string) error {
	return nil
}

func initMockServer() *TodoServer {
	svr := NewTodoServer(new(stubTodoStore))
	return svr
}

func TestGetAllTodos(t *testing.T) {
	svr := initMockServer()
	defer svr.Close()

	req, err := http.NewRequest(http.MethodGet, "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()

	ctx := context.Background()
	t.Run("正常系テスト", func(t *testing.T) {
		want := RespGetAllTodos{make([]RespTodo, testTodoNum)}
		todos, _ := svr.store.FindAll(ctx)
		for i, todo := range todos {
			want.Todos[i] = RespTodo{
				ID:          todo.ID,
				Title:       todo.Title,
				IsCompleted: todo.IsCompletedString(),
			}
		}

		svr.GetAllTodos(res, req)
		got := RespGetAllTodos{make([]RespTodo, 0)}
		if err := json.Unmarshal(res.Body.Bytes(), &got); err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %v, but got %v", want, got)
		}

		if res.Code != http.StatusOK {
			t.Errorf("status code expected %d, got %d", http.StatusOK, res.Code)
		}
	})
}

func TestCompleteTodo(t *testing.T) {
	svr := initMockServer()
	defer svr.Close()

	req, err := http.NewRequest(http.MethodGet, "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue(PathValueID, "00C867A8-2202-44E1-9FB9-3023549B5A11")
	res := httptest.NewRecorder()

	t.Run("correctly returns status ok", func(t *testing.T) {
		svr.CompleteTodo(res, req)

		want := http.StatusOK
		got := res.Code
		if want != got {
			t.Errorf("want %d, but got %d", want, got)
		}
	})
}

func TestDeleteTodo(t *testing.T) {

}

func TestCreateTodo(t *testing.T) {

}
