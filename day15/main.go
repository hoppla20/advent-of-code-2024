package main

import (
	"aoc2024/internal/utils"
	"fmt"
	"io"
	"log"
	"os"
)

func gpsCoordinate(position utils.Coordinate) int {
	return position.Y*100 + position.X
}

func part1(inputPath string) {
	warehouse, moves := loadInputFile(inputPath)

	for _, m := range moves {
		warehouse.moveRobot(m)

		log.Printf("Warehouse State:\n%v\n", warehouse)
	}

	result := 0
	for y := range warehouse.height {
		for x := range warehouse.width {
			c := utils.Coordinate{X: x, Y: y}
			if warehouse.spaceHolder(c) == box {
				result += gpsCoordinate(c)
			}
		}
	}

	fmt.Println("Result:", result)
}

func main() {
	var inputPath string
	switch len(os.Args) {
	case 1:
		inputPath = "example1.txt"
	case 2:
		inputPath = os.Args[1]
	default:
		log.Fatal("FATAL : Expected none or exactly one positional argument!")
	}

	// disable logs
	log.SetOutput(io.Discard)

	part1(inputPath)
}
