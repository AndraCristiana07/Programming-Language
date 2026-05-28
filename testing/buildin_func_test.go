package testing

import (
	"fmt"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLen(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Len on simple arr",
			input: `
				var arr = [1, 2, 4]
				var testResult = len(arr)
			`,
			expected: 3,
		},
		{
			name: "Len on comp arr",
			input: `
				var arr = [[1,3], [3,4], [5,1]]
				var testResult = len(arr[0])
			`,
			expected: 2,
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

func TestAppend(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    any
		shouldPanic bool
	}{
		{
			name: "Append on simple arr",
			input: `
                var arr = [1, 2, 3]
                var testResult = append(arr, 5)
            `,
			expected:    &[]any{1, 2, 3, 5},
			shouldPanic: false,
		},
		{
			name: "Append on comp arr mismatch should fail",
			input: `
                var arr = [[1,3], [2,6], [5,7]]
                var testResult = append(arr, 5) 
            `,
			expected:    nil,
			shouldPanic: true,
		},
		{
			name: "Append on comp arr mismatch should fail",
			input: `
                var arr = [[1,3], [2,6], [5,7]]
                var testResult = append(arr, [5]) 
            `,
			expected: &[]any{
				&[]any{1, 3},
				&[]any{2, 6},
				&[]any{5, 7},
				&[]any{5},
			},
			shouldPanic: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				r := recover()
				if tc.shouldPanic && r == nil {
					t.Errorf("Expected code to panic with a TypeError, but it completed successfully.")
				}
				if !tc.shouldPanic && r != nil {
					t.Errorf("Unexpected panic: %v", r)
				}
			}()

			env := runInterpreter(tc.input)

			if !tc.shouldPanic {
				result, ok := env.Lookup("testResult")
				if !ok {
					t.Fatalf("Variable 'testResult' was missing")
				}

				if !reflect.DeepEqual(result, tc.expected) {
					t.Errorf("Expected %v (%T), got %v (%T)", tc.expected, tc.expected, result, result)
				}
			}
		})
	}
}

func TestIntCast(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    any
		shouldPanic bool
	}{
		{
			name: "Int cast on str",
			input: `
                var testResult = "6"
				testResult = int(testResult)
            `,
			expected:    6,
			shouldPanic: false,
		},
		{
			name: "Int cast on str that's not nr",
			input: `
                 var testResult = "hello"
				testResult = int(testResult)
            `,
			expected:    nil,
			shouldPanic: true,
		},
		{
			name: "Int cast on bool",
			input: `
                var testResult = false
				testResult = int(testResult)
            `,
			expected:    0,
			shouldPanic: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				r := recover()
				if tc.shouldPanic && r == nil {
					t.Errorf("Expected code to panic with a TypeError, but it completed successfully.")
				}
				if !tc.shouldPanic && r != nil {
					t.Errorf("Unexpected panic: %v", r)
				}
			}()

			env := runInterpreter(tc.input)

			if !tc.shouldPanic {
				result, ok := env.Lookup("testResult")
				if !ok {
					t.Fatalf("Variable 'testResult' was missing")
				}

				if !reflect.DeepEqual(result, tc.expected) {
					t.Errorf("Expected %v (%T), got %v (%T)", tc.expected, tc.expected, result, result)
				}
			}
		})
	}
}

func TestStrCast(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    any
		shouldPanic bool
	}{
		{
			name: "Str cast on int",
			input: `
                var testResult = 6
				testResult = str(testResult)
            `,
			expected:    "6",
			shouldPanic: false,
		},
		{
			name: "Str cast on bool",
			input: `
                 var testResult = true
				testResult = str(testResult)
            `,
			expected:    "true",
			shouldPanic: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				r := recover()
				if tc.shouldPanic && r == nil {
					t.Errorf("Expected code to panic with a TypeError, but it completed successfully.")
				}
				if !tc.shouldPanic && r != nil {
					t.Errorf("Unexpected panic: %v", r)
				}
			}()

			env := runInterpreter(tc.input)

			if !tc.shouldPanic {
				result, ok := env.Lookup("testResult")
				if !ok {
					t.Fatalf("Variable 'testResult' was missing")
				}

				if !reflect.DeepEqual(result, tc.expected) {
					t.Errorf("Expected %v (%T), got %v (%T)", tc.expected, tc.expected, result, result)
				}
			}
		})
	}
}

