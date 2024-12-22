package main

import (
	"strconv"
	"testing"
)

func TestNext(t *testing.T) {
	want := []int64{
		15887950,
		16495136,
		527345,
		704524,
		1553684,
		12683156,
		11100544,
		12249484,
		7753432,
		5908254,
	}
	secret := int64(123)
	for i := range 10 {
		have := next(secret)

		if have != want[i] {
			t.Fatalf("unexpected value at %d generation: want: %d, have: %d", i+1, want[i], have)
		}

		secret = have
	}
}

func TestSolve_Part1(t *testing.T) {
	cases := []struct {
		secrets []int64
		want    int64
	}{
		{
			secrets: []int64{
				1,
				10,
				100,
				2024,
			},
			want: 37327623,
		},
	}

	for i, tt := range cases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			if have, _ := solve2(tt.secrets); have != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, have)
			}
		})
	}
}

func TestSolve_Part2(t *testing.T) {
	cases := []struct {
		secrets []int64
		want    int64
	}{
		{
			secrets: []int64{
				1,
				2,
				3,
				2024,
			},
			want: 23,
		},
	}

	for i, tt := range cases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			if _, have := solve2(tt.secrets); have != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, have)
			}
		})
	}
}
