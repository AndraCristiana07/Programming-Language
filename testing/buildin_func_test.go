package testing

import (
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
