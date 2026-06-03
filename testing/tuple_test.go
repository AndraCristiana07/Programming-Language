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
		{
			name: "Tuple concatenation",
			input: `
				var tuple1 = ("apple", "banana", "cherry")
				var tuple2 = (1,3,5)
				var testResult = tuple1 + tuple2
            `,
			expected: &ast.Tuple{Elements: []any{"apple", "banana", "cherry", 1, 3, 5}},
		},
		{
			name: "Tuple multiply",
			input: `
				var tuple1 = ("apple", 2)
				var testResult = tuple1 * 3

			`,
			expected: &ast.Tuple{Elements: []any{"apple", 2, "apple", 2, "apple", 2}},
		},
		{
			name: "Tuple concatenation multiplied",
			input: `
				var tuple1 = ("apple", "pear")
				var tuple2 = (true, false)
				var testResult = (tuple1 + tuple2) * 2
			`,
			expected: &ast.Tuple{Elements: []any{"apple", "pear", true, false, "apple", "pear", true, false}},
		},
		{
			name: "Tuple multiply and concatenation after",
			input: `
				var tuple1 = ("apple", "pear")
				var tuple2 = (true, false)
				var testResult = tuple1 * 2 + tuple2
			`,
			expected: &ast.Tuple{Elements: []any{"apple", "pear", "apple", "pear", true, false}},
		},
		{
			name: "Tuple count",
			input: `
				var tuple1 = ("apple", "pear", "apple", "kiwi")
				var testResult = count(tuple1, "apple")
			`,
			expected: 2,
		},
		{
			name: "Tuple index",
			input: `
				var tuple1 = ("apple", "pear", "apple", "kiwi")
				var testResult = index(tuple1, "pear")
			`,
			expected: 1,
		},
		{
			name: "Tuple count after concatenation",
			input: `
				var tuple1 = ("apple", "pear", "apple", "kiwi")
				var tuple2 = ("pear", "apple", "kiwi")
				var mixed = tuple1 + tuple2

				var testResult = count(mixed, "apple")
			`,
			expected: 3,
		},
		{
			name: "Tuple count after multiplification",
			input: `
				var tuple1 = ("apple", "pear", "apple", "kiwi")
				var mul = tuple1 * 2

				var testResult = count(mul, "apple")
			`,
			expected: 4,
		},
		{
			name: "Tuple min",
			input: `
				var numbers = (1, 2, 3, 4, 5, 6)
				var testResult = min(numbers)
			`,
			expected: 1,
		},
		{
			name: "Tuple max",
			input: `
				var numbers = (1, 2, 3, 4, 5, 6)
				var testResult = max(numbers)
			`,
			expected: 6,
		},
		{
			name: "Tuple sum",
			input: `
				var numbers = (1, 2, 3, 4, 5, 6)
				var testResult = sum(numbers)
			`,
			expected: 21,
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
