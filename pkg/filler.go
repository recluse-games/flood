package fill

// Filler flood fills a board state
type Filler struct {
	grid Grid
}

func (f Filler) flood(x int, y int, filledID string, blockedID string) {
	if f.grid.nodes()[x][y].id() != filledID && f.grid.nodes()[x][y].id() != blockedID {
		node := f.grid.nodes()[x][y].clone()
		node.setID(filledID)
		f.grid.setNode(x, y, node)

		directions := []Direction{North, South, East, West}

		for direction := range directions {
			switch direction {
			case North.Int():
				if node.y()+1 < len(f.grid.nodes()[x]) {
					f.flood(node.x(), node.y()+1, filledID, blockedID)
				}
			case South.Int():
				if node.y()-1 > 0 {
					f.flood(node.x(), node.y()-1, filledID, blockedID)
				}
			case East.Int():
				if node.x()+1 < len(f.grid.nodes()) {
					f.flood(node.x()+1, node.y(), filledID, blockedID)
				}
			case West.Int():
				if node.x()-1 > 0 {
					f.flood(node.x()-1, node.y(), filledID, blockedID)
				}
			}
		}
	}
}
