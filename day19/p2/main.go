package main

import "fmt"

type node struct {
	e int
	p *node
	n *node
}

func main() {
	size := 3017957
	ptr := &node{e: 1, p: nil, n: nil}
	var last *node
	for e, p := 2, ptr; e <= size; e, p = e+1, p.n {
		n := &node{e: e, p: p, n: nil}
		p.n = n
		last = n
	}
	last.n, ptr.p = ptr, last

	for size > 1 {
		fmt.Println("size: ", size)
		rm := size / 2
		p := ptr
		for i := 0; i < rm; i++ {
			p = p.n
		}
		p.p.n, p.n.p = p.n, p.p
		ptr = ptr.n
		size--
	}

	fmt.Println(ptr.e)

}
