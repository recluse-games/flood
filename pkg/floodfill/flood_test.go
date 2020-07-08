package floodfill

import (
	"testing"
)

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

// A simple initilizer function to provide a default empty x by x grid state.
func emptyGridState(width int, height int) [][]Node {
	nodes := make([][]Node, height)

	for x := 0; x < width; x++ {
		row := make([]Node, width)

		for y := 0; y < height; y++ {
			node := GenericNode{}
			node.SetPoint(Point{x, y})
			node.SetID("0000")

			row[y] = &node
		}

		nodes[x] = row
	}

	return nodes
}

// A simple initilizer function to provide a default grid with a set of blocked nodes.
func blockedGridState(width int, height int) [][]Node {
	nodes := make([][]Node, height)

	for x := 0; x < width; x++ {
		row := make([]Node, width)

		for y := 0; y < height; y++ {
			node := GenericNode{}
			node.SetPoint(Point{x, y})
			node.SetID("0000")

			// Create an artifical wall here to block flood filling.
			if x == 2 {
				node.SetID("0002")
			}

			row[y] = &node
		}

		nodes[x] = row
	}

	return nodes
}

func TestFill(t *testing.T) {
	grid := GenericGrid{}
	grid.nodes = emptyGridState(4, 4)

	flood := Flood{
		Origin:      Point{0, 0},
		Grid:        &grid,
		MaxDistance: 8,
	}

	flood.Fill(Point{0, 0}, "0001", "0002")

	// Test full empty fill

	if flood.Grid.Nodes()[0][0].ID() != "0001" {
		t.Logf("Origin node not filled.")
		t.Fail()
	}

	if flood.Grid.Nodes()[3][3].ID() != "0001" {
		t.Logf("Node not filled.")
		t.Fail()
	}

	// Test blocked fill
	grid = GenericGrid{}
	grid.nodes = blockedGridState(4, 4)

	flood = Flood{
		Origin:      Point{0, 0},
		Grid:        &grid,
		MaxDistance: 3,
	}

	flood.Fill(Point{0, 0}, "0001", "0002")

	if flood.Grid.Nodes()[0][0].ID() != "0001" {
		t.Logf("Origin node not filled.")
		t.Fail()
	}

	if flood.Grid.Nodes()[3][3].ID() != "0000" {
		t.Logf("Blocking not correctly respected.")
		t.Fail()
	}
}
