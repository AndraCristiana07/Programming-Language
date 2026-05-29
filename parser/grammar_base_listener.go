// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

// BaseGrammarListener is a complete listener for a parse tree produced by GrammarParser.
type BaseGrammarListener struct{}

var _ GrammarListener = &BaseGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseGrammarListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseGrammarListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseGrammarListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseGrammarListener) ExitStatement(ctx *StatementContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseGrammarListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseGrammarListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterTerminator is called when production terminator is entered.
func (s *BaseGrammarListener) EnterTerminator(ctx *TerminatorContext) {}

// ExitTerminator is called when production terminator is exited.
func (s *BaseGrammarListener) ExitTerminator(ctx *TerminatorContext) {}

// EnterFuncStmt is called when production funcStmt is entered.
func (s *BaseGrammarListener) EnterFuncStmt(ctx *FuncStmtContext) {}

// ExitFuncStmt is called when production funcStmt is exited.
func (s *BaseGrammarListener) ExitFuncStmt(ctx *FuncStmtContext) {}

// EnterExprStmt is called when production exprStmt is entered.
func (s *BaseGrammarListener) EnterExprStmt(ctx *ExprStmtContext) {}

// ExitExprStmt is called when production exprStmt is exited.
func (s *BaseGrammarListener) ExitExprStmt(ctx *ExprStmtContext) {}

// EnterTryCatchStmt is called when production tryCatchStmt is entered.
func (s *BaseGrammarListener) EnterTryCatchStmt(ctx *TryCatchStmtContext) {}

// ExitTryCatchStmt is called when production tryCatchStmt is exited.
func (s *BaseGrammarListener) ExitTryCatchStmt(ctx *TryCatchStmtContext) {}

// EnterThrowStmt is called when production throwStmt is entered.
func (s *BaseGrammarListener) EnterThrowStmt(ctx *ThrowStmtContext) {}

// ExitThrowStmt is called when production throwStmt is exited.
func (s *BaseGrammarListener) ExitThrowStmt(ctx *ThrowStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *BaseGrammarListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *BaseGrammarListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *BaseGrammarListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *BaseGrammarListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterSwitchStmt is called when production switchStmt is entered.
func (s *BaseGrammarListener) EnterSwitchStmt(ctx *SwitchStmtContext) {}

// ExitSwitchStmt is called when production switchStmt is exited.
func (s *BaseGrammarListener) ExitSwitchStmt(ctx *SwitchStmtContext) {}

// EnterCaseBlock is called when production caseBlock is entered.
func (s *BaseGrammarListener) EnterCaseBlock(ctx *CaseBlockContext) {}

// ExitCaseBlock is called when production caseBlock is exited.
func (s *BaseGrammarListener) ExitCaseBlock(ctx *CaseBlockContext) {}

// EnterDefaultBlock is called when production defaultBlock is entered.
func (s *BaseGrammarListener) EnterDefaultBlock(ctx *DefaultBlockContext) {}

// ExitDefaultBlock is called when production defaultBlock is exited.
func (s *BaseGrammarListener) ExitDefaultBlock(ctx *DefaultBlockContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseGrammarListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseGrammarListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterAssignStmt is called when production assignStmt is entered.
func (s *BaseGrammarListener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production assignStmt is exited.
func (s *BaseGrammarListener) ExitAssignStmt(ctx *AssignStmtContext) {}

// EnterArrayAssignStmt is called when production arrayAssignStmt is entered.
func (s *BaseGrammarListener) EnterArrayAssignStmt(ctx *ArrayAssignStmtContext) {}

// ExitArrayAssignStmt is called when production arrayAssignStmt is exited.
func (s *BaseGrammarListener) ExitArrayAssignStmt(ctx *ArrayAssignStmtContext) {}

// EnterCompoundAssignStmt is called when production compoundAssignStmt is entered.
func (s *BaseGrammarListener) EnterCompoundAssignStmt(ctx *CompoundAssignStmtContext) {}

// ExitCompoundAssignStmt is called when production compoundAssignStmt is exited.
func (s *BaseGrammarListener) ExitCompoundAssignStmt(ctx *CompoundAssignStmtContext) {}

// EnterPrintStmt is called when production printStmt is entered.
func (s *BaseGrammarListener) EnterPrintStmt(ctx *PrintStmtContext) {}

// ExitPrintStmt is called when production printStmt is exited.
func (s *BaseGrammarListener) ExitPrintStmt(ctx *PrintStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseGrammarListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseGrammarListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterWhileStmt is called when production whileStmt is entered.
func (s *BaseGrammarListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production whileStmt is exited.
func (s *BaseGrammarListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BaseGrammarListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BaseGrammarListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterBlockStmt is called when production blockStmt is entered.
func (s *BaseGrammarListener) EnterBlockStmt(ctx *BlockStmtContext) {}

// ExitBlockStmt is called when production blockStmt is exited.
func (s *BaseGrammarListener) ExitBlockStmt(ctx *BlockStmtContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BaseGrammarListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BaseGrammarListener) ExitForInit(ctx *ForInitContext) {}

// EnterForPost is called when production forPost is entered.
func (s *BaseGrammarListener) EnterForPost(ctx *ForPostContext) {}

// ExitForPost is called when production forPost is exited.
func (s *BaseGrammarListener) ExitForPost(ctx *ForPostContext) {}

// EnterPostfixStmt is called when production postfixStmt is entered.
func (s *BaseGrammarListener) EnterPostfixStmt(ctx *PostfixStmtContext) {}

// ExitPostfixStmt is called when production postfixStmt is exited.
func (s *BaseGrammarListener) ExitPostfixStmt(ctx *PostfixStmtContext) {}

// EnterNull is called when production Null is entered.
func (s *BaseGrammarListener) EnterNull(ctx *NullContext) {}

// ExitNull is called when production Null is exited.
func (s *BaseGrammarListener) ExitNull(ctx *NullContext) {}

// EnterBitOr is called when production BitOr is entered.
func (s *BaseGrammarListener) EnterBitOr(ctx *BitOrContext) {}

// ExitBitOr is called when production BitOr is exited.
func (s *BaseGrammarListener) ExitBitOr(ctx *BitOrContext) {}

// EnterOr is called when production Or is entered.
func (s *BaseGrammarListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production Or is exited.
func (s *BaseGrammarListener) ExitOr(ctx *OrContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseGrammarListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseGrammarListener) ExitAddSub(ctx *AddSubContext) {}

// EnterBitShift is called when production BitShift is entered.
func (s *BaseGrammarListener) EnterBitShift(ctx *BitShiftContext) {}

// ExitBitShift is called when production BitShift is exited.
func (s *BaseGrammarListener) ExitBitShift(ctx *BitShiftContext) {}

// EnterExponential is called when production Exponential is entered.
func (s *BaseGrammarListener) EnterExponential(ctx *ExponentialContext) {}

// ExitExponential is called when production Exponential is exited.
func (s *BaseGrammarListener) ExitExponential(ctx *ExponentialContext) {}

// EnterString is called when production String is entered.
func (s *BaseGrammarListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BaseGrammarListener) ExitString(ctx *StringContext) {}

// EnterArrayLiteral is called when production ArrayLiteral is entered.
func (s *BaseGrammarListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production ArrayLiteral is exited.
func (s *BaseGrammarListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterArrayIndex is called when production ArrayIndex is entered.
func (s *BaseGrammarListener) EnterArrayIndex(ctx *ArrayIndexContext) {}

// ExitArrayIndex is called when production ArrayIndex is exited.
func (s *BaseGrammarListener) ExitArrayIndex(ctx *ArrayIndexContext) {}

// EnterUnary is called when production Unary is entered.
func (s *BaseGrammarListener) EnterUnary(ctx *UnaryContext) {}

// ExitUnary is called when production Unary is exited.
func (s *BaseGrammarListener) ExitUnary(ctx *UnaryContext) {}

// EnterMapLiteral is called when production MapLiteral is entered.
func (s *BaseGrammarListener) EnterMapLiteral(ctx *MapLiteralContext) {}

// ExitMapLiteral is called when production MapLiteral is exited.
func (s *BaseGrammarListener) ExitMapLiteral(ctx *MapLiteralContext) {}

// EnterMulDivMod is called when production MulDivMod is entered.
func (s *BaseGrammarListener) EnterMulDivMod(ctx *MulDivModContext) {}

// ExitMulDivMod is called when production MulDivMod is exited.
func (s *BaseGrammarListener) ExitMulDivMod(ctx *MulDivModContext) {}

// EnterIdentifier is called when production Identifier is entered.
func (s *BaseGrammarListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production Identifier is exited.
func (s *BaseGrammarListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterBitXor is called when production BitXor is entered.
func (s *BaseGrammarListener) EnterBitXor(ctx *BitXorContext) {}

// ExitBitXor is called when production BitXor is exited.
func (s *BaseGrammarListener) ExitBitXor(ctx *BitXorContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseGrammarListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseGrammarListener) ExitNumber(ctx *NumberContext) {}

// EnterComparison is called when production Comparison is entered.
func (s *BaseGrammarListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production Comparison is exited.
func (s *BaseGrammarListener) ExitComparison(ctx *ComparisonContext) {}

// EnterAnd is called when production And is entered.
func (s *BaseGrammarListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production And is exited.
func (s *BaseGrammarListener) ExitAnd(ctx *AndContext) {}

// EnterBitAnd is called when production BitAnd is entered.
func (s *BaseGrammarListener) EnterBitAnd(ctx *BitAndContext) {}

// ExitBitAnd is called when production BitAnd is exited.
func (s *BaseGrammarListener) ExitBitAnd(ctx *BitAndContext) {}

// EnterFunctionCall is called when production FunctionCall is entered.
func (s *BaseGrammarListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production FunctionCall is exited.
func (s *BaseGrammarListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterBoolean is called when production Boolean is entered.
func (s *BaseGrammarListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production Boolean is exited.
func (s *BaseGrammarListener) ExitBoolean(ctx *BooleanContext) {}

// EnterFieldAccess is called when production FieldAccess is entered.
func (s *BaseGrammarListener) EnterFieldAccess(ctx *FieldAccessContext) {}

// ExitFieldAccess is called when production FieldAccess is exited.
func (s *BaseGrammarListener) ExitFieldAccess(ctx *FieldAccessContext) {}

// EnterParentheses is called when production Parentheses is entered.
func (s *BaseGrammarListener) EnterParentheses(ctx *ParenthesesContext) {}

// ExitParentheses is called when production Parentheses is exited.
func (s *BaseGrammarListener) ExitParentheses(ctx *ParenthesesContext) {}

// EnterMapEntry is called when production mapEntry is entered.
func (s *BaseGrammarListener) EnterMapEntry(ctx *MapEntryContext) {}

// ExitMapEntry is called when production mapEntry is exited.
func (s *BaseGrammarListener) ExitMapEntry(ctx *MapEntryContext) {}
