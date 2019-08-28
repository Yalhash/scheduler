package algorithms

import "github.com/Yalhash/schedule/datastructures"

func isEndNode(qNode *datastructures.Task, taskGraph datastructures.Graph) bool {
	//a node is an end node if it is not required for any other node
	for k, _ := range taskGraph {
		if qNode != k {
			for _, req := range k.Requirements {
				if qNode == req {
					return false
				}
			}
		}
	}
	return true
}


func backRecur(currentTask *datastructures.Task, criticalPaths *map[*datastructures.Task]uint){
	for _, req := range currentTask.Requirements {
		if criticalPaths[req] < criticalPaths[currentTask] + req.Time {
			criticalPaths[req] = criticalPaths[currentTask] + req.Time
			backRecur(currentTask, criticalPaths)
		}
	}
}

//Backflow takes  a graph of tasks and applies
//the Backflow Algorithm on to create a graph
//with each Node's critical path weight to it
func Backflow(taskGraph datastructures.Graph) (criticalPaths map[*datastructures.Task] uint){
	criticalPaths = make(map[*datastructures.Task] uint)
	//find all end nodes
	for k, _ := range taskGraph {
		if isEndNode(k, taskGraph) {

			//set it's own critical path equal to its time
			criticalPaths[k] = k.Time
			backRecur(k, &criticalPaths)
		}
	}
	return criticalPaths
}