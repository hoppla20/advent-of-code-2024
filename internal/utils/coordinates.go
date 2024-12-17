package utils

import (
	"fmt"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type CoordinateSystemSize Coordinate

type CoordinateSystem struct {
	size CoordinateSystemSize
}

type OutOfBoundsError struct {
	state      CoordinateSystem
	coordinate Coordinate
}

var state CoordinateSystem

func SetCoordinateSystemSize(size CoordinateSystemSize) {
	state.size = size
}

func (e OutOfBoundsError) Error() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Coordinate %v not in coordinate system with size %v", e.coordinate, e.state.size))

	return sb.String()
}

func (c Coordinate) String() string {
	s := fmt.Sprintf("{X: %d, Y: %d}", c.X, c.Y)

	return s
}

func (c Coordinate) Up() (Coordinate, error) {
	if c.Y-1 < 0 {
		return Coordinate{}, &OutOfBoundsError{state, Coordinate{X: c.X, Y: c.Y - 1}}
	}

	return Coordinate{X: c.X, Y: c.Y - 1}, nil
}

func (c Coordinate) Right() (Coordinate, error) {
	if c.X+1 >= state.size.X {
		return Coordinate{}, &OutOfBoundsError{state, Coordinate{X: c.X + 1, Y: c.Y}}
	}

	return Coordinate{X: c.X + 1, Y: c.Y}, nil
}

func (c Coordinate) Down() (Coordinate, error) {
	if c.Y+1 >= state.size.Y {
		return Coordinate{}, &OutOfBoundsError{state, Coordinate{X: c.X, Y: c.Y + 1}}
	}

	return Coordinate{X: c.X, Y: c.Y + 1}, nil
}

func (c Coordinate) Left() (Coordinate, error) {
	if c.X-1 < 0 {
		return Coordinate{}, &OutOfBoundsError{state, Coordinate{X: c.X - 1, Y: c.Y}}
	}

	return Coordinate{X: c.X - 1, Y: c.Y}, nil
}
