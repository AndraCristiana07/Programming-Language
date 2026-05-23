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

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseGrammarListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseGrammarListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterAssignStmt is called when production assignStmt is entered.
func (s *BaseGrammarListener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production assignStmt is exited.
func (s *BaseGrammarListener) ExitAssignStmt(ctx *AssignStmtContext) {}

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

// EnterMulDivMod is called when production MulDivMod is entered.
func (s *BaseGrammarListener) EnterMulDivMod(ctx *MulDivModContext) {}

// ExitMulDivMod is called when production MulDivMod is exited.
func (s *BaseGrammarListener) ExitMulDivMod(ctx *MulDivModContext) {}

// EnterIdentifier is called when production Identifier is entered.
func (s *BaseGrammarListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production Identifier is exited.
func (s *BaseGrammarListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseGrammarListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseGrammarListener) ExitNumber(ctx *NumberContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseGrammarListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseGrammarListener) ExitAddSub(ctx *AddSubContext) {}

// EnterComparison is called when production Comparison is entered.
func (s *BaseGrammarListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production Comparison is exited.
func (s *BaseGrammarListener) ExitComparison(ctx *ComparisonContext) {}

// EnterExponential is called when production Exponential is entered.
func (s *BaseGrammarListener) EnterExponential(ctx *ExponentialContext) {}

// ExitExponential is called when production Exponential is exited.
func (s *BaseGrammarListener) ExitExponential(ctx *ExponentialContext) {}

// EnterString is called when production String is entered.
func (s *BaseGrammarListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BaseGrammarListener) ExitString(ctx *StringContext) {}

// EnterBoolean is called when production Boolean is entered.
func (s *BaseGrammarListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production Boolean is exited.
func (s *BaseGrammarListener) ExitBoolean(ctx *BooleanContext) {}
