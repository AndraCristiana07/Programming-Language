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

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterTerminator is called when entering the terminator production.
	EnterTerminator(c *TerminatorContext)

	// EnterFuncStmt is called when entering the funcStmt production.
	EnterFuncStmt(c *FuncStmtContext)

	// EnterExprStmt is called when entering the exprStmt production.
	EnterExprStmt(c *ExprStmtContext)

	// EnterTryCatchStmt is called when entering the tryCatchStmt production.
	EnterTryCatchStmt(c *TryCatchStmtContext)

	// EnterThrowStmt is called when entering the throwStmt production.
	EnterThrowStmt(c *ThrowStmtContext)

	// EnterBreakStmt is called when entering the breakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the continueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterSwitchStmt is called when entering the switchStmt production.
	EnterSwitchStmt(c *SwitchStmtContext)

	// EnterCaseBlock is called when entering the caseBlock production.
	EnterCaseBlock(c *CaseBlockContext)

	// EnterDefaultBlock is called when entering the defaultBlock production.
	EnterDefaultBlock(c *DefaultBlockContext)

	// EnterForInStmt is called when entering the forInStmt production.
	EnterForInStmt(c *ForInStmtContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterAssignStmt is called when entering the assignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterArrayAssignStmt is called when entering the arrayAssignStmt production.
	EnterArrayAssignStmt(c *ArrayAssignStmtContext)

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

	// EnterNull is called when entering the Null production.
	EnterNull(c *NullContext)

	// EnterBitOr is called when entering the BitOr production.
	EnterBitOr(c *BitOrContext)

	// EnterOr is called when entering the Or production.
	EnterOr(c *OrContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterBitShift is called when entering the BitShift production.
	EnterBitShift(c *BitShiftContext)

	// EnterExponential is called when entering the Exponential production.
	EnterExponential(c *ExponentialContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// EnterArrayLiteral is called when entering the ArrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterArrayIndex is called when entering the ArrayIndex production.
	EnterArrayIndex(c *ArrayIndexContext)

	// EnterUnary is called when entering the Unary production.
	EnterUnary(c *UnaryContext)

	// EnterTernaryOp is called when entering the TernaryOp production.
	EnterTernaryOp(c *TernaryOpContext)

	// EnterMapLiteral is called when entering the MapLiteral production.
	EnterMapLiteral(c *MapLiteralContext)

	// EnterMulDivMod is called when entering the MulDivMod production.
	EnterMulDivMod(c *MulDivModContext)

	// EnterIdentifier is called when entering the Identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterBitXor is called when entering the BitXor production.
	EnterBitXor(c *BitXorContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterComparison is called when entering the Comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterAnd is called when entering the And production.
	EnterAnd(c *AndContext)

	// EnterBitAnd is called when entering the BitAnd production.
	EnterBitAnd(c *BitAndContext)

	// EnterMembership is called when entering the Membership production.
	EnterMembership(c *MembershipContext)

	// EnterFunctionCall is called when entering the FunctionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterBoolean is called when entering the Boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterFieldAccess is called when entering the FieldAccess production.
	EnterFieldAccess(c *FieldAccessContext)

	// EnterParentheses is called when entering the Parentheses production.
	EnterParentheses(c *ParenthesesContext)

	// EnterMapEntry is called when entering the mapEntry production.
	EnterMapEntry(c *MapEntryContext)

	// EnterStandardArray is called when entering the StandardArray production.
	EnterStandardArray(c *StandardArrayContext)

	// EnterListComprehension is called when entering the ListComprehension production.
	EnterListComprehension(c *ListComprehensionContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitTerminator is called when exiting the terminator production.
	ExitTerminator(c *TerminatorContext)

	// ExitFuncStmt is called when exiting the funcStmt production.
	ExitFuncStmt(c *FuncStmtContext)

	// ExitExprStmt is called when exiting the exprStmt production.
	ExitExprStmt(c *ExprStmtContext)

	// ExitTryCatchStmt is called when exiting the tryCatchStmt production.
	ExitTryCatchStmt(c *TryCatchStmtContext)

	// ExitThrowStmt is called when exiting the throwStmt production.
	ExitThrowStmt(c *ThrowStmtContext)

	// ExitBreakStmt is called when exiting the breakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the continueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitSwitchStmt is called when exiting the switchStmt production.
	ExitSwitchStmt(c *SwitchStmtContext)

	// ExitCaseBlock is called when exiting the caseBlock production.
	ExitCaseBlock(c *CaseBlockContext)

	// ExitDefaultBlock is called when exiting the defaultBlock production.
	ExitDefaultBlock(c *DefaultBlockContext)

	// ExitForInStmt is called when exiting the forInStmt production.
	ExitForInStmt(c *ForInStmtContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitAssignStmt is called when exiting the assignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitArrayAssignStmt is called when exiting the arrayAssignStmt production.
	ExitArrayAssignStmt(c *ArrayAssignStmtContext)

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

	// ExitNull is called when exiting the Null production.
	ExitNull(c *NullContext)

	// ExitBitOr is called when exiting the BitOr production.
	ExitBitOr(c *BitOrContext)

	// ExitOr is called when exiting the Or production.
	ExitOr(c *OrContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitBitShift is called when exiting the BitShift production.
	ExitBitShift(c *BitShiftContext)

	// ExitExponential is called when exiting the Exponential production.
	ExitExponential(c *ExponentialContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

	// ExitArrayLiteral is called when exiting the ArrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitArrayIndex is called when exiting the ArrayIndex production.
	ExitArrayIndex(c *ArrayIndexContext)

	// ExitUnary is called when exiting the Unary production.
	ExitUnary(c *UnaryContext)

	// ExitTernaryOp is called when exiting the TernaryOp production.
	ExitTernaryOp(c *TernaryOpContext)

	// ExitMapLiteral is called when exiting the MapLiteral production.
	ExitMapLiteral(c *MapLiteralContext)

	// ExitMulDivMod is called when exiting the MulDivMod production.
	ExitMulDivMod(c *MulDivModContext)

	// ExitIdentifier is called when exiting the Identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitBitXor is called when exiting the BitXor production.
	ExitBitXor(c *BitXorContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitComparison is called when exiting the Comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitAnd is called when exiting the And production.
	ExitAnd(c *AndContext)

	// ExitBitAnd is called when exiting the BitAnd production.
	ExitBitAnd(c *BitAndContext)

	// ExitMembership is called when exiting the Membership production.
	ExitMembership(c *MembershipContext)

	// ExitFunctionCall is called when exiting the FunctionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitBoolean is called when exiting the Boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitFieldAccess is called when exiting the FieldAccess production.
	ExitFieldAccess(c *FieldAccessContext)

	// ExitParentheses is called when exiting the Parentheses production.
	ExitParentheses(c *ParenthesesContext)

	// ExitMapEntry is called when exiting the mapEntry production.
	ExitMapEntry(c *MapEntryContext)

	// ExitStandardArray is called when exiting the StandardArray production.
	ExitStandardArray(c *StandardArrayContext)

	// ExitListComprehension is called when exiting the ListComprehension production.
	ExitListComprehension(c *ListComprehensionContext)
}
