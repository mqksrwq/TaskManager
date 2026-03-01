package ui

import (
	. "TaskManager/manager"
	"bufio"
	"fmt"
	"os"
)

type UI struct {
	manager *Manager
	scanner *bufio.Scanner
}

func NewUI(manager *Manager) *UI {
	return &UI{
		manager: manager,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (ui *UI) Run() {
	fmt.Println("----- МЕНЕДЖЕР ЗАДАЧ -----")
	ui.menu()
	fmt.Print("Выберите действие: ")
	ui.scanner.Scan()
	line := ui.scanner.Text()
	switch line {
	case "1":
		_ = line
	case "2":
		_ = line
	case "3":
		_ = line
	case "4":
		_ = line
	case "5":
		_ = line
	case "6":
		return
	}
}

func (ui *UI) menu() {
	fmt.Print("\n1. Добавить задачу\n" +
		"2. Показать все задачи\n" +
		"3. Изменить статус задачи\n" +
		"4. Редактировать задачу\n" +
		"5. Удалить задачу\n" +
		"6. Выход\n")
}
