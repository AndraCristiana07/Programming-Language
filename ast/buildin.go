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
			if arr, ok := args[0].([]any); ok {
				return len(arr)
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
			arr, ok := args[0].([]any)
			if !ok {
				panic("InvalidArgument: append() expects an array as the first argument")
			}

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
					if expectedName == "[]interface {}" || expectedName == "[]main.any" {
						expectedName = "array"
					}
					if gotName == "[]interface {}" || gotName == "[]main.any" {
						gotName = "array"
					}

					panic(fmt.Sprintf("TypeError: Cannot append %s to an array of %s", gotName, expectedName))
				}
			}

			return append(arr, newItem)
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
			return fmt.Sprintf("%v", args[0])
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
			case []any:
				// empty arra is false
				return len(v) > 0
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
			case []any:
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
				return result
			}

			result := make([]any, end-start)
			for i := 0; i < (end - start); i++ {
				result[i] = start + i
			}
			return result
		},
	})

	return globals
}
