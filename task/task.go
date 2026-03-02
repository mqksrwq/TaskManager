package task

// nextId - ID задачи
var nextId int

// Task - структура задачи
type Task struct {
	id          int
	name        string
	description string
	status      bool
}

// NewTask - метод для создания новой задачи
func NewTask(_name string, _description string) *Task {
	nextId++
	return &Task{
		id:          nextId,
		name:        _name,
		description: _description,
		status:      false,
	}
}

// GetId - метод для получения ID задачи
func (t *Task) GetId() int {
	return t.id
}

// GetName - метод для получения названия задачи
func (t *Task) GetName() string {
	return t.name
}

// SetName - метод для установки названия задачи
func (t *Task) SetName(_name string) {
	t.name = _name
}

// GetDescription - метод для получения описания задачи
func (t *Task) GetDescription() string {
	return t.description
}

// SetDescription - метод для установки описания задачи
func (t *Task) SetDescription(_description string) {
	t.description = _description
}

// GetStatus - метод для получения статуса задачи
func (t *Task) GetStatus() bool {
	return t.status
}

// SetStatus - метод для установки статуса задачи
func (t *Task) SetStatus() {
	t.status = !t.status
}
