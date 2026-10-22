package goals

import (
	"errors"
	"time"
)

type Goal struct {
	description string
	status      string
	timelines   time.Time
}

// create a new goal and check if the mandatory fields are empty.
func NewGoal(description string, status string, timelines time.Time) (*Goal, error) {
	if description == "" || timelines.IsZero() {
		return nil, errors.New("these fields cannot be empty")
	}
	return &Goal{
		description: description,
		status:      status,
		timelines:   timelines,
	}, nil
}

// Get goals by description.
func (g *Goal) GetByDescription() string {
	return g.description
}
func (g *Goal) GetGoalByStatus() string {
	return g.status
}
