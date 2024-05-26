package routes

import (
	"golang-react-todo/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/tasks/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/tasks/deleteAll/{id}", middleware.DeleteAllTasks).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/tasks/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	return router
}
