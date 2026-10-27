package todo

import (
	"sort"
)

type List struct {
	tasks map[string]Task
}

func (l *List) AddTask(task Task) error {
	if _, exists := l.tasks[task.Title]; exists {
		return errTaskExists
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

func (l *List) DoneTask(title string) error {
	task, ok := l.tasks[title]
	if !ok {
		return errTaskNotFound
	}

	task.Done()

	l.tasks[title] = task

	return nil
}

func (l *List) DeleteTask(title string) error {
	if _, ok := l.tasks[title]; !ok {
		return errTaskNotFound
	}

	delete(l.tasks, title)

	return nil
}
