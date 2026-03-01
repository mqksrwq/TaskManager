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
	for {
		ui.menu()
		fmt.Print("Выберите действие: ")
		ui.scanner.Scan()
		line := ui.scanner.Text()
		switch line {
		case "1":
			ui.add()
		case "2":
			ui.show()
		case "3":
			ui.status()
		case "4":
			ui.edit()
		case "5":
			ui.delete()
		case "6":
			return
		}
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

func (ui *UI) add() {
	fmt.Print("\n----- ДОБАВЛЕНИЕ ЗАДАЧИ -----\n")
	fmt.Print("Введите название задачи:")
	ui.scanner.Scan()
	name := ui.scanner.Text()
	fmt.Print("Введите описание задачи:")
	ui.scanner.Scan()
	description := ui.scanner.Text()
	ui.manager.AddTask(name, description)
	fmt.Println("Задача добавлена")
}

func (ui *UI) show() {
	fmt.Println(ui.manager.ShowTasks())
}

func (ui *UI) status() {
	fmt.Print("Введите ID задачи для изменения статуса: ")
	ui.scanner.Scan()
	id := ui.scanner.Text()
	ui.manager.ChangeStatus(id)
	fmt.Printf("Статус задачи с ID %s изменен\n", id)
}

func (ui *UI) edit() {
	fmt.Print("\n----- ИЗМЕНЕНИЕ ЗАДАЧИ -----\n")
	fmt.Print("Введите ID задачи для изменения:")
	ui.scanner.Scan()
	id := ui.scanner.Text()
	fmt.Print("Введите название задачи:")
	ui.scanner.Scan()
	name := ui.scanner.Text()
	fmt.Print("Введите описание задачи:")
	ui.scanner.Scan()
	description := ui.scanner.Text()
	ui.manager.EditTask(id, name, description)
	fmt.Printf("Задача с ID %s изменена\n", id)
}

func (ui *UI) delete() {
	fmt.Print("Введите ID задачи для удаления:")
	ui.scanner.Scan()
	id := ui.scanner.Text()
	ui.manager.DeleteTask(id)
	fmt.Printf("Задача с ID %s удалена\n", id)
}
