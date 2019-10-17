package backflow

import (
	"github.com/Yalhash/schedule/taskgraph"
	"testing"
)

func TestBackflow(t *testing.T) {
	var inputGraph taskgraph.TaskGraph = map[string]taskgraph.Task{
		"": taskgraph.Task{
			"",
			0,
			[]string{"A", "B"},
			End,
		},
		"A": taskgraph.Task{
			"A",
			4,
			[]string{},
			Ready,
		},
		"B": taskgraph.Task{
			"B",
			3,
			[]string{"C"},
			Ready,
		},
		"C": taskgraph.Task{
			"C",
			2,
			[]string{"C"},
			Ready,
		},
	}

	testAnswer := Backflow(inputGraph)
	if testAnswer != {"" : 0, "A" : 4, "B" : 3, "C" : 5} {
		t.Errorf("Backflow returned wrong answer, expected: %v\n got %v",
		{"" : 0, "A" : 4, "B" : 3, "C" : 5},
		testAnswer) 

	}
}
