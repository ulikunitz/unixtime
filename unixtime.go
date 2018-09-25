// Package unixtime provides helper function to converts the time values
// to the the Unix time in milliseconds and microseconds.
//
// It also provides the types S, MS, US and NS that can be used to parse
// or create JSON structures that represent time values as Unix time in
// seconds, milliseconds, microseconds and nanoseconds as integers.
//
// The package provides functionality requested in golang issues #27782
// and #19835.
//
//    https://github.com/golang/go/issues/27782
//    https://github.com/golang/go/issues/18935
package unixtime

import (
	"encoding/json"
	"time"
)

// Micro converts a time value to the Unix time in microseconds.
func Micro(t time.Time) int64 {
	s := t.Unix() * 1e6
	µs := int64(t.Nanosecond()) / 1e3
	return s + µs
}

// FromMicro converts the Unix time in microseconds to a time value.
func FromMicro(µs int64) time.Time {
	s := µs / 1e6
	ns := (µs - s*1e6) * 1e3
	return time.Unix(s, ns)
}

// Milli converts a time value to the Unix time in milliseconds.
func Milli(t time.Time) int64 {
	s := t.Unix() * 1e3
	ms := int64(t.Nanosecond()) / 1e6
	return s + ms
}

// FromMilli converts the Unix time in milliseconds to a time value.
func FromMilli(ms int64) time.Time {
	s := ms / 1e3
	ns := (ms - s*1e3) * 1e6
	return time.Unix(s, ns)
}

// S can be used to represent time in a JSON structure as Unix time in
// seconds.
type S struct {
	time.Time
}

// MarshalJSON represents the time value as Unix time in seconds.
func (t S) MarshalJSON() ([]byte, error) { return json.Marshal(t.Unix()) }

// UnmarshalJSON decodes the Unix time in seconds.
func (t *S) UnmarshalJSON(b []byte) error {
	var s int64
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t.Time = time.Unix(s, 0)
	return nil
}

// MS is provided to represent time in JSON structures as Unix time in
// milliseconds.
type MS struct {
	time.Time
}

// MarshalJSON converts the time value to the JSON value of the Unix
// time in milliseconds.
func (t MS) MarshalJSON() ([]byte, error) {
	return json.Marshal(Milli(t.Time))
}

// UnmarshalJSON converts the JSON integer and stores the correspinding
// time value in t.
func (t *MS) UnmarshalJSON(b []byte) error {
	var ms int64
	if err := json.Unmarshal(b, &ms); err != nil {
		return err
	}
	t.Time = FromMilli(ms)
	return nil
}

// US is provided to represent time in JSON structures as Unix time in
// microseconds.
type US struct {
	time.Time
}

// MarshalJSON converts the time value to a JSON integer providing the
// Unix time in microseconds.
func (t US) MarshalJSON() ([]byte, error) {
	return json.Marshal(Micro(t.Time))
}

// UmarshalJSON converts the JSON integer value to a time value and
// stores it in t.
func (t *US) UnmarshalJSON(b []byte) error {
	var µs int64
	if err := json.Unmarshal(b, &µs); err != nil {
		return err
	}
	t.Time = FromMicro(µs)
	return nil
}

// NS is provided to represent time in JSON structures as Unix time in
// nanoseconds.
type NS struct {
	time.Time
}

// MarshalJSON converts the time value stored in t to a JSON integer
// providing the Unix time in nanoseconds.
func (t NS) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.UnixNano())
}

// UnmarshalJSON converts the JSON integer value to a time value and
// stores it in t.
func (t *NS) UnmarshalJSON(b []byte) error {
	var ns int64
	if err := json.Unmarshal(b, &ns); err != nil {
		return err
	}
	t.Time = time.Unix(0, ns)
	return nil
}
