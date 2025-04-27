package simple

import "fmt"

type SGraph interface {
	fmt.Stringer
	DegreePtr(v *Vertex) uint
	NeighborPtr(v *Vertex) uint

	// DegreeName will return -1 if v is not found in graph
	DegreeName(v byte) int
	// NeighborName will return -1 if v is not found in graph
	NeighborName(v byte) int
}
