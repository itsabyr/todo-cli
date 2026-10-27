package scanner

import (
	"fmt"

	"github.com/itsabyr/todo-cli/todo"
)

func printResult(err error) {
	fmt.Println(err)
}

func printPrompt() {
	fmt.Print("Введите команду: ")
}

func printExit() {
	fmt.Println("Завершение работы... До скорого!")
}

func printAdd(title string) {
	fmt.Printf("Задача %q успешно добавлена!\n", title)
}

func printTasks(tasks []todo.Task) {
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст")
		return
	}

	fmt.Println("Список задач:")
	fmt.Println()

	for i, task := range tasks {
		status := "○"
		if task.IsDone {
			status = "✓"
		}

		fmt.Printf("%s [%d] %s\n", status, i+1, task.Title)
		fmt.Printf("   %s\n", task.Text)
		fmt.Printf("   Создана: %s\n", task.CreatedAt.Format("02.01.2006 15:04"))

		if task.IsDone && task.DoneAt != nil {
			fmt.Printf("   Выполнена: %s\n", task.DoneAt.Format("02.01.2006 15:04"))
		}

		fmt.Println()
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
		fmt.Println("Событий пока нет")
		return
	}

	fmt.Println("История событий:")
	fmt.Println()

	for i, event := range events {
		fmt.Printf("[%d] %s\n", i+1, event.CreatedAt.Format("02.01.2006 15:04:05"))
		fmt.Printf("    Команда: %s\n", event.UserInput)

		if event.Description != nil {
			fmt.Printf("    Ошибка: %s\n", event.Description)
		} else {
			fmt.Println("    Статус: успешно")
		}

		fmt.Println()
	}
}
