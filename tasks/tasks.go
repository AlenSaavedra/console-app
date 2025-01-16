package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Task is a struct that represents a task in the task list
type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

// ListTasks lists all the tasks in the task list
func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("You have no tasks")
		return
	}

	for _, task := range tasks {

		status := ""

		if task.Complete {
			status = "✔"
		}
		fmt.Printf("%d: %s %s\n", task.ID, task.Name, status)
	}

}

// AddTask adds a task to the task list
func AddTask(tasks []Task, name string) []Task {
	newTask := Task{
		ID:       GetNextID(tasks),
		Name:     name,
		Complete: false,
	}

	tasks = append(tasks, newTask)
	return tasks

}

func CompleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Complete = true
			break
		}
	}
	return tasks
}

// DeleteTask deletes a task from the task list
func DeleteTask(tasks []Task, id int) []Task {
	// Loop through the tasks to find the task with the given ID
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks
		}
	}
	return tasks
}

// SaveTasks saves the tasks to a file
func SaveTasks(file *os.File, tasks []Task) {
	// Convert the tasks to a JSON string
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0) // Move the cursor to the start of the file
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0) // Clear the file
	if err != nil {
		panic(err)
	}

	// Create a writer
	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	// Flush the writer
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}

// GetNextID gets the next ID for a task
func GetNextID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}
