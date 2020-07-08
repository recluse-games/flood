package floodfill

// Point A Point holds a cartesian coordinate pair with x and y as standard integer values.
type Point struct {
	X int
	Y int
}

// SetX sets the x value of a Point
func (p *Point) SetX(x int) {
	p.X = x
}

// SetY sets the y value of a Point
func (p *Point) SetY(y int) {
	p.Y = y
}

// Distance Calculates the Manhattan distance from this point to a given point in euclidan geometric
// space and returns it as an integer value.
func (p *Point) Distance(point Point) int {
	distance := abs(p.X-point.X) + abs(p.Y-point.Y)

	return distance
}

// abs calculates the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
