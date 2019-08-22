package parentpointertree

//Node is the node of the parent pointer tree containing
//a value and pointers
type Node struct {
	parent *Node
	value  interface{}
}


//SetValue sets the value of the node
func (n *Node) SetValue(val interface{}) {
	n.value = val
}

//GetValue gets the value of the node
func (n *Node) GetValue() {
	return n.value
}

//SetParent sets the parent node for a child
func (n *Node) SetParent(parent *Node) {
	n.parent = parent
}

//GetParent gets the parent for the node
func (n *Node) GetParent() *Node {
	return n.parent
}

//PPT is a parent pointer tree
type PPT struct {
	rootChildren []*Node
}

//AddChild adds a node pointer to the slice of children
func (p *PPT) AddChild(newChildValue interface{}) {
	p.rootChildren = append(p.rootChildren, &Node{value: newChildValue})
}

//GetChildren gets the child of the index
func (p PPT) GetChildren(index int) *Node {
	defer func() {
		//recover and custom error
		if recover() != nil {
			panic("No child node exists at this index!")
		}
	}()
	return p.rootChildren[index]

}

//Walkthrough takes the index of the child chain
//returns a function that will return the next node in the chain
func (p PPT) Walkthrough(index int) func() *Node {
	currNode := p.GetChildren(index)
	return func() *Node {
		retNode := currNode
		if currNode != nil {
			currNode = currNode.parent
		}
		return retNode
	}
}

//AddToChain adds a parent to the end of a children chain
func (p *PPT) AddToChain(index int, newValue interface{}) {
	//get the index of one of the root children
	currNode := p.GetChildren(index)
	//go to the end of the chain
	for currNode.parent != nil {
		currNode = currNode.parent
	}
	currNode.parent = &Node{value: newValue}
}

//AttachChains takes two chains, the parent of
//the tailChainIndex attatches to the parent of the headChainIndex
func (p *PPT) AttachChains(headChainIndex, tailChainIndex int) {
	headNode := p.GetChildren(headChainIndex)
	if headNode == nil {
		panic("headChain is empty!")
	}
	for headNode.parent != nil {
		headNode = headNode.parent
	}
	//headNode is at the front of the headChain
	tailNode := p.GetChildren(tailChainIndex)
	if tailNode == nil {
		panic("headChain is empty!")
	}
	for tailNode.parent != nil {
		tailNode = tailNode.parent
	}
	//set the tailNode parent to the headNode
	tailNode.parent = headNode

}
