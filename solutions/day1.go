package solutions

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

func Day1Part1() int {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc/"+"day1.txt"))
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	count := 0
	for _, line := range strings.Split(str, "\n") {
		if len(line) > 0 {
			lineNum, _ := strconv.Atoi(string(line))
			count += lineNum
		}
	}
	return count
}

func Day1Part2() (int, error) {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc/"+"day1.txt"))
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)
	count := 0
	seen := make(map[int]bool)
	for {
		for _, line := range strings.Split(str, "\n") {
			if len(line) > 0 {
				lineNum, _ := strconv.Atoi(string(line))
				count += lineNum
				if seen[count] {
					return count, nil
				}
				seen[count] = true
			}
		}
	}
	//return 0, errors.New("duplicate not found")
}
