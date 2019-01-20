package test

import (
	"testing"

	// "github.com/google/go-cmp/cmp"
	// "github.com/google/go-cmp/cmp/cmpopts"

	// "github.com/raphi011/scores"
)


func Check(t testing.TB, message string, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf(message, err)
	}
}

func Assert(t testing.TB, message string, condition bool, args ...interface{}) {
	t.Helper()
	if condition {
		t.Fatalf(message, args...)
	}
}

// var compareOptions = cmp.Options{ cmpopts.IgnoreUnexported(volleynet.Player) }

func Compare(t testing.TB, message string, first, second interface{}) {
	t.Helper()
	// if diff := cmp.Diff(first, second, cmp.Options{ cmpopts.IgnoreUnexported(scores.Tracked{}) }); diff != "" {
	// 	t.Fatalf(message, diff)
	// }
}