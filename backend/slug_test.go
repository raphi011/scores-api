package scores

import "testing"

var tests = []struct {
	Input    string
	Expected string
}{
	{
		Input:    "abc def",
		Expected: "abc-def",
	},
	{
		Input:    "abc      def",
		Expected: "abc-def",
	},
	{
		Input:    "AMATEUR TOUR",
		Expected: "amateur-tour",
	},
	{
		Input:    "ABV Tour AMATEUR 1",
		Expected: "abv-tour-amateur-1",
	},
}

func TestSlug(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.Expected, func(t *testing.T) {
			result := Sluggify(tt.Input)
			if result != tt.Expected {
				t.Fatalf("sluggify(%s) should be %q, got: %q", tt.Input, tt.Expected, result)
			}
		})
	}

}
