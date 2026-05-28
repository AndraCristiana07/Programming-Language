package testing

import (
	"reflect"
	"testing"
)

func TestMapLiteralsAndBuiltins(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Construct and fetch keys from map",
			input: `
                var user = {"name": "Alice", "age": 30}
                var testResult = keys(user)
            `,
			expected: &[]any{"name", "age"},
		},
		{
			name: "Construct and fetch values from map",
			input: `
                var user = {"name": "Alice", "age": 30}
                var testResult = values(user)
            `,
			expected: &[]any{"Alice", 30},
		},
		{
			name: "Clear drops all elements inside map",
			input: `
                var data = {"status": "ok"}
                clear(data)
                var testResult = data
            `,
			expected: &map[string]any{},
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
