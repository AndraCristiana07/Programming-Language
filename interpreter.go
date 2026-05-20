package main

import "fmt"

func Eval(node Stmt, env *Environment) any {
	switch n := node.(type) {
	case *Program:
		// var result int
		for _, stmt := range n.Body {
			Eval(stmt, env)
		}
		return nil
	case *VarDeclaration:
		value := Eval(n.Value, env)
		env.vars[n.Identifier] = value
		return value
	case *PrintStmt:
		value := Eval(n.Value, env)
		fmt.Println(value)
		return nil

	case *BinaryExpr:
		leftVal := Eval(n.Left, env)
		rightVal := Eval(n.Right, env)

		// TODO: handle type checking and string concatenation
		_, leftIsString := leftVal.(string)
		_, rightIsString := rightVal.(string)

		if leftIsString || rightIsString {
			if n.Operator == "+" {
				return fmt.Sprintf("%v%v", leftVal, rightVal)
			} else {
				panic("Unsupported operator for strings: " + n.Operator)
			}
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
		default:
			panic("Unknown operator: " + n.Operator)
		}
	case *Identifier:
		val, ok := env.vars[n.Symbol]
		if !ok {
			panic("Undefined variable: " + n.Symbol)
		}
		return val
	case *NumericLiteral:
		return n.Value
	case *StringLiteral:
		return n.Value
	case *Assignment:
		value := Eval(n.Value, env)
		if _, ok := env.vars[n.Identifier]; !ok {
			panic("Undefined variable: " + n.Identifier)
		}
		env.vars[n.Identifier] = value
		return value
	default:
		panic("Unknown node type: " + n.GetType())
	}
}
