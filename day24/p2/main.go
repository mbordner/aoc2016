package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common"
	"github.com/mbordner/aoc2016/common/file"
	"sort"
	"strings"
)

type Grid [][]byte
type POIMap map[byte]Pos
type POI map[Pos]byte

type State struct {
	y int
	x int
	a string
}

func (s State) String() string {
	return fmt.Sprintf("{%d,%d,%s}", s.x, s.y, s.a)
}

type Pos struct {
	y int
	x int
}

var (
	deltas = []Pos{{y: -1, x: 0}, {y: 0, x: 1}, {y: 1, x: 0}, {y: 0, x: -1}}
)

func main() {
	grid, poim := getData("../data.txt")

	start := State{y: poim['0'].y, x: poim['0'].x, a: ""}
	poi := make(POI)
	for n, p := range poim {
		poi[p] = n
	}

	queue := make(common.Queue[State], 0, 100)
	visited := make(map[State]bool)
	prev := make(map[State]State)

	visited[start] = true
	queue.Enqueue(start)

	var solution []State

	for !queue.Empty() {
		cur := *(queue.Dequeue())
		cp := Pos{y: cur.y, x: cur.x}
		if len(cur.a) == len(poi) && poi[cp] == '0' {
			solution = []State{cur}
			for p := prev[cur]; p != start; p = prev[p] {
				solution = append([]State{p}, solution...)
			}
			break
		} else {
			for _, d := range deltas {
				np := Pos{y: cur.y + d.y, x: cur.x + d.x}
				if grid[np.y][np.x] == '.' {
					na := cur.a
					if _, e := poi[np]; e {
						if poi[np] != '0' || len(na) == len(poi)-1 {
							if !strings.Contains(na, string(poi[np])) {
								nabs := []byte(na)
								nabs = append(nabs, poi[np])
								sort.Slice(nabs, func(i, j int) bool {
									return nabs[i] < nabs[j]
								})
								na = string(nabs)
							}
						}
					}
					ns := State{y: np.y, x: np.x, a: na}
					if _, e := visited[ns]; !e {
						visited[ns] = true
						prev[ns] = cur
						queue.Enqueue(ns)
					}
				}
			}
		}
	}

	for i, s := range solution {
		sp := Pos{y: s.y, x: s.x}
		c := "."
		if b, e := poi[sp]; e {
			c = string(b)
		}
		fmt.Printf("%d: %d, %d %s\n", i, s.x, s.y, c)
	}
	fmt.Println(len(solution))
}

func getData(f string) (Grid, POIMap) {
	lines, _ := file.GetLines(f)
	grid := make(Grid, len(lines))
	poi := make(POIMap)
	for y, line := range lines {
		grid[y] = []byte(line)
		for x := range line {
			if grid[y][x] == '#' || grid[y][x] == '.' {
				continue
			}
			poi[grid[y][x]] = Pos{y: y, x: x}
			grid[y][x] = '.'
		}
	}

	return grid, poi
}
