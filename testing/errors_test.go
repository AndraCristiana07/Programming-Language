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
		{
			name: "Catch int Unsupported type for conversion",
			input: `
				var testResult = (1,5)
				try {
					var x = int(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "Unsupported type for conversion",
		},
		{
			name: "Catch min not enough args",
			input: `
				var testResult = 2
				try {
					var x = min()
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "min() expects at least 1 argument",
		},
		{
			name: "Catch max require numbers",
			input: `
				var testResult = 2
				try {
					var x = max(testResult, "aaaaa")
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "max() comparisons require numeric values",
		},
		{
			name: "Catch ord require single char",
			input: `
				var testResult = "sdc"
				try {
					var x = ord(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "ord() expects a single-character string",
		},
		{
			name: "Catch chr require ASCII int",
			input: `
				var testResult = "sdc"
				try {
					var x = chr(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "chr() expects an integer ASCII code (0-255)",
		},
		{
			name: "Catch range require numbers message",
			input: `
				var testResult = "sdc"
				try {
					var x = range(1, testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "range() arguments must be integers",
		},
		{
			name: "Catch range require 1/2 args",
			input: `
				var testResult = "sdc"
				try {
					var x = range(1, testResult, 5)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "range() expects 1 or 2 arguments",
		},
		{
			name: "Catch set requires array",
			input: `
				var testResult = "sdc"
				try {
					var x = set(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "set() expects an array as an argument",
		},
		{
			name: "Catch lower requires string",
			input: `
				var testResult = [1,4,5]
				try {
					var x = lower(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "lower() expects a string",
		},
		{
			name: "Catch upper requires string",
			input: `
				var testResult = [1,4,5]
				try {
					var x = upper(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "upper() expects a string",
		},
		{
			name: "Catch pop requires array",
			input: `
				var testResult = "sss"
				try {
					var x = pop(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "pop() expects an array reference",
		},
		{
			name: "Catch pop on empty array",
			input: `
				var testResult = []
				try {
					var x = pop(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "pop from empty array",
		},
		{
			name: "Catch find on string requires string",
			input: `
				var testResult = "hello"
				try {
					var x = find(testResult, 3)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "find() on a string expects a string target",
		},
		{
			name: "Catch find requires string or array",
			input: `
				var testResult = 51
				try {
					var x = find(testResult, 3)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "find() expects a string or array",
		},
		{
			name: "Catch reverse requires array",
			input: `
				var testResult = 51
				try {
					var x = reverse(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "reverse() expects an array",
		},
		{
			name: "Catch all requires array, tuple, or map",
			input: `
				var testResult = 51
				try {
					var x = all(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "all() expects an array, tuple, or map, got int",
		},
		{
			name: "Catch filter requires function",
			input: `
				var testResult = 51
				var arr = [4,6,1]
				try {
					var x = filter(testResult, arr)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "First argument to filter() must be a callable function",
		},
		{
			name: "Catch filter requires array/tuple on second arg",
			input: `
				func myfct(x) {
					return x % 2
				}

				var testResult = 4
				try {
					var x = filter(myfct, testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "Second argument to filter() must be an array or tuple, got int",
		},
		{
			name: "Catch float requires number string",
			input: `
				var testResult = "hello"
				try {
					var x = float(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "Cannot convert string hello to float",
		},
		{
			name: "Catch float unsupported",
			input: `
				var testResult = [1,5,6]
				try {
					var x = float(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "Unsupported type for float() conversion",
		},
		{
			name: "Catch list requires string/array",
			input: `
				var testResult = 531
				try {
					var x = list(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "list() expects a string or an array",
		},
		{
			name: "Catch map requires function",
			input: `
				var testResult = 531
				var arr = [4, 6, 1]
				try {
					var x = map(testResult, arr)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "First argument to map() must be a callable function",
		},
		{
			name: "Catch map requires array/tuple on second arg",
			input: `
				func myfct(x) {
					return x % 2
				}

				var testResult = 4
				try {
					var x = map(myfct, testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "Second argument to map() must be an array or tuple, got int",
		},
		{
			name: "Catch round requires number",
			input: `
				var testResult = [4, 6, 1]
				try {
					var x = round(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "round() expects a number",
		},
		{
			name: "Catch sorted requires array/tuple/map",
			input: `
				var testResult = 6
				try {
					var x = sorted(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "sorted() expects an array, tuple, or map, got int",
		},
		{
			name: "Catch sorted requires seq of int/strings",
			input: `
				var testResult = [true, false, false]
				try {
					var x = sorted(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "sorted() only supports sequences of integers or strings",
		},
		{
			name: "Catch sum requires array/tuple",
			input: `
				var testResult = 53
				try {
					var x = sum(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "sum() expects an array or a tuple, got int",
		},
		{
			name: "Catch split requires strings",
			input: `
				var testResult = 53
				try {
					var x = split(testResult, 6)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "split() expects two string arguments",
		},
		{
			name: "Catch insert requires array and int",
			input: `
				var testResult = 53
				try {
					var x = insert(testResult, 6,1)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "insert() expects an array and an integer index",
		},
		{
			name: "Catch strip requires array and int",
			input: `
				var testResult = 53
				try {
					var x = strip(testResult)
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "strip() expects a string argument",
		},
		{
			name: "Catch startswith requires strings",
			input: `
				var testResult = 53
				try {
					var x = startswith(testResult, "str")
				} catch (err) {
					testResult = err.message;
				}
			`,
			expected: "startswith() expects two string arguments",
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
