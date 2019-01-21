package scrape

import (
	"os"
	"testing"
)

func TestParseUniqueWriteCode(t *testing.T) {
	response, _ := os.Open("../testdata/signup.html")

	expected := "0.77187700 1534784156"

	code, err := UniqueWriteCode(response)

	if err != nil {
		t.Errorf("UniqueWriteCode() err: %s", err)
	}

	if code != expected {
		t.Errorf("UniqueWriteCode() want: %s, got: %s ", expected, code)
	}
}
