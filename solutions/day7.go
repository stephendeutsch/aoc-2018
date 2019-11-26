package solutions

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

type op struct {
	Char    byte
	ReqOps  map[byte]*op
	NextOps map[byte]*op
	Time    int
	Working bool
}

// Day7Part1 ...
func Day7Part1() string {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc/"+"day7.txt"))
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	lines := strings.Split(str, "\n")

	var opsMap = make(map[byte]*op)
	var seq = make([]byte, 0)
	for _, line := range lines {
		reqChar := line[5]
		opChar := line[36]
		if _, ok := opsMap[reqChar]; !ok {
			opsMap[reqChar] = &op{
				Char:    reqChar,
				ReqOps:  make(map[byte]*op),
				NextOps: make(map[byte]*op),
			}
		}
		if _, ok := opsMap[opChar]; !ok {
			opsMap[opChar] = &op{
				Char:    opChar,
				ReqOps:  make(map[byte]*op),
				NextOps: make(map[byte]*op),
			}
		}
		opsMap[opChar].ReqOps[reqChar] = opsMap[reqChar]
		opsMap[reqChar].NextOps[opChar] = opsMap[opChar]
	}
	for len(opsMap) > 0 {
		var poss = make([]byte, 0)
		for c, op := range opsMap {
			if len(op.ReqOps) == 0 {
				poss = append(poss, c)
			}
		}
		sort.Slice(poss, func(i int, j int) bool { return poss[i] < poss[j] })
		nextOp := poss[0]
		seq = append(seq, nextOp)
		for _, op := range opsMap[nextOp].NextOps {
			delete(op.ReqOps, nextOp)
		}
		delete(opsMap, nextOp)
	}
	return string(seq)
}

// Day7Part2 ...
func Day7Part2() string {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/aoc/"+"day7.txt"))
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	lines := strings.Split(str, "\n")

	var opsMap = make(map[byte]*op)
	var seq = make([]byte, 0)
	for _, line := range lines {
		reqChar := line[5]
		opChar := line[36]
		if _, ok := opsMap[reqChar]; !ok {
			opsMap[reqChar] = &op{
				Char:    reqChar,
				ReqOps:  make(map[byte]*op),
				NextOps: make(map[byte]*op),
				Time:    int(reqChar) - 64 + 60,
			}
		}
		if _, ok := opsMap[opChar]; !ok {
			opsMap[opChar] = &op{
				Char:    opChar,
				ReqOps:  make(map[byte]*op),
				NextOps: make(map[byte]*op),
				Time:    int(opChar) - 64 + 60,
			}
		}
		opsMap[opChar].ReqOps[reqChar] = opsMap[reqChar]
		opsMap[reqChar].NextOps[opChar] = opsMap[opChar]
	}
	currentSecond := 0
	fmt.Println("Time\tW1\tW2\tW3\tW4\tW5\tSeq")
	var workers = make([]*op, 5)
	for len(opsMap) > 0 {
		currentSecond++
		fmt.Printf("%d", currentSecond)
		var poss = make([]byte, 0)
		for c, op := range opsMap {
			if len(op.ReqOps) == 0 && !op.Working {
				poss = append(poss, c)
			}
		}
		sort.Slice(poss, func(i int, j int) bool { return poss[i] < poss[j] })
		for i, worker := range workers {
			if len(poss) == 0 {
				break
			}
			if worker == nil {
				workers[i] = opsMap[poss[0]]
				workers[i].Working = true
				poss = poss[1:]
			}
		}
		for i, worker := range workers {
			if worker != nil {
				fmt.Printf("\t%s", string(worker.Char))
				worker.Time--
				if worker.Time == 0 {
					seq = append(seq, worker.Char)
					for _, op := range worker.NextOps {
						delete(op.ReqOps, worker.Char)
					}
					delete(opsMap, worker.Char)
					workers[i] = nil
				}
			} else {
				fmt.Printf("\t.")
			}
		}
		fmt.Printf("\t%s\n", string(seq))
	}
	return strconv.Itoa(currentSecond)
}
