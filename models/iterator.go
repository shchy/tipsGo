package models

import "time"

// Iterator 作業を含む期間
type Iterator struct {
	Identity
	StartDate time.Time
	EndDate   time.Time
	Tasks     []Task
}
