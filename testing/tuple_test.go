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
		{
			name: "Tuple unpacking in for...in",
			input: `
				var dataMatrix = ((10, "A"), (20, "B"), (30, "C"))
				var testResult = 0
				for ((value, label) in dataMatrix) {
					testResult += value
				}
			`,
			expected: 60,
		},
		{
			name: "Tuple with slicing",
			input: `
				var dataMatrix = (1,4,7,4,1,7,3)
				var testResult = dataMatrix[1:3]
			`,
			expected: &ast.Tuple{Elements: []any{4, 7}},
		},
		{
			name: "Tuple with slicing from beginning",
			input: `
				var dataMatrix = (1,4,7,4,1,7,3)
				var testResult = dataMatrix[:3]
			`,
			expected: &ast.Tuple{Elements: []any{1, 4, 7}},
		},
		{
			name: "Tuple with negative index for last",
			input: `
				var dataMatrix = (1,4,7,4,1,7,3)
				var testResult = dataMatrix[-1]
			`,
			expected: 3,
		},
		{
			name: "Tuple with negative index for -2",
			input: `
				var dataMatrix = (1,4,7,4,1,7,3)
				var testResult = dataMatrix[-2]
			`,
			expected: 7,
		},
		{
			name: "Tuple with slicing with negative with range ",
			input: `
				var dataMatrix = (1,4,7,4,1,7,3)
				var testResult = dataMatrix[-3:-1]
			`,
			expected: &ast.Tuple{Elements: []any{1, 7}},
		},
		{
			name: "Tuple with slicing with negative ",
			input: `
				var dataMatrix = (1,4,7,4,1,7,3)
				var testResult = dataMatrix[:-3]
			`,
			expected: &ast.Tuple{Elements: []any{1, 4, 7, 4}},
		},
		{
			name: "Tuple with slicing and step ",
			input: `
				var dataMatrix = (1,4,7,4,1,7,3)
				var testResult = dataMatrix[1:5:2]
			`,
			expected: &ast.Tuple{Elements: []any{4, 4}},
		},
		{
			name: "Tuple reverse with slicing ",
			input: `
				var dataMatrix = (1,4,7,4,1,7,3)
				var testResult = dataMatrix[::-1]
			`,
			expected: &ast.Tuple{Elements: []any{3, 7, 1, 4, 7, 4, 1}},
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
