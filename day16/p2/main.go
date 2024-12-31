package main

import "fmt"

func main() {
	fmt.Println(string(fillDisk(35651584, []byte(`10001001100000001`))))
}

func fillDisk(l int, state []byte) []byte {
	a := make([]byte, len(state))
	copy(a, state)

	for len(a) < l {
		a = dragonCurve(a)
	}

	a = a[0:l]
	return checkSum(a)
}

func checkSum(a []byte) []byte {
	r := make([]byte, 0, len(a)/2)
	for i := 0; i <= len(a)-1; i += 2 {
		if a[i] == a[i+1] {
			r = append(r, '1')
		} else {
			r = append(r, '0')
		}
	}
	for len(r)%2 == 0 {
		r = checkSum(r)
	}
	return r
}

func dragonCurve(a []byte) []byte {
	b := make([]byte, len(a))
	copy(b, a)
	for i, j, h := 0, len(b)-1, len(b)/2; i < h; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	for i := range b {
		if b[i] == '0' {
			b[i] = '1'
		} else {
			b[i] = '0'
		}
	}
	r := make([]byte, len(a)*2+1)
	copy(r, a)
	r[len(a)] = '0'
	copy(r[len(a)+1:], b)
	return r
}
