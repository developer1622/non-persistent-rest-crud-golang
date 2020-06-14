package utils

import "reflect"

// Task holds all the info about task :(.
type Task struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}

// Tasks holds all tasks.
var Tasks = []Task{}

// NewTask is constructor for creating new task.
func NewTask(id, name, desc string) Task { return Task{ID: id, Name: name, Description: desc} }

// TaskbyID fetches the task by ID.
func TaskbyID(id string) Task {
	for _, task := range Tasks {
		if task.ID == id {
			return task
		}
	}

	// could not find anything , so returning empty task
	return Task{}
}

// Equals compares 2 IDs.
func Equals(t, t2 Task) bool {
	return reflect.DeepEqual(t, t2)
}

// UpdateTask updates the task.
func UpdateTask(id, name, desc string) Task {
	for _, task := range Tasks {
		if task.ID == id {
			task.Name = name
			task.Description = desc
			return task
		}
	}
	return NewTask("1", "name", "desc, this is error")
}
