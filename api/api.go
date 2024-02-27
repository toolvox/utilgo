// Package api provides the interface which the structs in the library implement.
package api

// Unit is an empty struct. An efficient way to represent nothing.
type Unit = struct{}

// U is the unit. The single instance of the type nothing.
var U = Unit{}
