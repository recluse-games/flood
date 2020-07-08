package generic

import "github.com/recluse-games/flood/pkg/flood"

func main() {
	grid := GenericGrid{}
	nodes := make([][]flood.Node, 10)

	// Generate an empty 10x10 grid of nodes
	for x := 0; x < 10; x++ {
		row := make([]flood.Node, 10)

		for y := 0; y < 10; y++ {
			node := GenericNode{}
			node.SetPoint(flood.Point{X: x, Y: y})
			node.SetID("0000")

			row[y] = &node
		}

		nodes[x] = row
	}

	// Create a new flood instance with an origin point and max taxi-cab distance to fill to.
	filler := flood.NewFiller(flood.Point{X: 0, Y: 0}, grid, 8)

	// Fill the grid
	filler.Fill(flood.Point{X: 0, Y: 0}, "0001", "0002")

	// Return your now filled in nodes
	filler.Grid.Nodes()[0][0].ID()
}
