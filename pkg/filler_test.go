package fill

import (
	"fmt"
	"testing"
)

// GenericNode is an example node structure that implements all required traits for the flood filling.
type GenericNode struct {
	identifier string
	xPos       int
	yPos       int
}

func (n GenericNode) id() string {
	return n.identifier
}

func (n *GenericNode) setID(id string) {
	n.identifier = id
}

func (n *GenericNode) x() int {
	return n.xPos
}

func (n *GenericNode) setX(x int) {
	n.xPos = x
}

func (n *GenericNode) y() int {
	return n.yPos
}

func (n *GenericNode) setY(y int) {
	n.yPos = y
}

func (n GenericNode) clone() Node {
	nodeCopy := GenericNode{
		identifier: n.identifier,
		xPos:       n.xPos,
		yPos:       n.yPos,
	}

	typeCastedNode := nodeCopy

	return &typeCastedNode
}

type GenericBoard struct {
	boardNodes [][]Node
}

func (n GenericBoard) nodes() [][]Node {
	return n.boardNodes
}

func (n *GenericBoard) setNodes() {
	nodes := make([][]Node, 10)

	for y := 0; y < 10; y++ {
		row := make([]Node, 10)

		for x := 0; x < 10; x++ {
			node := GenericNode{}
			node.setX(x)
			node.setY(y)
			node.setID("0000")

			row[x] = &node
		}

		nodes[y] = row
	}
	n.boardNodes = nodes
}

func (n GenericBoard) setNode(x int, y int, node Node) {
	n.boardNodes[x][y] = node
}

func TestFiller(t *testing.T) {
	newBoard := GenericBoard{}
	newBoard.setNodes()

	newFiller := Filler{
		grid: &newBoard,
	}

	newFiller.flood(0, 0, "0001", "0002")

	fmt.Printf("%v", newFiller.grid.nodes()[0][2])

}
