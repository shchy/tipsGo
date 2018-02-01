package models

// TaskImpl 最小単位の作業
type TaskImpl struct {
	id    int
	name  string
	value int
}

// NewTask ...
func NewTask(name string, value int) Task {
	return &TaskImpl{id: -1, name: name, value: value}
}

// GetID ...
func (t *TaskImpl) GetID() int {
	return t.id
}

// GetName ...
func (t *TaskImpl) GetName() string {
	return t.name
}

// GetValue ...
func (t *TaskImpl) GetValue() int {
	return t.value
}

// SetValue ...
func (t *TaskImpl) SetValue(value int) {
	t.value = value
}

func (t *TaskImpl) setId(id int) {
	t.id = id
}
