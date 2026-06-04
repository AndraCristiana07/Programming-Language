package testing

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Array indexing",
			input: `
                var myArr = [1,4]
				var testResult = myArr[0]
            `,
			expected: 1,
		},
		{
			name: "Array negative indexing",
			input: `
                var myArr = [1,4,6,2]
				var testResult = myArr[-2]
            `,
			expected: 6,
		},
		{
			name: "Array slicing from beginning",
			input: `
                var myArr = [1,4,6,2]
				var testResult = myArr[:2]
            `,
			expected: &[]any{1, 4},
		},
		{
			name: "Array slicing to end",
			input: `
                var myArr = [1,4,6,2]
				var testResult = myArr[2:]
            `,
			expected: &[]any{6, 2},
		},
		{
			name: "Array slicing with range",
			input: `
                var myArr = [1,4,6,2,7]
				var testResult = myArr[1:4]
            `,
			expected: &[]any{4, 6, 2},
		},

		{
			name: "Array slicing with range and step",
			input: `
                var myArr = [1,4,6,2,7]
				var testResult = myArr[1:4:2]
            `,
			expected: &[]any{4, 2},
		},
		{
			name: "Array slicing for reverse",
			input: `
                var myArr = [1,4,6]
				var testResult = myArr[::-1]
            `,
			expected: &[]any{6, 4, 1},
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
