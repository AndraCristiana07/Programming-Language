package testing

import (
	"my_language/my_language"
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
			tokens := my_language.Lex(tc.input)
			parser := my_language.NewParser(tokens)
			program := parser.Parse()

			env := my_language.NewEnvironment()
			my_language.Eval(program, env)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if result != tc.expected {
				t.Errorf("Expected %v (%T), got %v (%T) for loop execution block",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}
