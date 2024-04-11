package models

import "time"

type TmuxSession struct {
	Created           *time.Time
	LastAttached      *time.Time
	Activity          *time.Time
	Group             string
	ID                string
	Name              string
	Path              string
	Alerts            []int
	GroupList         []string
	GroupAttachedList []string
	AttachedList      []string
	Stack             []int
	Windows           int
	GroupSize         int
	GroupAttached     int
	Attached          int
	Grouped           bool
	ManyAttached      bool
	Marked            bool
	GroupManyAttached bool
	Format            bool
}
