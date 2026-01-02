package hexutil

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/toolvox/utilgo/pkg/sets"
)

type HexShape string

const (
	RingShape  HexShape = "Ring"
	RangeShape HexShape = "Range"
	StarShape  HexShape = "Star"
)

func (s HexShape) GetShapeTiles(origin Hex3, value uint) []Hex3 {
	switch s {
	case RingShape:
		return Ring(origin, value)
	case RangeShape:
		return Range(origin, value)
	case StarShape:
		return Star(origin, value)
	default:
		panic(fmt.Errorf("ERROR: unrecognized shape: %s", s))
	}
}

func Ring(origin Hex3, radius uint) []Hex3 {
	var r uint
	ring := []Hex3{origin}
	visited := sets.NewSet[Hex3]()
	for r < radius {
		var ringCandidates []Hex3
		for _, hex := range ring {
			visited.Add(hex)
			ringCandidates = append(ringCandidates, hex.Neighbors()...)
		}
		var newRing []Hex3
		for _, hex := range ringCandidates {
			if _, ok := visited[hex]; !ok {
				newRing = append(newRing, hex)
			}
		}
		ring = newRing
		r++
	}
	return Unique(ring...)
}

func Range(origin Hex3, radius uint) []Hex3 {
	results := []Hex3{origin}
	for i := uint(1); i <= radius; i++ {
		results = append(results, Ring(origin, i)...)
	}
	return results
}

func Arrow(origin, direction Hex3, limitLength uint) []Hex3 {
	var result []Hex3
	next := origin.Add(direction)
	result = append(result, next)
	for limitLength == 0 || len(result) < int(limitLength) {
		next = next.Add(direction)
		result = append(result, next)
	}
	return result
}

func Star(origin Hex3, limitLength uint) []Hex3 {
	var result []Hex3
	for _, d := range Directions {
		dir := d.ToUnitHex3()
		result = append(result, Arrow(origin, dir, limitLength)...)
	}
	return result
}

type Shape struct {
	Shape HexShape
	Value uint
}

func ParseShape(text string) Shape {
	parts := strings.Split(text, " ")

	if len(parts) == 0 {
		panic(fmt.Errorf("ERROR: illegal shape: %s", text))
	}

	if len(parts) >= 2 {
		val, err := strconv.ParseInt(parts[1], 0, 0)
		if err != nil {
			panic(fmt.Errorf("ERROR: illegal value %s: %v", parts[1], err))
		}
		return Shape{HexShape(parts[0]), uint(val)}
	}

	return Shape{HexShape(text), 0}
}

func (s Shape) GetTiles(origin Hex3) []Hex3 {
	return s.Shape.GetShapeTiles(origin, s.Value)
}

func (s Shape) Exists() bool {
	return len(s.Shape) != 0
}

func (s Shape) String() string {
	return fmt.Sprintf("%s %d", s.Shape, s.Value)
}
