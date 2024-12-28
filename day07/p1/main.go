package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common/file"
	"regexp"
	"strings"
)

var (
	reSequences = regexp.MustCompile(`(\[?[^\[\]]+]?)`)
)

type Address struct {
	str          string
	normalSeqs   []string
	hyperNetSeqs []string
}

func (a *Address) supportsTransportLayerSnooping() bool {
	for _, seq := range a.hyperNetSeqs {
		if hasAutonomousBridgeBypassAnnotation(seq) {
			return false
		}
	}
	for _, seq := range a.normalSeqs {
		if hasAutonomousBridgeBypassAnnotation(seq) {
			return true
		}
	}
	return false
}

// ABBA
func hasAutonomousBridgeBypassAnnotation(s string) bool {
	for i := 0; i <= len(s)-4; i++ {
		if s[i] == s[i+3] && s[i] != s[i+1] && s[i+1] == s[i+2] {
			return true
		}
	}
	return false
}

// 102 too low
func main() {
	addresses := getData("../data.txt")

	count := 0
	for _, addr := range addresses {
		if addr.supportsTransportLayerSnooping() {
			count++
		}
	}

	fmt.Println(count)
}

func getData(f string) []Address {
	lines, _ := file.GetLines(f)
	addresses := make([]Address, 0, len(lines))
	for _, line := range lines {
		line = strings.Split(line, " ")[0]
		addr := Address{str: line, normalSeqs: []string{}, hyperNetSeqs: []string{}}
		matches := reSequences.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			m := match[0]
			if m[0] == '[' {
				addr.hyperNetSeqs = append(addr.hyperNetSeqs, m[1:len(m)-1])
			} else {
				addr.normalSeqs = append(addr.normalSeqs, m)
			}
		}
		addresses = append(addresses, addr)
	}
	return addresses
}
