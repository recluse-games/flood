package fill

// Node A generic interface that supports all required traits to be flood filled.
type Node interface {
	id() string
	setID(string)
	x() int
	setX(int)
	y() int
	setY(int)
	clone() Node
}
