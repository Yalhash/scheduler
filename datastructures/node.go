package datastructures

//Node is the node of the parent pointer tree containing
//a value and pointers
type Node struct {
	Parents []*Node
	Value  interface{}
}


//AddParent adds appends a parent to the slice this node's parents
func (n *Node) AddParent(parent *Node) {
	n.Parents = append(n.Parents, parent)
}

//SetParents sets the parent node for a child
func (n *Node) SetParents(parents []*Node) {
	n.Parents = parents
}

//GetParents gets the parent for the node
func (n *Node) GetParents() []*Node {
	return n.Parents
}
