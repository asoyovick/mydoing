package goals

import (
	"time"
	"errors"
)
type Goal struct {
	description	string
	status string
	timelines	time.Time
}
// create a new goal and check if the mandatory fields are empty.
func NewGoal(description string, status string, timelines time.Time)(error, *Goal){
	if description == "" || timelines == "" {
		return errors.New("these fields cannot be empty"),
	}
	return &Goal {
		description: string,
		status: string,
		timeline: time.Time,
	}
}

func(g *Goal)GetByDescription()string {

}
