package testing

import (
	"reflect"
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
		{
			name: "Pointer on array item",
			input: `
                var numbers = [1, 2, 3]
    			var ptr = &numbers[1]
				var testResult = *ptr
            `,
			expected: 2,
		},
		{
			name: "Pointer on array",
			input: `
                var numbers = [1, 2, 3]
    			var ptr = &numbers[1]
				*ptr = 4
				var testResult = numbers
            `,
			expected: &[]any{1, 4, 3},
		},
		{
			name: "Double Pointer",
			input: `
                var x = 1
				var ptr = &x
				var pptr = &ptr 

				var testResult = **pptr 
            `,
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' missing")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
