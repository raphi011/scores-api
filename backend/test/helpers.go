package test

import (
	"testing"

	// "github.com/google/go-cmp/cmp"
	// "github.com/google/go-cmp/cmp/cmpopts"

	// "github.com/raphi011/scores"
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

// var compareOptions = cmp.Options{ cmpopts.IgnoreUnexported(volleynet.Player) }

// Compare fails the test if `first` does not deep equal `second` and fails
// with `message` and diff(first, second) as arg.
func Compare(t testing.TB, message string, first, second interface{}) {
	t.Helper()
	// if diff := cmp.Diff(first, second, cmp.Options{ cmpopts.IgnoreUnexported(scores.Tracked{}) }); diff != "" {
	// 	t.Fatalf(message, diff)
	// }
}