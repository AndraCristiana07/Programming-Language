package ast

import (
	"bufio"
	"fmt"
	"io"
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

type FileHandle struct {
	File *os.File
	Mode string // "r", "w", "a" (read, write, append)
}

func NewGlobalEnvironment() *Environment {
	globals := NewEnvironment()

	//  len(container)
	globals.Define("len", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			if args[0] == nil {
				return 0
			}
			// check for array slice
			if arrPtr, ok := args[0].(*[]any); ok {
				return len(*arrPtr) // dereference pointer
			}
			// check for string type
			if str, ok := args[0].(string); ok {
				return len(str)
			}
			// check for tuple
			if tpl, ok := args[0].(*Tuple); ok {
				return len(tpl.Elements)
			}
			// check map
			if mapPtr, ok := args[0].(*map[string]any); ok {
				return len(*mapPtr)
			}
			panic(RuntimeError("TypeError", "len() expects a string, array, tuple, or map type", v.currCtx))
		},
	})

	// append(arr, item)
	globals.Define("append", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			arrPtr, ok := args[0].(*[]any)
			if !ok {
				panic(RuntimeError("TypeError", "append() expects an array as the first argument", v.currCtx))
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

					panic(RuntimeError("TypeError", fmt.Sprintf("TypeError: Cannot append %s to an array of %s", gotName, expectedName), v.currCtx))
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
			switch val := args[0].(type) {
			case string:
				if i, err := strconv.Atoi(val); err == nil {
					return i
				}
				panic(RuntimeError("TypeError", fmt.Sprintf("Cannot convert string '%v' to int", val), v.currCtx))

			case bool:
				if val {
					return 1
				}
				return 0
			case int:
				return val
			default:
				panic(RuntimeError("TypeError", "Unsupported type for conversion", v.currCtx))
			}
		},
	})

	// str(value)
	globals.Define("str", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			return stringify(args[0])
		},
	})

	globals.Define("bool", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			val := args[0]
			if val == nil {
				return false
			}

			switch vl := val.(type) {
			case bool:
				return vl
			case int:
				// 0  false
				return vl != 0
			case string:
				// empty string "" is false
				return vl != ""
			case *[]any:
				// empty array is false
				return len(*vl) > 0
			case *Tuple:
				return len(vl.Elements) > 0
			case *map[string]any:
				// empty map is false
				return len(*vl) > 0
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
			case *Tuple:
				return "tuple"
			case *map[string]any:
				return "map"
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
				panic(RuntimeError("TypeError", "bin() expects an integer", v.currCtx))
			}

			return "0b" + strconv.FormatInt(int64(val), 2)
		},
	})

	globals.Define("hex", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			val, ok := args[0].(int)
			if !ok {
				panic(RuntimeError("TypeError", "hex() expects an integer", v.currCtx))
			}

			return "0x" + strconv.FormatInt(int64(val), 16)
		},
	})

	globals.Define("oct", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			val, ok := args[0].(int)
			if !ok {
				panic(RuntimeError("TypeError", "oct() expects an integer", v.currCtx))
			}
			return "0o" + strconv.FormatInt(int64(val), 8)
		},
	})

	globals.Define("abs", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			val, ok := args[0].(int)
			if !ok {
				panic(RuntimeError("TypeError", "abs() expects an integer", v.currCtx))
			}
			if val < 0 {
				return -val
			}
			return val
		},
	})

	// min(item1, item2)
	globals.Define("min", &NativeFunction{
		ArgsCount: -1,
		Body: func(v *Visitor, args []any) any {
			if len(args) == 0 {
				panic(RuntimeError("ValueError", "min() expects at least 1 argument", v.currCtx))
			}

			var elements []any

			// array or tuple
			if len(args) == 1 {
				switch collection := args[0].(type) {
				case *[]any:
					if collection != nil {
						elements = *collection
					}
				case *Tuple:
					if collection != nil {
						elements = collection.Elements
					}
				default:
					return args[0]
				}
			} else {
				// simple min(10, 20)
				elements = args
			}

			if len(elements) == 0 {
				panic(RuntimeError("ValueError", "min() arg is an empty sequence", v.currCtx))
			}

			smallest := elements[0]
			for _, item := range elements {
				curr, ok1 := item.(int)
				minSoFar, ok2 := smallest.(int)
				if !ok1 || !ok2 {
					panic(RuntimeError("TypeError", "min() comparisons require numeric values", v.currCtx))
				}
				if curr < minSoFar {
					smallest = item
				}
			}
			return smallest
		},
	})

	// max(item1, item2)
	globals.Define("max", &NativeFunction{
		ArgsCount: -1,
		Body: func(v *Visitor, args []any) any {
			if len(args) == 0 {
				panic(RuntimeError("ValueError", "max() expects at least 1 argument", v.currCtx))
			}

			var elements []any

			// array or tuple
			if len(args) == 1 {
				switch collection := args[0].(type) {
				case *[]any:
					if collection != nil {
						elements = *collection
					}
				case *Tuple:
					if collection != nil {
						elements = collection.Elements
					}
				default:
					return args[0]
				}
			} else {
				// simple max(10, 20)
				elements = args
			}

			if len(elements) == 0 {
				panic(RuntimeError("ValueError", "max() arg is an empty sequence", v.currCtx))
			}

			biggest := elements[0]
			for _, item := range elements {
				curr, ok1 := item.(int)
				maxSoFar, ok2 := biggest.(int)
				if !ok1 || !ok2 {
					panic(RuntimeError("TypeError", "max() comparisons require numeric values", v.currCtx))
				}
				if curr > maxSoFar {
					biggest = item
				}
			}
			return biggest
		},
	})

	globals.Define("input", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			prompt, ok := args[0].(string)
			if !ok {
				panic(RuntimeError("TypeError", "input() expects a string prompt", v.currCtx))
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
				panic(RuntimeError("TypeError", "ord() expects a single-character string", v.currCtx))
			}
			return int(str[0])
		},
	})

	globals.Define("chr", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			val, ok := args[0].(int)
			if !ok || val < 0 || val > 255 {
				panic(RuntimeError("TypeError", "chr() expects an integer ASCII code (0-255)", v.currCtx))
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
					panic(RuntimeError("TypeError", "range() arguments must be integers", v.currCtx))
				}
				end = val
			} else if len(args) == 2 {
				sVal, ok1 := args[0].(int)
				eVal, ok2 := args[1].(int)
				if !ok1 || !ok2 {
					panic(RuntimeError("TypeError", "range() arguments must be integers", v.currCtx))
				}
				start = sVal
				end = eVal
			} else {
				panic(RuntimeError("ValueError", "range() expects 1 or 2 arguments", v.currCtx))
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
				panic(RuntimeError("TypeError", "set() expects an array as an argument", v.currCtx))
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
				panic(RuntimeError("TypeError", "lower() expects a string", v.currCtx))
			}
			return strings.ToLower(str)
		},
	})

	globals.Define("upper", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic(RuntimeError("TypeError", "upper() expects a string", v.currCtx))
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
				panic(RuntimeError("TypeError", "pop() expects an array reference", v.currCtx))
			}

			arr := *arrPtr
			if len(arr) == 0 {
				panic(RuntimeError("IndexError", "pop from empty array", v.currCtx))
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
					panic(RuntimeError("TypeError", "find() on a string expects a string target", v.currCtx))
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
				panic(RuntimeError("TypeError", "find() expects a string or array", v.currCtx))
			}
		},
	})

	globals.Define("reverse", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			arrPtr, ok := args[0].(*[]any)
			if !ok {
				panic(RuntimeError("TypeError", "reverse() expects an array", v.currCtx))
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
			var elements []any

			switch val := args[0].(type) {
			case *[]any:
				elements = *val
			case *Tuple:
				elements = val.Elements
			case *map[string]any:
				m := *val
				for k := range m {
					elements = append(elements, k)
				}
			default:
				panic(RuntimeError("TypeError", fmt.Sprintf("all() expects an array, tuple, or map, got %T", args[0]), v.currCtx))
			}

			boolFunc, _ := globals.Lookup("bool")
			boolCallable := boolFunc.(Callable)

			for _, item := range elements {
				truth := boolCallable.Call(v, []any{item}).(bool)
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
			var elements []any

			switch val := args[0].(type) {
			case *[]any:
				elements = *val
			case *Tuple:
				elements = val.Elements
			case *map[string]any:
				m := *val
				for k := range m {
					elements = append(elements, k)
				}
			default:
				panic(RuntimeError("TypeError", fmt.Sprintf("any() expects an array, tuple, or map, got %T", args[0]), v.currCtx))
			}

			boolFunc, _ := globals.Lookup("bool")
			boolCallable := boolFunc.(Callable)

			for _, item := range elements {
				truth := boolCallable.Call(v, []any{item}).(bool)
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
				panic(RuntimeError("TypeError", "First argument to filter() must be a callable function", v.currCtx))
			}

			// track item type
			var elements []any
			var isTuple bool

			switch collection := args[1].(type) {
			case *[]any:
				if collection != nil {
					elements = *collection
				}
				isTuple = false
			case *Tuple:
				if collection != nil {
					elements = collection.Elements
				}
				isTuple = true
			default:
				panic(RuntimeError("TypeError", fmt.Sprintf("Second argument to filter() must be an array or tuple, got %T", args[1]), v.currCtx))
			}

			boolFunc, _ := globals.Lookup("bool")
			boolCallable := boolFunc.(Callable)

			filteredElements := make([]any, 0)

			for _, item := range elements {
				res := predicate.Call(v, []any{item})

				if boolCallable.Call(v, []any{res}).(bool) {
					filteredElements = append(filteredElements, item)
				}
			}

			// return correct matching
			if isTuple {
				return &Tuple{Elements: filteredElements}
			}
			return &filteredElements
		},
	})

	globals.Define("float", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			switch val := args[0].(type) {
			case float64:
				return val
			case int:
				return float64(val)
			case string:
				if f, err := strconv.ParseFloat(val, 64); err == nil {
					return f
				}
				panic(RuntimeError("ValueError", fmt.Sprintf("Cannot convert string %s to float", val), v.currCtx))

			case bool:
				if val {
					return 1.0
				}
				return 0.0
			default:
				panic(RuntimeError("ValueError", "Unsupported type for float() conversion", v.currCtx))
			}
		},
	})

	globals.Define("list", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			switch val := args[0].(type) {
			case string:
				// convert string to an array of individual characters
				result := []any{}
				for _, r := range val {
					result = append(result, string(r))
				}
				return &result
			case *Tuple:
				newSlice := make([]any, len(val.Elements))
				copy(newSlice, val.Elements)
				return &newSlice
			case *[]any:
				// create copy of an existing array pointer
				result := make([]any, len(*val))
				copy(result, *val)
				return &result
			case *map[string]any:
				m := *val
				sortedKeys := make([]string, 0, len(m))
				for k := range m {
					sortedKeys = append(sortedKeys, k)
				}
				sort.Strings(sortedKeys)

				result := make([]any, len(sortedKeys))
				for i, k := range sortedKeys {
					result[i] = k
				}
				return &result
			default:
				panic(RuntimeError("TypeError", "list() expects a string or an array", v.currCtx))
			}
		},
	})

	globals.Define("map", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			predicate, ok := args[0].(Callable)
			if !ok {
				panic(RuntimeError("TypeError", "First argument to map() must be a callable function", v.currCtx))
			}

			var elements []any
			var isTuple bool

			switch collection := args[1].(type) {
			case *[]any:
				if collection != nil {
					elements = *collection
				}
				isTuple = false
			case *Tuple:
				if collection != nil {
					elements = collection.Elements
				}
				isTuple = true
			default:
				panic(RuntimeError("TypeError", fmt.Sprintf("Second argument to map() must be an array or tuple, got %T", args[1]), v.currCtx))
			}

			mappedElements := make([]any, 0, len(elements))
			for _, item := range elements {
				res := predicate.Call(v, []any{item})
				mappedElements = append(mappedElements, res)
			}

			if isTuple {
				return &Tuple{Elements: mappedElements}
			}
			return &mappedElements
		},
	})

	globals.Define("round", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			switch val := args[0].(type) {
			case int:
				return val
			case float64:
				// basic round logic for floats to integers
				if val < 0 {
					return int(val - 0.5)
				}
				return int(val + 0.5)
			default:
				panic(RuntimeError("TypeError", "round() expects a number", v.currCtx))
			}
		},
	})

	globals.Define("sorted", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			var elements []any

			switch val := args[0].(type) {
			case *[]any:
				elements = *val
			case *Tuple:
				elements = val.Elements
			case *map[string]any:
				m := *val
				sortedKeys := make([]string, 0, len(m))
				for k := range m {
					sortedKeys = append(sortedKeys, k)
				}
				sort.Strings(sortedKeys)

				result := make([]any, len(sortedKeys))
				for i, k := range sortedKeys {
					result[i] = k
				}
				return &result
			default:
				panic(RuntimeError("TypeError", fmt.Sprintf("sorted() expects an array, tuple, or map, got %T", args[0]), v.currCtx))
			}

			if len(elements) == 0 {
				return &[]any{}
			}

			sortedElements := make([]any, len(elements))
			copy(sortedElements, elements)

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
				panic(RuntimeError("TypeError", "sorted() only supports sequences of integers or strings", v.currCtx))
			}

			return &sortedElements
		},
	})

	globals.Define("sum", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {

			var elements []any
			switch collection := args[0].(type) {
			case *[]any:
				if collection != nil {
					elements = *collection
				}
			case *Tuple:
				if collection != nil {
					elements = collection.Elements
				}
			default:
				panic(RuntimeError("TypeError", fmt.Sprintf("sum() expects an array or a tuple, got %T", args[0]), v.currCtx))
			}

			var totalInt int = 0
			var totalFloat float64 = 0.0
			hasFloat := false

			for _, item := range elements {
				switch num := item.(type) {
				case int:
					if hasFloat {
						totalFloat += float64(num)
					} else {
						totalInt += num
					}
				case float64:
					if !hasFloat {
						hasFloat = true
						totalFloat = float64(totalInt)
					}
					totalFloat += num
				default:
					panic(RuntimeError("TypeError", "sum() collection elements must be numbers", v.currCtx))
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
				panic(RuntimeError("TypeError", "split() expects two string arguments", v.currCtx))
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
				panic(RuntimeError("TypeError", "insert() expects an array and an integer index", v.currCtx))
			}

			arr := *arrPtr
			if idx < 0 || idx > len(arr) {
				panic(RuntimeError("IndexError", fmt.Sprintf("insert index %d out of bounds for length %d", idx, len(arr)), v.currCtx))
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
				panic(RuntimeError("TypeError", "strip() expects a string argument", v.currCtx))
			}
			return strings.TrimSpace(str)
		},
	})

	globals.Define("startswith", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			str, ok1 := args[0].(string)
			prefix, ok2 := args[1].(string)
			if !ok1 || !ok2 {
				panic(RuntimeError("TypeError", "startswith() expects two string arguments", v.currCtx))
			}
			return strings.HasPrefix(str, prefix)
		},
	})

	globals.Define("endswith", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			str, ok1 := args[0].(string)
			suffix, ok2 := args[1].(string)
			if !ok1 || !ok2 {
				panic(RuntimeError("TypeError", "endswith() expects two string arguments", v.currCtx))
			}
			return strings.HasSuffix(str, suffix)
		},
	})

	globals.Define("isalnum", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic(RuntimeError("TypeError", "isalnum() expects a string", v.currCtx))
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
				panic(RuntimeError("TypeError", "isalpha() expects a string", v.currCtx))
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
				panic(RuntimeError("TypeError", "isdigit() expects a string", v.currCtx))
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
				panic(RuntimeError("TypeError", "islower() expects a string", v.currCtx))
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
				panic(RuntimeError("TypeError", "isupper() expects a string", v.currCtx))
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
				panic(RuntimeError("TypeError", "isspace() expects a string", v.currCtx))
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
				panic(RuntimeError("TypeError", "replace() expects three string arguments", v.currCtx))
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
				panic(RuntimeError("TypeError", "rfind() expects two string arguments", v.currCtx))
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
				panic(RuntimeError("TypeError", "rindex() expects two string arguments", v.currCtx))
			}

			idx := strings.LastIndex(str, substr)
			if idx == -1 {
				panic(RuntimeError("ValueError", fmt.Sprintf("Substring '%s' not found in rindex()", substr), v.currCtx))
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
				panic(RuntimeError("TypeError", "divmod() expects two integers", v.currCtx))
			}
			if b == 0 {
				panic(RuntimeError("ZeroDivisionError", "integer division or modulo by zero", v.currCtx))
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
				panic(RuntimeError("TypeError", "pow() expects two integer arguments", v.currCtx))
			}

			return v.power(base, exp)
		},
	})

	globals.Define("count", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			switch collection := args[0].(type) {
			case string:
				sub, ok := args[1].(string)
				if !ok {
					panic(RuntimeError("TypeError", "count() expects two strings", v.currCtx))
				}
				return strings.Count(collection, sub)

			case *Tuple:
				occurrences := 0
				for _, val := range collection.Elements {
					if reflect.DeepEqual(val, args[1]) {
						occurrences++
					}
				}
				return occurrences
			case *[]any:
				occurrences := 0
				for _, val := range *collection {
					if reflect.DeepEqual(val, args[1]) {
						occurrences++
					}
				}
				return occurrences

			default:
				panic(RuntimeError("TypeError", fmt.Sprintf("count() expects a string, array, or tuple as the first argument, got %T", args[0]), v.currCtx))
			}

		},
	})

	globals.Define("enumerate", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			var elements []any

			switch val := args[0].(type) {
			case *[]any:
				elements = *val
			case *Tuple:
				elements = val.Elements
			case *map[string]any:
				m := *val
				sortedKeys := make([]string, 0, len(m))
				for k := range m {
					sortedKeys = append(sortedKeys, k)
				}
				sort.Strings(sortedKeys)

				elements = make([]any, len(sortedKeys))
				for i, k := range sortedKeys {
					elements[i] = k
				}
			default:
				panic(RuntimeError("TypeError", fmt.Sprintf("enumerate() expects an array, tuple, or map, got %T", args[0]), v.currCtx))
			}

			result := make([]any, len(elements))
			for i, val := range elements {
				pair := []any{i, val}
				result[i] = &pair
			}
			return &result
		},
	})

	globals.Define("index", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			switch collection := args[0].(type) {
			case *[]any:
				for i, val := range *collection {
					if reflect.DeepEqual(val, args[1]) {
						return i
					}
				}
			case *Tuple:
				for i, val := range collection.Elements {
					if reflect.DeepEqual(val, args[1]) {
						return i
					}
				}
			default:
				panic(RuntimeError("TypeError", fmt.Sprintf("index() expects an array or tuple as the first argument, got %T", args[0]), v.currCtx))
			}

			panic(RuntimeError("ValueError", fmt.Sprintf("element '%v' is not in collection", args[1]), v.currCtx))
		},
	})

	globals.Define("isinstance", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			targetType, ok := args[1].(string)
			if !ok {
				panic(RuntimeError("TypeError", "isinstance() second argument must be a type string", v.currCtx))
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
			case *Tuple:
				actualType = "tuple"
			case *map[string]any:
				actualType = "map"
			default:
				if _, ok := args[0].(Callable); ok {
					actualType = "function"
				} else if args[0] == nil {
					actualType = "nil"
				} else {
					actualType = "unknown"
				}
			}

			return actualType == targetType
		},
	})

	globals.Define("eval", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			code, ok := args[0].(string)
			if !ok {
				panic(RuntimeError("TypeError", "eval() expects a string argument", v.currCtx))
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
				panic(RuntimeError("TypeError", "exec() expects a string argument", v.currCtx))
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
				panic(RuntimeError("TypeError", "format() specifier must be a string", v.currCtx))
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
			extractElements := func(arg any) []any {
				switch val := arg.(type) {
				case *[]any:
					return *val
				case *Tuple:
					return val.Elements
				case *map[string]any:
					m := *val
					keys := make([]string, 0, len(m))
					for k := range m {
						keys = append(keys, k)
					}
					sort.Strings(keys)

					res := make([]any, len(keys))
					for i, k := range keys {
						res[i] = k
					}
					return res
				default:
					panic(RuntimeError("TypeError", fmt.Sprintf("zip() arguments must be an array, tuple, or map, got %T", arg), v.currCtx))
				}
			}

			arr1 := extractElements(args[0])
			arr2 := extractElements(args[1])

			minLen := len(arr1)
			if len(arr2) < minLen {
				minLen = len(arr2)
			}

			result := make([]any, minLen)
			for i := 0; i < minLen; i++ {
				result[i] = &Tuple{Elements: []any{arr1[i], arr2[i]}}
			}
			return &Tuple{Elements: result}
		},
	})

	globals.Define("join", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			var elements []any
			switch val := args[0].(type) {
			case *[]any:
				elements = *val
			case *Tuple:
				elements = val.Elements
			default:
				panic(RuntimeError("TypeError", fmt.Sprintf("join() expects an array or tuple as the first argument, got %T", args[0]), v.currCtx))
			}

			delim, ok2 := args[1].(string)
			if !ok2 {
				panic(RuntimeError("TypeError", "join() second argument must be a delimiter string", v.currCtx))
			}

			strParts := make([]string, len(elements))
			for i, val := range elements {
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
				panic(RuntimeError("TypeError", "rsplit() expects two strings and an integer maxsplit count", v.currCtx))
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
				panic(RuntimeError("TypeError", "lstrip() expects a string argument", v.currCtx))
			}

			return strings.TrimLeftFunc(str, unicode.IsSpace)
		},
	})

	globals.Define("rstrip", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic(RuntimeError("TypeError", "rstrip() expects a string argument", v.currCtx))
			}

			return strings.TrimRightFunc(str, unicode.IsSpace)
		},
	})

	globals.Define("swapcase", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			str, ok := args[0].(string)
			if !ok {
				panic(RuntimeError("TypeError", "swapcase() expects a string argument", v.currCtx))
			}

			runes := []rune(str)
			for i, r := range runes {
				if unicode.IsUpper(r) {
					runes[i] = unicode.ToLower(r)
				} else if unicode.IsLower(r) {
					runes[i] = unicode.ToUpper(r)
				}
			}
			return string(runes)
		},
	})

	globals.Define("remove", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			arrPtr, ok := args[0].(*[]any)
			if !ok {
				panic(RuntimeError("TypeError", "remove() expects an array pointer as the first argument", v.currCtx))
			}

			arr := *arrPtr
			target := args[1]
			foundIdx := -1

			// first occurrence matching the target
			for i, val := range arr {
				if reflect.DeepEqual(val, target) {
					foundIdx = i
					break
				}
			}

			if foundIdx == -1 {
				panic(RuntimeError("ValueError", "remove(x): x not in array", v.currCtx))
			}

			// re-assign the slice back to the pointer
			*arrPtr = append(arr[:foundIdx], arr[foundIdx+1:]...)
			return nil
		},
	})

	globals.Define("open", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			filename, ok1 := args[0].(string)
			mode, ok2 := args[1].(string)
			if !ok1 || !ok2 {
				panic(RuntimeError("TypeError", "open() expects a filename string and a mode string", v.currCtx))
			}

			var flags int
			switch mode {
			case "r":
				flags = os.O_RDONLY
			case "w":
				flags = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
			case "a":
				flags = os.O_WRONLY | os.O_CREATE | os.O_APPEND
			default:
				panic(RuntimeError("ValueError", "mode for open must be 'r', 'w', or 'a'", v.currCtx))
			}

			file, err := os.OpenFile(filename, flags, 0666)
			if err != nil {
				panic(RuntimeError("IOError", fmt.Sprintf("Could not open file '%s': %v", filename, err), v.currCtx))
			}

			return &FileHandle{File: file, Mode: mode}
		},
	})

	globals.Define("read", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			handle, ok := args[0].(*FileHandle)
			if !ok {
				panic(RuntimeError("TypeError", "read() expects a valid FileHandle object", v.currCtx))
			}
			if handle.Mode != "r" {
				panic(RuntimeError("IOError", "file not opened for reading", v.currCtx))
			}

			content, err := io.ReadAll(handle.File)
			if err != nil {
				panic(RuntimeError("IOError", fmt.Sprintf("Failed reading file: %v", err), v.currCtx))
			}
			return string(content)
		},
	})

	globals.Define("write", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			handle, ok1 := args[0].(*FileHandle)
			content, ok2 := args[1].(string)
			if !ok1 || !ok2 {
				panic(RuntimeError("TypeError", "write() expects a FileHandle object and a content string", v.currCtx))
			}
			if handle.Mode == "r" {
				panic(RuntimeError("TypeError", "file opened as read-only", v.currCtx))
			}

			_, err := handle.File.WriteString(content)
			if err != nil {
				panic(RuntimeError("TypeError", fmt.Sprintf("Failed writing to file: %v", err), v.currCtx))
			}
			return nil
		},
	})

	globals.Define("close", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			handle, ok := args[0].(*FileHandle)
			if !ok {
				panic(RuntimeError("TypeError", "close() expects a valid FileHandle object", v.currCtx))
			}

			err := handle.File.Close()
			if err != nil {
				panic(RuntimeError("IOError", fmt.Sprintf("Failed closing file cleanly: %v", err), v.currCtx))
			}
			return nil
		},
	})

	globals.Define("keys", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			mPtr, ok := args[0].(*map[string]any)
			if !ok {
				panic(RuntimeError("TypeError", "keys() expects a map pointer", v.currCtx))
			}

			m := *mPtr

			//gather keys as plain strings
			sortedKeys := make([]string, 0, len(m))
			for k := range m {
				sortedKeys = append(sortedKeys, k)
			}

			sort.Strings(sortedKeys)

			res := make([]any, len(sortedKeys))
			for i, k := range sortedKeys {
				res[i] = k
			}
			return &res
		},
	})

	globals.Define("values", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			mPtr, ok := args[0].(*map[string]any)
			if !ok {
				panic(RuntimeError("TypeError", "values() expects a map pointer", v.currCtx))
			}

			m := *mPtr

			//gather keys as plain strings
			sortedKeys := make([]string, 0, len(m))
			for k := range m {
				sortedKeys = append(sortedKeys, k)
			}

			sort.Strings(sortedKeys)

			res := make([]any, len(sortedKeys))
			for i, k := range sortedKeys {
				res[i] = m[k]
			}
			return &res
		},
	})

	globals.Define("clear", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			if mapPtr, ok := args[0].(*map[string]any); ok && mapPtr != nil {
				*mapPtr = make(map[string]any)
				return nil
			}
			if arrPtr, ok := args[0].(*[]any); ok && arrPtr != nil {
				*arrPtr = []any{}
				return nil
			}
			panic(RuntimeError("TypeError", "clear() expects an array or map pointer", v.currCtx))
		},
	})

	// asert on interface
	globals.Define("assert", &NativeFunction{
		ArgsCount: 2,
		Body: func(v *Visitor, args []any) any {
			// ensure the parameters are correct types
			mapPtr, isStruct := args[0].(*map[string]any)
			interfaceName, isString := args[1].(string)

			if !isString {
				panic(RuntimeError("TypeError", "assert() second argument must be an interface name string", v.currCtx))
			}

			if !isStruct || mapPtr == nil {
				panic(RuntimeError("TypeError", fmt.Sprintf("Value does not satisfy interface '%s' (not a struct object)", interfaceName), v.currCtx))
			}

			// execute the implicit structural match contract scan
			passed, reason := v.satisfiesInterface(mapPtr, interfaceName)
			if !passed {
				panic(RuntimeError("TypeError", fmt.Sprintf("Interface contract violation! %s", reason), v.currCtx))
			}

			// return the object if it matches
			return mapPtr
		},
	})

	globals.Define("tuple", &NativeFunction{
		ArgsCount: 1,
		Body: func(v *Visitor, args []any) any {
			switch obj := args[0].(type) {
			case *[]any:
				elements := make([]any, len(*obj))
				copy(elements, *obj)
				return &Tuple{Elements: elements}

			case *Tuple:
				elements := make([]any, len(obj.Elements))
				copy(elements, obj.Elements)
				return &Tuple{Elements: elements}

			case *map[string]any:
				m := *obj
				sortedKeys := make([]string, 0, len(m))
				for k := range m {
					sortedKeys = append(sortedKeys, k)
				}
				sort.Strings(sortedKeys)

				elements := make([]any, len(sortedKeys))
				for i, k := range sortedKeys {
					elements[i] = k
				}
				return &Tuple{Elements: elements}

			default:
				panic(RuntimeError("TypeError", fmt.Sprintf("Object of type %T cannot be converted to a tuple", args[0]), v.currCtx))
			}
		},
	})

	globals.Define("printf", &NativeFunction{
		ArgsCount: -1,
		Body: func(v *Visitor, args []any) any {
			if len(args) < 1 {
				panic(RuntimeError("TypeError", "printf() expects at least 1 argument", v.currCtx))
			}
			formatStr, ok := args[0].(string)
			if !ok {
				panic(RuntimeError("TypeError", "printf() first argument must be a string", v.currCtx))
			}

			// clean strings
			processedArgs := make([]any, len(args[1:]))
			for i, arg := range args[1:] {
				processedArgs[i] = stringify(arg)
			}

			finalOutput := fmt.Sprintf(formatStr, processedArgs...)

			// print cleanly message
			fmt.Print(finalOutput)
			return nil
		},
	})

	globals.Define("sprintf", &NativeFunction{
		ArgsCount: -1,
		Body: func(v *Visitor, args []any) any {
			if len(args) < 1 {
				panic(RuntimeError("TypeError", "sprintf() expects at least 1 argument", v.currCtx))
			}
			formatStr, ok := args[0].(string)
			if !ok {
				panic(RuntimeError("TypeError", "sprintf() first argument must be a string", v.currCtx))
			}

			processedArgs := make([]any, len(args[1:]))
			for i, arg := range args[1:] {
				processedArgs[i] = stringify(arg)
			}

			return fmt.Sprintf(formatStr, processedArgs...)
		},
	})

	return globals
}

