package graph

type Keyed interface {
	Key() string
}

type MultiDag[Node, Edge Keyed] struct {
	Nodes     []Node
	NodeMap   map[string]int
	Edges     []Edge
	EdgeMap   map[string]int
	EdgeNodes map[int]map[int]int
}

func (g *MultiDag[Node, Edge]) AddEdge(a, b Node, e Edge) {
	ida, ok := g.NodeMap[a.Key()]
	if !ok {
		ida = len(g.Nodes)
		g.NodeMap[a.Key()] = ida
		g.Nodes = append(g.Nodes, a)
	}
	idb, ok := g.NodeMap[b.Key()]
	if !ok {
		idb = len(g.Nodes)
		g.NodeMap[b.Key()] = idb
		g.Nodes = append(g.Nodes, b)
	}
	ide, ok := g.EdgeMap[e.Key()]
	if !ok {
		ide = len(g.Edges)
		g.EdgeMap[e.Key()] = ide
		g.Edges = append(g.Edges, e)
	}

	if _, ok = g.EdgeNodes[ida]; !ok {
		g.EdgeNodes[ida] = make(map[int]int)
	}
	g.EdgeNodes[ida][ide] = idb
}

func (g *MultiDag[Node, Edge]) AddBothEdges(a, b Node, e Edge) {
	g.AddEdge(a, b, e)
	g.AddEdge(b, a, e)
}

func (g MultiDag[Node, Edge]) AllNodes() []Node {
	return g.Nodes
}

func (g MultiDag[Node, Edge]) EdgesNeighbors(a Node) ([]Edge, []Node) {
	var resultEs []Edge
	var resultNs []Node
	ida := g.NodeMap[a.Key()]
	for ide, idb := range g.EdgeNodes[ida] {
		resultEs = append(resultEs, g.Edges[ide])
		resultNs = append(resultNs, g.Nodes[idb])
	}
	return resultEs, resultNs
}
