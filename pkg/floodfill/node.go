package floodfill

// Node A Node represents a point in a 2D matrix with identifying information
type Node interface {
	//ID Returns a valid identifier for this nood to be filled
	ID() string

	//SetID Sets the ID value of a structure conforming to the Node type
	SetID(string)

	//Point Returns a pair of cartesian coordinates
	Point() Point

	//SetPoint Sets the point where this node currently exists
	SetPoint(Point)

	//Clone Creates a deep copy of this node
	Clone() Node
}
