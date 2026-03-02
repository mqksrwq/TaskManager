package manager

import (
	. "TaskManager/task"
	"TaskManager/utils"
	"fmt"
	"strconv"
	"strings"
)

// Manager - структура менеджера
type Manager struct {

	// tasks - динамический массив задач
	tasks []*Task
}

// NewManager - метод создания нового менеджера
func NewManager() *Manager {
	return &Manager{}
}

// AddTask - метод добавления новой задачи
func (m *Manager) AddTask(_name, _description string) error {
	if strings.TrimSpace(_name) == "" {
		return fmt.Errorf("ошибка: название задачи не" +
			"может быть пустым")
	}
	if strings.TrimSpace(_description) == "" {
		return fmt.Errorf("ошибка: описание задачи не" +
			"может быть пустым")
	}
	task := NewTask(strings.TrimSpace(_name), strings.TrimSpace(_description))
	m.tasks = append(m.tasks, task)
	return nil
}

// ShowTasks - метод для вывода списка всех задач
func (m *Manager) ShowTasks() (string, error) {
	if len(m.tasks) == 0 {
		return "", fmt.Errorf("ошибка: список задач пуст")
	}
	tasks := utils.ConvertToString(m.tasks)
	return tasks, nil
}

// ChangeStatus - метод для изменения статуса выполнения задачи
func (m *Manager) ChangeStatus(_id string) error {
	id, err := strconv.Atoi(_id)
	if err != nil {
		return fmt.Errorf("ошибка: ID должен " +
			"представлять число")
	}
	if id < 1 || id > len(m.tasks) {
		return fmt.Errorf("ошибка: задача с "+
			"ID %d не найдена", id)
	}
	m.tasks[id-1].SetStatus()
	return nil
}

// EditTask - метод для редактирования задачи
func (m *Manager) EditTask(_id, _name, _description string) error {
	id, err := strconv.Atoi(_id)
	if err != nil {
		return fmt.Errorf("ошибка: ID должен " +
			"представлять число")
	}
	if id < 1 || id > len(m.tasks) {
		return fmt.Errorf("ошибка: задача с "+
			"ID %d не найдена", id)
	}

	if strings.TrimSpace(_name) == "" {
		return fmt.Errorf("ошибка: название задачи не" +
			"может быть пустым")
	}
	if strings.TrimSpace(_description) == "" {
		return fmt.Errorf("ошибка: описание задачи не" +
			"может быть пустым")
	}

	m.tasks[id-1].SetName(_name)
	m.tasks[id-1].SetDescription(_description)
	return nil
}

// DeleteTask - метод для удаления задачи
func (m *Manager) DeleteTask(_id string) error {
	id, err := strconv.Atoi(_id)
	if err != nil {
		return fmt.Errorf("ошибка: ID должен " +
			"представлять число")
	}
	if id < 1 || id > len(m.tasks) || m.tasks[id-1] == nil {
		return fmt.Errorf("ошибка: задача с "+
			"ID %d не найдена", id)
	}

	m.tasks = append(m.tasks[:id-1], m.tasks[id:]...)
	return nil
}
