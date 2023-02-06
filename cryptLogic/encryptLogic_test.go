package cryptLogic

import (
	"testing"
)

func TestEncodecod(t *testing.T) {

	testCases := []struct {
		input    string
		expected string
	}{
		{"wewe", "2(we)"},
		{"ssffgg", "2s2f2g"},
		{"qqqq", "4q"},
	}

	for _, test := range testCases {

		t.Run(test.input, func(t *testing.T) {

			for _, test = range testCases {

				got := Encode(test.input)
				if got != test.expected {
					t.Errorf("Expected response body %q but got %q", test.expected, Encode(test.input))
				}
			}
		})
	}
}
