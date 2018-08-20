package volleynet

import (
	"os"
	"testing"
)

func TestParseUniqueWriteCode(t *testing.T) {
	response, _ := os.Open("testdata/signup.html")

	expected := "0.77187700 1534784156"

	code, err := parseUniqueWriteCode(response)

	if err != nil {
		t.Errorf("parseUniqueWriteCode() err: %s", err)
	}

	if code != expected {
		t.Errorf("parseUniqueWriteCode() want: %s, got: %s ", expected, code)
	}
}
