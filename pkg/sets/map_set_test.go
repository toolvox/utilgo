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

func TestSetUnion(t *testing.T) {
	s1 := sets.NewSet(1, 2)
	s2 := sets.NewSet(2, 3)
	union := s1.SetUnion(s2)
	if len(union) != 3 || !union.Contains(1) || !union.Contains(2) || !union.Contains(3) {
		t.Errorf("Union does not contain the correct elements")
	}
}

func TestUnion(t *testing.T) {
	s := sets.NewSet(1, 2)
	union := s.Union(3, 4)
	if len(union) != 4 || !union.Contains(1) || !union.Contains(4) {
		t.Errorf("Union does not contain the correct elements")
	}
}

func TestSetIntersection(t *testing.T) {
	s1 := sets.NewSet(1, 2, 3)
	s2 := sets.NewSet(2, 3, 4)
	intersection := s1.SetIntersection(s2)
	if len(intersection) != 2 || !intersection.Contains(2) || !intersection.Contains(3) {
		t.Errorf("Intersection does not contain the correct elements")
	}
}

func TestIntersection(t *testing.T) {
	s := sets.NewSet(1, 2, 3)
	intersection := s.Intersection(2, 3, 4)
	if len(intersection) != 2 || !intersection.Contains(2) || !intersection.Contains(3) {
		t.Errorf("Intersection does not contain the correct elements")
	}
}

func TestSetDifference(t *testing.T) {
	s1 := sets.NewSet(1, 2, 3)
	s2 := sets.NewSet(2, 3, 4)
	difference := s1.SetDifference(s2)
	if len(difference) != 1 || !difference.Contains(1) {
		t.Errorf("Difference does not contain the correct elements")
	}
}

func TestDifference(t *testing.T) {
	s := sets.NewSet(1, 2, 3)
	difference := s.Difference(2, 3)
	if len(difference) != 1 || !difference.Contains(1) {
		t.Errorf("Difference does not contain the correct elements")
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
