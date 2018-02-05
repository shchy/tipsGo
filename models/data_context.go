package models

// DataContext is DataStore
type DataContext interface {
	GetTasks() Tasks
	UpdateTask(task *Task) error
	DeleteTask(id int) error

	GetUsers() []User
	UpdateUser(user *User) error
	DeleteUser(id int) error

	GetIterators() []Iterator
	UpdateIterator(it *Iterator) error
	DeleteIterator(id int) error

	GetProjects() []Project
	UpdateProject(it *Project) error
	DeleteProject(id int) error
}
