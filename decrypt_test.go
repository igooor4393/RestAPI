package main

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestDecrypt(t *testing.T) {

	testCases := []struct {
		input    string
		expected string
	}{
		{`{"decrypt": "A2(A2BC3(F4B)D)E"}`, "Decrypted string: AABBCFBBBBFBBBBFBBBBDABBCFBBBBFBBBBFBBBBDE"},
		{`{"decrypt": "A2(BB)"}`, "Decrypted string: ABBBB"},
		{`{"decrypt": "A2(BB)DD2(A)"}`, "Decrypted string: ABBBBDDAA"},
		//{`{"decrypt": "A2(BB)"}`, "Decrypted string: ABBBB"},
	}

	for _, test := range testCases {

		t.Run(test.input, func(t *testing.T) {
			reqBody := []byte(test.input)
			req := httptest.NewRequest("POST", "http://localhost/:8080/decrypt", bytes.NewBuffer(reqBody))

			res := httptest.NewRecorder()

			decrypt(res, req)

			got := res.Body.String()
			if got != test.expected {
				t.Errorf("Expected response body %q but got %q", test.expected, res.Body.String())
			}
		})
	}

}
