package goals

import (
	"fmt"
	"time"
)

// create a struc that wraps the goals
type ServiceGoal struct {
	ID   int
	goal *Goal
}

// create a struct that holds the database and tracks the global id counter.
type GoalService struct {
	goal   []*ServiceGoal
	NextID int
	// Get goals with ID tracks the that are sav
	// Get goals with ID tracks the that are sav
}

// create a new goal-service
func NewGoalService() *GoalService {
	return &GoalService{
		goal:   make([]*ServiceGoal, 0),
		NextID: 1,
	}
}

// create a new goal service
func (s *GoalService) CreateServiceGoal(description string, status string, timeline time.Time) (*ServiceGoal, error) {
	// construct new goal
	rawGoal := NewGoal(description, status, timeline)
	if rawGoal == nil {
		return nil, fmt.Errorf("failed to create a new goal, all the fields cannot be empty")
	}
	// wrap created the newly created goal with a unique ID.
	serviceGoal := &ServiceGoal{
		ID:   s.NextID,
		goal: rawGoal,
	}
	// save the newly created goal.
	s.goal = append(s.goal, serviceGoal)
	s.NextID++
	return serviceGoal, nil
}

// Get goals with ID tracks the that are saved in the serviceGoal.
func (s *GoalService) GetGoalByID(id int) (*ServiceGoal, error) {
	for _, sg := range s.goal {
		if sg.ID == id {
			return sg, nil
		}
	}
	return nil, fmt.Errorf("goal ID %d is not found")
}

// list all the goals by ID
func (s *GoalService) ListGoalByID() []*ServiceGoal {
	return s.goal
}
