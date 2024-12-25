package server

type TodoTitleRequest struct {
	Title string `json:"title"`
}

type TodoIDResponse struct {
	ID string `json:"id"`
}

type TodoResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed string `json:"completed"`
}

type TodoListResponse struct {
	Todos []TodoResponse `json:"todos"`
}

const PathValueID = "id"
