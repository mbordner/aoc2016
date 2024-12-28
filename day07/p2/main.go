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
	superNetSeqs []string
	hyperNetSeqs []string
}

func (a *Address) supportsTransportLayerSnooping() bool {
	for _, seq := range a.hyperNetSeqs {
		if hasAutonomousBridgeBypassAnnotation(seq) {
			return false
		}
	}
	for _, seq := range a.superNetSeqs {
		if hasAutonomousBridgeBypassAnnotation(seq) {
			return true
		}
	}
	return false
}

func (a *Address) supportsSuperSecretListening() bool {
	abaS := []string{}
	for _, seq := range a.superNetSeqs {
		abaS = append(abaS, getAreaBroadcastAccessors(seq)...)
	}
	if len(abaS) == 0 {
		return false
	}
	for _, aba := range abaS {
		bab := getByteAllocationBlock(aba)
		for _, seq := range a.hyperNetSeqs {
			if strings.Contains(seq, bab) {
				return true
			}
		}
	}
	return false
}

// ABA
func getAreaBroadcastAccessors(s string) []string {
	abaS := make([]string, 0, len(s))
	for i := 0; i <= len(s)-3; i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			abaS = append(abaS, s[i:i+3])
		}
	}
	return abaS
}

// BAB
func getByteAllocationBlock(aba string) string {
	bab := make([]byte, len(aba))
	bab[0], bab[1], bab[2] = aba[1], aba[0], aba[1]
	return string(bab)
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
		if addr.supportsSuperSecretListening() {
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
		addr := Address{str: line, superNetSeqs: []string{}, hyperNetSeqs: []string{}}
		matches := reSequences.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			m := match[0]
			if m[0] == '[' {
				addr.hyperNetSeqs = append(addr.hyperNetSeqs, m[1:len(m)-1])
			} else {
				addr.superNetSeqs = append(addr.superNetSeqs, m)
			}
		}
		addresses = append(addresses, addr)
	}
	return addresses
}
