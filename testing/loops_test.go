package testing

import (
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
                while (testResult < 5) { 
                    testResult = testResult + 1 
                }
            `,
			expected: 5,
		},
		{
			name: "While loop with multiplication",
			input: `
                var testResult = 1 
                while (testResult < 100) { 
                    testResult = testResult * 2 
                }
            `,
			expected: 128,
		},
		{
			name: "While loop with division",
			input: `
                var testResult = 100 
                while (testResult > 25) { 
                    testResult = testResult / 2 
                }
            `,
			expected: 25,
		},
		{
			name: "While loop with break",
			input: `
                var testResult = 0 
				var i = 0
                while (i < 25) { 
                    i++
					if (i >= 10){
						break
					}
					testResult += i
                }

            `,
			expected: 45,
		},
		{
			name: "While loop with continue",
			input: `
                var testResult = 0 
				var i = 0
                while (i < 12) { 
                    i++
					if (i == 10){
						continue
					}
					testResult += i
                }

            `,
			expected: 68,
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
				while (testResult < 5 and testResult != 3) { 
					testResult = testResult + 1 
				}`,
			expected: 3,
		},
		{
			name: "While loop with logical OR",
			input: `
				var testResult = 0 
				while (testResult < 5 or testResult == 3) { 
					testResult = testResult + 1 
				}`,
			expected: 5,
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

// for loop tests
func TestForLoop(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Simple for loop",
			input: `
                var testResult = 0 
                for (var i = 1; i <= 3; i = i + 1) {
					testResult += 1
				}
            `,
			expected: 3,
		},
		{
			name: "For loop with multiplication",
			input: `
                var testResult = 1 
                for (var i = 1; i <= 3; i = i + 1) {
                    testResult = testResult * 2 
                }
            `,
			expected: 8,
		},
		{
			name: "For loop with inc ++",
			input: `
                var testResult = 1 
                for (var i = 1; i <= 3; i++) {
                    testResult = testResult * 2 
                }
            `,
			expected: 8,
		},
		{
			name: "For loop with dec --",
			input: `
                var testResult = 1 
                for (var i = 3; i > 0; i--) {
                    testResult = testResult * 2 
                }
            `,
			expected: 8,
		},
		{
			name: "For loop with continue",
			input: `
			 	var testResult = 0
				for (var i = 1; i < 4; i = i + 1) {
					if (i == 3) {
    					continue
					}
					testResult += i
				}

			`,
			expected: 3,
		},
		{
			name: "For loop with break",
			input: `
				var testResult = 0
				for (var i = 1; i < 10; i = i + 1) {
					if (i == 3){
						break
					}
					testResult += i
				}
			`,
			expected: 3,
		},
		{
			name: "For loop with brak and continue",
			input: `
				var testResult = 0
				for (var i = 0; i < 10; i = i + 1) {
					if (i == 3) {
						continue
					}
					if (i == 6) {
						break
					}
					testResult += i
				}
			`,
			expected: 12,
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
