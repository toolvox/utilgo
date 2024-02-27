package sets_test

import (
	"testing"

	"utilgo/pkg/sets"
)

func TestNewSet(t *testing.T) {
	s := sets.NewSet(1, 2, 3)
	if len(s) != 3 {
		t.Errorf("Expected set size of 3, got %d", len(s))
	}
}

func TestAdd(t *testing.T) {
	s := sets.NewSet[int]()
	s.Add(1)
	if !s.Contains(1) {
		t.Errorf("Set should contain element 1")
	}
}

func TestUnion(t *testing.T) {
	s1 := sets.NewSet(1, 2)
	s2 := sets.NewSet(2, 3)
	UnionWith := s1.Union(s2)
	if len(UnionWith) != 3 || !UnionWith.Contains(1) || !UnionWith.Contains(2) || !UnionWith.Contains(3) {
		t.Errorf("UnionWith does not contain the correct elements")
	}
}

func TestUnionWith(t *testing.T) {
	s := sets.NewSet(1, 2)
	UnionWith := s.UnionWith(3, 4)
	if len(UnionWith) != 4 || !UnionWith.Contains(1) || !UnionWith.Contains(4) {
		t.Errorf("UnionWith does not contain the correct elements")
	}
}

func TestIntersection(t *testing.T) {
	s1 := sets.NewSet(1, 2, 3)
	s2 := sets.NewSet(2, 3, 4)
	IntersectionWith := s1.Intersection(s2)
	if len(IntersectionWith) != 2 || !IntersectionWith.Contains(2) || !IntersectionWith.Contains(3) {
		t.Errorf("IntersectionWith does not contain the correct elements")
	}
}

func TestIntersectionWith(t *testing.T) {
	s := sets.NewSet(1, 2, 3)
	IntersectionWith := s.IntersectionWith(2, 3, 4)
	if len(IntersectionWith) != 2 || !IntersectionWith.Contains(2) || !IntersectionWith.Contains(3) {
		t.Errorf("IntersectionWith does not contain the correct elements")
	}
}

func TestDifference(t *testing.T) {
	s1 := sets.NewSet(1, 2, 3)
	s2 := sets.NewSet(2, 3, 4)
	DifferenceWith := s1.Difference(s2)
	if len(DifferenceWith) != 1 || !DifferenceWith.Contains(1) {
		t.Errorf("DifferenceWith does not contain the correct elements")
	}
}

func TestDifferenceWith(t *testing.T) {
	s := sets.NewSet(1, 2, 3)
	DifferenceWith := s.DifferenceWith(2, 3)
	if len(DifferenceWith) != 1 || !DifferenceWith.Contains(1) {
		t.Errorf("DifferenceWith does not contain the correct elements")
	}
}

func TestThreeWay(t *testing.T) {
	s1 := sets.NewSet(1, 2)
	s2 := sets.NewSet(2, 3)
	threeWay := s1.ThreeWay(s2)
	if len(threeWay[0]) != 1 || len(threeWay[1]) != 1 || len(threeWay[2]) != 1 {
		t.Errorf("ThreeWay does not contain the correct elements")
	}
}

func TestContains(t *testing.T) {
	s := sets.NewSet(1, 2, 3)
	if !s.Contains(1) {
		t.Errorf("Set should contain element 1")
	}
	if s.Contains(4) {
		t.Errorf("Set should not contain element 4")
	}
}

func TestString(t *testing.T) {
	s := sets.NewSet(1, 2, 3)
	expected := "{ 1, 2, 3 }"
	result := s.String()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
