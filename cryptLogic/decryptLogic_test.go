package cryptLogic

import (
	"testing"
)

func TestDecod(t *testing.T) {

	testCases := []struct {
		input    string
		expected string
	}{
		{"A2(A2BC3(F4B)D)E", "AABBCFBBBBFBBBBFBBBBDABBCFBBBBFBBBBFBBBBDE"},
		{"A2(BB)", "ABBBB"},
	}

	for _, test := range testCases {

		t.Run(test.input, func(t *testing.T) {

			for _, test = range testCases {

				got := Decod(test.input)
				if got != test.expected {
					t.Errorf("Expected response body %q but got %q", test.expected, Decod(test.input))
				}
			}
		})
	}
}
