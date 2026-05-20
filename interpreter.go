package main

import "fmt"

func Eval(node Stmt, env *Environment) int {
	switch n := node.(type) {
	case *Program:
		var result int
		for _, stmt := range n.Body {
			result = Eval(stmt, env)
		}
		return result
	case *VarDeclaration:
		value := Eval(n.Value, env)
		env.vars[n.Identifier] = value
		return value
	case *PrintStmt:
		value := Eval(n.Value, env)
		fmt.Println(value)
		return 0

	case *BinaryExpr:
		left := Eval(n.Left, env)
		right := Eval(n.Right, env)
		switch n.Operator {
		case "+":
			return left + right
		case "-":
			return left - right
		case "*":
			return left * right
		case "/":
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
	default:
		panic("Unknown node type: " + n.GetType())
	}
}
