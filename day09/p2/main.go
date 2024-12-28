package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common"
	"github.com/mbordner/aoc2016/common/file"
	"regexp"
	"strings"
)

var (
	reMarkerStart = regexp.MustCompile(`^(\((\d+)x(\d+)\))`)
	reMarkerAny   = regexp.MustCompile(`(\((\d+)x(\d+)\))`)
)

func decompressedLength(s string) int {
	dl := 0
	for ptr := 0; ptr < len(s); ptr++ {
		if reMarkerStart.MatchString(s[ptr:]) {
			matches := reMarkerStart.FindStringSubmatch(s[ptr:])
			charCount := common.StrToA(matches[2])
			repeatCount := common.StrToA(matches[3])
			ptr += len(matches[1])
			chars := s[ptr : ptr+charCount]
			subdl := len(chars)
			if reMarkerAny.MatchString(chars) {
				subdl = decompressedLength(chars)
			}
			ptr += charCount - 1
			for r := 0; r < repeatCount; r++ {
				dl += subdl
			}
		} else {
			dl++
		}
	}
	return dl
}

func main() {

	contentBytes, _ := file.GetContent("../data.txt")

	//contentBytes = []byte(`(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN`)
	//contentBytes = []byte(`(27x12)(20x12)(13x14)(7x10)(1x12)A`)

	data := string(contentBytes)
	data = strings.TrimSpace(data)

	dl := decompressedLength(data)

	fmt.Println(dl)
}
