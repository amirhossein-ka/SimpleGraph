package simple

import (
	"fmt"
	"strings"
)

type (
	Graph struct {
		Vertices Vertices
		Edges    Edges
	}
)

func (s *Graph) String() string {
	var b = new(strings.Builder)
	b.WriteString(fmt.Sprintf("SimpleGraph(V,E):\n"))
	b.WriteString(fmt.Sprintf("\tVertices: {"))
	// append vertices:
	for i, vertex := range s.Vertices {
		b.WriteByte(vertex.Name)
		if i < len(s.Vertices)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString("}\n")

	b.WriteString("\tEdges: {")
	// append edges:
	for i, e := range s.Edges {
		b.WriteString(fmt.Sprintf("(%c,%c)", e.Endpoints[0].Name, e.Endpoints[1].Name))
		if i < len(s.Edges)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString("}\n")

	return b.String()
}

func NewSimpleGraph(v Vertices, e Edges) *Graph {
	return &Graph{
		Edges:    e,
		Vertices: v,
	}
}
