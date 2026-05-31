package testing

import (
	"reflect"
	"testing"
)

func TestStruct(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput any
		expectError    bool
	}{
		{
			name: "Basic Declaration and Initialization",
			input: `
                struct User { name; age;}
                var testResult = User { name: "Hunter", age: 24 }
            `,
			expectedOutput: &map[string]any{
				"__type__": "User",
				"name":     "Hunter",
				"age":      24,
			},
			expectError: false,
		},
		{
			name: "Property get and set",
			input: `
                struct User { name; age; active; }
                var player = User { name: "Hunter", age: 24 }
                player.active = true;
                player.age = player.age + 1; 
                var testResult = player.name + " is now " + player.age;
            `,
			expectedOutput: `Hunter is now 25`,
			expectError:    false,
		},
		{
			name: "Struct Not Defined",
			input: `
				var testResult = FakeStruct { name: "Hunter" }
			`,
			expectError: true,
		},
		{
			name: "Invalid Field Placement",
			input: `
				struct User { name; }
				var testResult = User { name: "Hunter", cheatCode: 999 }
			`,
			expectError: true,
		},
		{
			name: "Overwriting Internal Meta Field",
			input: `
				struct User { name; }
				var testResult = User { name: "Hunter" }
				testResult.__type__ = "Hacker";
			`,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if tc.expectError && r == nil {
					t.Errorf("Expected code to panic with a TypeError, but it completed successfully.")
				}
				if !tc.expectError && r != nil {
					t.Errorf("Unexpected panic: %v", r)
				}
			}()

			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expectedOutput) {
				t.Errorf("Data State Mismatch!\nExpected: %#v\nGot:      %#v",
					tc.expectedOutput, result)
			}

		})
	}
}
