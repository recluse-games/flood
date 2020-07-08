package floodfill

// Directions A slice representing the four cardinal directions.
var Directions = []Direction{North, South, East, West}

// Flood flood fills a board state
type Flood struct {
	Origin      Point
	Grid        Grid
	MaxDistance int
}

func (f Flood) validateDistance(point Point) bool {
	if f.Origin.Distance(point) < f.MaxDistance {
		return true
	}

	return false
}

// Fill Runs flood fill for a grid based on starting coordinates and IDs for fillage and blockage.
func (f Flood) Fill(point Point, filledID string, blockedID string) {
	node := f.Grid.Nodes()[point.X][point.Y].Clone()

	if node.ID() != filledID && node.ID() != blockedID && f.validateDistance(point) {
		node.SetID(filledID)
		f.Grid.SetNode(point.X, point.Y, node)

		for direction := range Directions {
			switch direction {
			case North.Int():
				if point.Y+1 < len(f.Grid.Nodes()[point.X]) {
					f.Fill(Point{point.X, point.Y + 1}, filledID, blockedID)
				}
			case South.Int():
				if 0 < point.Y-1 {
					f.Fill(Point{point.X, point.Y - 1}, filledID, blockedID)
				}
			case East.Int():
				if point.X+1 < len(f.Grid.Nodes()) {
					f.Fill(Point{point.X + 1, point.Y}, filledID, blockedID)
				}
			case West.Int():
				if 0 < point.X-1 {
					f.Fill(Point{point.X - 1, point.Y}, filledID, blockedID)
				}
			}
		}
	}
}
