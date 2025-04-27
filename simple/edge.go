package simple

import "fmt"

type (
	// Edge is the line between vertices (points)
	// each Edge is a set containing two vertices that shows the line between those two
	Edge struct {
		Endpoints [2]*Vertex
	}
	Edges []*Edge
)

func NewEdges(v Vertices, endpoints Endpoint) Edges {
	maxEdges := len(v) * (len(v) - 1) / 2
	(&endpoints).UniquePairs()

	fmt.Println(maxEdges)
	if len(endpoints) > maxEdges {
		panic(fmt.Errorf("Invalid endpoints: Maximum number of endpoints for given Vertices is %v\n", maxEdges))
	}

	e := make([]*Edge, len(endpoints))
	for i := range e {
		v0 := v.FindVertex(endpoints[i][0])
		v1 := v.FindVertex(endpoints[i][1])
		if v0 == nil || v1 == nil {
			panic(fmt.Errorf("invalid endpoint name: %c or %c does not exists in vertices\n",
				endpoints[i][0], endpoints[i][1]))
		}

		e[i] = &Edge{Endpoints: [2]*Vertex{v0, v1}}
	}

	return e
}
