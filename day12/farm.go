package main

import (
	"aoc2024/internal/utils"
	"fmt"
	"log"
	"strings"
)

type plot struct {
	plantType rune
	regionId  uint
}

type idGenerator func() uint

type farm struct {
	regionIdGenerator idGenerator
	plots             [][]plot
	assignedRegions   []uint
	size              coordinateSystemSize
}

func newContinuousIdGenerator() func() uint {
	var nextId uint = 1
	return func() uint {
		result := nextId
		nextId++
		return result
	}
}

func loadFarmFromFile(path string) *farm {
	f := new(farm)
	f.regionIdGenerator = newContinuousIdGenerator()
	f.assignedRegions = make([]uint, 0)

	rows := utils.ReadFile(path)

	f.size = coordinateSystemSize{x: uint(len(rows[1])), y: uint(len(rows))}

	f.plots = make([][]plot, f.size.y)
	for i := range f.plots {
		f.plots[i] = make([]plot, f.size.x)
	}

	for y, row := range rows {
		for x, plantType := range row {
			f.plots[y][x] = plot{
				plantType: plantType,
				regionId:  0,
			}
		}
	}

	return f
}

func (f *farm) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "Farm dimensions: %dx%d\n", f.size.x, f.size.y)

	for i, row := range f.plots {
		for _, plot := range row {
			sb.WriteRune(plot.plantType)
		}

		if i != len(f.plots)-1 {
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}

func (f *farm) printRegionIds() {
	var sb strings.Builder

	fmt.Fprintln(&sb, "Farm regions:")

	for i, row := range f.plots {
		for _, plot := range row {
			sb.WriteRune(rune(plot.regionId + 48))
		}

		if i != len(f.plots)-1 {
			sb.WriteRune('\n')
		}
	}

	log.Println(sb.String())
}

func (f *farm) plot(c coordinate) *plot {
	return &f.plots[c.y][c.x]
}

func (f *farm) changeRegionId(posInRegion coordinate, id uint) {
	plotAtPos := f.plot(posInRegion)
	if plotAtPos.regionId != id {
		oldRegionId := plotAtPos.regionId
		plotAtPos.regionId = id
		for _, neighbor := range posInRegion.neighbors(f.size) {
			if f.plot(neighbor).regionId == oldRegionId {
				f.changeRegionId(neighbor, id)
			}
		}
	}
}

func (f *farm) assignRegionNumbers() {
	for y := range f.size.y {
		for x := range f.size.x {
			curPos := coordinate{x: x, y: y}
			curPlot := f.plot(curPos)

			foundNeighboringRegion := false
			for _, neighborPos := range curPos.neighbors(f.size) {
				neighborPlot := f.plot(neighborPos)

				// only neighbors that have already been assigned a region
				if neighborPlot.regionId != 0 {
					if !foundNeighboringRegion {
						// take first neighboring region id with the same plant type
						if curPlot.plantType == neighborPlot.plantType {
							curPlot.regionId = neighborPlot.regionId
							foundNeighboringRegion = true
						}
					} else {
						// merge two different regions if they have the same plant type
						if curPlot.plantType == neighborPlot.plantType && curPlot.regionId != neighborPlot.regionId {
							f.changeRegionId(neighborPos, curPlot.regionId)
						}
					}
				}
			}

			if !foundNeighboringRegion {
				curPlot.regionId = f.regionIdGenerator()
				f.assignedRegions = append(f.assignedRegions, curPlot.regionId)
			}
		}
	}
}
