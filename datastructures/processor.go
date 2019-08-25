package datastructures


var processorCount int


//Processor is the object that can work on Task objects
type Processor struct {
	ID int
	Name string
}


//NewGenericProcessor creates a generic processor without a name
func NewGenericProcessor() Processor {
	newProc := Processor {ID: processorCount}
	processorCount++
	return newProc
}

//NewNamedProcessor creates a named processor
func NewNamedProcessor(processName string) Processor {
	newProc := Processor {
		ID: processorCount,
		Name: processName,
	}
	processorCount++
	return newProc
}