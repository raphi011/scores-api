package test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/raphi011/scores"
)

// Check fails the test if `err` != `nil` with the `message` and arg `err`.
func Check(t testing.TB, message string, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf(message, err)
	}
}

// Assert fails the test with the `message` and args `args` if `condition` is false.
func Assert(t testing.TB, message string, condition bool, args ...interface{}) {
	t.Helper()
	if !condition {
		t.Fatalf(message, args...)
	}
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

var compareOptions = cmp.Options{
	cmpopts.IgnoreUnexported(scores.Track{}),
	cmp.Comparer(func(x, y time.Time) bool {
		// since some databases don't have the same time precision
		// as go's time.Time we will only compare the Unix timestamps
		// ignoring rounding errors
		return abs(x.Unix()-y.Unix()) <= 1
	}),
}

// Compare fails the test if `first` does not deep equal `second` and fails
// with `message` and diff(first, second) as arg.
func Compare(t testing.TB, message string, first, second interface{}) {
	t.Helper()
	if equal := cmp.Equal(first, second, compareOptions); !equal {
		t.Fatalf(message, cmp.Diff(first, second, compareOptions))
	}
}

// Equal compares `expected` and `actual` and fails with message `message` and args `expected` and `actual` if they are not equal.
func Equal(t testing.TB, message string, expected, actual interface{}) {
	t.Helper()
	if expected != actual {
		t.Fatalf(message, expected, actual)
	}
}
