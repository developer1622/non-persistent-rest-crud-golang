package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", SayHello).Methods("GET")
	router.HandleFunc("/api/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", GetTaskByID).Methods("GET")
	router.HandleFunc("/api/tasks", AddTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", UpdateTask)

	http.ListenAndServe(":8443", router)
}
