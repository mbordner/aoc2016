package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common/file"
	"regexp"
	"sort"
	"strconv"
)

var (
	reNodeConfig = regexp.MustCompile(`/dev/grid/node-x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T`)
)

type NodeConfigs []*NodeConfig

type NodeConfig struct {
	x int
	y int
	s int
	u int
}

func (n *NodeConfig) size() int {
	return n.s
}

func (n *NodeConfig) used() int {
	return n.u
}

func (n *NodeConfig) avail() int {
	return n.s - n.u
}

func (n *NodeConfig) name() string {
	return fmt.Sprintf("/dev/grid/node-x%d-y%d", n.x, n.y)
}

type ViablePair struct {
	n *NodeConfig
	o *NodeConfig
}

func main() {
	ncs := getNodeData("../data.txt")

	avail := make(NodeConfigs, len(ncs))
	copy(avail, ncs)

	sort.Slice(avail, func(i, j int) bool {
		return avail[i].avail() > avail[j].avail()
	})

	used := make(NodeConfigs, len(ncs))
	copy(used, ncs)

	sort.Slice(used, func(i, j int) bool {
		return used[i].used() < used[j].used()
	})

	vps := make([]*ViablePair, 0, len(avail))
	for a := range avail {
		for u := range used {
			n, o := used[u], avail[a]
			if n != o {
				if n.used() > 0 && n.used() <= o.avail() {
					vps = append(vps, &ViablePair{n, o})
				} else {
					break
				}
			}
		}
	}

	fmt.Println(len(vps))
}

func getNodeData(f string) NodeConfigs {
	lines, _ := file.GetLines(f)
	nodeConfigs := make(NodeConfigs, 0, len(lines)-2)
	for _, line := range lines[2:] {
		matches := reNodeConfig.FindStringSubmatch(line)
		nc := &NodeConfig{x: getIntVal(matches[1]), y: getIntVal(matches[2]), s: getIntVal(matches[3]), u: getIntVal(matches[4])}
		//avail := getIntVal(matches[5])
		//if avail != nc.avail() {
		//	panic("avail isn't matching")
		//}
		nodeConfigs = append(nodeConfigs, nc)
	}
	return nodeConfigs
}

func getIntVal(s string) int {
	val, _ := strconv.ParseInt(s, 10, 64)
	return int(val)
}
