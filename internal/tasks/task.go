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

// create a new task and chack if any of the fields are empty.
func NewTask(title, description, goal string, deadline time.Time) *Task {
	if (title == "") || (goal == "") {
		return nil
	}
	return &Task{
		title:       title,
		description: description,
		goal:        goal,
		deadline:    deadline,
	}
}

// Use methods to get the tasks fields, use this when you want to display the task's fields which are to be visible.
// GetTitle returns the title of the task.
func (t *Task) GetTitle() string {
	return t.title
}

// GetDescription returns the description of the task.
func (t *Task) GetDescription() string {
	return t.description
}

// GetGoal returns the goal of the task.
func (t *Task) GetGoal() string {
	return t.goal
}

// GetDeadline returns the deadline of the task.
func (t *Task) GetDeadline() time.Time {
	return t.deadline
}

// UpdateTask updates the task's fields.
func (task *Task) UpdateTask(title, description, goal string, deadline time.Time) {
	if title != "" {
		task.title = title
	}
	if description != "" {
		task.description = description
	}
	if goal != "" {
		task.goal = goal
	}
	if !deadline.IsZero() {
		task.deadline = deadline
	}
}
