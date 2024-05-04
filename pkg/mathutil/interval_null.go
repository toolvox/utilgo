package mathutil

import (
	"github.com/toolvox/utilgo/api"
)

type nullInterval[N Number] struct{}

func NewEmpty[N Number]() api.Interval[N] {
	return nullInterval[N]{}
}

func (ni nullInterval[N]) Min() N { return N(0) }
func (ni nullInterval[N]) Max() N { return N(0) }
func (ni nullInterval[N]) Len() N { return N(0) }

func (ni nullInterval[N]) IsEmpty() bool     { return true }
func (ni nullInterval[N]) IsSingleton() bool { return false }
func (ni nullInterval[N]) IsCompound() bool  { return false }

func (ni nullInterval[N]) Enumerate(_ N) []N            { return nil }
func (ni nullInterval[N]) Intervals() []api.Interval[N] { return nil }

func (ni nullInterval[N]) Contains(_ N) bool               { return false }
func (ni nullInterval[N]) Overlaps(_ api.Interval[N]) bool { return false }
func (ni nullInterval[N]) Equals(other api.Interval[N]) bool {
	return other == nil || other.IsEmpty()
}

func (ni nullInterval[N]) Union(other api.Interval[N]) api.Interval[N]    { return other }
func (ni nullInterval[N]) Intersection(_ api.Interval[N]) api.Interval[N] { return ni }
func (ni nullInterval[N]) Difference(_ api.Interval[N]) api.Interval[N]   { return ni }

func (ni nullInterval[N]) Resize(_ N, _ api.GrowFlags) api.Interval[N]      { return ni }
func (ni nullInterval[N]) Scale(_ float64, _ api.GrowFlags) api.Interval[N] { return ni }
func (ni nullInterval[N]) Translate(_ N, _ bool) api.Interval[N]            { return ni }
