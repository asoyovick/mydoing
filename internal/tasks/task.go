package tasks

import (
	"time"
)

// User accepted prediction:
type Task struct {
	title       string
	description string
	goal        string
	deadline    time.Time
}

// create a new task and chack if the title is empty.
func NewTask(title, description, goal string, deadline time.Time) *Task {
	if title == "" {
		return nil
	}
	return &Task{
		title:       title,
		description: description,
		goal:        goal,
		deadline:    deadline,
	}
}
