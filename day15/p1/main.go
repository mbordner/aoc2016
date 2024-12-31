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
			break
		}
		t++
	}

	fmt.Printf("Part 1: %d\n", t)
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
