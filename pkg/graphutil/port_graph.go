package graphutil

import (
	"slices"

	"github.com/toolvox/utilgo/pkg/sets"
)

type NodePort struct {
	Node int
	Port byte
}

type hyper_edge = sets.TinySet[NodePort]

func hyper_edge_cmp(base, candidate hyper_edge) (match bool, exact bool) {
	// fmt.Printf("  %v - %v\n", base, candidate)
	var find_count int
	for _, np := range candidate {
		var found bool
		for _, base_np := range base {
			found = np.Node == base_np.Node && np.Port == base_np.Port
			if found {
				break
			}
		}
		if found {
			find_count++
		}
	}
	return find_count > 0, find_count == len(base)
}

type PortGraph struct {
	NodePorts  []byte
	HyperEdges []hyper_edge

	FreeNodes []int
	FreeEdges []int
}

func NewPortGraph() PortGraph {
	return PortGraph{
		NodePorts:  []byte{},
		HyperEdges: []hyper_edge{},
		FreeNodes:  []int{},
		FreeEdges:  []int{},
	}
}

func (pg *PortGraph) NextNodeId(consume bool) int {
	if len(pg.FreeNodes) == 0 {
		return len(pg.NodePorts)
	}
	result := pg.FreeNodes[0]
	if consume {
		pg.FreeEdges = pg.FreeEdges[1:]
	}
	return result
}

func (pg *PortGraph) NextEdgeId(consume bool) int {
	if len(pg.FreeEdges) == 0 {
		return len(pg.HyperEdges)
	}
	result := pg.FreeEdges[0]
	if consume {
		pg.FreeEdges = pg.FreeEdges[1:]
	}
	return result
}

func (pg *PortGraph) AddNode(maxPort byte) int {
	nodeId := pg.NextNodeId(true)
	if nodeId >= len(pg.NodePorts) {
		pg.NodePorts = append(pg.NodePorts, maxPort)
	} else {
		pg.NodePorts[nodeId] = maxPort
	}
	return nodeId
}

func (pg *PortGraph) FindEdge(nodePorts ...NodePort) (found bool, matchId int, exact bool) {
	for edgeId, edge := range pg.HyperEdges {
		if slices.Contains(pg.FreeEdges, edgeId) {
			continue
		}
		found, exact = hyper_edge_cmp(edge, nodePorts)
		if found {
			matchId = edgeId
			return
		}
	}
	return false, -1, false
}

func (pg *PortGraph) AddEdge(nodePorts ...NodePort) int {
	found, edgeId, _ := pg.FindEdge(nodePorts...)
	// fmt.Printf("%t, %d, %t\n", found, edgeId, exact)
	if !found {
		edgeId = pg.NextEdgeId(true)
	}
	if edgeId >= len(pg.HyperEdges) {
		pg.HyperEdges = append(pg.HyperEdges, nodePorts)
	} else {
		pg.HyperEdges[edgeId].Add(nodePorts...)
	}
	return edgeId
}
