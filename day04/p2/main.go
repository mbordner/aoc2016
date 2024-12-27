package main

import (
	"bytes"
	"fmt"
	"github.com/mbordner/aoc2016/common"
	"github.com/mbordner/aoc2016/common/file"
	"regexp"
	"sort"
	"strings"
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

func rotate(b byte, sn int) byte {
	if b == '-' {
		return ' '
	}
	letters := []byte(`abcdefghijklmnopqrstuvwxyz`)
	r := (bytes.IndexByte(letters, b) + sn) % len(letters)
	return letters[r]
}

func main() {
	rds := getData("../data.txt")

	decrypted := make([]string, len(rds))

	for j, rd := range rds {
		if rd.IsReal() {
			data := make([]byte, len(rd.encryptedName))
			for i, b := range rd.encryptedName {
				data[i] = rotate(byte(b), rd.sectionNumber)
			}
			decrypted[j] = string(data)
		}
	}

	for i, d := range decrypted {
		if strings.Contains(d, "north") && strings.Contains(d, "pole") && strings.Contains(d, "object") {
			fmt.Println(rds[i].sectionNumber, d)
		}
	}

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
