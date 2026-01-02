# API

* `api.go`:
  * `type Unit : struct{}`
  * `var U = Unit{}`
* `sets.go`:
  ```go
  type BasicSet[C comparable] interface {
    String() string
    Len() int
    Elements() []C
    Add(elements ...C)
    Contains(elements ...C) bool
    Remove(elements ...C)
  }

  type ElementSet[C comparable] interface {
    BasicSet[C]

    Union(other ElementSet[C]) ElementSet[C]
    Intersection(other ElementSet[C]) ElementSet[C]
    Difference(other ElementSet[C]) ElementSet[C]

    ThreeWay(other ElementSet[C]) [3]ElementSet[C]

    UnionWith(elements ...C) ElementSet[C]
    IntersectionWith(elements ...C) ElementSet[C]
    DifferenceWith(elements ...C) ElementSet[C]
  }
  ```
* `number_types.go`:
  ```go
  type Number interface{ Integer | FloatingPoint }
  type Integer interface { SignedInteger | UnsignedInteger }
  type SignedInteger interface { ~int | ~int8 | ~int16 | ~int32 | ~int64 }
  type UnsignedInteger interface { ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr }
  type FloatingPoint interface{ ~float32 | ~float64 }
  ```
* `interval.go`:
  * `type GrowFlags int`
  * `type Interval[T] interface { ... }` [#TODO: Is this needed?]
