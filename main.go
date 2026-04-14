package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const fileName = "tasks.json"

// Load tasks from file
func loadTasks() []Task {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return []Task{}
	}

	var tasks []Task
	json.Unmarshal(data, &tasks)
	return tasks
}

// Save tasks to file
func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(fileName, data, 0644)
}

// Add task
func addTask(desc string) {
	tasks := loadTasks()

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	task := Task{
		ID:          id,
		Description: desc,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, task)
	saveTasks(tasks)

	fmt.Println("Task added successfully (ID:", id, ")")
}

// Update task
func updateTask(id int, desc string) {
	tasks := loadTasks()

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = desc
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Println("Task updated")
			return
		}
	}

	fmt.Println("Task not found")
}

// Delete task
func deleteTask(id int) {
	tasks := loadTasks()

	newTasks := []Task{}
	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		}
	}

	saveTasks(newTasks)
	fmt.Println("Task deleted")
}

// Update status
func updateStatus(id int, status string) {
	tasks := loadTasks()

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Println("Status updated")
			return
		}
	}

	fmt.Println("Task not found")
}

// List tasks
func listTasks(filter string) {
	tasks := loadTasks()

	for _, t := range tasks {
		if filter == "" || t.Status == filter {
			fmt.Printf("[%d] %s (%s)\n", t.ID, t.Description, t.Status)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command>")
		return
	}

	switch os.Args[1] {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide task description")
			return
		}
		addTask(os.Args[2])

	case "update":
		id, _ := strconv.Atoi(os.Args[2])
		updateTask(id, os.Args[3])

	case "delete":
		id, _ := strconv.Atoi(os.Args[2])
		deleteTask(id)

	case "mark-in-progress":
		id, _ := strconv.Atoi(os.Args[2])
		updateStatus(id, "in-progress")

	case "mark-done":
		id, _ := strconv.Atoi(os.Args[2])
		updateStatus(id, "done")

	case "list":
		if len(os.Args) > 2 {
			listTasks(os.Args[2])
		} else {
			listTasks("")
		}

	default:
		fmt.Println("Invalid command")
	}
}
