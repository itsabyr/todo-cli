package main

import (
	"github.com/itsabyr/todo-cli/scanner"
	"github.com/itsabyr/todo-cli/todo"
)

func main() {
	todoList := todo.NewList()

	scanner := scanner.NewScanner(todoList)

	scanner.Start()
}
