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

type ReqCreateTodo struct {
	Title string `json:"title"`
}

type RespCreateTodo struct {
	ID string `json:"id"`
}
