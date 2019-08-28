package datastructures

//linked list holds bool values
type linkedlist struct {
	head *Node
	tail *Node
	Length int
}

func (l *linkedlist) Add(val interface{}) {
	if l.head == nil {
		l.head = &Node{Value: val}
		l.tail = l.head
	}else{
		l.tail.Next = &Node{Value: val}
		l.tail = l.tail.Next
	}
	l.Length++
}

func (l *linkedlist) GetNode(index int) interface{} {
	if index >= l.Length || index  < 0 {
		panic("index out of range")
	}
	currNode := l.head
	for i := 0; i < index; i++{
		currNode = currNode.Next
	}
	return currNode.Value
}


//Graph holds an adjacancy list containing all nodes and their connections
//cannot repeat an ID value
type Graph map[*Task]*linkedlist

//AddNode adds a node to the graph
func (g Graph) AddNode(newTask *Task) {
	
	if _, check := g[newTask]; !check {
		g[newTask] = new(linkedlist)
	}
}
//AddDirectedEdge adds an edge between startTask and endTask
//such that startTask points to endTask
func (g Graph) AddDirectedEdge(startTask, endTask *Task) {
	g[startTask].Add(endTask)
}

//MakeGraphFromTasks takes a slice of tasks and
//creates a directed graph
func MakeGraphFromTasks(tasks []*Task) Graph {
	var newGraph Graph = make(map[*Task]*linkedlist)
	//add all nodes
	for _, t := range tasks {
		//if the node has not been added yet add it
		if _, check :=  newGraph[t]; !check {
			newGraph.AddNode(t)
		}
		//add all of the edges associated with it
		for _, reqTask := range t.Requirements {
			//if the node has not been added yet add it
			if _, check :=  newGraph[reqTask]; !check {
				newGraph.AddNode(reqTask)
			}
			newGraph.AddDirectedEdge(reqTask, t)
		}
	}
	return newGraph
}

func isEndTask(checkTask *Task, taskList []*Task) bool{
	for _, t := range taskList {
		if t != checkTask {
			for _, req := t.Requirements {
				if req == checkTask {
					return false
				}
			}
		}
	}
	return true
}

