package task

var nextId int

type Task struct {
	id          int
	name        string
	description string
	status      bool
}

func NewTask(_name string, _description string) *Task {
	nextId++
	return &Task{
		id:          nextId,
		name:        _name,
		description: _description,
		status:      false,
	}
}

func (t *Task) GetId() int {
	return t.id
}

func (t *Task) GetName() string {
	return t.name
}

func (t *Task) SetName(_name string) {
	t.name = _name
}

func (t *Task) GetDescription() string {
	return t.description
}

func (t *Task) SetDescription(_description string) {
	t.description = _description
}

func (t *Task) GetStatus() bool {
	return t.status
}

func (t *Task) SetStatus(_status bool) {
	t.status = _status
}
