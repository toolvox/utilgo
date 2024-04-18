// Package test holds testing helpers.
package test

// TypedTestData is an interface that allows getting an element from a type, by index.
//
// Usage:
//
//	// Write a "typed" test function like:
//	func TypedTest_MyThing[MyT any](t *testing.T, data TypedTestData[MyT]) {
//	    // now you can use data.Element(i) to get the i-th element from the function.
//	    _ = data.Element(0)
//	    // ...
//	}
//	// Call the function above from your normal test functions, like:
//	func Test_Normal(t *testing.T) {
//	    TypedTest_MyThing[string](t, stringData)
//	    TypedTest_MyThing[myType](t, myTypeData)
//	}
//
// Definition:
type TypedTestData[A any] interface {
	Element(i int) A
}

// TestDataFor implements [TypedTestData] by holding a func.
//
// For a simple case of int:
//
//	var intData = TestDataFor[int]{ ElementFunc: func(i int) int { return i } }
//
// For a case of string:
//
//	var stringData = TestDataFor[string]{ ElementFunc: func(i int) string { return fmt.Sprint(i) } }
//
// Definition:
type TestDataFor[A any] struct {
	ElementFunc func(i int) A
}

// Element gets an indexed element of the type by invoking the function.
func (d TestDataFor[A]) Element(i int) A {
	return d.ElementFunc(i)
}

// TypedTestConstructor is an interface for creating new instances of a type.
//
// Usage:
//
//	type MyTypeConstructor struct{}
//
//	func (MyTypeConstructor) New() MyType {
//	    return MyType{} // Return a new MyType instance
//	}
//
// Definition:
type TypedTestConstructor[A any] interface {
	New() A
}

// TestConstructorFor provides a concrete implementation of TypedTestConstructor.
//
// Usage:
//
//	var intConstructor = TestConstructorFor[int]{ NewFunc: func() int { return 0 } }
//
// Definition:
type TestConstructorFor[A any] struct {
	NewFunc func() A
}

// New calls NewFunc to create and return a new instance of A.
func (c TestConstructorFor[A]) New() A {
	return c.NewFunc()
}
