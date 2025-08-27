package main

import (
	"log"
	"net/http"
	"task-manager/internal/db"
	"task-manager/internal/handlers"
	"task-manager/internal/repo"
	"task-manager/internal/routes"

	"github.com/gorilla/mux"
)

func main() {
	conn := db.InitDB()
	db.RunMigrations()
	defer conn.Close()

	taskRepo := repo.NewTaskRepository(conn)
	taskHandler := &handlers.TaskHandler{Repo: taskRepo}

	router := mux.NewRouter()
	routes.SetupTaskRoutes(router, taskHandler)

	router.HandleFunc("/register", handlers.RegisterHandler(conn)).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler(conn)).Methods("POST")

	log.Println("Connected to PostgreSQL!")
	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
