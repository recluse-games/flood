package flood

// Direction an integer value that represents a cardinal direction
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

// String returns the string representation of a direction.
func (d Direction) String() string {
	return [...]string{"Nort", "South", "East", "West"}[d]
}

// Int returns the integer representation of a direction.
func (d Direction) Int() int {
	return [...]int{0, 1, 2, 3}[d]
}
