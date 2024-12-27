package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common"
	"github.com/mbordner/aoc2016/common/file"
	"strings"
)

type Dir common.Pos
type Turn int

const (
	L Turn = -1
	R Turn = 1
)

var (
	N Dir = Dir(common.Pos{Y: -1, X: 0})
	E Dir = Dir(common.Pos{Y: 0, X: 1})
	S Dir = Dir(common.Pos{Y: 1, X: 0})
	W Dir = Dir(common.Pos{Y: 0, X: -1})
)

func StrToTurn(s string) Turn {
	if s == "L" {
		return L
	}
	return R
}

func (d Dir) Turn(t Turn) Dir {
	if t == L {
		switch d {
		case N:
			return W
		case E:
			return N
		case S:
			return E
		case W:
			return S
		}
	} else if t == R {
		switch d {
		case N:
			return E
		case E:
			return S
		case S:
			return W
		case W:
			return N
		}
	}
	return d
}

func main() {

	p := common.Pos{X: 0, Y: 0}

	d := N

	content, _ := file.GetContent("../data.txt")

	//content = []byte(`R8, R4, R4, R8, L00`)

	seen := make(common.PosContainer)

	moves := strings.Split(string(content), ", ")
outer:
	for _, move := range moves {
		d = d.Turn(StrToTurn(string(move[0])))
		s := common.StrToA(move[1:])
		for i := 0; i < s; i++ {
			n := p.Add(common.Pos(d))
			p = n
			if seen.Has(n) {
				break outer
			}
			seen[n] = true
		}
	}

	fmt.Println(common.Abs(p.Y) + common.Abs(p.X))
}
