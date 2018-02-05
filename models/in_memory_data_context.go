package models

import (
	linq "github.com/ahmetb/go-linq"
)

// InMemoryDataContext is debug use
type InMemoryDataContext struct {
	tasks     Tasks
	users     []User
	iterators []Iterator
	projects  []Project
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
func (c *InMemoryDataContext) UpdateTask(task *Task) error {
	if task.ID == -1 {
		if len(c.tasks) == 0 {
			task.ID = 0
		} else {
			task.ID = linq.From(c.tasks).SelectT(func(t Task) int { return t.ID }).Max().(int) + 1
		}

		c.tasks = append(c.tasks, *task)
	} else {
		c.DeleteTask(task.ID)
		c.tasks = append(c.tasks, *task)
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
func (c *InMemoryDataContext) UpdateUser(user *User) error {
	if user.ID == -1 {
		if len(c.users) == 0 {
			user.ID = 0
		} else {
			user.ID = linq.From(c.users).SelectT(func(u User) int { return u.ID }).Max().(int) + 1
		}

		c.users = append(c.users, *user)
	} else {
		c.DeleteUser(user.ID)
		c.users = append(c.users, *user)
	}
	return nil
}

// DeleteUser ...
func (c *InMemoryDataContext) DeleteUser(id int) error {
	linq.From(c.users).WhereT(func(x User) bool { return x.ID != id }).ToSlice(&c.users)
	return nil
}

// GetIterators ...
func (c *InMemoryDataContext) GetIterators() []Iterator {
	return c.iterators
}

// UpdateIterator ...
func (c *InMemoryDataContext) UpdateIterator(it *Iterator) error {
	if it.ID == -1 {
		if len(c.iterators) == 0 {
			it.ID = 0
		} else {
			it.ID = linq.From(c.iterators).SelectT(func(i Iterator) int { return i.ID }).Max().(int) + 1
		}

		c.iterators = append(c.iterators, *it)
	} else {
		c.DeleteIterator(it.ID)
		c.iterators = append(c.iterators, *it)
	}
	return nil
}

// DeleteIterator ...
func (c *InMemoryDataContext) DeleteIterator(id int) error {
	linq.From(c.iterators).WhereT(func(x Iterator) bool { return x.ID != id }).ToSlice(&c.iterators)
	return nil
}

// GetProjects ...
func (c *InMemoryDataContext) GetProjects() []Project {
	return c.projects
}

// UpdateProject ...
func (c *InMemoryDataContext) UpdateProject(p *Project) error {
	if p.ID == -1 {
		if len(c.projects) == 0 {
			p.ID = 0
		} else {
			p.ID = linq.From(c.projects).SelectT(func(x Project) int { return x.ID }).Max().(int) + 1
		}

		c.projects = append(c.projects, *p)
	} else {
		c.DeleteProject(p.ID)
		c.projects = append(c.projects, *p)
	}
	return nil
}

// DeleteProject ...
func (c *InMemoryDataContext) DeleteProject(id int) error {
	linq.From(c.projects).WhereT(func(x Project) bool { return x.ID != id }).ToSlice(&c.projects)
	return nil
}
