package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common"
	"github.com/mbordner/aoc2016/common/file"
)

func main() {
	grid := common.Grid{{'1', '2', '3'}, {'4', '5', '6'}, {'7', '8', '9'}}
	button := common.Pos{Y: 1, X: 1}

	lines, _ := file.GetLines("../data.txt")
	buttons := make([]byte, len(lines))
	for i, line := range lines {
		for _, d := range line {
			nb := button
			if d == 'U' {
				nb = nb.Add(common.DU)
			} else if d == 'R' {
				nb = nb.Add(common.DR)
			} else if d == 'D' {
				nb = nb.Add(common.DD)
			} else if d == 'L' {
				nb = nb.Add(common.DL)
			}
			if grid.ContainsPos(nb) {
				button = nb
			}
		}
		buttons[i] = grid[button.Y][button.X]
	}

	fmt.Println(string(buttons))
}
