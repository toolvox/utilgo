// Package timeutil provides helpers for working with [pkg/time.Time].
package timeutil

import "time"

const (
	TS_FORMAT = "2006_01_02_15_04_05"
)

// Timestamp returns a timestamp string based on the input [pkg/time.Time].
//
// Uses the: "2006_01_02_15_04_05" template for formatting.
func Timestamp(dt time.Time) string {
	return dt.Format(TS_FORMAT)
}

// TimestampNow returns a timestamp string based on [pkg/time.Now]().
//
// Uses the: "2006_01_02_15_04_05" template for formatting.
func TimestampNow() string {
	return time.Now().Format(TS_FORMAT)
}
