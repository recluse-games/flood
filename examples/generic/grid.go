package generic

import "github.com/recluse-games/floodfill/pkg/floodfill"

// GenericGrid structure for 2d games
type GenericGrid struct {
	nodes [][]floodfill.Node
}

func (n GenericGrid) Nodes() [][]floodfill.Node {
	return n.nodes
}

func (n GenericGrid) SetNode(x int, y int, node floodfill.Node) {
	n.nodes[x][y] = node
}
