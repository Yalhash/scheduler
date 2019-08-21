package main

import (
	"fmt"
	"scheduler"
)

func main() {
	thing := new(parentpointertree.PPT)

	thing.AddChild(1)
	thing.AddToChain(0, "sseee")
	thing.AddToChain(0, 00)
	thing.AddToChain(0, 0x11)

	thing.AddChild(16)
	thing.AddToChain(1, 0.123)
	thing.AddToChain(1, 1023)
	thing.AddToChain(1, "YEEETSSSTS")
	thing.AttachChains(0, 1)
	thing.AddToChain(0, 15)
	thing.AddToChain(1, "hehe")

	walker := thing.Walkthrough(0)
	fmt.Println("First chain:")
	for currNode := walker(); currNode != nil; currNode = walker() {
		fmt.Println(currNode.Value, ",")
	}

	walker = thing.Walkthrough(1)
	fmt.Println("Second chain:")
	for currNode := walker(); currNode != nil; currNode = walker() {
		fmt.Println(currNode.Value, ",")
	}
}
