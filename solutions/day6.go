package solutions

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type cell struct {
	Loc  [2]int
	Dist int
}

// Day6Part1 ...
func Day6Part1() int {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc-2018/"+"day6.txt"))
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	lines := strings.Split(str, "\n")

	var locs [][2]int
	notInfiniteLocs := make(map[[2]int]bool)
	sizes := make(map[[2]int]int)
	for _, line := range lines {
		coordinates := strings.Split(line, ", ")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		locs = append(locs, [2]int{x, y})
		notInfiniteLocs[[2]int{x, y}] = true
	}
	const fieldSize = 375
	var field [fieldSize][fieldSize]cell
	for x, row := range field {
		for y := range row {
			var minLoc [2]int
			minDistance := 999
			isEquidistant := false
			for _, loc := range locs {
				distance := Abs(x-loc[0]) + Abs(y-loc[1])
				if distance < minDistance {
					isEquidistant = false
					minLoc = loc
					minDistance = distance
				} else if distance == minDistance {
					isEquidistant = true
				}
			}
			if !isEquidistant {
				field[x][y].Loc = minLoc
				field[x][y].Dist = minDistance
				sizes[minLoc]++
			}
		}
	}
	for i := 0; i < fieldSize; i++ {
		notInfiniteLocs[field[0][i].Loc] = false
		notInfiniteLocs[field[i][0].Loc] = false
		notInfiniteLocs[field[fieldSize-1][i].Loc] = false
		notInfiniteLocs[field[i][fieldSize-1].Loc] = false
	}
	maxSize := 0
	for k, v := range notInfiniteLocs {
		if v {
			if sizes[k] > maxSize {
				maxSize = sizes[k]
			}
		}
	}
	return maxSize
}

// Day6Part2 ...
func Day6Part2() int {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc-2018/"+"day6.txt"))
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	lines := strings.Split(str, "\n")

	var locs [][2]int
	for _, line := range lines {
		coordinates := strings.Split(line, ", ")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		locs = append(locs, [2]int{x, y})
	}
	const fieldSize = 375
	var field [fieldSize][fieldSize]cell
	lessThan10000Area := 0
	for x, row := range field {
		for y := range row {
			for _, loc := range locs {
				distance := Abs(x-loc[0]) + Abs(y-loc[1])
				field[x][y].Dist += distance
			}
			if field[x][y].Dist < 10000 {
				lessThan10000Area++
			}
		}
	}
	return lessThan10000Area
}

// Abs ...
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
