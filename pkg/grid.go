package fill

type Grid interface {
	nodes() [][]Node
	setNode(int, int, Node)
	setNodes()
}
