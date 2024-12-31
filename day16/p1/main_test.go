package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FillDisk(t *testing.T) {
	testCases := []struct {
		a string
		l int
		r string
	}{
		{
			a: "10000",
			l: 20,
			r: "01100",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.a, func(t *testing.T) {

			r := fillDisk(tc.l, []byte(tc.a))

			assert.Equal(t, len(r), len(tc.r))
			for i := range r {
				assert.Equal(t, byte(r[i]), byte(tc.r[i]))
			}

		})
	}
}

func Test_CheckSum(t *testing.T) {
	testCases := []struct {
		a string
		r string
	}{
		{
			a: "110010110100",
			r: "100",
		},
		{
			a: "10000011110010000111",
			r: "01100",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.a, func(t *testing.T) {

			r := checkSum([]byte(tc.a))

			assert.Equal(t, len(r), len(tc.r))
			for i := range r {
				assert.Equal(t, byte(r[i]), byte(tc.r[i]))
			}

		})
	}
}

func Test_DragonCurve(t *testing.T) {
	testCases := []struct {
		a string
		r string
	}{
		{
			a: "1",
			r: "100",
		},
		{
			a: "0",
			r: "001",
		},
		{
			a: "11111",
			r: "11111000000",
		},
		{
			a: "111100001010",
			r: "1111000010100101011110000",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.a, func(t *testing.T) {

			r := dragonCurve([]byte(tc.a))

			assert.Equal(t, len(r), len(tc.r))
			for i := range r {
				assert.Equal(t, byte(r[i]), byte(tc.r[i]))
			}

		})
	}
}
