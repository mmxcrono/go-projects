package db

import (
	"testing"
)

func TestGetClientProfileById(t *testing.T) {
	testCases := []struct{
		name string
		input string
		expected bool
	}{
		{
			name: "Get existing user",
			input: "user1",
			expected: true,
		},
		{
			name: "User not found",
			input: "user123",
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, ok := Database[testCase.input]

			actual := ok

			if actual != testCase.expected {
				expectation := "exist"

				if !testCase.expected {
					expectation = "not exist"
				}

				t.Errorf("Failed expecting user to %s", expectation)
			}
		})
	}
}