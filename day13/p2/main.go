package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common"
)

type State struct {
	p common.Pos
	s int
}

func main() {

	favNum := 1352
	startPos := common.Pos{X: 1, Y: 1}

	queue := make(common.Queue[State], 0, 100)
	visited := make(common.PosContainer)

	queue.Enqueue(State{p: startPos, s: 0})
	visited[startPos] = true

	deltas := common.Positions{common.DN, common.DE, common.DS, common.DW}

	for !queue.Empty() {
		cur := *(queue.Dequeue())
		if cur.s < 50 {
			for _, d := range deltas {
				p := cur.p.Add(d)
				if p.X >= 0 && p.Y >= 0 && !isWall(p, favNum) {
					if !visited.Has(p) {
						visited[p] = true
						queue.Enqueue(State{p: p, s: cur.s + 1})
					}
				}
			}
		}

	}

	fmt.Println(len(visited))

}

func isWall(p common.Pos, favNum int) bool {
	v := p.X*p.X + 3*p.X + 2*p.X*p.Y + p.Y + p.Y*p.Y
	v += favNum
	c := count1Bits(v)
	return c%2 != 0
}

func count1Bits(n int) int {
	c := 0
	for n != 0 {
		n &= n - 1
		c++
	}
	return c
}
