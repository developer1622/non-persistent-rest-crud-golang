package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var tasks = []Task{}

// GetTasks fetches all the tasks that are availble.
func GetTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In GetTasks !!")
	writeResponse(w, Tasks)
}

// GetTaskByID fetch the task by ID.
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In GetTaskByID page")
	writeResponse(w, TaskbyID(mux.Vars(r)["id"]))
}

// AddTask is to add taks.
func AddTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In AddTask page")

	var t Task
	json.NewDecoder(r.Body).Decode(&t)
	Tasks = append(Tasks, t)
	writeResponse(w, Tasks[len(Tasks)-1])
}

// UpdateTask updates the task.
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In UpdateTask Page")

	var t Task
	json.NewDecoder(r.Body).Decode(&t)

	if Equals(Task{}, TaskbyID((mux.Vars(r)["id"]))) {
		writeResponse(w, "we could not find the ID you want to update !!")
	}
	q := r.URL.Query()

	writeResponse(w, ModelUpdateTask((mux.Vars(r)["id"]), q.Get("name"), q.Get("desc")))
}

// SayHello is starting of the page.
func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In SayHello Page !!")

	writeResponse(w, "If you see this , server is running is fine !")
}

// this is util function.
func writeResponse(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
