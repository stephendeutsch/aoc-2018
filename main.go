package main

import (
	"aoc/solutions"
	"flag"
	"fmt"
)

func main() {
	dayPtr := flag.String("day", "", "Choose the day")
	flag.Parse()
	switch *dayPtr {
	case "1":
		fmt.Println("Day 1 Part 1")
		fmt.Println(solutions.Day1Part1())
		fmt.Println("Day 1 Part 2")
		result, err := solutions.Day1Part2()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	case "2":
		fmt.Println("Day 2 Part 1")
		fmt.Println(solutions.Day2Part1())
		fmt.Println("Day 2 Part 2")
		fmt.Println(solutions.Day2Part2())
	case "3":
		fmt.Println("Day 3 Part 1")
		fmt.Println(solutions.Day3Part1())
		fmt.Println("Day 3 Part 2")
		fmt.Println(solutions.Day3Part2())
	case "4":
		fmt.Println("Day 4 Part 1")
		fmt.Println(solutions.Day4Part1())
		fmt.Println("Day 4 Part 2")
		fmt.Println(solutions.Day4Part2())
	case "5":
		fmt.Println("Day 5 Part 1")
		fmt.Println(solutions.Day5Part1())
		fmt.Println("Day 5 Part 2")
		fmt.Println(solutions.Day5Part2())
	case "6":
		fmt.Println("Day 6 Part 1")
		fmt.Println(solutions.Day6Part1())
		fmt.Println("Day 6 Part 2")
		fmt.Println(solutions.Day6Part2())
	case "7":
		//fmt.Println("Day 7 Part 1")
		//fmt.Println(solutions.Day7Part1())
		//fmt.Println("Day 7 Part 2")
		//fmt.Println(solutions.Day7Part2())
	case "8":
		//fmt.Println("Day 8 Part 1")
		//fmt.Println(solutions.Day8Part1())
		//fmt.Println("Day 8 Part 2")
		//fmt.Println(solutions.Day8Part2())
	case "9":
		//fmt.Println("Day 9 Part 1")
		//fmt.Println(solutions.Day9Part1())
		//fmt.Println("Day 9 Part 2")
		//fmt.Println(solutions.Day9Part2())
	case "10":
		//fmt.Println("Day 10 Part 1")
		//fmt.Println(solutions.Day10Part1())
		//fmt.Println("Day 10 Part 2")
		//fmt.Println(solutions.Day10Part2())
	case "11":
		//fmt.Println("Day 11 Part 1")
		//fmt.Println(solutions.Day11Part1())
		//fmt.Println("Day 11 Part 2")
		//fmt.Println(solutions.Day11Part2())
	case "12":
		//fmt.Println("Day 12 Part 1")
		//fmt.Println(solutions.Day12Part1())
		//fmt.Println("Day 12 Part 2")
		//fmt.Println(solutions.Day12Part2())
	case "13":
		//fmt.Println("Day 13 Part 1")
		//fmt.Println(solutions.Day13Part1())
		//fmt.Println("Day 13 Part 2")
		//fmt.Println(solutions.Day13Part2())
	case "14":
		//fmt.Println("Day 14 Part 1")
		//fmt.Println(solutions.Day14Part1())
		//fmt.Println("Day 14 Part 2")
		//fmt.Println(solutions.Day14Part2())
	case "15":
		//fmt.Println("Day 15 Part 1")
		//fmt.Println(solutions.Day15Part1())
		//fmt.Println("Day 15 Part 2")
		//fmt.Println(solutions.Day15Part2())
	case "16":
		//fmt.Println("Day 16 Part 1")
		//fmt.Println(solutions.Day16Part1())
		//fmt.Println("Day 16 Part 2")
		//fmt.Println(solutions.Day16Part2())
	case "17":
		//fmt.Println("Day 17 Part 1")
		//fmt.Println(solutions.Day17Part1())
		//fmt.Println("Day 17 Part 2")
		//fmt.Println(solutions.Day17Part2())
	case "18":
		//fmt.Println("Day 18 Part 1")
		//fmt.Println(solutions.Day18Part1())
		//fmt.Println("Day 18 Part 2")
		//fmt.Println(solutions.Day18Part2())
	case "19":
		//fmt.Println("Day 19 Part 1")
		//fmt.Println(solutions.Day19Part1())
		//fmt.Println("Day 19 Part 2")
		//fmt.Println(solutions.Day19Part2())
	case "20":
		//fmt.Println("Day 20 Part 1")
		//fmt.Println(solutions.Day20Part1())
		//fmt.Println("Day 20 Part 2")
		//fmt.Println(solutions.Day20Part2())
	case "21":
		//fmt.Println("Day 21 Part 1")
		//fmt.Println(solutions.Day21Part1())
		//fmt.Println("Day 21 Part 2")
		//fmt.Println(solutions.Day21Part2())
	case "22":
		//fmt.Println("Day 22 Part 1")
		//fmt.Println(solutions.Day22Part1())
		//fmt.Println("Day 22 Part 2")
		//fmt.Println(solutions.Day22Part2())
	case "23":
		//fmt.Println("Day 23 Part 1")
		//fmt.Println(solutions.Day23Part1())
		//fmt.Println("Day 23 Part 2")
		//fmt.Println(solutions.Day23Part2())
	case "24":
		//fmt.Println("Day 24 Part 1")
		//fmt.Println(solutions.Day24Part1())
		//fmt.Println("Day 24 Part 2")
		//fmt.Println(solutions.Day24Part2())
	case "25":
		//fmt.Println("Day 25 Part 1")
		//fmt.Println(solutions.Day25Part1())
		//fmt.Println("Day 25 Part 2")
		//fmt.Println(solutions.Day25Part2())
	default:
		fmt.Println("Usage: --day <# of day>")
	}
}
