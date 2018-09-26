package unixtime

import (
	"math"
	"testing"
)

func TestMilli(t *testing.T) {
	tests := []int64{math.MinInt64, math.MinInt64 + 1, -1e9, -1, 0, 1, 1e9,
		math.MaxInt64 - 1e9, math.MaxInt64 - 1, math.MaxInt64}
	for _, tc := range tests {
		ti := FromMilli(tc)
		ms := Milli(ti)
		if ms != tc {
			t.Errorf("Milli got %d; want %d", ms, tc)
		}
	}
}

func TestMicro(t *testing.T) {
	tests := []int64{math.MinInt64, math.MinInt64 + 1, -1e12, -1e9, -1, 0,
		1, 1e9, 1e12, math.MaxInt64 - 1e9, math.MaxInt64 - 1,
		math.MaxInt64}
	for _, tc := range tests {
		ti := FromMicro(tc)
		ms := Micro(ti)
		if ms != tc {
			t.Errorf("Milli got %d; want %d", ms, tc)
		}
	}
}
