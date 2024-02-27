// Package api provides the interfaces which the structs in the library implement, and common type definitions.
package api

// Unit is an empty struct. An efficient way to represent nothing.
type Unit = struct{}

// U is the unit. The single instance of the type nothing.
var U = Unit{}
