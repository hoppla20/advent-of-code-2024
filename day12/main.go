package main

import (
	"fmt"
	"log"
	"os"
)

func part1(inputPath string) {
	f := loadFarmFromFile(inputPath)
	// log.Println(f)

	f.assignRegionNumbers()

	// log.Println("Assigned regions:", f.assignedRegions)
	// f.printRegionIds()

	type regionStat struct {
		area      uint
		perimeter uint
	}
	regionStats := make(map[uint]regionStat)
	for _, id := range f.assignedRegions {
		regionStats[id] = regionStat{}
	}

	for y := range f.size.y {
		for x := range f.size.x {
			curPos := coordinate{x: x, y: y}
			curPlot := f.plot(curPos)
			curPlotRegion := regionStats[curPlot.regionId]

			curPlotRegion.area++

			if curPlotRegion.perimeter == 0 {
				curPlotRegion.perimeter = 4
			} else {
				var regionNeighbors uint = 0
				for _, neighborPos := range curPos.neighbors(f.size) {
					if neighborPos.x < curPos.x || neighborPos.y < curPos.y {
						neighborPlot := f.plot(neighborPos)
						if neighborPlot.regionId == curPlot.regionId {
							regionNeighbors++
						}
					}
				}
				curPlotRegion.perimeter += 4 - 2*regionNeighbors
			}

			regionStats[curPlot.regionId] = curPlotRegion
		}
	}

	// for regionId, regionStat := range regionStats {
	// 	log.Printf("Region %s stats: Area %d ; Perimeter %d\n", string(regionId+48), regionStat.area, regionStat.perimeter)
	// }

	var result uint = 0
	for _, regionStat := range regionStats {
		result += regionStat.area * regionStat.perimeter
	}

	fmt.Println("Result Part 1:", result)
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

	part1(inputPath)
}
