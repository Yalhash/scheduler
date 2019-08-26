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
type Graph map[string]*linkedlist

//AddNode adds a node to the graph
func (g Graph) AddNode(nodeID string) {
	
	if _, check := g[nodeID]; !check {
		g[nodeID] = new(linkedlist)
	}
}
//AddDirectedEdge adds an edge between startID and endID,
//such that startID points to endID
func (g Graph) AddDirectedEdge(startID, endID string) {
	g[startID].Add(endID)
}

//MakeGraphFromTasks takes a slice of tasks and
//creates a directed graph
func MakeGraphFromTasks(tasks []Task) Graph {
	var newGraph Graph = make(map[string]*linkedlist)
	//add all nodes
	for _, t := range tasks {
		//if the node has not been added yet add it
		if _, check :=  newGraph[t.ID]; !check {
			newGraph.AddNode(t.ID)
		}
		//add all of the edges associated with it
		for _, reqTask := range t.Requirements {
			//if the node has not been added yet add it
			if _, check :=  newGraph[reqTask.ID]; !check {
				newGraph.AddNode(reqTask.ID)
			}
			newGraph.AddDirectedEdge(reqTask.ID, t.ID)
		}
		
	}
	return newGraph
}

