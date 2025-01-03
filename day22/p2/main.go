package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common"
	"github.com/mbordner/aoc2016/common/file"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	reNodeConfig = regexp.MustCompile(`/dev/grid/node-x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T`)
)

type NodeConfig struct {
	x int
	y int
	s int
	u int
}

type NodeUsage int
type NodeUsages []NodeUsage
type NetworkUsage struct {
	network *Network
	usages  []NodeUsages
}

func (nu NetworkUsage) Clone() NetworkUsage {
	onu := NetworkUsage{}
	onu.network = nu.network
	onu.usages = make([]NodeUsages, len(nu.usages))
	for i, us := range nu.usages {
		onu.usages[i] = make(NodeUsages, len(us))
		copy(onu.usages[i], us)
	}
	return onu
}

func (nu NetworkUsage) MoveData(vp ViablePair) {
	nu.usages[vp.o.y][vp.o.x] += nu.usages[vp.n.y][vp.n.x]
	nu.usages[vp.n.y][vp.n.x] = 0
}

func (nu NetworkUsage) ViablePairs() []ViablePair {
	viablePairs := make([]ViablePair, 0, 10)

	for y := range nu.usages {
		for x := range nu.usages[y] {
			a := nu.usages[y][x]
			bs := make(NodeConfigs, 0, 4)
			if y > 0 { // top
				bs = append(bs, (*nu.network)[y-1][x])
			}
			if x < len(nu.usages[y])-1 { // right
				bs = append(bs, (*nu.network)[y][x+1])
			}
			if y < len(nu.usages)-1 { // below
				bs = append(bs, (*nu.network)[y+1][x])
			}
			if x > 0 { // left
				bs = append(bs, (*nu.network)[y][x-1])
			}
			for _, b := range bs {
				if a > 0 && (b.size()-int(nu.usages[b.y][b.x])) >= int(a) {
					viablePairs = append(viablePairs, ViablePair{n: (*nu.network)[y][x], o: b})
				}
			}
		}
	}

	return viablePairs
}

type NodeConfigs []NodeConfig
type Network []NodeConfigs

type ViablePair struct {
	n NodeConfig
	o NodeConfig
}

func (n NodeConfig) size() int {
	return n.s
}

func (n NodeConfig) used() int {
	return n.u
}

func (n NodeConfig) avail() int {
	return n.s - n.u
}

func (n NodeConfig) name() string {
	return fmt.Sprintf("/dev/grid/node-x%d-y%d", n.x, n.y)
}

func (n Network) State() string {
	used := make([]string, 0, len(n)*len(n[0]))
	for y := range n {
		for x := range n[y] {
			used = append(used, fmt.Sprintf("%d", n[y][x].used()))
		}
	}
	return strings.Join(used, ",")
}

func (n Network) From(s string) Network {
	tokens := strings.Split(s, ",")
	used := make([]int, len(tokens))
	for u := range tokens {
		used[u] = getIntVal(tokens[u])
	}
	up := 0
	on := make(Network, len(n))
	for y := range on {
		on[y] = make(NodeConfigs, len(n[y]))
		for x := 0; x < len(n[y]); x, up = x+1, up+1 {
			nc := NodeConfig{x: x, y: y, s: n[y][x].s, u: used[up]}
			on[y][x] = nc
		}
	}
	return on
}

func (n Network) Clone() Network {
	return n.From(n.State())
}

func (n Network) MoveData(vp ViablePair) {
	n[vp.o.y][vp.o.x].u += n[vp.n.y][vp.n.x].u
	n[vp.n.y][vp.n.x].u = 0
}

func (n Network) ViablePairs() []ViablePair {
	viablePairs := make([]ViablePair, 0, 10)

	for y := range n {
		for x := range n[y] {
			a := n[y][x]
			bs := make(NodeConfigs, 0, 4)
			if y > 0 { // top
				bs = append(bs, n[y-1][x])
			}
			if x < len(n[y])-1 { // right
				bs = append(bs, n[y][x+1])
			}
			if y < len(n)-1 { // below
				bs = append(bs, n[y+1][x])
			}
			if x > 0 { // left
				bs = append(bs, n[y][x-1])
			}
			for _, b := range bs {
				if a.used() > 0 && b.avail() >= a.used() {
					viablePairs = append(viablePairs, ViablePair{n: a, o: b})
				}
			}
		}
	}

	return viablePairs
}

