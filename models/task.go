package models

// Task 作業
type Task interface {
	Identity
	GetValue() int
	SetValue(value int)
	setId(id int)
}

// Tasks ...
type Tasks []Task

// ToIds ...
func (ts Tasks) ToIds() IDs {
	xs := make([]Identity, len(ts))
	for i, v := range ts {
		xs[i] = Identity(v)
	}
	return xs
}
