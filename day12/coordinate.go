package main

import "log"

type coordinate struct {
	x uint
	y uint
}

type coordinateSystemSize coordinate

func (c coordinate) neighbors(size coordinateSystemSize) []coordinate {
	if c.x >= size.x || c.y >= size.y {
		log.Fatal("Coordinate ", c, " is not in a coordinate system of size ", size)
	}

	result := make([]coordinate, 0, 4)

	// top
	if c.y > 0 {
		result = append(result, coordinate{x: c.x, y: c.y - 1})
	}

	// right
	if c.x < size.x-1 {
		result = append(result, coordinate{x: c.x + 1, y: c.y})
	}

	// bottom
	if c.y < size.y-1 {
		result = append(result, coordinate{x: c.x, y: c.y + 1})
	}

	// left
	if c.x > 0 {
		result = append(result, coordinate{x: c.x - 1, y: c.y})
	}

	return result
}
