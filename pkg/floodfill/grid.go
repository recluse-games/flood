package floodfill

// Grid represents a 2D matrix comprised of nodes.
type Grid interface {
	//Nodes contains a 2D matrix of nodes that make up this 2D Grid
	Nodes() [][]Node

	//SetNode Updates a Node at a particular X and Y coordinate in this 2D Grid
	SetNode(int, int, Node)
}
