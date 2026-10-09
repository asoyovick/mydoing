package main

import (
	"time"
	"mydoing/internal/tasks"
)

func main(){
	// Check if the command line argument is provided
	deadline := time.Now().add(24*time.Hour)

	task, err := tasks.NewTask(
		"Complete Project",
		"Finish the Backen API endpoint",
		"Ship it!",
		deadline,
	)
	if err != nil {
		fmt.Printf("failed to create a new task: %v\n", err)
		return
	}
	fmt.Printf("Successfuly created a new task: %s (Goal: %s) due by %s\n"
		task.Title, task.Goal, Task.Deadline
	)

}
