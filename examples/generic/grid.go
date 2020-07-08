package generic

// Generic Grid structure for 2d games
type GenericGrid struct {
	nodes [][]Node
}

func (n GenericGrid) Nodes() [][]Node {
	return n.nodes
}

func (n GenericGrid) SetNode(x int, y int, node Node) {
	n.nodes[x][y] = node
}
