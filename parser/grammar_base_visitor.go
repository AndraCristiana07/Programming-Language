// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

type BaseGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGrammarVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitTerminator(ctx *TerminatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitFuncStmt(ctx *FuncStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitExprStmt(ctx *ExprStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitTryCatchStmt(ctx *TryCatchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitThrowStmt(ctx *ThrowStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitSwitchStmt(ctx *SwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitCaseBlock(ctx *CaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitDefaultBlock(ctx *DefaultBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitForInStmt(ctx *ForInStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitVarDecl(ctx *VarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitAssignStmt(ctx *AssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitArrayAssignStmt(ctx *ArrayAssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitCompoundAssignStmt(ctx *CompoundAssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitPrintStmt(ctx *PrintStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitBlockStmt(ctx *BlockStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructStmt(ctx *StructStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructField(ctx *StructFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitIfInit(ctx *IfInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitForInit(ctx *ForInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitForPost(ctx *ForPostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitPostfixStmt(ctx *PostfixStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitNull(ctx *NullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitOr(ctx *OrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitExponential(ctx *ExponentialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitArrayIndex(ctx *ArrayIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStruct(ctx *StructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitMulDivMod(ctx *MulDivModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitMembership(ctx *MembershipContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitParentheses(ctx *ParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitBitOr(ctx *BitOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitBitShift(ctx *BitShiftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitUnary(ctx *UnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitTernaryOp(ctx *TernaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitMapLiteral(ctx *MapLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitBitXor(ctx *BitXorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitAnd(ctx *AndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitBitAnd(ctx *BitAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitFieldAccess(ctx *FieldAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructLiteral(ctx *StructLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitMapEntry(ctx *MapEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitMembershipOp(ctx *MembershipOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStandardArray(ctx *StandardArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitListComprehension(ctx *ListComprehensionContext) interface{} {
	return v.VisitChildren(ctx)
}
