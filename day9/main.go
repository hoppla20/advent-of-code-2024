package main

import (
	"aoc2024/internal/utils"
	"fmt"
	"os"
)

func sumDigits(input *string) int {
	result := 0

	for i := 0; i < len(*input); i++ {
		result += int((*input)[i]) - 48
	}

	return result
}

func parseDenseDiskMap(input *string) []int {
	result := make([]int, sumDigits(input))

	curPos := 0
	curIsFile := true
	curFileId := 0

	for i := 0; i < len(*input); i++ {
		curDigit := int((*input)[i] - 48)

		for j := 0; j < curDigit; j++ {
			if curIsFile {
				result[curPos+j] = curFileId
			} else {
				result[curPos+j] = -1
			}
		}

		curPos += curDigit
		curIsFile = !curIsFile
		if curIsFile {
			curFileId++
		}
	}

	return result
}

func compactDiskMap1(diskMap []int) {
	curLastPos := len(diskMap) - 1

	for i := 0; i < len(diskMap) && i < curLastPos; i++ {
		if diskMap[i] < 0 {
			curLastDigit := diskMap[curLastPos]
			diskMap[curLastPos] = diskMap[i]
			diskMap[i] = curLastDigit

			for diskMap[curLastPos] < 0 {
				curLastPos--
			}
		}
	}
}

func compactDiskMap2(diskMap []int) {
	for i := len(diskMap) - 1; i >= 0; i-- {
		if diskMap[i] >= 0 {
			// calculate file size
			fileId := diskMap[i]
			fileEnd := i
			for ; i > 0; i-- {
				if diskMap[i-1] != fileId {
					break
				}
			}
			fileSize := fileEnd - i + 1

			// find position to move file to
			for j := 0; j < i; j++ {
				if diskMap[j] < 0 {
					// calculate free space size
					spaceStart := j
					for ; j < i-1; j++ {
						if diskMap[j+1] >= 0 {
							break
						}
					}
					spaceSize := j - spaceStart + 1

					if spaceSize >= fileSize {
						// copy file to beginning of free space
						for k := spaceStart; k < spaceStart+fileSize; k++ {
							diskMap[k] = fileId
						}

						// delete file at original position
						for k := fileEnd; k > fileEnd-fileSize; k-- {
							diskMap[k] = -1
						}

						break
					}
				}
			}
		}
	}
}

func checksumDiskMap(diskMap []int) int {
	result := 0

	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] >= 0 {
			result += i * diskMap[i]
		}
	}

	return result
}

func part1(input *string) int {
	diskMap := parseDenseDiskMap(input)

	// fmt.Println(diskMap)

	compactDiskMap1(diskMap)

	// fmt.Println(diskMap)

	return checksumDiskMap(diskMap)
}

func part2(input *string) int {
	diskMap := parseDenseDiskMap(input)

	// fmt.Println(diskMap)

	compactDiskMap2(diskMap)

	// fmt.Println(diskMap)

	return checksumDiskMap(diskMap)
}

func main() {
	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	} else {
		fileName = "example.txt"
	}
	input := utils.ReadFile(fileName)[0]

	// fmt.Println("Part 1:", part1(&input))
	fmt.Println("Part 2:", part2(&input))
}
