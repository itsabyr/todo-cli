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

	if cmd == "add" {
		return s.cmdAdd(fields)
	}

	if cmd == "list" {
		return s.cmdList(fields)
	}

	if cmd == "done" {
		return s.cmdDone(fields)
	}

	return nil
}

func (s *Scanner) cmdAdd(fields []string) error {
	if len(fields) < 3 {
		return errWrongArgs
	}

	title := fields[1]
	text := strings.Join(fields[2:], " ")

	task := todo.NewTask(title, text)

	s.todoList.AddTask(task)

	printAdd(title)

	return nil
}

func (s *Scanner) cmdList(fields []string) error {
	if len(fields) != 1 {
		return errWrongArgs
	}

	tasks := s.todoList.ListTasks()

	printTasks(tasks)

	return nil
}

func (s *Scanner) cmdDone(fields []string) error {
	if len(fields) != 2 {
		return errWrongArgs
	}

	title := fields[1]

	if err := s.todoList.DoneTask(title); err != nil {
		return err
	}

	printDone(title)

	return nil
}
