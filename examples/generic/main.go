package generic

func main() {
	grid := GenericGrid{}
	nodes := make([][]Node, height)

	// Generate an empty 10x10 grid of nodes
	for x := 0; x < 10; x++ {
		row := make([]Node, 10)

		for y := 0; y < 10; y++ {
			node := GenericNode{}
			node.SetPoint(Point{x, y})
			node.SetID("0000")

			row[y] = &node
		}

		nodes[x] = row
	}

	// Create a new flood instance with an origin point and max taxi-cab distance to fill to.
	flood := Flood{
		Origin:      Point{0, 0},
		Grid:        &grid,
		MaxDistance: 8,
	}

	// Fill the grid
	flood.Fill(Point{0, 0}, "0001", "0002")

	// Return your now filled in nodes
	return flood.Grid.Nodes()[0][0].ID()
}
