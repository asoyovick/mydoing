package tasks

import (
	"fmt"
	"time"
)

// ServiceTask wraps your Task with a unique ID tracker required by the service layer.
type ServiceTask struct {
	ID   int
	Task *Task
}

type TaskService struct {
	tasks  []*ServiceTask
	nextID int
}

// NewTaskService creates a new TaskService with an empty task list and nextID set to 1.
func NewTaskService() *TaskService {
	return &TaskService{
		tasks:  make([]*ServiceTask, 0),
		nextID: 1,
	}
}

// CreateTask builds a new task using your NewTask constructor, assigns an ID, and tracks it.
func (s *TaskService) CreateTask(title, description, goal string, deadline time.Time) (*ServiceTask, error) {
	// Use your package's NewTask constructor (which validates empty title/goal)
	rawTask := NewTask(title, description, goal, deadline)
	if rawTask == nil {
		return nil, fmt.Errorf("failed to create task: title and goal cannot be empty")
	}

	// Wrap it with an ID
	serviceTask := &ServiceTask{
		ID:   s.nextID,
		Task: rawTask,
	}

	s.tasks = append(s.tasks, serviceTask)
	s.nextID++

	return serviceTask, nil
}

// GetTaskByID looks up a wrapped task using its unique ID.
func (s *TaskService) GetTaskByID(id int) (*ServiceTask, error) {
	for _, st := range s.tasks {
		if st.ID == id {
			return st, nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", id)
}

// ListTasks returns all tasks tracked by the sevice
func (s *TaskService) ListTaskbyID() []*ServiceTask {
	return s.tasks
}

// UpdateTask update the tasks fields.

func (s *TaskService) UpdateTask(id int, title, description, goal string, deadline time.Time) error {
	for _, st := range s.tasks {
		if st.ID == id {
			st.Task.UpdateTask(title, description, goal, deadline)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}
