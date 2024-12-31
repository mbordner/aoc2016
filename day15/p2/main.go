package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common/file"
	"regexp"
	"strconv"
)

var (
	reDisc = regexp.MustCompile(`Disc #(\d+)\s+has\s+(\d+)\s+positions; at time=0, it is at position\s+(\d+)`)
)

type Disc struct {
	num             int
	positions       int
	initialPosition int
}

func main() {
	discs := getData("../data.txt")

	times := []int{}

	t := 0
	for {
		i := 0
		for i < len(discs) {
			d := discs[i]
			if (d.initialPosition+i+1+t)%d.positions != 0 {
				break
			}
			i++
		}
		if i == len(discs) {
			times = append(times, t)
			if len(times) == 3 {
				break
			}
		}
		t++
	}

	t = times[0]

	fmt.Printf("Part 1: %d\n", t)

	tr := times[2] - times[1]
	ip := 0
	ps := 11
	dt := len(discs) + 1
	i := 0
	for {
		if (ip+dt+t+(i*tr))%ps == 0 {
			break
		}
		i++
	}

	fmt.Printf("Part 2: %d\n", t+(i*tr))
}

func getIntVal(s string) int {
	val, _ := strconv.ParseInt(s, 10, 64)
	return int(val)
}

func getData(f string) []Disc {
	lines, _ := file.GetLines(f)
	discs := make([]Disc, len(lines))
	for i, line := range lines {
		matches := reDisc.FindStringSubmatch(line)
		discs[i] = Disc{num: getIntVal(matches[1]), positions: getIntVal(matches[2]), initialPosition: getIntVal(matches[3])}
	}
	return discs
}
