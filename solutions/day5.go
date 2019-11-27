package solutions

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

// Day5Part1 ...
func Day5Part1() int {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc-2018/"+"day5.txt"))
	if err != nil {
		fmt.Print(err)
	}
	// Golang regex flavor does not support backreferences, so just generate a regex
	// with an OR between groups of all uppercase/lowercase letter combinations
	var reArray []string
	for _, rune := range "abcdefghijklmnopqrstuvwxyz" {
		letter := string(rune)
		upperLetter := strings.ToUpper(letter)
		reArray = append(reArray, letter+upperLetter+"|"+upperLetter+letter)
	}
	reString := strings.Join(reArray, "|")
	re := regexp.MustCompile(reString)
	// We will only ever match one letter before the previous match, so advance a counter
	// instead of searching through the entire string every single time
	b = append([]byte(" "), b...)
	previousMatchIndex := 1
	for re.Match(b[previousMatchIndex-1:]) {
		matchIndex := re.FindIndex(b[previousMatchIndex-1:])
		b = append(b[0:matchIndex[0]+previousMatchIndex-1], b[matchIndex[1]+previousMatchIndex-1:]...)
		previousMatchIndex += matchIndex[0] - 1
	}
	return len(b) - 1 // -1 because of space added at beginning to prevent out of range
}

// Day5Part2 ...
func Day5Part2() int {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc-2018/"+"day5.txt"))
	if err != nil {
		fmt.Print(err)
	}
	var bArray [][]byte
	var reArray []string
	for _, rune := range "abcdefghijklmnopqrstuvwxyz" {
		letter := string(rune)
		upperLetter := strings.ToUpper(letter)
		reArray = append(reArray, letter+upperLetter+"|"+upperLetter+letter)
		bArray = append(bArray, bytes.Replace(bytes.Replace(b, []byte(letter), []byte(""), -1), []byte(upperLetter), []byte(""), -1))
	}
	reString := strings.Join(reArray, "|")
	re := regexp.MustCompile(reString)
	smallest := len(b)
	for _, bs := range bArray {
		bs = append([]byte(" "), bs...)
		previousMatchIndex := 1
		for re.Match(bs[previousMatchIndex-1:]) {
			matchIndex := re.FindIndex(bs[previousMatchIndex-1:])
			bs = append(bs[0:matchIndex[0]+previousMatchIndex-1], bs[matchIndex[1]+previousMatchIndex-1:]...)
			previousMatchIndex += matchIndex[0] - 1
		}
		if len(bs)-1 < smallest {
			smallest = len(bs) - 1
		}
	}
	return smallest
}
