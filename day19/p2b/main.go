package main

import "fmt"

// 26583 too low
func main() {

	size := 3017957
	elves := make([]int, size)
	for i := 0; i < size; i++ {
		elves[i] = i + 1
	}

	ptr := 0
	for len(elves) > 1 {
		cur := elves[ptr]
		//fmt.Println("size: ", len(elves), "cur: ", cur, elves)
		rm := (len(elves)/2 + ptr) % len(elves)
		if rm == len(elves)-1 {
			elves = elves[0:rm]
		} else {
			elves = append(elves[0:rm], elves[rm+1:]...)
		}
		if ptr < len(elves) && elves[ptr] == cur {
			ptr++
		}
		if ptr >= len(elves) {
			ptr = 0
		}

	}

	fmt.Println(elves[0])
}
