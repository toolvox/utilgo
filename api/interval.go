package api

type GrowFlags int

const (
	GROW_BOTH_OVERFLOW_RIGHT GrowFlags = 0
	GROW_LEFT_OVERFLOW_RIGHT GrowFlags = 1
	GROW_RIGHT_OVERFLOW_LEFT GrowFlags = 2
	GROW_BOTH_OVERFLOW_LEFT  GrowFlags = 3
	GROW_NO_OVERFLOW         GrowFlags = 4
)

type Interval[T Number] interface {
	Min() T
	Max() T
	Len() T

	IsEmpty() bool
	IsSingleton() bool
	IsCompound() bool

	Enumerate(step T) []T
	Intervals() []Interval[T]

	Contains(value T) bool
	Overlaps(other Interval[T]) bool
	Equals(other Interval[T]) bool

	Union(other Interval[T]) Interval[T]
	Intersection(other Interval[T]) Interval[T]
	Difference(other Interval[T]) Interval[T]

	Resize(newSize T, growMode GrowFlags) Interval[T]
	Scale(scale float64, growMode GrowFlags) Interval[T]
	Translate(offset T, back bool) Interval[T]
}
