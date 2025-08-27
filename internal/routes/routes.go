package routes

import (
	"github.com/gorilla/mux"
	"task-manager/internal/handlers"
)

func SetupTaskRoutes(router *mux.Router, taskHandler *handlers.TaskHandler) {
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskHandler.GetTaskByID).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
}
