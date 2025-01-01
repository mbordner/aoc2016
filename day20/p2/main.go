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

	vps := rs.ValuePairs()

	count := int64(0)
	if vps[0] > 0 {
		count = vps[0] - 1
	}
	for i := 2; i < len(vps); i += 2 {
		count += vps[i] - vps[i-1] - 1
	}

	count += 4294967295 - vps[len(vps)-1]

	fmt.Println(count)

}

func getIntVal(s string) int64 {
	val, _ := strconv.ParseInt(s, 10, 64)
	return val
}
