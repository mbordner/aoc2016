package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/mbordner/aoc2016/common"
	"slices"
	"strconv"
)

const (
	L = "L"
	R = "R"
	U = "U"
	D = "D"
)

type Pos struct {
	Y, X int
}

func (p Pos) Neighbors() ([]string, []Pos) {
	var moves []string
	var neighbors []Pos

	if p.X < 3 {
		moves = append(moves, R)
		neighbors = append(neighbors, Pos{X: p.X + 1, Y: p.Y})
	}
	if p.Y < 3 {
		moves = append(moves, D)
		neighbors = append(neighbors, Pos{X: p.X, Y: p.Y + 1})
	}
	if p.X > 0 {
		moves = append(moves, L)
		neighbors = append(neighbors, Pos{X: p.X - 1, Y: p.Y})
	}
	if p.Y > 0 {
		moves = append(moves, U)
		neighbors = append(neighbors, Pos{X: p.X, Y: p.Y - 1})
	}
	return moves, neighbors
}

type State struct {
	p    Pos
	path string
}

func main() {

	passcode := `pgflpeqp`

	queue := make(common.Queue[State], 0, 16)

	start := Pos{0, 0}
	goal := Pos{3, 3}
	queue.Enqueue(State{p: start, path: ""})

	for !queue.Empty() {
		cur := *(queue.Dequeue())
		if cur.p == goal {
			fmt.Println(cur.path)
			break
		} else {
			unlockedMoves := getUnlocked(passcode, cur.path)
			ms, ns := cur.p.Neighbors()
			for i := range ns {
				if slices.Contains(unlockedMoves, ms[i]) {
					queue.Enqueue(State{p: ns[i], path: cur.path + ms[i]})
				}
			}
		}
	}
}

func getUnlocked(passcode, path string) []string {
	order := []string{U, D, L, R}
	var moves []string

	data := []byte(fmt.Sprintf("%s%s", passcode, path))
	hash := md5.Sum(data)

	hashString := hex.EncodeToString(hash[:])
	for i := range order {
		v, _ := strconv.ParseInt(string(hashString[i]), 16, 32)
		if v > 10 {
			moves = append(moves, order[i])
		}
	}

	return moves
}
