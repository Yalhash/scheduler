package taskgraph

//Status of Task
type Status int

//status enum
const (
	_ = iota
	NotReady
	ReadyWorking
	Finished
	End
)

//Task is the object that needs to be scheduled
type Task struct {
	ID           string
	Time         uint
	Requirements []string
	WorkStatus   Status
}
