package backflow

import "github.com/Yalhash/scheduler/taskgraph"

//Backflow takes a task graph and outputs a map[string]uint of ID-> critical path
func Backflow(inGraph taskgraph.TaskGraph) map[string]uint {
	criticalPaths := make(map[string]uint)
	var queIndex int
	var queue []string
	var flag bool
	//start with end task
	currTask := inGraph[""].ID
	queue = append(queue, "")
	criticalPaths[currTask] = 0

	//transverse via BFS
	for queIndex < len(inGraph) {
		//go through all of the requirements of the current Task
		for _, currReq := range inGraph[currTask].Requirements {
			//check if the requirement is in the queue already
			flag = true
			for _, inReq := range queue {
				if inReq == currReq {
					flag = false
					break
				}
			}
			//if the requirement is not in the queue, add it
			if flag {
				queue = append(queue, currReq)
			}
			//update the critical paths of all of the requirements
			if cp, ok := criticalPaths[currReq]; ok {
				//if the value already exists in the critical Paths dict, check if this path is longer:
				if cp < inGraph[currReq].Time+criticalPaths[currTask] {
					//if so it needs to be replaced with the new critical path
					criticalPaths[currReq] = inGraph[currReq].Time + criticalPaths[currTask]
				}
			} else {
				//if the value doesn't exist, add it
				criticalPaths[currReq] = inGraph[currReq].Time + criticalPaths[currTask]
			}
		}
		//after going through all of the requirements, change the current Task to a new Task
		queIndex++
		currTask = queue[queIndex]
	}
	return criticalPaths
}
