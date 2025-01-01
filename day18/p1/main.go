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

type Tiles []string

func (t Tiles) Print() {
	for _, row := range t {
		fmt.Println(row)
	}
}

func main() {
	tiles := getTiles(getData("../data.txt"), 40)

	sum := 0
	for _, row := range tiles {
		sum += strings.Count(row, ".")
	}

	fmt.Println(sum)
}

func getTiles(firstRow string, rows int) Tiles {
	tiles := make(Tiles, rows)
	tiles[0] = firstRow

	for i := 1; i < rows; i++ {
		row := make([]byte, len(firstRow))
		pRow := []byte(tiles[i-1])
		for j := range row {
			l, c, r := false, false, false
			if j > 0 && pRow[j-1] == TRAP {
				l = true
			}
			if pRow[j] == TRAP {
				c = true
			}
			if j < len(firstRow)-1 && pRow[j+1] == TRAP {
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
		tiles[i] = string(row)
	}

	return tiles
}

func getData(f string) string {
	content, _ := file.GetContent(f)
	return strings.TrimSpace(string(content))
}
