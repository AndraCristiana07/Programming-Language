package ast

import (
	"fmt"
	"reflect"
	"strconv"
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

	return globals
}
