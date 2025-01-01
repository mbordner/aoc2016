package main

import "fmt"

func main() {
	elves := make([]int, 3017957)
	for i := range elves {
		elves[i] = i + 1
	}
	fmt.Println(whiteElephant(elves))
}

func whiteElephant(elves []int) int {
	if len(elves) == 1 {
		return elves[0]
	}

	next := make([]int, 0, len(elves)/2)

	for i := 0; i < len(elves); i += 2 {
		if i == len(elves)-1 {
			next = append(next[1:], elves[i])
		} else {
			next = append(next, elves[i])
		}
	}

	return whiteElephant(next)
}
