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
func (c *InMemoryDataContext) UpdateTask(task Task) error {
	if task.ID == -1 {
		task.ID = linq.From(c.tasks).SelectT(func(t Task) int { return t.ID }).Max().(int) + 1
		c.tasks = append(c.tasks, task)
	} else {
		c.DeleteTask(task.ID)
		c.tasks = append(c.tasks, task)
	}
	return nil
}

// DeleteTask ...
func (c *InMemoryDataContext) DeleteTask(id int) error {
	linq.From(c.tasks).WhereT(func(x Task) bool { return x.ID != id }).ToSlice(&c.tasks)
	return nil
}

// GetUsers ...
func (c *InMemoryDataContext) GetUsers() []User {
	return c.users
}

// UpdateUser ...
func (c *InMemoryDataContext) UpdateUser(user User) error {
	if user.ID == -1 {
		user.ID = linq.From(c.users).SelectT(func(u User) int { return u.ID }).Max().(int) + 1
		c.users = append(c.users, user)
	} else {
		c.DeleteUser(user.ID)
		c.users = append(c.users, user)
	}
	return nil
}

// DeleteUser ...
func (c *InMemoryDataContext) DeleteUser(id int) error {
	return nil
}

// GetIterators ...
func (c *InMemoryDataContext) GetIterators() []Iterator {
	return []Iterator{}
}

// UpdateIterator ...
func (c *InMemoryDataContext) UpdateIterator(it Iterator) error {
	return nil
}

// DeleteIterator ...
func (c *InMemoryDataContext) DeleteIterator(id int) error {
	return nil
}

// GetProjects ...
func (c *InMemoryDataContext) GetProjects() []Project {
	return []Project{}
}

// UpdateProject ...
func (c *InMemoryDataContext) UpdateProject(it *Project) error {
	return nil
}

// DeleteProject ...
func (c *InMemoryDataContext) DeleteProject(id int) error {
	return nil
}
