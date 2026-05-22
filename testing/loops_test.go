package testing

import (
	"my_language/my_language"
	"testing"
)

// while loops tests
func TestWhileLoop(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Simple while loop",
			input: `
                var testResult = 0 
                while testResult < 5 { 
                    testResult = testResult + 1 
                }
            `,
			expected: 5,
		},
		{
			name: "While loop with multiplication",
			input: `
                var testResult = 1 
                while testResult < 100 { 
                    testResult = testResult * 2 
                }
            `,
			expected: 128,
		},
		{
			name: "While loop with division",
			input: `
                var testResult = 100 
                while testResult > 25 { 
                    testResult = testResult / 2 
                }
            `,
			expected: 25,
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

func TestWhileLoopWithLogical(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "While loop with logical AND",
			input: `
				var testResult = 0 
				while testResult < 5 and testResult != 3 { 
					testResult = testResult + 1 
				}`,
			expected: 3,
		},
		{
			name: "While loop with logical OR",
			input: `
				var testResult = 0 
				while testResult < 5 or testResult == 3 { 
					testResult = testResult + 1 
				}`,
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
