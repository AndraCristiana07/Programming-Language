package testing

import (
	"testing"
)

func TestFunction(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Simple function ",
			input: `
                var testResult = 0 
                func addOne(x) {
					return x + 1
				}
				testResult = addOne(5)
            `,
			expected: 6,
		},
		{
			name: "Function with multiple parameters",
			input: `
				var testResult = 0
				func add(x, y) {
					return x + y
				}
				testResult = add(5, 10)
			`,
			expected: 15,
		},
		{
			name: "Function with nested function",
			input: `
				var testResult = 0
				func add(x, y) {
					func addInner(a, b) {
						return a + b
					}
					return addInner(x, y)
				}
				testResult = add(5, 10)
			`,
			expected: 15,
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

func TestLambdas(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Simple lambda function",
			input: `
				var x = lambda a : a + 10
				var testResult = x(5)
			`,
			expected: 15,
		},
		{
			name: "Lambda function",
			input: `
				func myfunc(n) {
					return lambda a : a * n
				}
				var double = myfunc(2)
				var testResult = double(11)
			`,
			expected: 22,
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
