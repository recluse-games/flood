package generic

// GenericNode A simple generic node implementation for basic 2d grids.
type GenericNode struct {
	id    string
	point Point
}

func (n *GenericNode) ID() string {
	return n.id
}

func (n *GenericNode) SetID(id string) {
	n.id = id
}

func (n *GenericNode) Point() Point {
	return n.point
}

func (n *GenericNode) SetPoint(point Point) {
	n.point = point
}

func (n *GenericNode) Clone() Node {
	nodeCopy := GenericNode{
		id:    n.id,
		point: n.point,
	}

	typeCastedNode := nodeCopy

	return &typeCastedNode
}
