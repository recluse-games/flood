package floodfill

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"Nort", "South", "East", "West"}[d]
}

func (d Direction) Int() int {
	return [...]int{0, 1, 2, 3}[d]
}
