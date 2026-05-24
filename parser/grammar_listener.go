// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

// GrammarListener is a complete listener for a parse tree produced by GrammarParser.
type GrammarListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterAssignStmt is called when entering the assignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterCompoundAssignStmt is called when entering the compoundAssignStmt production.
	EnterCompoundAssignStmt(c *CompoundAssignStmtContext)

	// EnterPrintStmt is called when entering the printStmt production.
	EnterPrintStmt(c *PrintStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterWhileStmt is called when entering the whileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterForStmt is called when entering the forStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterBlockStmt is called when entering the blockStmt production.
	EnterBlockStmt(c *BlockStmtContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterForPost is called when entering the forPost production.
	EnterForPost(c *ForPostContext)

	// EnterPostfixStmt is called when entering the postfixStmt production.
	EnterPostfixStmt(c *PostfixStmtContext)

	// EnterNot is called when entering the Not production.
	EnterNot(c *NotContext)

	// EnterMulDivMod is called when entering the MulDivMod production.
	EnterMulDivMod(c *MulDivModContext)

	// EnterIdentifier is called when entering the Identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterOr is called when entering the Or production.
	EnterOr(c *OrContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterComparison is called when entering the Comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterAnd is called when entering the And production.
	EnterAnd(c *AndContext)

	// EnterExponential is called when entering the Exponential production.
	EnterExponential(c *ExponentialContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// EnterBoolean is called when entering the Boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterParentheses is called when entering the Parentheses production.
	EnterParentheses(c *ParenthesesContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitAssignStmt is called when exiting the assignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitCompoundAssignStmt is called when exiting the compoundAssignStmt production.
	ExitCompoundAssignStmt(c *CompoundAssignStmtContext)

	// ExitPrintStmt is called when exiting the printStmt production.
	ExitPrintStmt(c *PrintStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitWhileStmt is called when exiting the whileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitForStmt is called when exiting the forStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitBlockStmt is called when exiting the blockStmt production.
	ExitBlockStmt(c *BlockStmtContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitForPost is called when exiting the forPost production.
	ExitForPost(c *ForPostContext)

	// ExitPostfixStmt is called when exiting the postfixStmt production.
	ExitPostfixStmt(c *PostfixStmtContext)

	// ExitNot is called when exiting the Not production.
	ExitNot(c *NotContext)

	// ExitMulDivMod is called when exiting the MulDivMod production.
	ExitMulDivMod(c *MulDivModContext)

	// ExitIdentifier is called when exiting the Identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitOr is called when exiting the Or production.
	ExitOr(c *OrContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitComparison is called when exiting the Comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitAnd is called when exiting the And production.
	ExitAnd(c *AndContext)

	// ExitExponential is called when exiting the Exponential production.
	ExitExponential(c *ExponentialContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

	// ExitBoolean is called when exiting the Boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitParentheses is called when exiting the Parentheses production.
	ExitParentheses(c *ParenthesesContext)
}
