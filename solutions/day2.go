package solutions

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Day2Part1 ...
func Day2Part1() int {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc/"+"day2.txt"))
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	duplicatesCount := make(map[int]int)
	for _, line := range strings.Split(str, "\n") {
		charmap := make(map[string]int)
		hasDuplicates := make(map[int]bool)
		for _, char := range strings.Split(line, "") {
			charmap[char]++
		}
		for _, charcount := range charmap {
			if charcount == 2 {
				hasDuplicates[2] = true
			}
			if charcount == 3 {
				hasDuplicates[3] = true
			}
		}
		if hasDuplicates[2] {
			duplicatesCount[2]++
		}
		if hasDuplicates[3] {
			duplicatesCount[3]++
		}
	}
	checksum := duplicatesCount[2] * duplicatesCount[3]
	return checksum
}

// Day2Part2 ...
func Day2Part2() string {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc/"+"day2.txt"))
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	lines := strings.Split(str, "\n")

	for _, line := range lines {
		for _, compLine := range lines {
			diff := 0
			diffIndex := -1
			for index, char := range line {
				if compLine[index] != byte(char) {
					diff++
					diffIndex = index
				}
			}
			if diff == 1 {
				return line[:diffIndex] + line[diffIndex+1:]
			}
		}
	}
	return "Not Found"
}
