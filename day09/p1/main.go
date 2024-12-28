package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common"
	"github.com/mbordner/aoc2016/common/file"
	"regexp"
	"strings"
)

var (
	reMarker = regexp.MustCompile(`^(\((\d+)x(\d+)\))`)
)

/*
ADVENT contains no markers and decompresses to itself with no changes, resulting in a decompressed length of 6.
A(1x5)BC repeats only the B a total of 5 times, becoming ABBBBBC for a decompressed length of 7.
(3x3)XYZ becomes XYZXYZXYZ for a decompressed length of 9.
A(2x2)BCD(2x2)EFG doubles the BC and EF, becoming ABCBCDEFEFG for a decompressed length of 11.
(6x1)(1x3)A simply becomes (1x3)A - the (1x3) looks like a marker, but because it's within a data section of another marker, it is not treated any differently from the A that comes after it. It has a decompressed length of 6.
X(8x2)(3x3)ABCY becomes X(3x3)ABC(3x3)ABCY (for a decompressed length of 18), because the decompressed data from the (8x2) marker (the (3x3)ABC) is skipped and not processed further.
*/
// 120942 too high
func main() {

	sb := strings.Builder{}
	contentBytes, _ := file.GetContent("../data.txt")

	//contentBytes = []byte(`(3x3)XYZ`)
	//contentBytes = []byte(`ADVENT`)
	//contentBytes = []byte(`A(1x5)BC`)
	//contentBytes = []byte(`A(2x2)BCD(2x2)EFG`)
	//contentBytes = []byte(`(6x1)(1x3)A`)
	//contentBytes = []byte(`X(8x2)(3x3)ABCY`)

	data := string(contentBytes)
	data = strings.TrimSpace(data)
	for ptr := 0; ptr < len(data); ptr++ {
		if reMarker.MatchString(data[ptr:]) {
			matches := reMarker.FindStringSubmatch(data[ptr:])
			charCount := common.StrToA(matches[2])
			repeatCount := common.StrToA(matches[3])
			ptr += len(matches[1])
			chars := data[ptr : ptr+charCount]
			ptr += charCount - 1
			for r := 0; r < repeatCount; r++ {
				sb.WriteString(chars)
			}
		} else {
			sb.WriteByte(data[ptr])
		}
	}

	fmt.Println(sb.String())
	fmt.Println(sb.Len())
}
