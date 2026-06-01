package testing

import (
	"reflect"
	"testing"
)

func TestInterface(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput any
		expectError    bool
	}{
		{
			name: "Simple one method interface",
			input: `
				interface User {
					GetName();
				}

				struct Student {
					name;
				}
				
				func (s Student) GetName() {
					return s.name
				}
				var st = Student {name: "Anna"}
				var testResult = st.GetName()
			`,
			expectedOutput: "Anna",
			expectError:    false,
		},
		{
			name: "Interface with 2 methods",
			input: `
				interface Shape {
					Area();
					Perimeter();
				}

				struct Circle { radius; }

				func (c Circle) Area() { return 3 * c.radius * c.radius }
				func (c Circle) Perimeter() { return 6 * c.radius }

				struct Rectangle { length; width; }
				func (r Rectangle) Area() { return r.length * r.width }
				func (r Rectangle) Perimeter() { return 2 * (r.length + r.width) }

				var c = Circle { radius: 5 }
				var r = Rectangle { length: 4, width: 5 }

				var testResult = c.Area() + r.Area()
			`,
			expectedOutput: 95,
			expectError:    false,
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
