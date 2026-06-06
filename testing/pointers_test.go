package testing

import (
	"testing"
)

func TestPointers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Simple pointer",
			input: `
                var x = 5
    			var ptr = &x 
				var testResult = *ptr
            `,
			expected: 5,
		},
		{
			name: "Pointer mutation",
			input: `
                var x = 5
    			var ptr = &x 
				*ptr = 7
				var testResult = *ptr
            `,
			expected: 7,
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
