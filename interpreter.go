package main

import "fmt"

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
	default:
		panic("Unknown node type: " + n.GetType())
	}
}
