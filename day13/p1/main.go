package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common"
)

func main() {
	params := []struct {
		favNum   int
		startPos common.Pos
		goalPos  common.Pos
	}{{
		favNum:   10,
		startPos: common.Pos{X: 1, Y: 1},
		goalPos:  common.Pos{X: 7, Y: 4},
	}, {
		favNum:   1352,
		startPos: common.Pos{X: 1, Y: 1},
		goalPos:  common.Pos{X: 31, Y: 39},
	}}
	param := params[1]
	startPos := param.startPos
	goalPos := param.goalPos

	queue := make(common.Queue[common.Pos], 0, 100)
	prev := make(common.PosLinker)
	visited := make(common.PosContainer)

	var solution common.Positions

	queue.Enqueue(startPos)
	visited[startPos] = true

	deltas := common.Positions{common.DN, common.DE, common.DS, common.DW}

	for !queue.Empty() {
		cur := *(queue.Dequeue())
		if cur == goalPos {
			solution = common.Positions{goalPos}
			for p := prev[goalPos]; p != startPos; p = prev[p] {
				solution = append(common.Positions{p}, solution...)
			}
			break
		} else {
			for _, d := range deltas {
				p := cur.Add(d)
				if p.X >= 0 && p.Y >= 0 && !isWall(p, param.favNum) {
					if !visited.Has(p) {
						visited[p] = true
						prev[p] = cur
						queue.Enqueue(p)
					}
				}
			}
		}
	}

	fmt.Println(len(solution))

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
