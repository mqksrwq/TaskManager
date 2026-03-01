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

func (m *Manager) ChangeStatus(_id string) {
	m.tasks[int([]rune(_id)[0])].SetStatus(!m.tasks[int([]rune(_id)[0])].GetStatus())
}

func (m *Manager) EditTask(_id, _name, _description string) {
	m.tasks[int([]rune(_id)[0])].SetName(_name)
	m.tasks[int([]rune(_id)[0])].SetDescription(_description)
}

func (m *Manager) DeleteTask(_id string) {
	m.tasks = append(m.tasks[:int([]rune(_id)[0])], m.tasks[int([]rune(_id)[0]):]...)
}
