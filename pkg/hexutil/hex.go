package hexutil

import (
	"github.com/toolvox/utilgo/api"
	"github.com/toolvox/utilgo/pkg/mathutil"
)

type Direction int

const (
	E Direction = iota
	SE
	SW
	W
	NW
	NE
)

var (
	Directions = [6]Direction{E, SE, SW, W, NW, NE}
)

func (d Direction) ToUnitHex3() Hex3 {
	switch d {
	case E:
		return FromXY(1, 0)
	case SE:
		return FromXY(1, -1)
	case SW:
		return FromXY(0, -1)
	case W:
		return FromXY(-1, 0)
	case NW:
		return FromXY(-1, 1)
	case NE:
		return FromXY(0, 1)
	default:
		panic("unknown direction")
	}
}

type Hex3 struct {
	X int `yaml:"x" json:"x"`
	Y int `yaml:"y" json:"y"`
	Z int `yaml:"z" json:"z"`
}

var HexOrigin = Hex3{0, 0, 0}

func FromXY(x, y int) Hex3 {
	return Hex3{X: x, Y: y, Z: -(x + y)}
}
func FromXZ(x, z int) Hex3 {
	return Hex3{X: x, Y: -(x + z), Z: z}
}
func FromYZ(y, z int) Hex3 {
	return Hex3{X: -(y + z), Y: y, Z: z}
}

func (o Hex3) Clone() Hex3 {
	return o
}

func (o Hex3) Add(by Hex3) Hex3 {
	return Hex3{
		o.X + by.X,
		o.Y + by.Y,
		o.Z + by.Z,
	}
}
func (o Hex3) Sub(by Hex3) Hex3 {
	return Hex3{
		o.X - by.X,
		o.Y - by.Y,
		o.Z - by.Z,
	}
}

func (o Hex3) Radius() uint {
	return uint(mathutil.Max(mathutil.Abs(o.X), mathutil.Abs(o.Y), mathutil.Abs(o.Z)))
}

func (o Hex3) Neighbors() []Hex3 {
	var results []Hex3
	for _, d := range Directions {
		od := o.Add(d.ToUnitHex3())
		results = append(results, od)
	}
	return results
}

func Unique(coords ...Hex3) []Hex3 {
	visited := make(map[Hex3]api.Unit)
	var uniq []Hex3
	for _, c := range coords {
		if _, ok := visited[c]; !ok {
			uniq = append(uniq, c)
			visited[c] = api.U
		}
	}
	return uniq
}
