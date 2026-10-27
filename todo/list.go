package todo

import "errors"

type List struct {
	tasks map[string]Task
}

var ErrTaskExists = errors.New("task with this title already exists")

func (l *List) AddTask(task Task) error {
	if _, exists := l.tasks[task.Title]; exists {
		return ErrTaskExists
	}

	l.tasks[task.Title] = task
	return nil
}

func (l *List) ListTasks() {

}

func (l *List) DoneTask(title string) {

}

func (l *List) DeleteTask(title string) {

}
