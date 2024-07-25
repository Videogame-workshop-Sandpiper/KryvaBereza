package main

import "strconv"

type Vector2 struct {
	x int
	y int
}

// Set vector's coordinates
func (v *Vector2) Init(X int, Y int) {
	v.x = X
	v.y = Y
}

// Create vector in one line
func NewV2(X int, Y int) Vector2 {
	var v Vector2
	v.Init(X, Y)
	return v
}

type Vector3 struct {
	x int
	y int
	z int
}

// Set vector's coordinates
func (v *Vector3) Init(X int, Y int, Z int) {
	v.x = X
	v.y = Y
	v.z = Z
}

// Create vector in one line
func NewV3(X int, Y int, Z int) Vector3 {
	var v Vector3
	v.Init(X, Y, Z)
	return v
}

//Convert from string to uint8 in one line
func StrToInt8(i string) uint8 {
	I, err := strconv.Atoi(i)
	if err != nil {
		return 0
	}
	return uint8(I)
}

//Convert from string to int in one line
func StrToInt(i string) int {
	I, err := strconv.Atoi(i)
	if err != nil {
		return 0
	}
	return I
}
