package test

// TypedTestData is an interface that allows getting an element from a type, by index.
//
// Usage:
//   // Write a "typed" test function like:
//   func TypedTest_MyThing[MyT any](t *testing.T, data TypedTestData[MyT]) {
//       // now you can use data.Element(i) to get the i-th element from the function.
//       _ = data.Element(0)
//       // ...
//   }
//   // Call the function above from your normal test functions, like:
//   func Test_Normal(t *testing.T) {
//       TypedTest_MyThing[string](t, stringData)
//       TypedTest_MyThing[myType](t, myTypeData)
//   }
// Definition:
type TypedTestData[A any] interface {
	Element(i int) A
}

// TestDataForAny implements [TypedTestData] by holding a func.
//
// For a simple case of int:
//   var intData = TestDataForAny[int]{ ElementFunc: func(i int) int { return i } }
// For a case of string:
//   var stringData = TestDataForAny[string]{ ElementFunc: func(i int) string { return fmt.Sprint(i) } }
// Definition:
type TestDataForAny[A any] struct {
	ElementFunc func(i int) A
}

// Element gets an indexed element of the type by invoking the function.
func (d TestDataForAny[A]) Element(i int) A {
	return d.ElementFunc(i)
}
