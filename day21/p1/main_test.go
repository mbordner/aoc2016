package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ScrambleInstructions(t *testing.T) {
	s := scrambleInstructions("abcde", "../test.txt")
	assert.Equal(t, `decab`, s)
}

func Test_Combination1(t *testing.T) {
	s := `abcde`
	s = swapPositions(s, 4, 0)
	assert.Equal(t, `ebcda`, s)
	s = swapLetters(s, "d", "b")
	assert.Equal(t, `edcba`, s)
	s = reverseThroughPositions(s, 0, 4)
	assert.Equal(t, `abcde`, s)
	s = rotateLeftSteps(s, 1)
	assert.Equal(t, `bcdea`, s)
	s = movePositions(s, 1, 4)
	assert.Equal(t, `bdeac`, s)
	s = movePositions(s, 3, 0)
	assert.Equal(t, `abdec`, s)
	s = rotateBasedOnLetterPosition(s, "b")
	assert.Equal(t, `ecabd`, s)
	s = rotateBasedOnLetterPosition(s, "d")
	assert.Equal(t, `decab`, s)
}

func Test_RotateBasedOnLetterPosition(t *testing.T) {
	tests := []struct {
		is string
		ib string
		rs string
	}{
		{
			is: "abcdefg",
			ib: "b",
			rs: "fgabcde",
		},
		{
			is: "abcdefg",
			ib: "e",
			rs: "bcdefga",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s", tt.is, tt.ib), func(t *testing.T) {
			rs := rotateBasedOnLetterPosition(tt.is, tt.ib)
			assert.Equal(t, tt.rs, rs)
		})
	}
}

func Test_ReverseThroughPositions(t *testing.T) {
	tests := []struct {
		is string
		ii int
		ij int
		rs string
	}{
		{
			is: "abcdefg",
			ii: 1,
			ij: 3,
			rs: "adcbefg",
		},
		{
			is: "abcdefgh",
			ii: 1,
			ij: 4,
			rs: "aedcbfgh",
		},
		{
			is: "abcdefg",
			ii: 0,
			ij: 6,
			rs: "gfedcba",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%d,%d", tt.is, tt.ii, tt.ij), func(t *testing.T) {
			rs := reverseThroughPositions(tt.is, tt.ii, tt.ij)
			assert.Equal(t, tt.rs, rs)
		})
	}
}

func Test_MovePositions(t *testing.T) {
	tests := []struct {
		is string
		ii int
		ij int
		rs string
	}{
		{
			is: "abcdefg",
			ii: 1,
			ij: 3,
			rs: "acdbefg",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%d,%d", tt.is, tt.ii, tt.ij), func(t *testing.T) {
			rs := movePositions(tt.is, tt.ii, tt.ij)
			assert.Equal(t, tt.rs, rs)
		})
	}
}

func Test_SwapLetters(t *testing.T) {
	tests := []struct {
		is string
		ia string
		ib string
		rs string
	}{
		{
			is: "abcdefg",
			ia: "b",
			ib: "e",
			rs: "aecdbfg",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s,%s", tt.is, tt.ia, tt.ib), func(t *testing.T) {
			rs := swapLetters(tt.is, tt.ia, tt.ib)
			assert.Equal(t, tt.rs, rs)
		})
	}
}

func Test_SwapPositions(t *testing.T) {
	tests := []struct {
		is string
		ii int
		ij int
		rs string
	}{
		{
			is: "abcdefg",
			ii: 1,
			ij: 4,
			rs: "aecdbfg",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%d,%d", tt.is, tt.ii, tt.ij), func(t *testing.T) {
			rs := swapPositions(tt.is, tt.ii, tt.ij)
			assert.Equal(t, tt.rs, rs)
		})
	}
}

func Test_RotateSteps(t *testing.T) {
	tests := []struct {
		is   string
		in   int
		idir string
		rs   string
	}{
		{
			is:   "abc",
			in:   1,
			idir: "right",
			rs:   "cab",
		},
		{
			is:   "abc",
			in:   1,
			idir: "left",
			rs:   "bca",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%d,%s", tt.is, tt.in, tt.idir), func(t *testing.T) {
			var rs string
			if tt.idir == "right" {
				rs = rotateRightSteps(tt.is, tt.in)
			} else {
				rs = rotateLeftSteps(tt.is, tt.in)
			}

			assert.Equal(t, tt.rs, rs)
		})
	}
}

func Test_InsertChar(t *testing.T) {
	tests := []struct {
		is string
		ii int
		ib byte
		rs string
	}{
		{
			is: "abc",
			ii: 0,
			ib: 'z',
			rs: "zabc",
		},
		{
			is: "abc",
			ii: 1,
			ib: 'z',
			rs: "azbc",
		},
		{
			is: "abc",
			ii: 3,
			ib: 'z',
			rs: "abcz",
		},
		{
			is: "abc",
			ii: 5,
			ib: 'z',
			rs: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%d,%s", tt.is, tt.ii, string(tt.ib)), func(t *testing.T) {
			r := insertChar(tt.is, tt.ii, tt.ib)

			assert.Equal(t, len(tt.rs), len(r))
		})
	}
}

func Test_RemoveChar(t *testing.T) {
	tests := []struct {
		is string
		ii int
		rs string
		rb byte
	}{
		{
			is: "abc",
			ii: 0,
			rs: "bc",
			rb: 'a',
		},
		{
			is: "abc",
			ii: 2,
			rs: "ab",
			rb: 'c',
		},
		{
			is: "abc",
			ii: 1,
			rs: "ac",
			rb: 'b',
		},
		{
			is: "abc",
			ii: -1,
			rs: "abc",
			rb: byte(0),
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%d", tt.is, tt.ii), func(t *testing.T) {
			rs, rb := removeChar(tt.is, tt.ii)

			assert.Equal(t, len(tt.rs), len(rs))
			assert.Equal(t, tt.rb, rb)
		})
	}
}
