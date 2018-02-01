package models

import (
	linq "github.com/ahmetb/go-linq"
)

// InMemoryDataContext is debug use
type InMemoryDataContext struct {
	tasks     Tasks
	users     []User
	iterators []Iterator
}

// NewInMemoryDataContext ...
func NewInMemoryDataContext() DataContext {
	dataContext := InMemoryDataContext{tasks: []Task{}, users: []User{}, iterators: []Iterator{}}
	return &dataContext
}

// GetTasks Taskの一覧を取得
func (c *InMemoryDataContext) GetTasks() Tasks {
	return c.tasks
}

// UpdateTask Taskの更新または新規登録
func (c *InMemoryDataContext) UpdateTask(task Task) bool {
	if task.GetID() == -1 {
		task.setId(c.tasks.ToIds().MaxID() + 1)
		c.tasks = append(c.tasks, task)
	} else {
		c.DeleteTask(task.GetID())
		c.tasks = append(c.tasks, task)
	}
	return true
}

// DeleteTask ...
func (c *InMemoryDataContext) DeleteTask(id int) bool {
	linq.From(c.tasks).WhereT(func(x Task) bool { return x.GetID() != id }).ToSlice(&c.tasks)
	return true
}

// GetUsers ...
func (c *InMemoryDataContext) GetUsers() []User {
	return []User{}
}

// UpdateUser ...
func (c *InMemoryDataContext) UpdateUser(user User) bool {
	return true
}

// DeleteUser ...
func (c *InMemoryDataContext) DeleteUser(id int) bool {
	return true
}

// GetIterators ...
func (c *InMemoryDataContext) GetIterators() []Iterator {
	return []Iterator{}
}

// UpdateIterator ...
func (c *InMemoryDataContext) UpdateIterator(it Iterator) bool {
	return true
}

// DeleteIterator ...
func (c *InMemoryDataContext) DeleteIterator(id int) bool {
	return true
}

// GetProjects ...
func (c *InMemoryDataContext) GetProjects() []Project {
	return []Project{}
}

// UpdateProject ...
func (c *InMemoryDataContext) UpdateProject(it *Project) bool {
	return true
}

// DeleteProject ...
func (c *InMemoryDataContext) DeleteProject(id int) bool {
	return true
}
