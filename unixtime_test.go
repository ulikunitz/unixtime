package unixtime

import (
	"encoding/json"
	"math"
	"testing"
	"time"
)

const t1 = "1961-04-12T09:06:59+03:00"

type T struct {
	I int
	T S
}

func TestS(t *testing.T) {
	dt, err := time.Parse(time.RFC3339Nano, t1)
	if err != nil {
		t.Fatalf("time.Parse(%q, %q) error %s",
			time.RFC3339Nano, t1, err)
	}
	dt = dt.Local()
	t.Logf("%q", dt.Format(time.RFC3339Nano))
	x := T{1, S{dt}}
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		t.Fatalf("MarshalIndent error %s", err)
	}
	t.Logf("json\n%s", b)
	var y T
	if err := json.Unmarshal(b, &y); err != nil {
		t.Fatalf("Unmarshal error %s", err)
	}
	u, v := y.T.Time, x.T.Time
	if !u.Equal(v) {
		t.Fatalf("Unmarshal got %s; want %s",
			u.Format(time.RFC3339Nano),
			v.Format(time.RFC3339Nano))
	}
}

func TestMilli(t *testing.T) {
	tests := []int64{math.MinInt64, math.MinInt64 + 1,
		-1e9, -1, 0, 1, 1e9, math.MaxInt64 - 1e9,
		math.MaxInt64 - 1, math.MaxInt64}
	for _, tc := range tests {
		ti := FromMilli(tc)
		ms := Milli(ti)
		if ms != tc {
			t.Errorf("Milli got %d; want %d", ms, tc)
		}
	}
}

func TestMicro(t *testing.T) {
	tests := []int64{math.MinInt64, math.MinInt64 + 1,
		-1e12, -1e9, -1, 0, 1, 1e9, 1e12, math.MaxInt64 - 1e9,
		math.MaxInt64 - 1, math.MaxInt64}
	for _, tc := range tests {
		ti := FromMicro(tc)
		ms := Micro(ti)
		if ms != tc {
			t.Errorf("Milli got %d; want %d", ms, tc)
		}
	}
}
