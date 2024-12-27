package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common"
	"github.com/mbordner/aoc2016/common/file"
	"regexp"
)

var (
	reValues = regexp.MustCompile(`^\s*(\d+)\s+(\d+)\s+(\d+)$`)
)

func main() {
	lines, _ := file.GetLines("../data.txt")
	values := make([][]int, len(lines))
	for i, line := range lines {
		if reValues.MatchString(line) {
			matches := reValues.FindStringSubmatch(line)
			values[i] = []int{common.StrToA(matches[1]), common.StrToA(matches[2]), common.StrToA(matches[3])}
		}
	}

	possibleTriangles := 0
	for _, vals := range values {
		dvals := append(vals, vals...)
		possible := true
		for i := 0; i < len(vals); i++ {
			if dvals[i]+dvals[i+1] <= dvals[i+2] {
				possible = false
				break
			}
		}
		if possible {
			possibleTriangles++
		}
	}

	fmt.Println(possibleTriangles)

}
