package main

import (
	"SimpleGraph/simple"
	"fmt"
)

func main() {
	V := simple.NewVertices([]byte{'A', 'B', 'C'})
	E := simple.NewEdges(V,
		simple.Endpoint{
			{'A', 'D'}, {'A', 'B'}, {'A', 'C'},
		},
	)

	_ = E
	G := simple.NewSimpleGraph(V, E)
	fmt.Print(G)
}
