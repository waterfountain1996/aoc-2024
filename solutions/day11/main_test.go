package main

import (
	"strconv"
	"testing"
)

var testInput = `2333133121414131402`

func TestBlink(t *testing.T) {
	testcases := []struct {
		input   string
		nblinks int
		want    int64
	}{
		{"0 1 10 99 999", 1, 7},
		{"125 17", 6, 22},
		{"125 17", 25, 55312},
	}

	for i, tt := range testcases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			stones := parseInput(tt.input)
			if have := blink(stones, tt.nblinks); have != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, have)
			}
		})
	}
}
