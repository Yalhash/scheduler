package datastructures


//PPT is a parent pointer tree
type PPT struct {
	RootChildren []*Node
}

//AddChild adds a node pointer to the slice of children
func (p *PPT) AddChild(newChildValue interface{}) {
	p.RootChildren = append(p.RootChildren, &Node{Value: newChildValue})
}

//GetChild gets the child of the index
func (p PPT) GetChild(index int) *Node {
	defer func() {
		//recover and custom error
		if recover() != nil {
			panic("No child node exists at this index!")
		}
	}()
	return p.RootChildren[index]

}
