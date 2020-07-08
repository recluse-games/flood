package generic

import "github.com/recluse-games/flood/pkg/flood"

type GenericNode struct {
	id    string
	point flood.Point
}

func (n *GenericNode) ID() string {
	return n.id
}

func (n *GenericNode) SetID(id string) {
	n.id = id
}

func (n *GenericNode) Point() flood.Point {
	return n.point
}

func (n *GenericNode) SetPoint(point flood.Point) {
	n.point = point
}

func (n *GenericNode) Clone() flood.Node {
	nodeCopy := GenericNode{
		id:    n.id,
		point: n.point,
	}

	typeCastedNode := nodeCopy

	return &typeCastedNode
}
