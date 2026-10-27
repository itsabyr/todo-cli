package todo

import "time"

type Task struct {
	Title     string
	Text      string
	CreatedAt time.Time
	IsDone    bool
	DoneAt    *time.Time
}

func NewTask(title, text string) Task {
	return Task{
		Title:     title,
		Text:      text,
		CreatedAt: time.Now(),
		IsDone:    false,
		DoneAt:    nil,
	}
}

func (t *Task) Done() {
	now := time.Now()
	t.IsDone = true
	t.DoneAt = &now
}
