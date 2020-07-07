package fill

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func (d Direction) Int() int {
	return [...]int{0, 1, 2, 3}[d]
}
