package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	todo "todo-rest-go/todo"
)

type APIServer struct {
	Manager *todo.TaskManager
}

func CreateAPIServer(manager *todo.TaskManager) *APIServer {
	return &APIServer{Manager: manager}
}

func (server *APIServer) HandleGetTodo(res http.ResponseWriter, req *http.Request) {
	taskList := server.Manager.Get()

	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(taskList); err != nil {
		http.Error(res, "Failed to retrieve data.", http.StatusInternalServerError)
	}
}

func (server *APIServer) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello world!")
	})

	mux.HandleFunc("GET /todos", server.HandleGetTodo)
}
