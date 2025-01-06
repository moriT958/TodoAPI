package server

type RespTodo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	IsCompleted string `json:"isCompleted"`
}

type RespGetAllTodos struct {
	Todos []RespTodo `json:"todos"`
}

const PathValueID = "id"

type RespDeleteTodo struct {
}

type ReqCreateTodo struct {
}

type RespCreateTodo struct {
}
