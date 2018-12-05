package solutions

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func Day3Part1() int {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc/"+"day3.txt"))
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	lines := strings.Split(str, "\n")

	var cloth [1001][1001]int
	for _, line := range lines {
		re := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
		matches := re.FindStringSubmatch(line)
		claimID, _ := strconv.Atoi(matches[1])
		leftEdge, _ := strconv.Atoi(matches[2])
		topEdge, _ := strconv.Atoi(matches[3])
		width, _ := strconv.Atoi(matches[4])
		height, _ := strconv.Atoi(matches[5])
		for y := topEdge; y < topEdge+height; y++ {
			for x := leftEdge; x < leftEdge+width; x++ {
				if cloth[x][y] == 0 {
					cloth[x][y] = claimID
				} else {
					cloth[x][y] = -1
				}
			}
		}
	}
	count := 0
	for y := 0; y <= 1000; y++ {
		for x := 0; x <= 1000; x++ {
			if cloth[x][y] == -1 {
				count++
			}
		}
	}
	return count
}

func Day3Part2() int {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc/"+"day3.txt"))
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	lines := strings.Split(str, "\n")

	var cloth [1001][1001]int
	for _, line := range lines {
		re := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
		matches := re.FindStringSubmatch(line)
		claimID, _ := strconv.Atoi(matches[1])
		leftEdge, _ := strconv.Atoi(matches[2])
		topEdge, _ := strconv.Atoi(matches[3])
		width, _ := strconv.Atoi(matches[4])
		height, _ := strconv.Atoi(matches[5])
		for y := topEdge; y < topEdge+height; y++ {
			for x := leftEdge; x < leftEdge+width; x++ {
				if cloth[x][y] == 0 {
					cloth[x][y] = claimID
				} else {
					cloth[x][y] = -1
				}
			}
		}
	}
	for _, line := range lines {
		re := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
		matches := re.FindStringSubmatch(line)
		claimID, _ := strconv.Atoi(matches[1])
		leftEdge, _ := strconv.Atoi(matches[2])
		topEdge, _ := strconv.Atoi(matches[3])
		width, _ := strconv.Atoi(matches[4])
		height, _ := strconv.Atoi(matches[5])
		doesOverlap := false
		for y := topEdge; y < topEdge+height; y++ {
			for x := leftEdge; x < leftEdge+width; x++ {
				if cloth[x][y] != claimID {
					doesOverlap = true
				}
			}
		}
		if !doesOverlap {
			return claimID
		}
	}
	return 0
}
