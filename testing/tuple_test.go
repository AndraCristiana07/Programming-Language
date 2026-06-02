package testing

import (
	"my_language/ast"
	"reflect"
	"testing"
)

func TestTuple(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Tuple indexing",
			input: `
                var myTuple = (10, "apple")
				var testResult = myTuple[0]
            `,
			expected: 10,
		},
		{
			name: "Tuple assign",
			input: `
                var testResult = (10, "apple")
            `,
			expected: &ast.Tuple{Elements: []any{10, "apple"}},
		},
		{
			name: "Tuple unpack",
			input: `
				var fruits = ("apple", "banana", "cherry")
				var green = ""
				var yellow = ""
				var red = ""
				(green, yellow, red) = fruits 
				var testResult = green
            `,
			expected: "apple",
		},
		{
			name: "Tuple reassign workaround",
			input: `
				var testResult = ("apple", "banana", "cherry")
				var y = list(testResult)
				y[1] = "kiwi"
				testResult = tuple(y)
            `,
			expected: &ast.Tuple{Elements: []any{"apple", "kiwi", "cherry"}},
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
