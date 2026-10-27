package scanner

import (
	"bufio"
	"errors"
	"os"
	"strings"

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

		line := scanner.Text()
		if err := s.process(line); err != nil {
			if errors.Is(err, errNeedExit) {
				printExit()
				return nil
			}
		}
	}
}

func (s *Scanner) process(line string) error {
	fields := strings.Fields(line)

	if len(fields) == 0 {
		return errEmptyInput
	}

	cmd := fields[0]

	if cmd == "exit" {
		return errNeedExit
	}

	return nil
}
