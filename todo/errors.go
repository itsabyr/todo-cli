package todo

import "errors"

var (
	errTaskExists   = errors.New("task with this title already exists")
	errTaskNotFound = errors.New("task with this title not found")
)
