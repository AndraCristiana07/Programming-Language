package main

import (
	"fmt"
	"my_language/parser"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	parser.BaseGrammarVisitor
	vars map[string]any
}

func NewVisitor() *Visitor {
	return &Visitor{
		vars: make(map[string]any),
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *Visitor) VisitProgram(ctx *parser.ProgramContext) any {
	for _, stmt := range ctx.AllStatement() {
		stmt.Accept(v)
	}
	return nil
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) any {
	if ctx.VarDecl() != nil {
		return ctx.VarDecl().Accept(v)
	} else if ctx.AssignStmt() != nil {
		return ctx.AssignStmt().Accept(v)
	} else if ctx.PrintStmt() != nil {
		return ctx.PrintStmt().Accept(v)
	} else if ctx.BlockStmt() != nil {
		return ctx.BlockStmt().Accept(v)
	} else if ctx.IfStmt() != nil {
		return ctx.IfStmt().Accept(v)
	} else if ctx.WhileStmt() != nil {
		return ctx.WhileStmt().Accept(v)
	} else if ctx.ForStmt() != nil {
		return ctx.ForStmt().Accept(v)
	}

	return nil
}

func (v *Visitor) VisitForInit(ctx *parser.ForInitContext) any {
	if ctx.VarDecl() != nil {
		return ctx.VarDecl().Accept(v)
	} else if ctx.AssignStmt() != nil {
		return ctx.AssignStmt().Accept(v)
	}
	return nil
}

func (v *Visitor) VisitVarDecl(ctx *parser.VarDeclContext) any {
	varName := ctx.IDENTIFIER().GetText()
	value := ctx.Expr().Accept(v)
	v.vars[varName] = value
	return nil
}

func (v *Visitor) VisitAssignStmt(ctx *parser.AssignStmtContext) any {
	varName := ctx.IDENTIFIER().GetText()
	value := ctx.Expr().Accept(v)
	v.vars[varName] = value
	return nil
}

func (v *Visitor) VisitNumber(ctx *parser.NumberContext) any {
	// convert into integer
	val, _ := strconv.Atoi(ctx.NUMBER().GetText())
	return val
}

func (v *Visitor) VisitIdentifier(ctx *parser.IdentifierContext) any {
	varName := ctx.IDENTIFIER().GetText()
	return v.vars[varName]
}

func (v *Visitor) VisitString(ctx *parser.StringContext) any {
	str := ctx.STRING().GetText()
	// remove the surrounding quotes
	return str[1 : len(str)-1]
}

func (v *Visitor) VisitAddSub(ctx *parser.AddSubContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)
	op := ctx.GetOp().GetText()

	_, leftIsString := left.(string)
	_, rightIsString := right.(string)

	if leftIsString || rightIsString {
		if op == "+" {
			return fmt.Sprintf("%v%v", left, right)
		} else {
			panic("Unsupported operator for strings: " + op)
		}
	}

	lVal, lOk := left.(int)
	rVal, rOk := right.(int)
	if !lOk || !rOk {
		panic("Both operands must be integers for operator: " + op)
	}

	switch op {
	case "+":
		return lVal + rVal
	case "-":
		return lVal - rVal
	default:
		panic("Unknown operator: " + op)
	}
}

func (v *Visitor) VisitMulDiv(ctx *parser.MulDivContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)
	op := ctx.GetOp().GetText()

	if left == nil || right == nil {
		panic("Undefined variable used in expression")
	}

	switch op {
	case "*":
		return left.(int) * right.(int)
	case "/":
		return left.(int) / right.(int)
	default:
		panic("Unknown operator: " + op)
	}
}

func (v *Visitor) VisitBoolean(ctx *parser.BooleanContext) any {
	val := ctx.GetVal().GetText()
	if val == "true" {
		return true
	}
	return false
}

func (v *Visitor) VisitComparison(ctx *parser.ComparisonContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)
	op := ctx.GetOp().GetText()

	switch op {
	case "==":
		return left == right
	case "!=":
		return left != right
	case "<":
		lVal, lOk := left.(int)
		rVal, rOk := right.(int)
		if lOk && rOk {
			return lVal < rVal
		}
		panic("Comparison operator '<' requires integer operands")
	case ">":
		lVal, lOk := left.(int)
		rVal, rOk := right.(int)
		if lOk && rOk {
			return lVal > rVal
		}
		panic("Comparison operator '>' requires integer operands")
	case "<=":
		lVal, lOk := left.(int)
		rVal, rOk := right.(int)
		if lOk && rOk {
			return lVal <= rVal
		}
		panic("Comparison operator '<=' requires integer operands")
	case ">=":
		lVal, lOk := left.(int)
		rVal, rOk := right.(int)
		if lOk && rOk {
			return lVal >= rVal
		}
		panic("Comparison operator '>=' requires integer operands")
	default:
		panic("Unknown operator: " + op)
	}
}

func (v *Visitor) VisitPrintStmt(ctx *parser.PrintStmtContext) any {
	value := ctx.Expr().Accept(v)
	fmt.Println(value)
	return nil
}

func (v *Visitor) VisitBlockStmt(ctx *parser.BlockStmtContext) any {
	for _, stmt := range ctx.AllStatement() {
		stmt.Accept(v)
	}
	return nil
}

func (v *Visitor) VisitIfStmt(ctx *parser.IfStmtContext) any {
	condition := ctx.Expr().Accept(v)
	conditionBool, ok := condition.(bool)
	if !ok {
		panic("Condition in if statement must be boolean")
	}

	if conditionBool {
		ctx.GetThenBranch().Accept(v)
	} else if ctx.GetElseBranch() != nil {
		ctx.GetElseBranch().Accept(v)
	}

	return nil
}

func (v *Visitor) VisitWhileStmt(ctx *parser.WhileStmtContext) any {
	for {
		condition := ctx.Expr().Accept(v)
		conditionBool, ok := condition.(bool)
		if !ok {
			panic("Condition in while statement must be boolean")
		}

		if !conditionBool {
			break
		}
		ctx.GetBody().Accept(v)
	}
	return nil
}

func (v *Visitor) VisitForStmt(ctx *parser.ForStmtContext) any {
	if ctx.GetInit() != nil {
		ctx.GetInit().Accept(v)
	}

	for {
		if ctx.GetCond() != nil {
			condition := ctx.GetCond().Accept(v)
			conditionBool, ok := condition.(bool)
			if !ok {
				panic("For loop condition must evaluate to a boolean expression")
			}
			if !conditionBool {
				break
			}
		}
		if ctx.GetBody() != nil {
			ctx.GetBody().Accept(v)
		}

		if ctx.GetPost() != nil {
			ctx.GetPost().Accept(v)
		}
	}

	return nil
}
