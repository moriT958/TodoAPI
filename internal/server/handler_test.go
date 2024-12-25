package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"just-do-it-2/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	res := httptest.NewRecorder()

	t.Run("prints hello msg", func(t *testing.T) {
		helloHandler(res, req)

		got := res.Body.String()
		want := "Hello, World!\n"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
}

var testTodo = model.Todo{
	ID:        "test-uuid-1",
	Title:     "test-todo-1",
	Completed: false,
}

func TestCreateTodoHandler(t *testing.T) {
	body := []byte(fmt.Sprintf(`{"title":%s}`, testTodo.Title))
	req, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	createTodoHandler(res, req)

	t.Run("return status created", func(t *testing.T) {
		got := res.Code
		want := http.StatusCreated
		if got != want {
			t.Errorf("response code is wrong, got %d want %d", got, want)
		}
	})

	t.Run("returns ID", func(t *testing.T) {
		var got TodoIDResponse
		if err := json.Unmarshal(res.Body.Bytes(), &got); err != nil {
			t.Log(res.Body)
			t.Fatal(err)
		}
		want := TodoIDResponse{
			ID: testTodo.ID,
		}
		if got != want {
			t.Errorf("response body is wrong, got %s want %s", got, want)
		}
	})
}
