package scanner

import (
	"fmt"

	"github.com/itsabyr/todo-cli/todo"
	"github.com/k0kubun/pp"
)

func printPrompt() {
	fmt.Print("Введите команду: ")
}

func printExit() {
	fmt.Print("Завершение работы... До скорого!")
}

func printAdd(title string) {
	fmt.Printf("Задача %q успешно добавлена!\n", title)
}

func printTasks(tasks []todo.Task) {
	if len(tasks) == 0 {
		pp.Println("Список задач пуст")
		return
	}

	pp.Println("Список задач:")
	pp.Println()

	for i, task := range tasks {
		status := "○"
		if task.IsDone {
			status = "✓"
		}

		pp.Printf("%s [%d] %s\n", status, i+1, task.Title)
		pp.Printf("   %s\n", task.Text)
		pp.Printf("   Создана: %s\n", task.CreatedAt.Format("02.01.2006 15:04"))

		if task.IsDone && task.DoneAt != nil {
			pp.Printf("   Выполнена: %s\n", task.DoneAt.Format("02.01.2006 15:04"))
		}

		pp.Println()
	}
}

func printDone(title string) {
	fmt.Printf("Задача %q помечана как выполненная!\n", title)
}

func printDel(title string) {
	fmt.Printf("Задача %q удалена!\n", title)
}
