package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common"
	"github.com/mbordner/aoc2016/common/file"
	"regexp"
	"sort"
)

var (
	reRoomData = regexp.MustCompile(`(([a-z]|\-)+)-(\d+)\[([a-z]+)\]`)
)

type RoomData struct {
	encryptedName string
	sectionNumber int
	checkSum      string
}

func (rd RoomData) IsReal() bool {
	check := []byte(rd.checkSum)

	if len(check) != 5 {
		panic("checksum length is not 5")
	}

	counts := make(map[byte]int)
	for _, b := range rd.encryptedName {
		if b == '-' {
			continue
		}
		if c, e := counts[byte(b)]; e {
			counts[byte(b)] = c + 1
		} else {
			counts[byte(b)] = 1
		}
	}

	type lc struct {
		b byte
		c int
	}

	lcs := make([]lc, 0, len(counts))

	for b, c := range counts {
		lcs = append(lcs, lc{b: b, c: c})
	}

	sort.Slice(lcs, func(i, j int) bool {
		if lcs[i].c == lcs[j].c {
			return lcs[i].b < lcs[j].b
		}
		return lcs[i].c > lcs[j].c
	})

	if len(lcs) < len(counts) {
		return false
	}

	for i := 0; i < len(check); i++ {
		if lcs[i].b != check[i] {
			return false
		}
	}

	return true
}

func main() {
	rds := getData("../data.txt")

	sum := 0
	for _, rd := range rds {
		if rd.IsReal() {
			sum += rd.sectionNumber
		}
	}

	fmt.Println(sum)
}

func getData(f string) []RoomData {
	lines, _ := file.GetLines(f)
	rd := make([]RoomData, 0, len(lines))

	for _, line := range lines {
		matches := reRoomData.FindStringSubmatch(line)
		rd = append(rd, RoomData{encryptedName: matches[1], sectionNumber: common.StrToA(matches[3]), checkSum: matches[4]})
	}

	return rd
}
