package scanner

import (
	"bufio"
	"os"

	"github.com/itsabyr/todo-cli/todo"
)

type Scanner struct {
	todoList *todo.List
}

func (s *Scanner) Start() error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printPrompt()

		if !scanner.Scan() {
			return scanner.Err()
		}
	}
}
