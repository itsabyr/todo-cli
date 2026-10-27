package scanner

import "time"

type Event struct {
	Description error
	UserInput   string
	CreatedAt   time.Time
}

func NewEvent(description error, userInput string) Event {
	return Event{
		Description: description,
		UserInput:   userInput,
		CreatedAt:   time.Now(),
	}
}
