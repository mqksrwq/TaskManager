package ui

import (
	. "TaskManager/manager"
	"bufio"
	"fmt"
	"os"
)

// UI - структура ui
type UI struct {

	// manager - менеджер
	manager *Manager

	// scanner - сканер
	scanner *bufio.Scanner
}

// NewUI - метод для создания нового ui
func NewUI(manager *Manager) *UI {
	return &UI{
		manager: manager,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

// Run - метод для запуска ui
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
		default:
			return
		}
	}
}

// menu - метод для показа интерфейса меню программы
func (ui *UI) menu() {
	fmt.Print("\n1. Добавить задачу\n" +
		"2. Показать все задачи\n" +
		"3. Изменить статус задачи\n" +
		"4. Редактировать задачу\n" +
		"5. Удалить задачу\n" +
		"6 / Enter. Выход\n")
}

// add - метод для показа интерфейса добавления задачи
func (ui *UI) add() {
	fmt.Print("\n----- ДОБАВЛЕНИЕ ЗАДАЧИ -----\n")
	fmt.Print("Введите название задачи: ")
	ui.scanner.Scan()
	name := ui.scanner.Text()
	fmt.Print("Введите описание задачи: ")
	ui.scanner.Scan()
	description := ui.scanner.Text()
	err := ui.manager.AddTask(name, description)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Задача добавлена")
	}
}

// show - метод для показа интерфейса вывода задач
func (ui *UI) show() {
	line, err := ui.manager.ShowTasks()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(line)
	}
}

// status - метод для показа интерфейса изменения статуса
func (ui *UI) status() {
	fmt.Print("Введите ID задачи для изменения статуса: ")
	ui.scanner.Scan()
	id := ui.scanner.Text()
	err := ui.manager.ChangeStatus(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Статус задачи с ID %s изменен\n", id)
	}
}

// edit - метод для показа интерфейса редактирования
func (ui *UI) edit() {
	fmt.Print("\n----- ИЗМЕНЕНИЕ ЗАДАЧИ -----\n")
	fmt.Print("Введите ID задачи для изменения: ")
	ui.scanner.Scan()
	id := ui.scanner.Text()
	fmt.Print("Введите название задачи: ")
	ui.scanner.Scan()
	name := ui.scanner.Text()
	fmt.Print("Введите описание задачи: ")
	ui.scanner.Scan()
	description := ui.scanner.Text()
	err := ui.manager.EditTask(id, name, description)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Задача с ID %s изменена\n", id)
	}
}

// delete - метод для показа интерфейса удаления
func (ui *UI) delete() {
	fmt.Print("Введите ID задачи для удаления: ")
	ui.scanner.Scan()
	id := ui.scanner.Text()
	err := ui.manager.DeleteTask(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Задача с ID %s удалена\n", id)
	}
}
