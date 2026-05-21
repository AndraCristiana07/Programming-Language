package my_language

import (
	"fmt"
)

type RuntimeFunction struct {
	Declaration *FuncDecl
	Env         *Environment
}

type ReturnValueSignal struct {
	Value any
}

func Eval(node Stmt, env *Environment) any {
	switch n := node.(type) {
	case *Program:
		for _, stmt := range n.Body {
			Eval(stmt, env)
		}
		return nil
	case *BooleanLiteral:
		return n.Value
	case *VarDeclaration:
		value := Eval(n.Value, env)
		env.Define(n.Identifier, value)
		return value
	case *PrintStmt:
		value := Eval(n.Value, env)
		fmt.Println(value)
		return nil
	case *BinaryExpr:
		leftVal := Eval(n.Left, env)
		rightVal := Eval(n.Right, env)

		// handle type checking and string concatenation
		_, leftIsString := leftVal.(string)
		_, rightIsString := rightVal.(string)

		if leftIsString || rightIsString {
			if n.Operator == "+" {
				return fmt.Sprintf("%v%v", leftVal, rightVal)
			} else {
				panic("Unsupported operator for strings: " + n.Operator)
			}
		}

		// handle boolean logic for "and" and "or"
		switch n.Operator {
		case "and":
			leftBool, ok1 := leftVal.(bool)
			rightBool, ok2 := rightVal.(bool)
			if !ok1 || !ok2 {
				panic("Both operands of 'and' must be boolean")
			}
			return leftBool && rightBool
		case "or":
			leftBool, ok1 := leftVal.(bool)
			rightBool, ok2 := rightVal.(bool)
			if !ok1 || !ok2 {
				panic("Both operands of 'or' must be boolean")
			}
			return leftBool || rightBool
		}

		// if not str, go back to int
		left := leftVal.(int)
		right := rightVal.(int)

		switch n.Operator {
		case "+":
			return left + right
		case "-":
			return left - right
		case "*":
			return left * right
		case "/":
			// check for division by zero
			if right == 0 {
				panic("Division by zero")
			}
			return left / right
		case "**":
			return power(left, right)
		case "%":
			// check for division by zero
			if right == 0 {
				panic("Modulo by zero")
			}
			return left % right
		case "++":
			return left + 1
		case "--":
			return left - 1
		case "<":
			return left < right
		case ">":
			return left > right
		case "==":
			return left == right
		case "!=":
			return left != right
		case "<=":
			return left <= right
		case ">=":
			return left >= right
		case "&":
			return left & right
		case "|":
			return left | right
		case "<<":
			return left << right
		case ">>":
			return left >> right
		case "^":
			return left ^ right

		default:
			panic("Unknown operator: " + n.Operator)
		}
	case *Identifier:
		val, ok := env.Lookup(n.Symbol)

		if !ok {
			panic("Undefined variable: " + n.Symbol)
		}
		return val
	case *NumericLiteral:
		return n.Value
	case *UnaryExpr:
		rightVal := Eval(n.Right, env)
		switch n.Operator {
		case "not":
			rightBool, ok := rightVal.(bool)
			if !ok {
				panic("Operand of 'not' must be boolean")
			}
			return !rightBool
		case "~":
			rightInt, ok := rightVal.(int)
			if !ok {
				panic("Operand of '~' must be an integer")
			}
			return ^rightInt
		default:
			panic("Unknown unary operator: " + n.Operator)
		}
	case *StringLiteral:
		return n.Value
	case *Assignment:
		value := Eval(n.Value, env)
		if !env.Assign(n.Identifier, value) {
			panic("Undefined variable: " + n.Identifier)
		}
		return value

	case *BlockStmt:
		blockEnv := NewScope(env)
		for _, stmt := range n.Body {
			Eval(stmt, blockEnv)
		}
		return nil

	case *IfStmt:
		// eval conditional leaf expression
		condVal := Eval(n.Condition, env)

		booleanCond, ok := condVal.(bool)
		if !ok {
			panic("Condition expression must evaluate to a boolean value")
		}

		if booleanCond {
			Eval(n.ThenBranch, env)
		} else if n.ElseBranch != nil {
			Eval(n.ElseBranch, env)
		}
		return nil

	case *WhileStmt:
		// evaluate the condition expression on every cycle
		for {
			condVal := Eval(n.Condition, env)
			booleanCond, ok := condVal.(bool)
			if !ok {
				panic("While loop condition must evaluate to a boolean")
			}

			// break the loop if it gets to false
			if !booleanCond {
				break
			}

			// exec the body block
			Eval(n.Body, env)
		}
		return nil
	case *ReturnStmt:
		var value any = nil
		if n.Value != nil {
			value = Eval(n.Value, env)
		}
		panic(ReturnValueSignal{Value: value})

	case *FuncDecl:
		// store the function decl and its env in a runtime func struct
		runtimeFunc := &RuntimeFunction{
			Declaration: n,
			Env:         env,
		}
		env.Define(n.Name, runtimeFunc)
		return nil
	case *CallExpr:
		// look up the function by name
		funcVal, ok := env.Lookup(n.Callee)
		if !ok {
			panic("Undefined function: " + n.Callee)
		}

		runtimeFunc, ok := funcVal.(*RuntimeFunction)
		if !ok {
			panic("Attempting to call a non-function value: " + n.Callee)
		}

		// check argument count
		if len(n.Arguments) != len(runtimeFunc.Declaration.Parameters) {
			panic(fmt.Sprintf("Expected %d arguments but got %d", len(runtimeFunc.Declaration.Parameters), len(n.Arguments)))
		}

		// create new env for the func call with the declaration env as parent
		funcEnv := NewScope(runtimeFunc.Env)

		// evaluate arguments and bind to parameter names in the new env
		for i, argExpr := range n.Arguments {
			argVal := Eval(argExpr, env)
			paramName := runtimeFunc.Declaration.Parameters[i]
			funcEnv.Define(paramName, argVal)
		}

		// return stmt
		var returnedVal any = nil
		func() {
			defer func() {
				if r := recover(); r != nil {
					// check if it's a return signal and extract the value
					if returnSignal, ok := r.(ReturnValueSignal); ok {
						returnedVal = returnSignal.Value
					} else {
						// if it's not a return signal, re-panic
						panic(r)
					}
				}
			}()

			// execute the function body in the new env
			Eval(runtimeFunc.Declaration.Body, funcEnv)
		}()

		return returnedVal

	case *ArrayLiteral:
		var evaluatedElements []any
		for _, expr := range n.Elements {
			evaluatedElements = append(evaluatedElements, Eval(expr, env))
		}
		return evaluatedElements

	case *IndexExpr:
		leftVal := Eval(n.Left, env)
		indexVal := Eval(n.Index, env)

		// left target -array slice
		array, ok := leftVal.([]any)
		if !ok {
			panic("Indexing operator [] can only be used on arrays")
		}

		// index evaluated int
		idx, ok := indexVal.(int)
		if !ok {
			panic("Array index subscript must be an integer")
		}

		// boundary check
		if idx < 0 || idx >= len(array) {
			panic(fmt.Sprintf("Index out of bounds. Array length is %d, got index %d", len(array), idx))
		}

		return array[idx]

	default:
		panic("Unknown node type: " + n.GetType())
	}
}

func power(base, exp int) int {
	if exp < 0 {
		panic("Negative exponent not supported")
	}
	result := 1
	for range exp {
		result *= base
	}
	return result
}
