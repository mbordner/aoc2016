package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetRepeatingByteStats(t *testing.T) {

	testCases := []struct {
		s              string
		results        []RepeatingByteStat
		expectedLength int
		min            int
	}{
		{
			s:              "aaabcd",
			results:        []RepeatingByteStat{{b: 'a', i: 0, c: 3}},
			expectedLength: 1,
			min:            3,
		},
		{
			s:              "zdeaaabcd",
			results:        []RepeatingByteStat{{b: 'a', i: 3, c: 3}},
			expectedLength: 1,
			min:            3,
		},
		{
			s:              "aaabcdaaa",
			results:        []RepeatingByteStat{{b: 'a', i: 0, c: 3}, {b: 'a', i: 6, c: 3}},
			expectedLength: 2,
			min:            3,
		},
		{
			s:              "bcdaaa",
			results:        []RepeatingByteStat{{b: 'a', i: 3, c: 3}},
			expectedLength: 1,
			min:            3,
		},
		{
			s:              "bcccccccdaaa",
			results:        []RepeatingByteStat{{b: 'c', i: 1, c: 7}, {b: 'a', i: 9, c: 3}},
			expectedLength: 2,
			min:            3,
		},
		{
			s:              "bcccccccdaaa",
			results:        []RepeatingByteStat{{b: 'c', i: 1, c: 7}},
			expectedLength: 1,
			min:            4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {

			stats := getRepeatingByteStats(tc.s, tc.min)

			for i := 0; i < len(stats); i++ {
				assert.Equal(t, tc.results[i], stats[i])
			}

			assert.Equal(t, tc.expectedLength, len(stats))
		})
	}
}
