package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
)

var (
	reZeroes = regexp.MustCompile(`^0{5}(.?)(.?)`)
)

func main() {

	// door id: abc
	// The first index which produces a hash that starts with five zeroes is 3231929, which we find by hashing abc3231929;
	///the sixth character of the hash, and thus the first character of the password, is 1.

	doorId := "cxdnnyjw"
	chars := make([]byte, 8)
	found := 0

	i := 0
	for {
		data := []byte(fmt.Sprintf("%s%d", doorId, i))
		hash := md5.Sum(data)
		hashString := hex.EncodeToString(hash[:])
		if reZeroes.MatchString(hashString) {
			matches := reZeroes.FindStringSubmatch(hashString)
			index, err := strconv.ParseInt(matches[1], 10, 64)
			if err == nil && int(index) < len(chars) {
				if chars[int(index)] == 0 {
					chars[int(index)] = byte(matches[2][0])
					found++
					if found == len(chars) {
						break
					}
				}
			}
		}
		i++
	}

	fmt.Println(string(chars))
}
