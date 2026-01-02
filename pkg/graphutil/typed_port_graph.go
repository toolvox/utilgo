package graphutil

type TypedPortGraph[TPortType ~byte, TEdgeType ~byte] struct {
	PortGraph
	NodePortTypes  [][]TPortType
	HyperEdgeTypes []TEdgeType
}

func NewTypedPortGraph[TPortType ~byte, TEdgeType ~byte]() TypedPortGraph[TPortType, TEdgeType] {
	return TypedPortGraph[TPortType, TEdgeType]{
		PortGraph:      NewPortGraph(),
		NodePortTypes:  [][]TPortType{},
		HyperEdgeTypes: []TEdgeType{},
	}
}

func (tpg *TypedPortGraph[TPortType, TEdgeType]) AddNode(ports ...TPortType) int {
	maxPort := byte(len(ports) - 1)
	nodeId := tpg.PortGraph.AddNode(maxPort)
	if nodeId >= len(tpg.NodePortTypes) {
		tpg.NodePortTypes = append(tpg.NodePortTypes, ports)
	} else {
		tpg.NodePortTypes[nodeId] = ports
	}
	return nodeId
}

func (tpg *TypedPortGraph[TPortType, TEdgeType]) AddEdge(edgeType TEdgeType, nodePorts ...NodePort) int {
	edgeId := tpg.PortGraph.AddEdge(nodePorts...)
	if edgeId >= len(tpg.HyperEdgeTypes) {
		tpg.HyperEdgeTypes = append(tpg.HyperEdgeTypes, edgeType)
	} else {
		tpg.HyperEdgeTypes[edgeId] = edgeType
	}
	return edgeId
}
