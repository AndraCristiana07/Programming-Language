package testing

import (
	"my_language/ast"
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
			expected: &[]any{"age", "name"},
		},
		{
			name: "Construct and fetch values from map",
			input: `
                var user = {"name": "Alice", "age": 30}
                var testResult = values(user)
            `,
			expected: &[]any{30, "Alice"},
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
		{
			name: "Reassign nested map in map",
			input: `
				var schools =  {"sch1":{"Alice":"class1", "Bob":"class2"}, "sch2":{"Anna":"class4", "Fred":"class2"}}
				schools["sch1"]["Bob"] = "class0"
				var testResult = schools
			`,
			expected: &map[string]any{
				"sch1": &map[string]any{"Alice": "class1", "Bob": "class0"},
				"sch2": &map[string]any{"Anna": "class4", "Fred": "class2"},
			},
		},
		{
			name: "Reassign multiple in nested map in map",
			input: `
				var schools =  {"sch1":{"Alice":"class1", "Bob":"class2"}, "sch2":{"Anna":"class4", "Fred":"class2"}}
				schools["sch1"]["Bob"] = "class0"
				schools["sch1"]["Bob"] = "class3"
				schools["sch2"]["Anna"] = "class0"
				schools["sch2"]["Fred"] = "class1"

				var testResult = schools
			`,
			expected: &map[string]any{
				"sch1": &map[string]any{"Alice": "class1", "Bob": "class3"},
				"sch2": &map[string]any{"Anna": "class0", "Fred": "class1"},
			},
		},
		{
			name: "Reassign nested array in map",
			input: `
				var grades = {"Anna":[6,8,10,7], "Bob":[6,10,6]}
				grades["Anna"][1] = 10
				var testResult = grades
			`,
			expected: &map[string]any{
				"Anna": &[]any{6, 10, 10, 7},
				"Bob":  &[]any{6, 10, 6},
			},
		},
		{
			name: "Reassign multiple in nested array in map",
			input: `
				var grades = {"Anna":[6,8,10,7], "Bob":[6,10,6]}
				grades["Anna"][1] = 10
				grades["Anna"][0] = 7
				grades["Bob"][2] = 8
				grades["Anna"][1] = 9

				var testResult = grades
			`,
			expected: &map[string]any{
				"Anna": &[]any{7, 9, 10, 7},
				"Bob":  &[]any{6, 10, 8},
			},
		},
		{
			name: "Reassign whole array in double nested array in map",
			input: `
				var grades = {"Anna":[[10,7,5], [5,3]], "Bob":[[6,7], [7,9]]}
				grades["Anna"][0] = [7,8,7]
				var testResult = grades

			`,
			expected: &map[string]any{
				"Anna": &[]any{
					&[]any{7, 8, 7},
					&[]any{5, 3},
				},
				"Bob": &[]any{
					&[]any{6, 7},
					&[]any{7, 9},
				},
			},
		},
		{
			name: "Reassign array values in double nested array in map",
			input: `
				var grades = {"Anna":[[10,7,5], [5,3]], "Bob":[[6,7], [7,9]]}
				grades["Anna"][0][1] = 5
				var testResult = grades

			`,
			expected: &map[string]any{
				"Anna": &[]any{
					&[]any{10, 5, 5},
					&[]any{5, 3},
				},
				"Bob": &[]any{
					&[]any{6, 7},
					&[]any{7, 9},
				},
			},
		},
		{
			name: "Reassign with dot field access on map",
			input: `
				var users = {"name":"Anna", "age":20}
				users.name = "Alice"
				var testResult = users
			`,
			expected: &map[string]any{"name": "Alice", "age": 20},
		},
		{
			name: "Reassign with dot field access on nested map",
			input: `
				var schools =  {"sch1":{"Alice":"class1", "Bob":"class2"}, "sch2":{"Anna":"class4", "Fred":"class2"}}
				schools.sch1.Alice = "class5"
				var testResult = schools
			`,
			expected: &map[string]any{
				"sch1": &map[string]any{"Alice": "class5", "Bob": "class2"},
				"sch2": &map[string]any{"Anna": "class4", "Fred": "class2"},
			},
		},
		{
			name: "Len on a map",
			input: `
				var users = {"name":"Anna", "age":20}
				var testResult = len(users)
			`,
			expected: 2,
		},
		{
			name: "Str on a map",
			input: `
				var users = {"name":"Anna", "age":20}
				var testResult = str(users)
			`,
			expected: `{"age" : 20, "name" : Anna}`,
		},
		{
			name: "Bool on map",
			input: `
				var users = {"name":"Anna", "age":20}
				var testResult = bool(users)
			`,
			expected: true,
		},
		{
			name: "Type on map",
			input: `
				var users = {"name":"Anna", "age":20}
				var testResult = type(users)
			`,
			expected: "map",
		},
		{
			name: "Any on map",
			input: `
				var bMap = {"was":true, "is":false}
				var testResult = any(bMap)
			`,
			expected: true,
		},
		{
			name: "All on map",
			input: `
				var bMap = {true : "apple", true : "orange"}
				var testResult = all(bMap)
			`,
			expected: true,
		},
		{
			name: "List on map",
			input: `
				var users = {"name":"Anna", "age":20}
				var testResult = list(users)
			`,
			expected: &[]any{"age", "name"},
		},
		{
			name: "Enumerate on map",
			input: `
				var users = {"name":"Anna", "age":20}
				var testResult = enumerate(users)
			`,
			expected: &[]any{
				&[]any{0, "age"},
				&[]any{1, "name"},
			},
		},
		{
			name: "Is instsnce on map",
			input: `
				var users = {"name":"Anna", "age":20}
				var testResult = isinstance(users, "map")
			`,
			expected: true,
		},
		{
			name: "Zip on maps",
			input: `
                var users = {"name":"Anna", "age":20}
                var school = {"year":2001, "principle":"Bob"}
                var testResult = tuple(zip(users, school))
            `,
			expected: &ast.Tuple{Elements: []any{
				&ast.Tuple{Elements: []any{"age", "principle"}},
				&ast.Tuple{Elements: []any{"name", "year"}},
			}},
		},
		{
			name: "Tuple on map",
			input: `
				var users = {"name":"Anna", "age":20}
				var testResult = tuple(users)
			`,
			expected: &ast.Tuple{Elements: []any{"age", "name"}},
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
