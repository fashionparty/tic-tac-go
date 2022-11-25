package main

// Symbol Represents one of three possible symbols on field: cross (`X`), circle (`0`) or empty field (`*`)
type Symbol int

const (
	Empty = iota
	Cross
	Circle
)
