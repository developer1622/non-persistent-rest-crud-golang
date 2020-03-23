package main

import (
	"net/http"

	"github.com/Sample_REST_APIs_App/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.SayHello).Methods("GET")
	router.HandleFunc("/api/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", controllers.GetTaskByID).Methods("GET")
	router.HandleFunc("/api/tasks", controllers.AddTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", controllers.UpdateTask)

	http.ListenAndServe(":8443", router)
}
