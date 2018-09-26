package unixtime_test

import (
	"fmt"
	"log"
	"time"

	"github.com/ulikunitz/unixtime"
)

func Example() {
	t, err := time.Parse(time.RFC3339Nano, "1961-04-12T09:06:59.7+03:00")
	if err != nil {
		log.Fatalf("Parse error %s", err)
	}

	ms := unixtime.Milli(t)
	fmt.Printf("Unix time: %d ms\n", ms)

	tms := unixtime.FromMilli(ms)
	fmt.Printf("FromMilli: %s\n", tms.Format(time.RFC3339Nano))

	µs := unixtime.Micro(t)
	fmt.Printf("Unix time: %d µs\n", µs)

	tµs := unixtime.FromMicro(µs)
	fmt.Printf("FromMicro: %s\n", tµs.Format(time.RFC3339Nano))

	// Output:
	// Unix time: -275248380300 ms
	// FromMilli: 1961-04-12T07:06:59.7+01:00
	// Unix time: -275248380300000 µs
	// FromMicro: 1961-04-12T07:06:59.7+01:00
}
