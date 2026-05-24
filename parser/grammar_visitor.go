// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GrammarParser.
type GrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GrammarParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by GrammarParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by GrammarParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by GrammarParser#assignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#compoundAssignStmt.
	VisitCompoundAssignStmt(ctx *CompoundAssignStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#printStmt.
	VisitPrintStmt(ctx *PrintStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#whileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#blockStmt.
	VisitBlockStmt(ctx *BlockStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#forInit.
	VisitForInit(ctx *ForInitContext) interface{}

	// Visit a parse tree produced by GrammarParser#forPost.
	VisitForPost(ctx *ForPostContext) interface{}

	// Visit a parse tree produced by GrammarParser#postfixStmt.
	VisitPostfixStmt(ctx *PostfixStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#BitOr.
	VisitBitOr(ctx *BitOrContext) interface{}

	// Visit a parse tree produced by GrammarParser#Or.
	VisitOr(ctx *OrContext) interface{}

	// Visit a parse tree produced by GrammarParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by GrammarParser#BitShift.
	VisitBitShift(ctx *BitShiftContext) interface{}

	// Visit a parse tree produced by GrammarParser#Exponential.
	VisitExponential(ctx *ExponentialContext) interface{}

	// Visit a parse tree produced by GrammarParser#String.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by GrammarParser#Unary.
	VisitUnary(ctx *UnaryContext) interface{}

	// Visit a parse tree produced by GrammarParser#MulDivMod.
	VisitMulDivMod(ctx *MulDivModContext) interface{}

	// Visit a parse tree produced by GrammarParser#Identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by GrammarParser#BitXor.
	VisitBitXor(ctx *BitXorContext) interface{}

	// Visit a parse tree produced by GrammarParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by GrammarParser#Comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by GrammarParser#And.
	VisitAnd(ctx *AndContext) interface{}

	// Visit a parse tree produced by GrammarParser#BitAnd.
	VisitBitAnd(ctx *BitAndContext) interface{}

	// Visit a parse tree produced by GrammarParser#Boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by GrammarParser#Parentheses.
	VisitParentheses(ctx *ParenthesesContext) interface{}
}
