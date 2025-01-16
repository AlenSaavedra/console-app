package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/AlenSaavedra/CRUD-GO/tasks"
)

func main() {
	// Open the file in read-write mode with the file being created if it doesn't exist
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)

	// Check for errors
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var tasks []task.Task // Create a slice of tasks

	// Get the file info
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	// If the file is not empty, read the contents
	if info.Size() != 0 {

		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)

		if err != nil {
			panic(err)
		}

	} else {
		tasks = []task.Task{}

	}

	// Check the command line arguments
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	// Switch on the command line arguments
	switch os.Args[1] {
	case "list":
		task.ListTasks(tasks)
	case "add":
		// Add a task
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter task name: ")

		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks = task.AddTask(tasks, name)
		fmt.Println("Task added")
		task.SaveTasks(file, tasks)
	case "delete":
		// Delete a task
		if len(os.Args) < 3 {
			fmt.Println("You must provide a task ID to delete")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID, please provide a number")
			return
		}
		tasks = task.DeleteTask(tasks, id)
		fmt.Println("Task deleted")
		task.SaveTasks(file, tasks)
	case "complete":
		// Mark a task as complete
		if len(os.Args) < 3 {
			fmt.Println("You must provide a task ID to complete")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID, please provide a number")
			return
		}
		tasks = task.CompleteTask(tasks, id)
		fmt.Println("Task completed")
		task.SaveTasks(file, tasks)
	default:
		printUsage()
	}

}

// Print the usage of the program
func printUsage() {
	fmt.Println("Usage: go-console-crud [list | add | complete | delete]")

}
