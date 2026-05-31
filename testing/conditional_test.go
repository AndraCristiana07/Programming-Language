package testing

import (
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
                if (testResult < 5) { 
                    testResult = testResult + 1 
                }
            `,
			expected: 1,
		},
		{
			name: "if else true",
			input: `
                var testResult = 0 
                if (testResult < 5) { 
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
                if (testResult > 5) { 
                    testResult = testResult + 1 
                } else {
					testResult = testResult + 3
				}
            `,
			expected: 3,
		},
		{
			name: "if elif else falling on elif",
			input: `
				var cond = 20
				var testResult = 0
				if (cond == 10){
					testResult = 10
				} elif (cond == 6) {
					testResult = 123
				} elif (cond == 20) {
					testResult = 11
				} elif (cond == 120) {
					testResult = 111
				} else {
					testResult =100
				}
			`,
			expected: 11,
		},
		{
			name: "if elif else falling on else",
			input: `
				var cond = 1
				var testResult = 0
				if (cond == 10){
					testResult = 10
				} elif (cond == 6) {
					testResult = 123
				} elif (cond == 20) {
					testResult = 11
				} elif (cond == 120) {
					testResult = 111
				} else {
					testResult =100
				}
			`,
			expected: 100,
		},
		{
			name: "if elif else falling on if",
			input: `
				var cond = 10
				var testResult = 0
				if (cond == 10){
					testResult = 10
				} elif (cond == 6) {
					testResult = 123
				} elif (cond == 20) {
					testResult = 11
				} elif (cond == 120) {
					testResult = 111
				} else {
					testResult =100
				}
			`,
			expected: 10,
		},
		{
			name: "if elif else falling on first elif",
			input: `
				var cond = 6
				var testResult = 0
				if (cond == 10){
					testResult = 10
				} elif (cond == 6) {
					testResult = 123
				} elif (cond == 20) {
					testResult = 11
				} elif (cond == 120) {
					testResult = 111
				} else {
					testResult =100
				}
			`,
			expected: 123,
		},
		{
			name: "if elif falling on elif",
			input: `
				var cond = 6
				var testResult = 0
				if (cond == 10){
					testResult = 10
				} elif (cond == 6) {
					testResult = 123
				} elif (cond == 20) {
					testResult = 11
				} elif (cond == 120) {
					testResult = 111
				}
			`,
			expected: 123,
		},
		{
			name: "if elif falling on none",
			input: `
				var cond = 100
				var testResult = 0
				if (cond == 10){
					testResult = 10
				} elif (cond == 6) {
					testResult = 123
				} elif (cond == 20) {
					testResult = 11
				} elif (cond == 120) {
					testResult = 111
				}
			`,
			expected: 0,
		},
		{
			name: "Ternary operation falling on if",
			input: `
				var cond = true
				var testResult = 10 if cond else 5
			`,
			expected: 10,
		},
		{
			name: "Ternary operation falling on else",
			input: `
				var cond = false
				var testResult = 10 if cond else 5
			`,
			expected: 5,
		},
		{
			name: "Nested ternary operation",
			input: `
				var num = -5
				var testResult = "Positive" if num > 0 else "Negative" if num < 0 else "Zero"
			`,
			expected: "Negative",
		},
		{
			name: "Direct variable initialization in if condition",
			input: `
				var testResult = 0
				if (var t = 95; t > 90) {
					testResult += t
				} else {
					testResult = 0
				}
			`,
			expected: 95,
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
				while (i < 5) {
					if (testResult < 5) { 
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
