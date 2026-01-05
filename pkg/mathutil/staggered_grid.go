package mathutil

type SGrid[TCell any, TEdge any] struct {
	Width, Height int
	Cells         [][]TCell
	Edges         [][][2]TEdge
}

func NewSGrid[TCell any, TEdge any](width, height int, defaultCell func() TCell, defaultEdge func() TEdge) *SGrid[TCell, TEdge] {
	g := SGrid[TCell, TEdge]{
		Width:  width,
		Height: height,
		Cells:  make([][]TCell, width),
		Edges:  make([][][2]TEdge, width+1),
	}
	for x := range g.Cells {
		g.Cells[x] = make([]TCell, height)
		g.Edges[x] = make([][2]TEdge, height+1)
		for y := 0; y < height; y++ {
			g.Cells[x][y] = defaultCell()
			g.Edges[x][y] = [2]TEdge{defaultEdge(), defaultEdge()}
		}
		g.Edges[x][height] = [2]TEdge{defaultEdge(), defaultEdge()}
	}
	g.Edges[width] = make([][2]TEdge, height+1)
	for y := 0; y < height; y++ {
		g.Edges[width][y] = [2]TEdge{defaultEdge(), defaultEdge()}
	}
	g.Edges[width][height] = [2]TEdge{defaultEdge(), defaultEdge()}
	return &g
}

func (g *SGrid[TCell, TEdge]) SetCell(x, y int, cell TCell) {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return
	}
	g.Cells[x][y] = cell
}

func (g *SGrid[TCell, TEdge]) GetCell(x, y int) TCell {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return *new(TCell)
	}
	return g.Cells[x][y]
}

func (g *SGrid[TCell, TEdge]) GetCellPt(pt Point2) TCell {
	return g.GetCell(pt.XY())
}

// GetNeighbors returns in NESW order
func (g *SGrid[TCell, TEdge]) GetNeighbors(x, y int) []TCell {

	return []TCell{
		g.GetCellPt(N.Offset(x, y)),
		g.GetCellPt(E.Offset(x, y)),
		g.GetCellPt(S.Offset(x, y)),
		g.GetCellPt(W.Offset(x, y)),
	}
}

// GetNeighborsMap returns in arbitrary order
func (g *SGrid[TCell, TEdge]) GetNeighborsMap(x, y int) map[Point2]TCell {
	return map[Point2]TCell{
		N.Offset(x, y): g.GetCellPt(N.Offset(x, y)),
		E.Offset(x, y): g.GetCellPt(E.Offset(x, y)),
		S.Offset(x, y): g.GetCellPt(S.Offset(x, y)),
		W.Offset(x, y): g.GetCellPt(W.Offset(x, y)),
	}
}

func (g *SGrid[TCell, TEdge]) SetEdges(x, y int, n, e, s, w TEdge) {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return
	}
	g.Edges[x][y][0] = w
	g.Edges[x][y][1] = s
	g.Edges[x][y+1][1] = n
	g.Edges[x+1][y][0] = e
}

// GetEdges returns (northEdge, eastEdge, southEdge, westEdge)
func (g *SGrid[TCell, TEdge]) GetEdges(x, y int) (TEdge, TEdge, TEdge, TEdge) {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return *new(TEdge), *new(TEdge), *new(TEdge), *new(TEdge)
	}
	return g.Edges[x][y+1][1], g.Edges[x+1][y][0], g.Edges[x][y][1], g.Edges[x][y][0]
}

// GetEdgesPt returns (northEdge, eastEdge, southEdge, westEdge)
func (g *SGrid[TCell, TEdge]) GetEdgesPt(pt Point2) (TEdge, TEdge, TEdge, TEdge) {
	return g.GetEdges(pt.XY())
}

func (g *SGrid[TCell, TEdge]) Swap(x1, y1, x2, y2 int) {
	if min(x1, x2) < 0 || max(x1, x2) >= g.Width || min(y1, y2) < 0 || max(y1, y2) >= g.Height {
		return
	}
	g.Cells[x1][y1], g.Cells[x2][y2] = g.Cells[x2][y2], g.Cells[x1][y1]
}
