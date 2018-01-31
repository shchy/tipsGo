package tipsGo

import (
	"time"
)

// Identity has id and name
type Identity interface {
	GetID() int
	GetName() string
}

// Task 作業
type Task interface {
	Identity
	GetValue() int64
}

// Iterator 作業を含む期間
type Iterator struct {
	Identity
	StartDate time.Time
	EndDate   time.Time
	Tasks     []Task
}

// User user
type User interface {
	Identity
}

// Project is []Iterator
type Project struct {
	Identity
	Iterators []Iterator
}

// TaskImpl 最小単位の作業
type TaskImpl struct {
	name  string
	value int
}
