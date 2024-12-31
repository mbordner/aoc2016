package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"sort"
)

type KeyCandidate struct {
	v string
	i int
}

type SearchKeys []KeyCandidate

type KeySearchMap map[byte]SearchKeys

type RepeatingByteStat struct {
	b byte
	i int
	c int
}

func (ksm KeySearchMap) PurgeExpired(num int) {
	for b, sk := range ksm {
		for i, kc := range sk {
			if num-kc.i < 1000 {
				if i > 0 {
					ksm[b] = sk[i:]
				}
				break
			}
		}
		if len(ksm[b]) > 0 && num-ksm[b][0].i >= 1000 {
			ksm[b] = SearchKeys{}
		}
	}
}

func (ksm KeySearchMap) Add(b byte, k KeyCandidate) {
	if _, e := ksm[b]; e {
		ksm[b] = append(ksm[b], k)
	} else {
		ksm[b] = []KeyCandidate{k}
	}
}

func getHash(salt string, index int) string {
	data := []byte(fmt.Sprintf("%s%d", salt, index))
	hash := md5.Sum(data)
	hashString := hex.EncodeToString(hash[:])
	for i := 0; i < 2016; i++ {
		hash = md5.Sum([]byte(hashString))
		hashString = hex.EncodeToString(hash[:])
	}
	return hashString
}

func main() {

	ksm := make(KeySearchMap)
	keys := make(SearchKeys, 0, 100)

	i := 0
	j := math.MaxInt64
	for i < j {

		hashString := getHash(`cuanljph`, i)

		repeatingBytes := getRepeatingByteStats(hashString, 3)

		// check for key candidates with these repeating chars > 5
		for _, rb := range repeatingBytes {
			if rb.c >= 5 {
				if kcs, e := ksm[rb.b]; e {
					for _, kc := range kcs {
						keys = append(keys, kc)
						sort.Slice(keys, func(i, j int) bool {
							return keys[i].i < keys[j].i
						})
						if len(keys) == 64 {
							j = i + 1000
						}
					}
					// clear these out, as they are no longer key candidates
					ksm[rb.b] = SearchKeys{}
				}

			}
		}

		for _, rb := range repeatingBytes {
			ksm.Add(rb.b, KeyCandidate{v: hashString, i: i})
			break // we're only adding first repeating group
		}

		ksm.PurgeExpired(i)

		i++
	}

	fmt.Println("The 64th key's index is: ", keys[63].i)

}

func getRepeatingByteStats(s string, min int) []RepeatingByteStat {

	rbs := make([]RepeatingByteStat, 0, 10)

	last := s[0]
	count := 0
	for i := 1; i < len(s); i++ {
		if s[i] == last {
			if count == 0 {
				count = 2
			} else {
				count++
			}
		} else if count > 0 {
			if count >= min {
				rbs = append(rbs, RepeatingByteStat{b: last, i: i - count, c: count})
			}
			count = 0
		}
		last = s[i]
	}

	if count >= min {
		rbs = append(rbs, RepeatingByteStat{b: last, i: len(s) - count, c: count})
	}

	return rbs
}
