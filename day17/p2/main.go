package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

func main() {

	passcode := `pgflpeqp`

	start := Pos{0, 0}
	goal := Pos{3, 3}

	path := getPath(goal, passcode, start, "")
	fmt.Println(path)
	fmt.Println(len(path))

}

func getPath(g Pos, passcode string, p Pos, path string) string {
	if p == g {
		return path
	}
	solution := ""
	unlockedMoves := getUnlocked(passcode, path)
	ms, ns := p.Neighbors()
	for i := range ns {
		if slices.Contains(unlockedMoves, ms[i]) {
			np := getPath(g, passcode, ns[i], path+ms[i])
			if len(np) > len(solution) {
				solution = np
			}
		}
	}
	return solution
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
