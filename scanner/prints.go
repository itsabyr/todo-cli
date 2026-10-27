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

func printHelp() {
	fmt.Println("Доступные команды:")
	fmt.Println()

	fmt.Println("  help                      Показать доступные команды и их формат")
	fmt.Println("  add {заголовок} {текст}   Добавить задачу")
	fmt.Println("  list                      Вывести список всех задач")
	fmt.Println("  done {заголовок}          Отметить задачу как выполненную")
	fmt.Println("  del {заголовок}           Удалить задачу по заголовку")
	fmt.Println("  events                    Вывести список всех событий")
	fmt.Println("  exit                      Завершить работу программы")
}

func printEvents(events []Event) {
	if len(events) == 0 {
		pp.Println("Событий пока нет")
		return
	}

	pp.Println("История событий:")
	pp.Println()

	for i, event := range events {
		pp.Printf("[%d] %s\n", i+1, event.CreatedAt.Format("02.01.2006 15:04:05"))
		pp.Printf("    Команда: %s\n", event.UserInput)

		if event.Description != nil {
			pp.Printf("    Ошибка: %s\n", event.Description)
		} else {
			pp.Println("    Статус: успешно")
		}

		pp.Println()
	}
}
