package unixtime_test

import (
	"fmt"
	"time"

	"github.com/ulikunitz/unixtime"
)

func Example_unixtime() {
	t, err := time.Parse(time.RFC3339Nano, "1961-04-12T09:06:59.7+03:00")
	if err != nil {
		fmt.Printf("Parse error %s\n", err)
		return
	}
	fmt.Printf("ISO8601   %s\n", t.Format(time.RFC3339Nano))
	ms := unixtime.Milli(t)
	µs := unixtime.Micro(t)
	ns := t.UnixNano()
	fmt.Printf("Unix time %d ms %d µs %d ns\n", ms, µs, ns)
	// Output:
	// ISO8601   1961-04-12T09:06:59.7+03:00
	// Unix time -275248380300 ms -275248380300000 µs -275248380300000000 ns
}
