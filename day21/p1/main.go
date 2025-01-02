package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common/file"
	"regexp"
	"strconv"
	"strings"
)

var (
	reSwapPositions               = regexp.MustCompile(`swap position (\d+) with position (\d+)`)
	reSwapLetters                 = regexp.MustCompile(`swap letter (\w) with letter (\w)`)
	reReverseThroughPositions     = regexp.MustCompile(`reverse positions (\d+) through (\d+)`)
	reRotateSteps                 = regexp.MustCompile(`rotate (left|right) (\d+) steps?`)
	reMovePositions               = regexp.MustCompile(`move position (\d+) to position (\d+)`)
	reRotateBasedOnLetterPosition = regexp.MustCompile(`rotate based on position of letter (\w)`)
)

// fdbgache is not right
func main() {
	fmt.Println(scrambleInstructions("abcdefgh", "../data.txt"))
}

func scrambleInstructions(password string, filename string) string {
	s := password
	lines, _ := file.GetLines(filename)
	for _, line := range lines {
		if reSwapPositions.MatchString(line) {
			matches := reSwapPositions.FindStringSubmatch(line)
			s = swapPositions(s, getIntVal(matches[1]), getIntVal(matches[2]))
		} else if reSwapLetters.MatchString(line) {
			matches := reSwapLetters.FindStringSubmatch(line)
			s = swapLetters(s, matches[1], matches[2])
		} else if reReverseThroughPositions.MatchString(line) {
			matches := reReverseThroughPositions.FindStringSubmatch(line)
			s = reverseThroughPositions(s, getIntVal(matches[1]), getIntVal(matches[2]))
		} else if reRotateSteps.MatchString(line) {
			matches := reRotateSteps.FindStringSubmatch(line)
			switch matches[1] {
			case "left":
				s = rotateLeftSteps(s, getIntVal(matches[2]))
			case "right":
				s = rotateRightSteps(s, getIntVal(matches[2]))
			}
		} else if reMovePositions.MatchString(line) {
			matches := reMovePositions.FindStringSubmatch(line)
			s = movePositions(s, getIntVal(matches[1]), getIntVal(matches[2]))
		} else if reRotateBasedOnLetterPosition.MatchString(line) {
			matches := reRotateBasedOnLetterPosition.FindStringSubmatch(line)
			s = rotateBasedOnLetterPosition(s, matches[1])
		} else {
			fmt.Println("didn't match:", line)
		}
	}
	return s
}

func rotateBasedOnLetterPosition(s, b string) string {
	index := strings.Index(s, b)
	n := index + 1
	if index >= 4 {
		n++
	}
	return rotateRightSteps(s, n)
}

func movePositions(s string, i, j int) string {
	ns, b := removeChar(s, i)
	ns = insertChar(ns, j, b)
	return ns
}

func swapPositions(s string, i, j int) string {
	bs := []byte(s)
	bs[i], bs[j] = bs[j], bs[i]
	return string(bs)
}

func swapLetters(s, a, b string) string {
	i := strings.Index(s, a)
	j := strings.Index(s, b)
	return swapPositions(s, j, i)
}

func reverseThroughPositions(s string, i, j int) string {
	bs := []byte(s)
	for k, l, h := i, j, (j-i+1)/2+i; k < h; k, l = k+1, l-1 {
		bs[k], bs[l] = bs[l], bs[k]
	}
	return string(bs)
}

func rotateLeftSteps(s string, n int) string {
	for i := 0; i < n; i++ {
		s = s[1:] + string(s[0])
	}
	return s
}

func rotateRightSteps(s string, n int) string {
	for i := 0; i < n; i++ {
		s = string(s[len(s)-1]) + s[:len(s)-1]
	}
	return s
}

func removeChar(s string, i int) (string, byte) {
	if i >= 0 && i < len(s) {
		b := byte(s[i])
		if i == 0 {
			return s[1:], b
		} else if i == len(s)-1 {
			return s[:i], b
		}
		return s[:i] + s[i+1:], b
	}
	return s, byte(0)
}

func insertChar(s string, i int, b byte) string {
	if i >= 0 && i <= len(s) {
		if i == 0 {
			return string(append([]byte{b}, s...))
		} else if i == len(s) {
			return s + string(b)
		}
		return s[:i] + string(b) + s[i:]
	}
	return s
}

func getIntVal(s string) int {
	val, _ := strconv.ParseInt(s, 10, 64)
	return int(val)
}
