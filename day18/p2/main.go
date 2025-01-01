package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common/file"
	"strings"
)

const (
	TRAP = '^'
	SAFE = '.'
)

func main() {
	firstRow := getData("../data.txt")

	sum := strings.Count(firstRow, string(SAFE))
	prevRow := firstRow
	for i := 1; i < 400000; i++ {
		row := getNextTileRow(prevRow)
		sum += strings.Count(row, string(SAFE))
		prevRow = row
	}

	fmt.Println(sum)
}

func getNextTileRow(prevRow string) string {
	pRow := []byte(prevRow)
	row := make([]byte, len(pRow))
	for j := range row {
		l, c, r := false, false, false
		if j > 0 && pRow[j-1] == TRAP {
			l = true
		}
		if pRow[j] == TRAP {
			c = true
		}
		if j < len(prevRow)-1 && pRow[j+1] == TRAP {
			r = true
		}
		trap := false
		if l && c && !r {
			trap = true
		}
		if c && r && !l {
			trap = true
		}
		if l && !c && !r {
			trap = true
		}
		if r && !c && !l {
			trap = true
		}
		if trap {
			row[j] = TRAP
		} else {
			row[j] = SAFE
		}
	}
	return string(row)
}

func getData(f string) string {
	content, _ := file.GetContent(f)
	return strings.TrimSpace(string(content))
}
