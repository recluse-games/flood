package generic

import "github.com/recluse-games/flood/pkg/flood"

type GenericGrid struct {
	nodes [][]flood.Node
}

func (n GenericGrid) Nodes() [][]flood.Node {
	return n.nodes
}

func (n GenericGrid) SetNode(x int, y int, node flood.Node) {
	n.nodes[x][y] = node
}
