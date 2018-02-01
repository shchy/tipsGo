package models

// Identity has id and name
type Identity interface {
	GetID() int
	GetName() string
}

// IDs ...
type IDs []Identity

// MaxID ...
func (xs IDs) MaxID() int {
	max := -1
	for _, i := range xs {
		id := i.GetID()
		if max < id {
			max = id
		}
	}
	return max
}

// FindByID ...
func (xs IDs) FindByID(id int) Identity {
	for _, i := range xs {
		if i.GetID() == id {
			return i
		}
	}
	return nil
}
