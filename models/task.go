package models

// Task 作業
type Task struct {
	Identity
	Value int
}

// Tasks ...
type Tasks []Task

func NewTask(name string, value int) Task {
	return Task{Identity: Identity{Name: name, ID: -1}, Value: value}
}
