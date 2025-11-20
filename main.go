package main

import (
	"fmt"
	"net/http"
	routes "todo-rest-go/routes"
	model "todo-rest-go/todo"
)

var serverPort uint16 = 8080

func main() {
	manager := model.TaskManager{}
	manager.SetFile("tasks.json")

	if err := manager.Load(); err != nil {
		fmt.Printf("Warning: Failed to load data with error: %v", err)
	}

	server := routes.CreateAPIServer(&manager)
	mux := http.NewServeMux()
	// cors handler
	corsHandler := routes.AllowCORS(mux)

	server.RegisterRoutes(mux)

	fmt.Printf("Server is running on port %d\n", serverPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), corsHandler); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
}
