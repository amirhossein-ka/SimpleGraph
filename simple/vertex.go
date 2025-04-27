package simple

type (
	// Vertex is the Node or Point in the graphs
	// each Vertex have a single char name
	// each Vertex have zero to many neighbor that is connected with Edges
	Vertex struct {
		Name byte
		data string
	}
	Vertices []*Vertex
)

// NewVertices
// names is the list of Vertices names, it's not a string (:
func NewVertices(names []byte) Vertices {
	v := make([]*Vertex, len(names))
	for i := range v {
		v[i] = &Vertex{Name: names[i]}
	}
	return v
}

func (v Vertices) FindVertex(name byte) *Vertex {
	for _, vertex := range v {
		if vertex.Name == name {
			return vertex
		}
	}
	return nil
}

func (v Vertices) Exists(name byte) bool {
	for _, vertex := range v {
		if vertex.Name == name {
			return true
		}
	}
	return false
}
