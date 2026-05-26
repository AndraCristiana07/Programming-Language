package ast

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type NativeFunction struct {
	ArgsCount int
	Body      func(args []any) any
}

func (n *NativeFunction) NrArgs() int {
	return n.ArgsCount
}

func (n *NativeFunction) Call(v *Visitor, args []any) any {
	return n.Body(args)
}

func NewGlobalEnvironment() *Environment {
	globals := NewEnvironment()

	//  len(arr)
	globals.Define("len", &NativeFunction{
		ArgsCount: 1,
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
			val, ok := args[0].(int)
			if !ok {
				panic("InvalidArgument: bin() expects an integer")
			}

			return "0b" + strconv.FormatInt(int64(val), 2)
		},
	})

	globals.Define("hex", &NativeFunction{
		ArgsCount: 1,
		Body: func(args []any) any {
			val, ok := args[0].(int)
			if !ok {
				panic("InvalidArgument: hex() expects an integer")
			}

			return "0x" + strconv.FormatInt(int64(val), 16)
		},
	})

	globals.Define("oct", &NativeFunction{
		ArgsCount: 1,
		Body: func(args []any) any {
			val, ok := args[0].(int)
			if !ok {
				panic("InvalidArgument: oct() expects an integer")
			}
			return "0o" + strconv.FormatInt(int64(val), 8)
		},
	})

	globals.Define("abs", &NativeFunction{
		ArgsCount: 1,
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
			str, ok := args[0].(string)
			if !ok || len(str) != 1 {
				panic("ord() expects a single-character string")
			}
			return int(str[0])
		},
	})

	globals.Define("chr", &NativeFunction{
		ArgsCount: 1,
		Body: func(args []any) any {
			val, ok := args[0].(int)
			if !ok || val < 0 || val > 255 {
				panic("chr() expects an integer ASCII code (0-255)")
			}
			return string(rune(val))
		},
	})

	globals.Define("range", &NativeFunction{
		ArgsCount: -1,
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic("lower() expects a string")
			}
			return strings.ToLower(str)
		},
	})

	globals.Define("upper", &NativeFunction{
		ArgsCount: 1,
		Body: func(args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic("upper() expects a string")
			}
			return strings.ToUpper(str)
		},
	})

	globals.Define("pop", &NativeFunction{
		ArgsCount: 1,
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
		Body: func(args []any) any {
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
	return globals
}