func (n Network) Usage() NetworkUsage {
	nu := NetworkUsage{network: &n, usages: make([]NodeUsages, len(n))}
	for y := range n {
		nu.usages[y] = make(NodeUsages, len(n[y]))
		for x := range n[y] {
			nu.usages[y][x] = NodeUsage(n[y][x].used())
		}
	}
	return nu
}

type State struct {
	dnX   int
	dnY   int
	srcX  int
	srcY  int
	destX int
	destY int
}

func main() {
	network := getNodeData("../data.txt")

	visited := make(map[State]State)

	stateQueue := make(common.Queue[State], 0, 1000)
	usageQueue := make(common.Queue[NetworkUsage], 0, 1000)

	initialState := State{dnX: len(network[0]) - 1, dnY: 0, srcX: 0, srcY: 0, destX: 0, destY: 0}

	stateQueue.Enqueue(initialState)
	usageQueue.Enqueue(network.Usage())

	visited[initialState] = initialState

	goal := false

	var solution []State

	for !stateQueue.Empty() {
		cur := *(stateQueue.Dequeue())
		usages := *(usageQueue.Dequeue())

		if !(cur.srcX == cur.destX && cur.srcY == cur.destY) {
			if cur.dnX+cur.dnY == 0 {
				goal = true
				solution = []State{cur}
				for p := visited[cur]; p != initialState; p = visited[p] {
					solution = append([]State{p}, solution...)
				}
			}
		}

		if !goal {
			vps := usages.ViablePairs()
			for _, vp := range vps {
				nextUsages := usages.Clone()
				nextUsages.MoveData(vp)
				nextDNX, nextDNY := cur.dnX, cur.dnY
				if vp.n.x == nextDNX && vp.n.y == nextDNY {
					nextDNX, nextDNY = vp.o.x, vp.o.y
					if int(nextUsages.usages[nextDNY][nextDNX]) > (*nextUsages.network)[0][0].s {
						continue
					}
				}
				nextState := State{dnX: nextDNX, dnY: nextDNY, srcX: vp.n.x, srcY: vp.n.y, destX: vp.o.x, destY: vp.o.y}

				if _, e := visited[nextState]; !e {
					visited[nextState] = cur
					stateQueue.Enqueue(nextState)
					usageQueue.Enqueue(nextUsages)
				}
			}
		}
	}

	fmt.Printf("Solution part 2: %d\n", len(solution))
}

func getNodeData(f string) Network {
	lines, _ := file.GetLines(f)
	nodeConfigs := make(NodeConfigs, 0, len(lines)-2)

	minX, minY, maxX, maxY := math.MaxUint32, math.MaxUint32, 0, 0

	for _, line := range lines[2:] {
		matches := reNodeConfig.FindStringSubmatch(line)
		nc := NodeConfig{x: getIntVal(matches[1]), y: getIntVal(matches[2]), s: getIntVal(matches[3]), u: getIntVal(matches[4])}
		minX, maxX, minY, maxY = min(minX, nc.x), max(maxX, nc.x), min(minY, nc.y), max(maxY, nc.y)

		//avail := getIntVal(matches[5])
		//if avail != nc.avail() {
		//	panic("avail isn't matching")
		//}
		nodeConfigs = append(nodeConfigs, nc)
	}

	network := make(Network, maxY+1)
	for y := range network {
		network[y] = make(NodeConfigs, maxX+1)
	}

	for _, nc := range nodeConfigs {
		network[nc.y][nc.x] = nc
	}

	return network
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getIntVal(s string) int {
	val, _ := strconv.ParseInt(s, 10, 64)
	return int(val)
}
