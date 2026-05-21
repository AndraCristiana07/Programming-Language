package testing

import (
	"my_language/my_language"
	"testing"
)

func TestAddition(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Basic positive addition",
			input:    "var testResult = 5 + 10",
			expected: 15,
		},
		{
			name:     "Chained additions",
			input:    "var testResult = 1 + 2 + 3",
			expected: 6,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tokens := my_language.Lex(tc.input)
			parser := my_language.NewParser(tokens)
			program := parser.Parse()

			env := my_language.NewEnvironment()

			my_language.Eval(program, env)

			// check the saved result
			result, ok := env.Lookup("testResult")

			if !ok {
				t.Fatalf("'testResult' variable was not created by the engine")
			}

			if result != tc.expected {
				t.Errorf("For %q: expected %v (%T), but got %v (%T)",
					tc.input, tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestSubstraction(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Basic positive substraction",
			input:    "var testResult = 10 - 7",
			expected: 3,
		},
		{
			name:     "Chained substraction",
			input:    "var testResult = 10 - 5 - 2",
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

			// check the saved result
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("'testResult' variable was not created by the engine")
			}

			if result != tc.expected {
				t.Errorf("For %q: expected %v (%T), but got %v (%T)",
					tc.input, tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestMultiplication(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Basic positive multiplication",
			input:    "var testResult = 10 * 7",
			expected: 70,
		},
		{
			name:     "Chained multiplication",
			input:    "var testResult = 10 * 2 * 10",
			expected: 200,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tokens := my_language.Lex(tc.input)
			parser := my_language.NewParser(tokens)
			program := parser.Parse()

			env := my_language.NewEnvironment()

			my_language.Eval(program, env)

			// check the saved result
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("'testResult' variable was not created by the engine")
			}

			if result != tc.expected {
				t.Errorf("For %q: expected %v (%T), but got %v (%T)",
					tc.input, tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestDivision(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Basic positive division",
			input:    "var testResult = 10 / 2",
			expected: 5,
		},
		{
			name:     "Chained division",
			input:    "var testResult = 20 / 10 / 2",
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tokens := my_language.Lex(tc.input)
			parser := my_language.NewParser(tokens)
			program := parser.Parse()

			env := my_language.NewEnvironment()

			my_language.Eval(program, env)

			// check the saved result
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("'testResult' variable was not created by the engine")
			}

			if result != tc.expected {
				t.Errorf("For %q: expected %v (%T), but got %v (%T)",
					tc.input, tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestMultipleOperations(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Addition and substraction",
			input:    "var testResult = 10 + 5 -2",
			expected: 13,
		},
		{
			name:     "Simple mutliplication and addition",
			input:    "var testResult = 20 * 2 + 5",
			expected: 45,
		},
		{
			name:     "Simple addition and mutliplication",
			input:    "var testResult = 5 + 20 * 2",
			expected: 45,
		},
		{
			name:     "Multiple addition and mutliplication",
			input:    "var testResult = 5 + 20 * 2 + 6 * 2",
			expected: 57,
		},
		{
			name:     "Simple addition and division",
			input:    "var testResult = 5 + 20 / 2",
			expected: 15,
		},
		{
			name:     "Simple multiplication and substraction",
			input:    "var testResult = 20 * 2 - 10",
			expected: 30,
		},
		{
			name:     "Simple division and substraction",
			input:    "var testResult = 20 / 2 - 10",
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tokens := my_language.Lex(tc.input)
			parser := my_language.NewParser(tokens)
			program := parser.Parse()

			env := my_language.NewEnvironment()

			my_language.Eval(program, env)

			// check the saved result
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("'testResult' variable was not created by the engine")
			}

			if result != tc.expected {
				t.Errorf("For %q: expected %v (%T), but got %v (%T)",
					tc.input, tc.expected, tc.expected, result, result)
			}
		})
	}
}
