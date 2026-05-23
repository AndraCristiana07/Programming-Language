package testing

import (
	"my_language/my_language"
	"testing"
)

func TestIf(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Simple if",
			input: `
                var testResult = 0 
                if testResult < 5 { 
                    testResult = testResult + 1 
                }
            `,
			expected: 1,
		},
		{
			name: "if else true",
			input: `
                var testResult = 0 
                if testResult < 5 { 
                    testResult = testResult + 1 
                } else {
					testResult = testResult + 3
				}
            `,
			expected: 1,
		},
		{
			name: "if else false",
			input: `
                var testResult = 0 
                if testResult > 5 { 
                    testResult = testResult + 1 
                } else {
					testResult = testResult + 3
				}
            `,
			expected: 3,
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

func TestIfWithLoops(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "If and while",
			input: `
                var testResult = 2 
				var i = 0
				while i < 5 {
					if testResult < 5 { 
						testResult = testResult + 1 
					}
					i++
				}
            `,
			expected: 5,
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
