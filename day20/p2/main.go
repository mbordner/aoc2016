package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common/file"
	"github.com/mbordner/aoc2016/common/ranges"
	"regexp"
	"strconv"
)

var (
	reRange = regexp.MustCompile(`^(\d+)-(\d+)$`)
)

func main() {

	rs := ranges.Collection[int64]{}

	lines, _ := file.GetLines("../data.txt")
	for _, line := range lines {
		matches := reRange.FindStringSubmatch(line)
		rs.Add(getIntVal(matches[1]), getIntVal(matches[2]))
	}

	fmt.Println(rs.ValuePairs()[1] + 1)

}

func getIntVal(s string) int64 {
	val, _ := strconv.ParseInt(s, 10, 64)
	return val
}
