package todo

import (
	"errors"
	"sort"
)

type List struct {
	tasks map[string]Task
}

var ErrTaskExists = errors.New("task with this title already exists")
var ErrTaskNotFound error

func (l *List) AddTask(task Task) error {
	if _, exists := l.tasks[task.Title]; exists {
		return ErrTaskExists
	}

	l.tasks[task.Title] = task
	return nil
}

func (l *List) ListTasks() []Task {
	tasks := make([]Task, 0, len(l.tasks))

	for _, t := range l.tasks {
		tasks = append(tasks, t)
	}

	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i].CreatedAt.Equal(tasks[j].CreatedAt) {
			return tasks[i].Title < tasks[j].Title
		}
		return tasks[i].CreatedAt.Before(tasks[j].CreatedAt)
	})

	return tasks
}

func (l *List) DoneTask(title string) {

}

func (l *List) DeleteTask(title string) {

}
