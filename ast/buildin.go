package ast

import (
	"bufio"
	"fmt"
	"my_language/parser"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)

type NativeFunction struct {
	ArgsCount int
	Body      func(v *Visitor, args []any) any
}

func (n *NativeFunction) NrArgs() int {
	return n.ArgsCount
}

func (n *NativeFunction) Call(v *Visitor, args []any) any {
	return n.Body(v, args)
}

func NewGlobalEnvironment() *Environment {
	globals := NewEnvironment()

	//  len(arr)
	globals.Define("len", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			// check for array slice
			if arrPtr, ok := args[0].(*[]any); ok {
				return len(*arrPtr) // dereference pointer
			}
			// check for string type
			if str, ok := args[0].(string); ok {
				return len(str)
			}
			panic("InvalidArgument: len() expects an array or string type")
		},
	})

	// append(arr, item)
	globals.Define("append", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			arrPtr, ok := args[0].(*[]any)
			if !ok {
				panic("InvalidArgument: append() expects an array as the first argument")
			}

			// grab the actual slice value
			arr := *arrPtr
			newItem := args[1]

			// type safe guard
			if len(arr) > 0 {
				firstItem := arr[0]

				firstType := reflect.TypeOf(firstItem)
				newType := reflect.TypeOf(newItem)

				if firstType != newType {
					var expectedName, gotName string = "unknown", "unknown"
					if firstType != nil {
						expectedName = firstType.String()
					}
					if newType != nil {
						gotName = newType.String()
					}

					// clean up Go-specific naming repr
					if expectedName == "[]interface {}" || expectedName == "*[]interface {}" || expectedName == "[]main.any" {
						expectedName = "array"
					}
					if gotName == "[]interface {}" || gotName == "*[]interface {}" || gotName == "[]main.any" {
						gotName = "array"
					}

					panic(fmt.Sprintf("TypeError: Cannot append %s to an array of %s", gotName, expectedName))
				}
			}

			*arrPtr = append(arr, newItem)
			return arrPtr
		},
	})

	// int(value)
	globals.Define("int", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			switch v := args[0].(type) {
			case string:
				if i, err := strconv.Atoi(v); err == nil {
					return i
				}
				panic("Cannot convert string '" + v + "' to int")
			case bool:
				if v {
					return 1
				}
				return 0
			case int:
				return v
			default:
				panic("Unsupported type for int() conversion")
			}
		},
	})

	// str(value)
	globals.Define("str", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			if args[0] == nil {
				return "nil"
			}

			switch v := args[0].(type) {
			case string:
				return v
			case *[]any:
				elements := *v
				var sb strings.Builder
				sb.WriteString("[")
				for i, element := range elements {
					fmt.Fprintf(&sb, "%v", element)
					if i < len(elements)-1 {
						sb.WriteString(", ")
					}
				}
				sb.WriteString("]")
				return sb.String()
			default:
				return fmt.Sprintf("%v", v)
			}
		},
	})

	globals.Define("bool", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			val := args[0]
			if val == nil {
				return false
			}

			switch v := val.(type) {
			case bool:
				return v
			case int:
				// 0  false
				return v != 0
			case string:
				// empty string "" is false
				return v != ""
			case *[]any:
				// empty array is false
				return len(*v) > 0
			default:
				// objects are true
				return true
			}
		},
	})

	globals.Define("type", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			if args[0] == nil {
				return "nil"
			}
			switch args[0].(type) {
			case int:
				return "int"
			case string:
				return "string"
			case bool:
				return "bool"
			case *[]any:
				return "array"
			default:
				if _, ok := args[0].(Callable); ok {
					return "function"
				}
				return "unknown"
			}
		},
	})

	globals.Define("bin", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			val, ok := args[0].(int)
			if !ok {
				panic("InvalidArgument: bin() expects an integer")
			}

			return "0b" + strconv.FormatInt(int64(val), 2)
		},
	})

	globals.Define("hex", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			val, ok := args[0].(int)
			if !ok {
				panic("InvalidArgument: hex() expects an integer")
			}

			return "0x" + strconv.FormatInt(int64(val), 16)
		},
	})

	globals.Define("oct", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			val, ok := args[0].(int)
			if !ok {
				panic("InvalidArgument: oct() expects an integer")
			}
			return "0o" + strconv.FormatInt(int64(val), 8)
		},
	})

	globals.Define("abs", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			val, ok := args[0].(int)
			if !ok {
				panic("abs() expects an integer")
			}
			if val < 0 {
				return -val
			}
			return val
		},
	})

	// min(item1, item2)
	globals.Define("min", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			a, ok1 := args[0].(int)
			b, ok2 := args[1].(int)
			if !ok1 || !ok2 {
				panic("min() expects two integers")
			}
			if a < b {
				return a
			}
			return b
		},
	})

	// max(item1, item2)
	globals.Define("max", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			a, ok1 := args[0].(int)
			b, ok2 := args[1].(int)
			if !ok1 || !ok2 {
				panic("max() expects two integers")
			}
			if a > b {
				return a
			}
			return b
		},
	})

	globals.Define("input", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			prompt, ok := args[0].(string)
			if !ok {
				panic("input() expects a string prompt")
			}

			fmt.Print(prompt)

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')

			// trim newline characters from the user input
			return strings.TrimRight(text, "\r\n")
		},
	})

	globals.Define("ord", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok || len(str) != 1 {
				panic("ord() expects a single-character string")
			}
			return int(str[0])
		},
	})

	globals.Define("chr", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			val, ok := args[0].(int)
			if !ok || val < 0 || val > 255 {
				panic("chr() expects an integer ASCII code (0-255)")
			}
			return string(rune(val))
		},
	})

	globals.Define("range", &NativeFunction{
		ArgsCount: -1,
		Body: func(v *Visitor, args []any) any {
			start := 0
			end := 0

			if len(args) == 1 {
				val, ok := args[0].(int)
				if !ok {
					panic("TypeError: range() arguments must be integers")
				}
				end = val
			} else if len(args) == 2 {
				sVal, ok1 := args[0].(int)
				eVal, ok2 := args[1].(int)
				if !ok1 || !ok2 {
					panic("TypeError: range() arguments must be integers")
				}
				start = sVal
				end = eVal
			} else {
				panic("ArgumentError: range() expects 1 or 2 arguments")
			}

			if start > end {
				result := []any{}
				for i := start; i > end; i-- {
					result = append(result, i)
				}
				return &result
			}

			result := make([]any, end-start)
			for i := 0; i < (end - start); i++ {
				result[i] = start + i
			}
			return &result
		},
	})

	globals.Define("set", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			arr, ok := args[0].(*[]any)
			if !ok {
				panic("TypeError: set() expects an array as an argument")
			}

			seen := make(map[any]bool)
			uniqueArr := []any{}

			for _, item := range *arr {
				if !seen[item] {
					seen[item] = true
					uniqueArr = append(uniqueArr, item)
				}
			}
			return &uniqueArr
		},
	})

	globals.Define("lower", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic("lower() expects a string")
			}
			return strings.ToLower(str)
		},
	})

	globals.Define("upper", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic("upper() expects a string")
			}
			return strings.ToUpper(str)
		},
	})

	globals.Define("pop", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			// grab shared pointer address
			arrPtr, ok := args[0].(*[]any)
			if !ok {
				panic("TypeError: pop() expects an array reference")
			}

			arr := *arrPtr
			if len(arr) == 0 {
				panic("IndexError: pop from empty array")
			}

			lastIdx := len(arr) - 1
			removedItem := arr[lastIdx]

			// cutthe shared array down by mutating the memory address
			*arrPtr = arr[:lastIdx]

			return removedItem
		},
	})

	// find(container, item)
	globals.Define("find", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			switch container := args[0].(type) {
			case string:
				target, ok := args[1].(string)
				if !ok {
					panic("find() on a string expects a string target")
				}
				return strings.Index(container, target)
			case *[]any:
				for i, val := range *container {
					if reflect.DeepEqual(val, args[1]) {
						return i
					}
				}
				return -1
			default:
				panic("TypeError: find() expects a string or array")
			}
		},
	})

	globals.Define("reverse", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			arrPtr, ok := args[0].(*[]any)
			if !ok {
				panic("reverse() expects an array")
			}
			arr := *arrPtr
			reversed := make([]any, len(arr))
			for i, item := range arr {
				reversed[len(arr)-1-i] = item
			}
			return &reversed
		},
	})

	globals.Define("all", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			arrPtr, ok := args[0].(*[]any)
			if !ok {
				panic("TypeError: all() expects an array as an argument")
			}

			boolFunc, _ := globals.Lookup("bool")
			boolCallable := boolFunc.(Callable)

			for _, item := range *arrPtr {
				truth := boolCallable.Call(nil, []any{item}).(bool)
				if !truth {
					return false
				}
			}
			return true
		},
	})

	globals.Define("any", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			arrPtr, ok := args[0].(*[]any)
			if !ok {
				panic("TypeError: any() expects an array as an argument")
			}

			boolFunc, _ := globals.Lookup("bool")
			boolCallable := boolFunc.(Callable)

			for _, item := range *arrPtr {
				truth := boolCallable.Call(nil, []any{item}).(bool)
				if truth {
					return true
				}
			}
			return false
		},
	})

	globals.Define("filter", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			// validate the predicate function
			predicate, ok := args[0].(Callable)
			if !ok {
				panic("TypeError: first argument to filter() must be a callable function")
			}

			// validate the target array pointer
			arrPtr, ok := args[1].(*[]any)
			if !ok {
				panic("TypeError: second argument to filter() must be an array")
			}

			boolFunc, _ := globals.Lookup("bool")
			boolCallable := boolFunc.(Callable)

			filteredElements := []any{}

			for _, item := range *arrPtr {
				res := predicate.Call(nil, []any{item})

				if boolCallable.Call(nil, []any{res}).(bool) {
					filteredElements = append(filteredElements, item)
				}
			}

			return &filteredElements
		},
	})

	globals.Define("float", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			switch v := args[0].(type) {
			case float64:
				return v
			case int:
				return float64(v)
			case string:
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					return f
				}
				panic("ValueError: Cannot convert string '" + v + "' to float")
			case bool:
				if v {
					return 1.0
				}
				return 0.0
			default:
				panic("TypeError: Unsupported type for float() conversion")
			}
		},
	})

	globals.Define("list", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			switch v := args[0].(type) {
			case string:
				// convert string to an array of individual characters
				result := []any{}
				for _, r := range v {
					result = append(result, string(r))
				}
				return &result
			case *[]any:
				// create copy of an existing array pointer
				result := make([]any, len(*v))
				copy(result, *v)
				return &result
			default:
				panic("TypeError: list() expects a string or an array")
			}
		},
	})

	globals.Define("map", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			transform, ok := args[0].(Callable)
			if !ok {
				panic("TypeError: first argument to map() must be a callable function")
			}

			arrPtr, ok := args[1].(*[]any)
			if !ok {
				panic("TypeError: second argument to map() must be an array")
			}

			mappedElements := make([]any, len(*arrPtr))
			for i, item := range *arrPtr {
				mappedElements[i] = transform.Call(nil, []any{item})
			}

			return &mappedElements
		},
	})

	globals.Define("round", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			switch v := args[0].(type) {
			case int:
				return v
			case float64:
				// basic round logic for floats to integers
				if v < 0 {
					return int(v - 0.5)
				}
				return int(v + 0.5)
			default:
				panic("TypeError: round() expects a number")
			}
		},
	})

	globals.Define("sorted", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			arrPtr, ok := args[0].(*[]any)
			if !ok {
				panic("TypeError: sorted() expects an array")
			}

			arr := *arrPtr
			if len(arr) == 0 {
				return &[]any{}
			}

			sortedElements := make([]any, len(arr))
			copy(sortedElements, arr)

			// type check
			switch sortedElements[0].(type) {
			case int:
				sort.Slice(sortedElements, func(i, j int) bool {
					return sortedElements[i].(int) < sortedElements[j].(int)
				})
			case string:
				sort.Slice(sortedElements, func(i, j int) bool {
					return sortedElements[i].(string) < sortedElements[j].(string)
				})
			default:
				panic("TypeError: sorted() only supports arrays of integers or strings")
			}

			return &sortedElements
		},
	})

	globals.Define("sum", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			arrPtr, ok := args[0].(*[]any)
			if !ok {
				panic("TypeError: sum() expects an array")
			}

			var totalInt int = 0
			var totalFloat float64 = 0.0
			hasFloat := false

			for _, item := range *arrPtr {
				switch v := item.(type) {
				case int:
					if hasFloat {
						totalFloat += float64(v)
					} else {
						totalInt += v
					}
				case float64:
					if !hasFloat {
						hasFloat = true
						totalFloat = float64(totalInt)
					}
					totalFloat += v
				default:
					panic("TypeError: sum() array elements must be numbers")
				}
			}

			if hasFloat {
				return totalFloat
			}
			return totalInt
		},
	})

	globals.Define("split", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			str, ok1 := args[0].(string)
			sep, ok2 := args[1].(string)
			if !ok1 || !ok2 {
				panic("TypeError: split() expects two string arguments")
			}

			parts := strings.Split(str, sep)
			result := make([]any, len(parts))
			for i, part := range parts {
				result[i] = part
			}
			return &result
		},
	})

	globals.Define("insert", &NativeFunction{
		ArgsCount: 3,
		Body: func(v *Visitor, args []any) any {
			arrPtr, ok1 := args[0].(*[]any)
			idx, ok2 := args[1].(int)
			if !ok1 || !ok2 {
				panic("TypeError: insert() expects an array and an integer index")
			}

			arr := *arrPtr
			if idx < 0 || idx > len(arr) {
				panic(fmt.Sprintf("IndexError: insert index %d out of bounds for length %d", idx, len(arr)))
			}

			// grow slice by appending a dummy element
			arr = append(arr, nil)
			// shift elements to the right by one slot
			copy(arr[idx+1:], arr[idx:])
			// insert the target item in the opened slot
			arr[idx] = args[2]

			*arrPtr = arr
			return nil
		},
	})
	globals.Define("strip", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic("TypeError: strip() expects a string argument")
			}
			return strings.TrimSpace(str)
		},
	})

	globals.Define("startswtih", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			str, ok1 := args[0].(string)
			prefix, ok2 := args[1].(string)
			if !ok1 || !ok2 {
				panic("TypeError: startswtih() expects two string arguments")
			}
			return strings.HasPrefix(str, prefix)
		},
	})
	globals.Define("endswitsh", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			str, ok1 := args[0].(string)
			suffix, ok2 := args[1].(string)
			if !ok1 || !ok2 {
				panic("TypeError: endswitsh() expects two string arguments")
			}
			return strings.HasSuffix(str, suffix)
		},
	})

	globals.Define("isalnum", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic("TypeError: isalnum() expects a string")
			}
			if len(str) == 0 {
				return false
			}

			for _, r := range str {
				if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
					return false
				}
			}
			return true
		},
	})

	globals.Define("isalpha", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic("TypeError: isalpha() expects a string")
			}
			if len(str) == 0 {
				return false
			}

			for _, r := range str {
				if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
					return false
				}
			}
			return true
		},
	})

	globals.Define("isdigit", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic("TypeError: isdigit() expects a string")
			}
			if len(str) == 0 {
				return false
			}

			for _, r := range str {
				if r < '0' || r > '9' {
					return false
				}
			}
			return true
		},
	})

	globals.Define("islower", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic("TypeError: islower() expects a string")
			}

			hasCased := false
			for _, r := range str {
				if r >= 'A' && r <= 'Z' {
					return false
				}
				if r >= 'a' && r <= 'z' {
					hasCased = true
				}
			}
			return hasCased
		},
	})

	globals.Define("isupper", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic("TypeError: isupper() expects a string")
			}

			hasCased := false
			for _, r := range str {
				if r >= 'a' && r <= 'z' {
					return false
				}
				if r >= 'A' && r <= 'Z' {
					hasCased = true
				}
			}
			return hasCased
		},
	})

	globals.Define("isspace", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic("TypeError: isspace() expects a string")
			}
			if len(str) == 0 {
				return false
			}

			for _, r := range str {
				if r != ' ' && r != '\t' && r != '\n' && r != '\r' {
					return false
				}
			}
			return true
		},
	})

	globals.Define("replace", &NativeFunction{
		ArgsCount: 3,
		Body: func(v *Visitor, args []any) any {
			str, ok1 := args[0].(string)
			old, ok2 := args[1].(string)
			newStr, ok3 := args[2].(string)
			if !ok1 || !ok2 || !ok3 {
				panic("TypeError: replace() expects three string arguments")
			}

			return strings.ReplaceAll(str, old, newStr)
		},
	})

	globals.Define("rfind", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			str, ok1 := args[0].(string)
			substr, ok2 := args[1].(string)
			if !ok1 || !ok2 {
				panic("TypeError: rfind() expects two string arguments")
			}

			return strings.LastIndex(str, substr)
		},
	})

	globals.Define("rindex", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			str, ok1 := args[0].(string)
			substr, ok2 := args[1].(string)
			if !ok1 || !ok2 {
				panic("TypeError: rindex() expects two string arguments")
			}

			idx := strings.LastIndex(str, substr)
			if idx == -1 {
				panic(fmt.Sprintf("ValueError: substring '%s' not found in rindex()", substr))
			}

			return idx
		},
	})

	globals.Define("divmod", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			a, ok1 := args[0].(int)
			b, ok2 := args[1].(int)
			if !ok1 || !ok2 {
				panic("TypeError: divmod() expects two integers")
			}
			if b == 0 {
				panic("ZeroDivisionError: integer division or modulo by zero")
			}

			result := []any{a / b, a % b}
			return &result
		},
	})

	globals.Define("pow", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			base, ok1 := args[0].(int)
			exp, ok2 := args[1].(int)
			if !ok1 || !ok2 {
				panic("TypeError: pow() expects two string arguments")
			}

			return power(base, exp)
		},
	})

	globals.Define("count", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			str, ok1 := args[0].(string)
			sub, ok2 := args[1].(string)
			if !ok1 || !ok2 {
				panic("TypeError: count() expects two strings")
			}
			return strings.Count(str, sub)
		},
	})

	globals.Define("enumerate", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			arrPtr, ok := args[0].(*[]any)
			if !ok {
				panic("TypeError: enumerate() expects an array")
			}

			arr := *arrPtr
			result := make([]any, len(arr))
			for i, val := range arr {
				pair := []any{i, val}
				result[i] = &pair
			}
			return &result
		},
	})

	globals.Define("index", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			arrPtr, ok := args[0].(*[]any)
			if !ok {
				panic("TypeError: index() expects an array as the first argument")
			}

			for i, val := range *arrPtr {
				if reflect.DeepEqual(val, args[1]) {
					return i
				}
			}
			panic(fmt.Sprintf("ValueError: element '%v' is not in array", args[1]))
		},
	})

	globals.Define("isinstance", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			targetType, ok := args[1].(string)
			if !ok {
				panic("TypeError: isinstance() second argument must be a type string")
			}

			var actualType string
			switch args[0].(type) {
			case int:
				actualType = "int"
			case float64:
				actualType = "float"
			case string:
				actualType = "string"
			case bool:
				actualType = "bool"
			case *[]any:
				actualType = "array"
			case Callable:
				actualType = "function"
			default:
				actualType = "nil"
			}

			return actualType == targetType
		},
	})

	globals.Define("eval", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			code, ok := args[0].(string)
			if !ok {
				panic("TypeError: eval() expects a string argument")
			}

			// givw dynamic string into ANTLR pipeline
			inputStream := antlr.NewInputStream(code)
			lexer := parser.NewGrammarLexer(inputStream)
			tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
			p := parser.NewGrammarParser(tokens)

			tree := p.Expr()

			return tree.Accept(v)
		},
	})

	globals.Define("exec", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			code, ok := args[0].(string)
			if !ok {
				panic("TypeError: exec() expects a string argument")
			}

			inputStream := antlr.NewInputStream(code)
			lexer := parser.NewGrammarLexer(inputStream)
			tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
			p := parser.NewGrammarParser(tokens)

			tree := p.Program()

			tree.Accept(v)
			return nil
		},
	})

	globals.Define("format", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			spec, ok := args[1].(string)
			if !ok {
				panic("TypeError: format() specifier must be a string")
			}

			switch spec {
			case "x": // hexadecimal
				if val, ok := args[0].(int); ok {
					return fmt.Sprintf("%x", val)
				}
			case "X": // hexadecimal uppercase
				if val, ok := args[0].(int); ok {
					return fmt.Sprintf("%X", val)
				}
			case "b": // binary
				if val, ok := args[0].(int); ok {
					return fmt.Sprintf("%b", val)
				}
			case "f": // float
				if val, ok := args[0].(float64); ok {
					return fmt.Sprintf("%f", val)
				}
			case "F": // float upper
				if val, ok := args[0].(float64); ok {
					return fmt.Sprintf("%F", val)
				}
			}

			// default formatting
			return fmt.Sprintf("%v", args[0])
		},
	})

	globals.Define("zip", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			arrPtr1, ok1 := args[0].(*[]any)
			arrPtr2, ok2 := args[1].(*[]any)
			if !ok1 || !ok2 {
				panic("TypeError: zip() expects two array")
			}

			arr1 := *arrPtr1
			arr2 := *arrPtr2

			minLen := len(arr1)
			if len(arr2) < minLen {
				minLen = len(arr2)
			}

			result := make([]any, minLen)
			for i := 0; i < minLen; i++ {
				pair := []any{arr1[i], arr2[i]}
				result[i] = &pair
			}
			return &result
		},
	})

	globals.Define("join", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			arrPtr, ok1 := args[0].(*[]any)
			delim, ok2 := args[1].(string)
			if !ok1 || !ok2 {
				panic("TypeError: join() expects an array pointer and a delimiter string")
			}

			arr := *arrPtr
			strParts := make([]string, len(arr))

			for i, val := range arr {
				switch v := val.(type) {
				case string:
					strParts[i] = v
				case int:
					strParts[i] = strconv.Itoa(v)
				case float64:
					strParts[i] = strconv.FormatFloat(v, 'g', -1, 64)
				case bool:
					strParts[i] = strconv.FormatBool(v)
				default:
					strParts[i] = fmt.Sprintf("%v", val)
				}
			}

			return strings.Join(strParts, delim)
		},
	})

	globals.Define("rsplit", &NativeFunction{
		ArgsCount: 3,
		Body: func(v *Visitor, args []any) any {
			str, ok1 := args[0].(string)
			sep, ok2 := args[1].(string)
			maxsplit, ok3 := args[2].(int)
			if !ok1 || !ok2 || !ok3 {
				panic("TypeError: rsplit() expects two strings and an integer maxsplit count")
			}

			var parts []string
			if maxsplit < 0 {
				parts = strings.Split(str, sep)
			} else {
				parts = strings.Split(str, sep)
				if len(parts) > maxsplit+1 {
					keepIdx := len(parts) - maxsplit
					mergedFirstPart := strings.Join(parts[:keepIdx], sep)

					finalParts := append([]string{mergedFirstPart}, parts[keepIdx:]...)
					parts = finalParts
				}
			}

			result := make([]any, len(parts))
			for i, part := range parts {
				result[i] = part
			}
			return &result
		},
	})

	globals.Define("lstrip", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic("TypeError: lstrip() expects a string argument")
			}

			return strings.TrimLeftFunc(str, unicode.IsSpace)
		},
	})

	globals.Define("rstrip", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic("TypeError: rstrip() expects a string argument")
			}

			return strings.TrimRightFunc(str, unicode.IsSpace)
		},
	})

	return globals
}
