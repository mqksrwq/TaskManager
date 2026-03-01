package manager

import (
	. "TaskManager/task"
	"TaskManager/utils"
)

type Manager struct {
	tasks []*Task
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) AddTask(_name, _description string) {
	task := NewTask(_name, _description)
	m.tasks = append(m.tasks, task)
}

func (m *Manager) ShowTasks() string {
	tasks := utils.ConvertToString(m.tasks)
	return tasks
}

func (m *Manager) EditTask(_id int, _name, _description string, _status bool) {
	m.tasks[_id].SetName(_name)
	m.tasks[_id].SetDescription(_description)
	m.tasks[_id].SetStatus(_status)
}

func (m *Manager) DeleteTask(_id int) {
	m.tasks = append(m.tasks[:_id], m.tasks[_id+1:]...)
}
