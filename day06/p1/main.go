package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common/file"
	"sort"
)

type ByteStat struct {
	b byte
	c int
}

func main() {
	_, stats := getData("../data.txt")
	data := make([]byte, len(stats))
	for i, s := range stats {
		data[i] = s[0].b
	}

	fmt.Println(string(data))
}

func getData(f string) ([]string, [][]ByteStat) {

	lines, _ := file.GetLines(f)
	counts := make([]map[byte]int, len(lines[0]))
	stats := make([][]ByteStat, len(counts))
	for c := range counts {
		counts[c] = make(map[byte]int)
		stats[c] = make([]ByteStat, 0, len(lines))
	}

	for _, line := range lines {
		for i, b := range line {
			if c, e := counts[i][byte(b)]; e {
				counts[i][byte(b)] = c + 1
			} else {
				counts[i][byte(b)] = 1
			}
		}
	}

	for i, cs := range counts {
		for b, c := range cs {
			stats[i] = append(stats[i], ByteStat{b: b, c: c})
		}
		sort.Slice(stats[i], func(a, b int) bool {
			return stats[i][a].c > stats[i][b].c
		})
	}

	return lines, stats
}