func stringify(val any) string {
	if val == nil {
		return "nil"
	}

	switch actual := val.(type) {
	case string:
		return actual
	case *[]any:
		elements := *actual
		var sb strings.Builder
		sb.WriteString("[")
		for i, element := range elements {
			sb.WriteString(stringify(element))
			if i < len(elements)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("]")
		return sb.String()

	case *Tuple:
		var sb strings.Builder
		sb.WriteString("(")
		for i, element := range actual.Elements {
			sb.WriteString(stringify(element))
			if i < len(actual.Elements)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString(")")
		return sb.String()

	case *map[string]any:
		targetMap := *actual

		// check if struct instance with a type tag metadata
		if t, exists := targetMap["__type__"].(string); exists {
			return fmt.Sprintf("<object Instance of %s>", t)
		}

		keys := make([]string, 0, len(targetMap))
		for k := range targetMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var sb strings.Builder
		sb.WriteString("{")
		for i, k := range keys {
			if strVal, isStr := targetMap[k].(string); isStr {
				fmt.Fprintf(&sb, `"%s" : "%s"`, k, strVal)
			} else {
				fmt.Fprintf(&sb, `"%s" : %s`, k, stringify(targetMap[k]))
			}
			if i < len(keys)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("}")
		return sb.String()

	default:
		return fmt.Sprintf("%v", val)
	}
}

// converts complex language objects to strings
func preprocessArgs(args []any) []any {
	processed := make([]any, len(args))
	for i, arg := range args {
		processed[i] = stringify(arg)
	}
	return processed
}

func (v *Visitor) satisfiesInterface(structInstanceMap *map[string]any, interfaceName string) (bool, string) {
	contract, exists := v.InterfaceRegistry[interfaceName]
	if !exists {
		return false, fmt.Sprintf("Interface '%s' is not defined", interfaceName)
	}

	// extract what structural blueprint this instance belongs to
	structType, ok := (*structInstanceMap)["__type__"].(string)
	if !ok {
		return false, "Value is not an instantiated struct object"
	}

	structMethods, _ := v.MethodRegistry[structType]

	// check if the struct satisfies required methods
	for reqMethod, reqArgs := range contract.Methods {
		methodCtx, hasMethod := structMethods[reqMethod]
		if !hasMethod {
			return false, fmt.Sprintf("Struct '%s' is missing required method '%s()'", structType, reqMethod)
		}

		// calculate arguments on the method structure
		actualArgs := len(methodCtx.AllIDENTIFIER()) - 1
		if actualArgs != reqArgs {
			return false, fmt.Sprintf("Struct '%s' method '%s' expects %d arguments, but interface '%s' requires %d",
				structType, reqMethod, actualArgs, interfaceName, reqArgs)
		}
	}

	return true, ""
}
