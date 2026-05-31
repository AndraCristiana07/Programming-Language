package testing

import "testing"

func TestSwitchCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Simple switch case to match value",
			input: `
				var testResult = 10
				switch (testResult) {
					case 4:
						testResult = 7
					case 7:
						testResult = 6
					case 10:
						testResult = 1
				}
			`,
			expected: 1,
		},
		{
			name: "Simple switch case with default case",
			input: `
				var testResult = 10
				switch (testResult) {
					case 4:
						testResult = 7
					case 7:
						testResult = 6
					case 10:
						testResult = 1
					default:
						testResult = 3
				}
			`,
			expected: 1,
		},
		{
			name: "Switch case falling on default case",
			input: `
				var testResult = 10
				switch (testResult) {
					case 1:
						testResult = 6
					case 3:
						testResult = 5
					default:
						testResult = 20	
				}
			`,
			expected: 20,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if result != tc.expected {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}
