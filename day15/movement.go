package main

import (
	"aoc2024/internal/utils"
	"fmt"
	"log"
)

type UnknownDirectionError struct {
	value int
}

func (e UnknownDirectionError) Error() string {
	return fmt.Sprintf("Unknown direction: %d", e.value)
}

func getNextPosition(curPos utils.Coordinate, dir direction) (nextPos utils.Coordinate, err error) {
	switch dir {
	case up:
		nextPos, err = curPos.Up()
	case right:
		nextPos, err = curPos.Right()
	case down:
		nextPos, err = curPos.Down()
	case left:
		nextPos, err = curPos.Left()
	default:
		err = &UnknownDirectionError{int(dir)}
	}

	return
}

func (w *warehouse) isBoxMovable(curPosition utils.Coordinate, dir direction, second bool) bool {
	curHolder := w.spaceHolder(curPosition)

	var secondBoxDir direction
	switch curHolder {
	case wall:
		return false
	case empty:
		return true
	case box_left:
		secondBoxDir = right
	case box_right:
		secondBoxDir = left
	}

	nextPosition, err := getNextPosition(curPosition, dir)
	if err != nil {
		log.Fatal(err)
	}

	switch dir {
	case left, right:
		return w.isBoxMovable(nextPosition, dir, false)
	case up, down:
		isCurMovable := w.isBoxMovable(nextPosition, dir, false)
		isOtherMovable := true

		if !second {
			secondBoxPosition, err := getNextPosition(curPosition, secondBoxDir)
			if err != nil {
				log.Fatal(err)
			}
			isOtherMovable = w.isBoxMovable(secondBoxPosition, dir, true)
		}

		return isCurMovable && isOtherMovable
	}

	return false
}

func (w *warehouse) moveBoxes(curPosition utils.Coordinate, dir direction, second bool) {
	curHolder := w.spaceHolder(curPosition)

	var secondBoxDir direction
	switch curHolder {
	case wall:
		log.Fatal("Box at position", curPosition, "not movable")
	case empty:
		return
	case box_left:
		secondBoxDir = right
	case box_right:
		secondBoxDir = left
	}

	nextPosition, err := getNextPosition(curPosition, dir)
	if err != nil {
		log.Fatal(err)
	}

	w.moveBoxes(nextPosition, dir, false)
	w.setSpaceHolder(curPosition, w.spaceHolder(nextPosition))
	w.setSpaceHolder(nextPosition, curHolder)

	if (dir == up || dir == down) && !second {
		secondBoxPosition, err := getNextPosition(curPosition, secondBoxDir)
		if err != nil {
			log.Fatal(err)
		}
		w.moveBoxes(secondBoxPosition, dir, true)
	}
}

func (w *warehouse) moveRobot(dir direction) bool {
	log.Println("Try to move robot in direction ", string(moveRuneLookup[dir]))

	var nextPosition utils.Coordinate
	var error error

	switch dir {
	case up:
		nextPosition, error = w.robotPosition.Up()
	case right:
		nextPosition, error = w.robotPosition.Right()
	case down:
		nextPosition, error = w.robotPosition.Down()
	case left:
		nextPosition, error = w.robotPosition.Left()
	default:
		log.Fatal("Unknown move:", dir)
	}

	if error != nil {
		log.Fatal(error)
	}

	switch h := w.spaceHolder(nextPosition); h {
	case wall:
		return false
	case empty:
		// nop
	case box_left, box_right:
		if !w.isBoxMovable(nextPosition, dir, false) {
			return false
		}

		w.moveBoxes(nextPosition, dir, false)
	default:
		log.Fatal("Unknown space holder:", h)
	}

	w.robotPosition = nextPosition

	return true
}
