package models

// DataContext is DataStore
type DataContext interface {
	GetTasks() Tasks
	UpdateTask(task Task) bool
	DeleteTask(id int) bool

	GetUsers() []User
	UpdateUser(user User) bool
	DeleteUser(id int) bool

	GetIterators() []Iterator
	UpdateIterator(it Iterator) bool
	DeleteIterator(id int) bool

	GetProjects() []Project
	UpdateProject(it *Project) bool
	DeleteProject(id int) bool
}
