package main

import (
	"aoc2024/internal/utils"
	"log"
)

func (w *warehouse) moveBoxes(boxPosition utils.Coordinate, d direction) bool {
	switch h := w.spaceHolder(boxPosition); h {
	case empty:
		return true
	case wall:
		return false
	case box:
		// nop
	default:
		log.Fatal("Unknown space holder", h, "while moving boxes")
	}

	var nextPosition utils.Coordinate
	var error error

	switch d {
	case up:
		nextPosition, error = boxPosition.Up()
	case right:
		nextPosition, error = boxPosition.Right()
	case down:
		nextPosition, error = boxPosition.Down()
	case left:
		nextPosition, error = boxPosition.Left()
	default:
		log.Fatal("Unknown move:", d)
	}

	if error != nil {
		log.Fatal(error)
	}

	if !w.moveBoxes(nextPosition, d) {
		return false
	}

	log.Println("Moving box at position", boxPosition, string(moveRuneLookup[d]))
	w.setSpaceHolder(boxPosition, empty)
	w.setSpaceHolder(nextPosition, box)

	return true
}

func (w *warehouse) moveRobot(d direction) bool {
	log.Println("Try to move robot in direction ", string(moveRuneLookup[d]))

	var nextPosition utils.Coordinate
	var error error

	switch d {
	case up:
		nextPosition, error = w.robotPosition.Up()
	case right:
		nextPosition, error = w.robotPosition.Right()
	case down:
		nextPosition, error = w.robotPosition.Down()
	case left:
		nextPosition, error = w.robotPosition.Left()
	default:
		log.Fatal("Unknown move:", d)
	}

	if error != nil {
		log.Fatal(error)
	}

	switch h := w.spaceHolder(nextPosition); h {
	case wall:
		return false
	case empty:
		// nop
	case box:
		if !w.moveBoxes(nextPosition, d) {
			return false
		}
	default:
		log.Fatal("Unknown space holder:", h)
	}

	w.robotPosition = nextPosition

	return true
}
