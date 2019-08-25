package datastructures



//Status of Task object
type Status int

//status enum
const (
	_ = iota
	UnFinished
	Working
	Finished
)

//Task is the object that needs to be scheduled
type Task struct {
	ID string
	Time uint
	Requirements []*Task
	WorkStatus Status
}

