package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sample_REST_APIs_App/utils"
	"github.com/gorilla/mux"
)

var tasks = []utils.Task{}

// GetTasks fetches all the tasks that are availble.
func GetTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In GetTasks !!")
	writeResponse(w, utils.Tasks)
}

// GetTaskByID fetch the task by ID.
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In GetTaskByID page")
	writeResponse(w, utils.TaskbyID(mux.Vars(r)["id"]))
}

// AddTask is to add taks.
func AddTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In AddTask page")

	var t utils.Task
	json.NewDecoder(r.Body).Decode(&t)
	utils.Tasks = append(utils.Tasks, t)
	writeResponse(w, utils.Tasks[len(utils.Tasks)-1])
}

// UpdateTask updates the task.
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In UpdateTask Page")

	var t utils.Task
	json.NewDecoder(r.Body).Decode(&t)

	if utils.Equals(utils.Task{}, utils.TaskbyID((mux.Vars(r)["id"]))) {
		writeResponse(w, "we could not find the ID you want to update !!")
	}
	q := r.URL.Query()

	writeResponse(w, utils.UpdateTask((mux.Vars(r)["id"]), q.Get("name"), q.Get("desc")))
}

// SayHello is starting of the page.
func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In SayHello Page !!")

	writeResponse(w, "Hello , this is my first program")
}

func writeResponse(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
