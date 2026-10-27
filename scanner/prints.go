package scanner

import "fmt"

func printPrompt() {
	fmt.Print("Введите команду: ")
}

func printExit() {
	fmt.Print("Завершение работы... До скорого!")
}

func printAdd(title string) {
	fmt.Printf("Задача %q успешно добавлена\n", title)
}
