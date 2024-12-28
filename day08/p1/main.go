package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common"
	"github.com/mbordner/aoc2016/common/file"
	"regexp"
)

var (
	reRect      = regexp.MustCompile(`^rect\s+(\d+)x(\d+)`)
	reRotateCol = regexp.MustCompile(`^rotate column x=(\d+) by (\d+)`)
	reRotateRow = regexp.MustCompile(`^rotate row y=(\d+) by (\d+)`)
)

func main() {
	//s := NewScreen(7, 3)
	//s.Rect(3, 2)
	//s.RotateCol(1, 1)
	//s.RotateRow(0, 4)
	//s.RotateCol(1, 1)
	//s.Print()
	s := NewScreen(50, 6)
	//s = NewScreen(7, 3)

	lines, _ := file.GetLines("../data.txt")
	for _, line := range lines {
		if reRect.MatchString(line) {
			matches := reRect.FindStringSubmatch(line)
			w, h := common.StrToA(matches[1]), common.StrToA(matches[2])
			s.Rect(w, h)
		} else if reRotateCol.MatchString(line) {
			matches := reRotateCol.FindStringSubmatch(line)
			x, n := common.StrToA(matches[1]), common.StrToA(matches[2])
			s.RotateCol(x, n)
		} else if reRotateRow.MatchString(line) {
			matches := reRotateRow.FindStringSubmatch(line)
			y, n := common.StrToA(matches[1]), common.StrToA(matches[2])
			s.RotateRow(y, n)
		}
	}

	s.Print()

	count := 0
	for y := 0; y < s.h; y++ {
		for x := 0; x < s.w; x++ {
			if s.grid[y][x] == '#' {
				count++
			}
		}
	}

	fmt.Println(count)
}

type Screen struct {
	w    int
	h    int
	grid [][]byte
}

func (s *Screen) Print() {
	for y := 0; y < s.h; y++ {
		fmt.Println(string(s.grid[y]))
	}
}

func NewScreen(w, h int) *Screen {
	s := &Screen{w: w, h: h}
	s.grid = make([][]byte, h)
	for y := 0; y < h; y++ {
		s.grid[y] = make([]byte, w)
		for x := 0; x < w; x++ {
			s.grid[y][x] = ' '
		}
	}
	return s
}

func (s *Screen) Rect(w, h int) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			s.grid[y][x] = '#'
		}
	}
}

func (s *Screen) RotateCol(x, n int) {
	col := s.getCol(x)
	col = s.rotate(col, n)
	s.setCol(x, col)
}

func (s *Screen) RotateRow(y, n int) {
	row := s.getRow(y)
	row = s.rotate(row, n)
	s.setRow(y, row)
}

func (s *Screen) setRow(y int, r []byte) {
	for x := 0; x < s.w; x++ {
		s.grid[y][x] = r[x]
	}
}

func (s *Screen) setCol(x int, r []byte) {
	for y := 0; y < s.h; y++ {
		s.grid[y][x] = r[y]
	}
}

func (s *Screen) getRow(y int) []byte {
	r := make([]byte, s.w)
	copy(r, s.grid[y])
	return r
}

func (s *Screen) getCol(x int) []byte {
	c := make([]byte, 0, s.h)
	for y := 0; y < s.h; y++ {
		c = append(c, s.grid[y][x])
	}
	return c
}

func (s *Screen) rotate(a []byte, n int) []byte {
	for i := 0; i < n; i++ {
		a = append([]byte{a[len(a)-1]}, a[0:len(a)-1]...)
	}
	return a
}
