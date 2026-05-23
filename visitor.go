package main

import (
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

func (v *Visitor) VisitAddSub(ctx *parser.AddSubContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)
	op := ctx.GetOp().GetText()

	if left == nil || right == nil {
		panic("Undefined variable used in expression")
	}

	switch op {
	case "+":
		return left.(int) + right.(int)
	case "-":
		return left.(int) - right.(int)
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