func TestType(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Type on int",
			input: `
				var someInt = 7
				var testResult = type(someInt)
			`,
			expected: "int",
		},
		{
			name: "Type on str",
			input: `
				var someStr = "hello"
				var testResult = type(someStr)
			`,
			expected: "string",
		},
		{
			name: "Type on bool",
			input: `
				var someBool = true
				var testResult = type(someBool)
			`,
			expected: "bool",
		},
		{
			name: "Type on arr",
			input: `
				var arr = [1,5]
				var testResult = type(arr)
			`,
			expected: "array",
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

func TestBin(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Simple bin",
			input: `
				var someInt = 1
				var testResult = bin(someInt)
			`,
			expected: "0b1",
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

func TestHex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Hex test example",
			input:    "var testResult = hex(255)",
			expected: "0xff",
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
func TestOct(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Oct test example",
			input:    "var testResult = oct(255)",
			expected: "0o377",
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

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Abs with negative number",
			input:    "var testResult = abs(-10)",
			expected: 10,
		},
		{
			name:     "Abs with positive number",
			input:    "var testResult = abs(10)",
			expected: 10,
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

func TestMinMax(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Minimum of 2",
			input:    "var testResult = min(3, 7)",
			expected: 3,
		},
		{
			name:     "Maximum from 2",
			input:    "var testResult = max(10, 100)",
			expected: 100,
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

func TestOrd(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Ord of b",
			input:    `var testResult = ord("b")`,
			expected: 98,
		},
		{
			name:     "Ord of B",
			input:    `var testResult = ord("B")`,
			expected: 66,
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

func TestChr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Ord of 98",
			input:    `var testResult = chr(98)`,
			expected: "b",
		},
		{
			name:     "Ord of 66",
			input:    `var testResult = chr(66)`,
			expected: "B",
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

func TestBoolCast(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Bool on non-zero int",
			input: `
                var testResult = bool(42)
            `,
			expected: true,
		},
		{
			name: "Bool on 0",
			input: `
                var testResult = bool(0)
            `,
			expected: false,
		},
		{
			name: "Bool on string",
			input: `
                var testResult = bool("hello")
            `,
			expected: true,
		},
		{
			name: "Bool on empty string",
			input: `
                var testResult = bool("")
            `,
			expected: false,
		},
		{
			name: "Bool on array",
			input: `
                var testResult = bool([1, 2])
            `,
			expected: true,
		},
		{
			name: "Bool on empty array",
			input: `
                var testResult = bool([])
            `,
			expected: false,
		},
		{
			name: "Bool on already boolean",
			input: `
                var testResult = bool(false)
            `,
			expected: false,
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

func TestRange(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Range of 10",
			input:    `var testResult = range(1,10)`,
			expected: &[]any{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Range of 0-10",
			input:    `var testResult = range(10)`,
			expected: &[]any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "set for simple arr",
			input: `
				var arr = [1,1, 2, 4, 6, 7, 1, 2]
				var testResult = set(arr)
			`,
			expected: &[]any{1, 2, 4, 6, 7},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestLowerUpper(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Lower of HEllo",
			input:    `var testResult = lower("HEllo")`,
			expected: "hello",
		},
		{
			name:     "Upper of heLlo",
			input:    `var testResult = upper("heLlo")`,
			expected: "HELLO",
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

func TestPop(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Pop simple array and test popped",
			input: `
				var arr = [1, 4, 6]
				var testResult = pop(arr)`,
			expected: 6,
		},
		{
			name: "Pop simple array and test popped",
			input: `
				var arr = [1, 4, 6]
				pop(arr)
				var testResult = arr`,
			expected: &[]any{1, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Find in array",
			input: `
				var arr = [1, 4, 6]
				var testResult = find(arr, 4)`,
			expected: 1,
		},
		{
			name: "Find in string",
			input: `
				var findStr = "Hello"
				var testResult = find(findStr, "e")`,
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Reverse simple array",
			input: `
				var arr = [1, 4, 6]
				var testResult = reverse(arr)`,
			expected: &[]any{6, 4, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestAllAny(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "All on simple array",
			input: `
				var arr = [true, false, true, true]
				var testResult = all(arr)`,
			expected: false,
		},
		{
			name: "Any on simple array",
			input: `
				var arr = [true, false, true, true]
				var testResult = any(arr)`,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Filter even numbers",
			input: `
                func isEven(x) {
                    return x % 2 == 0
                }
                var arr = [1, 2, 3, 4, 5, 6]
                var testResult = filter(isEven, arr)`,
			expected: &[]any{2, 4, 6},
		},
		{
			name: "Filter words longer than 3 characters",
			input: `
                func isLong(s) {
                    return len(s) > 3
                }
                var words = ["go", "antlr", "code"]
                var testResult = filter(isLong, words)`,
			expected: &[]any{"antlr", "code"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Float on int",
			input: ` var testResult = float(6)
               `,
			expected: 6.0,
		},
		{
			name: "Float on string",
			input: ` var testResult = float("6")
               `,
			expected: 6.0,
		},
		{
			name: "Float on bool",
			input: ` var testResult = float(true)
               `,
			expected: 1.0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestList(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "List on simple string",
			input: ` var testResult = list("here")
               `,
			expected: &[]any{"h", "e", "r", "e"},
		},
		{
			name: "List on string",
			input: ` var testResult = list("here I am")
               `,
			expected: &[]any{"h", "e", "r", "e", " ", "I", " ", "a", "m"},
		},

		{
			name: "List on array",
			input: ` 
				var arr = [1, 2, 4, 6, 7]
				var testResult = list(arr)
               `,
			expected: &[]any{1, 2, 4, 6, 7},
		},
		{
			name: "List on comp array",
			input: ` 
				var arr = [[1,2], [3,6], [7,5], [5,2]]
				var testResult = list(arr)
               `,
			expected: &[]any{
				&[]any{1, 2},
				&[]any{3, 6},
				&[]any{7, 5},
				&[]any{5, 2},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Map squares of integers",
			input: `
                func square(x) {
                    return x * x
                }
                var nums = [1, 2, 3, 4]
                var testResult = map(square, nums)`,
			expected: &[]any{1, 4, 9, 16},
		},
		{
			name: "Map strings to uppercase",
			input: `
                func goUppercase(word) {
                    return upper(word) + "!"
                }
                var words = ["hello", "world"]
                var testResult = map(goUppercase, words)`,
			expected: &[]any{"HELLO!", "WORLD!"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Round on int",
			input: ` var testResult = round(6)
               `,
			expected: 6,
		},
		{
			name: "Round on float",
			input: ` var testResult = round(5.3)
               `,
			expected: 5,
		},
		{
			name: "Float on bool",
			input: ` var testResult = float(true)
               `,
			expected: 1.0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestSorted(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Sorting on simple arr",
			input: `
				var arr = [1,1, 2, 4, 6, 7, 1, 2]
				var testResult = sorted(arr)
			`,
			expected: &[]any{1, 1, 1, 2, 2, 4, 6, 7},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Sum on simple array",
			input: `
				var arr = [1, 3, 6]
				var testResult = sum(arr)`,
			expected: 10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Split string on -",
			input: `
				var phrase = "hello-world-here"
				var testResult = split(phrase, "-")
			`,
			expected: &[]any{"hello", "world", "here"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Insert in simple arr",
			input: `
				var arr = [1, 4, 7, 3]
				insert(arr, 2, 9)
				var testResult = arr
			`,
			expected: &[]any{1, 4, 9, 7, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestStrip(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Strip string",
			input: `
				var phrase = "   hello-world   "
				var testResult = strip(phrase)
			`,
			expected: "hello-world",
		},
		{
			name: "Lstrip on string",
			input: `
				var phrase = "   hi there left side   "
				var testResult = lstrip(phrase)
			`,
			expected: "hi there left side   ",
		},
		{
			name: "Rstrip on string",
			input: `
				var phrase = "   hi there left side   "
				var testResult = rstrip(phrase)
			`,
			expected: "   hi there left side",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestStartsEndswtih(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Startswtih string",
			input: `
				var filename = "script.exe"
				var testResult = startswtih(filename, "script")  
			`,
			expected: true,
		},
		{
			name: "Startswtih string false",
			input: `
				var filename = "script.exe"
				var testResult = startswtih(filename, "file")  
			`,
			expected: false,
		},
		{
			name: "Endswtih string",
			input: `
				var filename = "script.exe"
				var testResult = endswitsh(filename, "exe")  
			`,
			expected: true,
		},
		{
			name: "Endswtih string",
			input: `
				var filename = "script.exe"
				var testResult = endswitsh(filename, "txt")  
			`,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestStringValidators(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Alphanumeric true",
			input:    `var testResult = isalnum("Agent007")`,
			expected: true,
		},
		{
			name:     "Alphanumeric false due to space",
			input:    `var testResult = isalnum("Agent 007")`,
			expected: false,
		},
		{
			name:     "Is digit true",
			input:    `var testResult = isdigit("12345")`,
			expected: true,
		},
		{
			name:     "Is lower true",
			input:    `var testResult = islower("hello")`,
			expected: true,
		},
		{
			name:     "Is upper false mixed case",
			input:    `var testResult = isupper("Hello")`,
			expected: false,
		},
		{
			name:     "Is space false",
			input:    `var testResult = isspace(" ana ")`,
			expected: false,
		},
		{
			name:     "Is space true",
			input:    `var testResult = isspace(" \t\n ")`,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' missing entirely")
			}
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestRightStringFinders(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Replace a with o",
			input:    `var testResult = replace("banana", "a", "o")`,
			expected: "bonono",
		},
		{
			name:     "Rfind match last occurrence index world",
			input:    `var testResult = rfind("hello world world", "world")`,
			expected: 12,
		},
		{
			name:     "Rfind not finding",
			input:    `var testResult = rfind("hello", "missing")`,
			expected: -1,
		},
		{
			name:     "Rindex matches last occurrence index apple",
			input:    `var testResult = rindex("apple banana apple", "apple")`,
			expected: 13,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' missing entirely")
			}
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestDivPow(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Divmod 6/2",
			input: `
				var testResult = divmod(6, 2)  
			`,
			expected: &[]any{3, 0},
		},
		{
			name: "Pow 2^3",
			input: `
				var testResult = pow(2, 3)  
			`,
			expected: 8,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Count on string",
			input: `
				var countStr = "one apple per apple per person"
				var testResult = count(countStr, "apple")`,
			expected: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' missing entirely")
			}
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestEnumerate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Enumerate simple array",
			input: `
				var arr = ["apple", "pear", "kiwi"]
				var testResult = enumerate(arr)  
			`,
			expected: &[]any{
				&[]any{0, "apple"},
				&[]any{1, "pear"},
				&[]any{2, "kiwi"},
			},
		},
		{
			name: "Pow 2^3",
			input: `
				var testResult = pow(2, 3)  
			`,
			expected: 8,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Count on string",
			input: `
				var arr = ["apple", "pear", "kiwi"]
				var testResult = index(arr, "apple")`,
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' missing entirely")
			}
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestIsinstance(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "isinstance on int",
			input: `
				var testResult = isinstance(5, "int")`,
			expected: true,
		},
		{
			name: "isinstance on string",
			input: `
				var testResult = isinstance("hello", "string")`,
			expected: true,
		},
		{
			name: "isinstance on bool",
			input: `
				var testResult = isinstance(true, "bool")`,
			expected: true,
		},
		{
			name: "isinstance on array",
			input: `
				var testResult = isinstance([4,6,4], "array")`,
			expected: true,
		},
		{
			name: "isinstance on function",
			input: `
				func add(a,b) {
					return a+b
				}
				var testResult = isinstance(add, "function")`,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' missing entirely")
			}
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestExecEval(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Exec assignment",
			input: `
				var targetVar = 0
				var xExec = "targetVar = 50" 
				exec(xExec)
				var testResult = targetVar  
			`,
			expected: 50,
		},
		{
			name: "Eval 50+5",
			input: `
				var xEval = "50 + 5"
				var testResult = eval(xEval) 
			`,
			expected: 55,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' missing entirely")
			}
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Format hex",
			input: `
                var targetVar = 15
                var testResult = format(targetVar, "x")  
            `,
			expected: "f",
		},
		{
			name: "Format binary",
			input: `
                var targetVar = 5
                var testResult = format(targetVar, "b")  
            `,
			expected: "101",
		},
		{
			name: "Format fixed float",
			input: `
                var targetVar = float(5)
                var testResult = format(targetVar, "f")  
            `,
			expected: "5.000000",
		},
		{
			name: "Format uppercase hex",
			input: `
				var targetVar = 14
				var testResult = format(targetVar, "X")  
			`,
			expected: "E",
		},
		{
			name: "Format with uppercase F",
			input: `
				var targetVar = float(7)
				var testResult = format(targetVar, "F")  
    		`,
			expected: "7.000000",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' missing entirely")
			}
			if result != tc.expected {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestZip(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Zip with simple arrays",
			input: `
				var arr1 = [1, 4, 7, 3]
				var arr2 = [5, 9, 17, 13]

				var testResult = zip(arr1, arr2)
			`,
			expected: &[]any{
				&[]any{1, 5},
				&[]any{4, 9},
				&[]any{7, 17},
				&[]any{3, 13},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Join simple str array",
			input: `
                var arr = ["a", "l", "l"]
                var testResult = join(arr, "")  
            `,
			expected: "all",
		},
		{
			name: "Join str array with -",
			input: `
                var arr = ["here", "I", "am"]
                var testResult = join(arr, "-")  
            `,
			expected: "here-I-am",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' missing entirely")
			}
			if result != tc.expected {
				t.Errorf("Expected %v (%T), got %v (%T)",
					tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestSwapcase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "Swapcase flips cases completely",
			input:    `var testResult = swapcase("Hello World 123")`,
			expected: "hELLO wORLD 123",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' missing entirely")
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)", tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Remove drops the first match from array",
			input: `
                var arr = [10, 20, 30, 20]
                remove(arr, 20)
                var testResult = arr
            `,
			expected: &[]any{10, 30, 20},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' missing entirely")
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v (%T), got %v (%T)", tc.expected, tc.expected, result, result)
			}
		})
	}
}

func TestFileIOBuiltins(t *testing.T) {
	tmpDir := t.TempDir()
	testFilePath := filepath.ToSlash(filepath.Join(tmpDir, "sample.txt"))

	inputScript := fmt.Sprintf(`
        var fileW = open("%s", "w")
        write(fileW, "Hello from the engine!")
        close(fileW)

        var fileR = open("%s", "r")
        var testResult = read(fileR)
        close(fileR)
    `, testFilePath, testFilePath)

	env := runInterpreter(inputScript)
	result, ok := env.Lookup("testResult")
	if !ok {
		t.Fatalf("Variable 'testResult' missing entirely from environment evaluation")
	}

	expected := "Hello from the engine!"
	if result != expected {
		t.Errorf("Expected exact file contents '%s', instead got '%v'", expected, result)
	}
}
