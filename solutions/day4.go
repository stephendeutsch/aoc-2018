package solutions

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// TimelineEntity ...
type TimelineEntity struct {
	TimeString  string
	Minutes     int
	Action      string
	GuardNumber string
}

// Day4Part1 ...
func Day4Part1() int {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc-2018/"+"day4.txt"))
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	var timeline []*TimelineEntity
	for _, line := range strings.Split(str, "\n") {
		timelineEntity := &TimelineEntity{}
		re := regexp.MustCompile(`\[(\d+-\d+-\d+ \d+:(\d+))] (.+)`)
		guardRe := regexp.MustCompile(`Guard #(\d+) begins shift`)
		matches := re.FindStringSubmatch(line)
		text := matches[3]
		if strings.Contains(text, "begins shift") {
			guardMatch := guardRe.FindStringSubmatch(text)
			timelineEntity.Action = "begins shift"
			timelineEntity.GuardNumber = guardMatch[1]
		} else {
			timelineEntity.Action = text
		}
		timelineEntity.TimeString = matches[1]
		timelineEntity.Minutes, _ = strconv.Atoi(matches[2])
		timeline = append(timeline, timelineEntity)
	}
	sort.Slice(timeline, func(i, j int) bool { return timeline[i].TimeString < timeline[j].TimeString })
	guardSleep := make(map[string][60]int)
	var currentGuardNum string
	for i, entity := range timeline {
		if len(entity.GuardNumber) != 0 {
			currentGuardNum = entity.GuardNumber
		}
		entity.GuardNumber = currentGuardNum
		if entity.Action == "wakes up" {
			sleepTime := timeline[i-1].Minutes
			wakeTime := entity.Minutes
			sleepMinutes := guardSleep[entity.GuardNumber]
			for m := sleepTime; m < wakeTime; m++ {
				sleepMinutes[m]++
			}
			guardSleep[entity.GuardNumber] = sleepMinutes
		}
	}
	maxSleep := 0
	maxGuard := ""
	maxMinute := -1
	for k, v := range guardSleep {
		guardSum := 0
		localMax := 0
		localMaxMinute := -1
		for i, m := range v {
			guardSum += m
			if m > localMax {
				localMax = m
				localMaxMinute = i
			}
		}
		if guardSum > maxSleep {
			maxGuard = k
			maxSleep = guardSum
			maxMinute = localMaxMinute
		}
	}
	maxGuardInt, _ := strconv.Atoi(maxGuard)
	return maxGuardInt * maxMinute
}

// Day4Part2 ...
func Day4Part2() int {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc-2018/"+"day4.txt"))
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	var timeline []*TimelineEntity
	for _, line := range strings.Split(str, "\n") {
		timelineEntity := &TimelineEntity{}
		re := regexp.MustCompile(`\[(\d+-\d+-\d+ \d+:(\d+))] (.+)`)
		guardRe := regexp.MustCompile(`Guard #(\d+) begins shift`)
		matches := re.FindStringSubmatch(line)
		text := matches[3]
		if strings.Contains(text, "begins shift") {
			guardMatch := guardRe.FindStringSubmatch(text)
			timelineEntity.Action = "begins shift"
			timelineEntity.GuardNumber = guardMatch[1]
		} else {
			timelineEntity.Action = text
		}
		timelineEntity.TimeString = matches[1]
		timelineEntity.Minutes, _ = strconv.Atoi(matches[2])
		timeline = append(timeline, timelineEntity)
	}
	sort.Slice(timeline, func(i, j int) bool { return timeline[i].TimeString < timeline[j].TimeString })
	guardSleep := make(map[string][60]int)
	var currentGuardNum string
	for i, entity := range timeline {
		if len(entity.GuardNumber) != 0 {
			currentGuardNum = entity.GuardNumber
		}
		entity.GuardNumber = currentGuardNum
		if entity.Action == "wakes up" {
			sleepTime := timeline[i-1].Minutes
			wakeTime := entity.Minutes
			sleepMinutes := guardSleep[entity.GuardNumber]
			for m := sleepTime; m < wakeTime; m++ {
				sleepMinutes[m]++
			}
			guardSleep[entity.GuardNumber] = sleepMinutes
		}
	}
	maxSleep := 0
	maxGuard := ""
	maxMinute := -1
	for k, v := range guardSleep {
		localMax := 0
		localMaxMinute := -1
		for i, m := range v {
			if m > localMax {
				localMax = m
				localMaxMinute = i
			}
		}
		if localMax > maxSleep {
			maxGuard = k
			maxSleep = localMax
			maxMinute = localMaxMinute
		}
	}
	maxGuardInt, _ := strconv.Atoi(maxGuard)
	return maxGuardInt * maxMinute
}
