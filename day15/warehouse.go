package main

import (
	"aoc2024/internal/utils"
	"log"
	"strings"
)

const (
	empty_rune = '.'
	wall_rune  = '#'
	robot_rune = '@'
	box_rune   = 'O'
)

type spaceHolder int

const (
	empty = iota
	wall
	robot
	box
)

var holderRuneLookup = []rune{
	empty_rune,
	wall_rune,
	robot_rune,
	box_rune,
}

type direction int

const (
	up_rune    = '^'
	right_rune = '>'
	down_rune  = 'v'
	left_rune  = '<'
)

const (
	up = iota
	right
	down
	left
)

var moveRuneLookup = []rune{
	up_rune,
	right_rune,
	down_rune,
	left_rune,
}

type warehouse struct {
	spaces        [][]spaceHolder
	width         int
	height        int
	robotPosition utils.Coordinate
}

func loadInputFile(filePath string) (*warehouse, []direction) {
	lines := utils.ReadFile(filePath)

	// parse warehouse

	w := new(warehouse)

	var warehouseLines []string
	for lines[0] != "" {
		warehouseLines = append(warehouseLines, lines[0])
		lines = lines[1:]
	}
	lines = lines[1:]

	log.Printf("Warehouse Lines:\n%s\n", strings.Join(warehouseLines, "\n"))

	w.height = len(warehouseLines)
	w.width = len(warehouseLines[0])

	utils.SetCoordinateSystemSize(utils.CoordinateSystemSize{
		Y: w.height,
		X: w.width,
	})

	w.spaces = make([][]spaceHolder, w.height)
	for i := range w.height {
		w.spaces[i] = make([]spaceHolder, w.width)
	}

	for y := range w.height {
		for x := range w.width {
			switch r := warehouseLines[y][x]; r {
			case wall_rune:
				w.spaces[y][x] = wall
			case robot_rune:
				w.robotPosition = utils.Coordinate{X: x, Y: y}
			case box_rune:
				w.spaces[y][x] = box
			case empty_rune:
				w.spaces[y][x] = empty
			default:
				log.Fatal("Unkown space holder rune:", r)
			}
		}
	}

	log.Printf("Robot position: %v\n", w.robotPosition)
	log.Printf("Parsed Warehouse:\n%v\n", w)

	// parse moves

	moves := make([]direction, 0)

	log.Printf("Moves lines:\n%s\n", strings.Join(lines, "\n"))

	for _, l := range lines {
		for i := range len(l) {
			switch r := l[i]; r {
			case up_rune:
				moves = append(moves, up)
			case right_rune:
				moves = append(moves, right)
			case down_rune:
				moves = append(moves, down)
			case left_rune:
				moves = append(moves, left)
			default:
				log.Fatal("Unknown move rune:", r)
			}
		}
	}

	log.Println("Parsed Moves:")
	log.Println(moves)

	return w, moves
}

func (w *warehouse) String() string {
	var sb strings.Builder

	for y := range w.height {
		for x := range w.width {
			if x == w.robotPosition.X && y == w.robotPosition.Y {
				sb.WriteRune('@')
			} else {
				sb.WriteRune(holderRuneLookup[w.spaces[y][x]])
			}
		}
		if y != w.height-1 {
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}

func (w *warehouse) spaceHolder(position utils.Coordinate) spaceHolder {
	return w.spaces[position.Y][position.X]
}

func (w *warehouse) setSpaceHolder(position utils.Coordinate, holder spaceHolder) {
	w.spaces[position.Y][position.X] = holder
}
