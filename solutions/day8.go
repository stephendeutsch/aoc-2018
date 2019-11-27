package solutions

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type el struct {
	Index    int
	Children []*el
	Metadata []int
}

var values []int

// Day8Part1 ...
func Day8Part1() string {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc-2018/"+"day8.txt"))
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	strValues := strings.Split(str, " ")
	values = make([]int, len(strValues))
	for i, val := range strValues {
		values[i], _ = strconv.Atoi(val)
	}
	root, _ := parseEl(0)
	sum := getSumMetadata(root)
	return strconv.Itoa(sum)
}

// Day8Part2 ...
func Day8Part2() string {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc-2018/"+"day8.txt"))
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	strValues := strings.Split(str, " ")
	values = make([]int, len(strValues))
	for i, val := range strValues {
		values[i], _ = strconv.Atoi(val)
	}
	root, _ := parseEl(0)
	sum := calculateValues(root)
	return strconv.Itoa(sum)
}

func parseEl(index int) (*el, int) {
	thisEl := &el{
		Index:    index,
		Children: make([]*el, 0),
		Metadata: make([]int, 0),
	}
	numMetadata := values[index+1]
	lastIndex := index + 2
	for i := 0; i < values[index]; i++ {
		var childEl *el
		childEl, lastIndex = parseEl(lastIndex)
		thisEl.Children = append(thisEl.Children, childEl)
	}
	thisEl.Metadata = values[lastIndex : lastIndex+numMetadata]
	return thisEl, lastIndex + numMetadata
}

func getSumMetadata(thisEl *el) int {
	sum := 0
	for _, num := range thisEl.Metadata {
		sum += num
	}
	for _, child := range thisEl.Children {
		sum += getSumMetadata(child)
	}
	return sum
}

func calculateValues(thisEl *el) int {
	val := 0
	if len(thisEl.Children) == 0 {
		for _, num := range thisEl.Metadata {
			val += num
		}
		return val
	}
	for _, num := range thisEl.Metadata {
		i := num - 1
		if len(thisEl.Children) > i {
			val += calculateValues(thisEl.Children[i])
		}
	}
	return val
}
