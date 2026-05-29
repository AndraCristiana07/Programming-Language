package testing

import (
	"testing"
)

func TestErrorHandlingProperties(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Catch Undefined Variable Type",
			input: `
				var testResult = "failed";
				try {
					print nonExistentVar;
				} catch (err) {
					testResult = err.type;
				}
			`,
			expected: "ReferenceError",
		},
		{
			name: "Catch Division by Zero Message",
			input: `
				var testResult = "failed";
				try {
					var x = 10 / 0;
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "Cannot divide a number by zero",
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
