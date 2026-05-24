// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GrammarParser struct {
	*antlr.BaseParser
}

var GrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func grammarParserInit() {
	staticData := &GrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'var'", "'='", "'+'", "'-'", "'*'", "'/'", "'<'", "'>'", "'=='",
		"'!='", "'<='", "'>='", "'{'", "'('", "')'", "'}'", "'if'", "'else'",
		"'while'", "'for'", "';'", "'++'", "'--'", "'%'", "'**'", "'+='", "'-='",
		"'*='", "'/='", "'%='", "'**='", "'print'", "'true'", "'false'", "'and'",
		"'or'", "'not'",
	}
	staticData.SymbolicNames = []string{
		"", "VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "LBRACE", "LPAREN",
		"RPAREN", "RBRACE", "IF", "ELSE", "WHILE", "FOR", "SEMICOLON", "INC",
		"DEC", "MODULO", "EXPONENTIAL", "PLUSEQUAL", "MINUSEQUAL", "STAREQUAL",
		"SLASHEQUAL", "MODEQUAL", "EXPONENTIALEQUAL", "PRINT", "TRUE", "FALSE",
		"AND", "OR", "NOT", "IDENTIFIER", "NUMBER", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "varDecl", "assignStmt", "compoundAssignStmt",
		"printStmt", "ifStmt", "whileStmt", "forStmt", "blockStmt", "forInit",
		"forPost", "postfixStmt", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 41, 145, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 5, 0, 30, 8, 0, 10,
		0, 12, 0, 33, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 46, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 71, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 5, 9, 91, 8, 9, 10, 9, 12, 9, 94, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 3,
		10, 100, 8, 10, 1, 11, 1, 11, 3, 11, 104, 8, 11, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		3, 13, 120, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		5, 13, 140, 8, 13, 10, 13, 12, 13, 143, 9, 13, 1, 13, 0, 1, 26, 14, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 0, 6, 1, 0, 26, 31, 1,
		0, 22, 23, 1, 0, 33, 34, 2, 0, 5, 6, 24, 24, 1, 0, 3, 4, 1, 0, 7, 12, 154,
		0, 31, 1, 0, 0, 0, 2, 45, 1, 0, 0, 0, 4, 47, 1, 0, 0, 0, 6, 52, 1, 0, 0,
		0, 8, 56, 1, 0, 0, 0, 10, 60, 1, 0, 0, 0, 12, 63, 1, 0, 0, 0, 14, 72, 1,
		0, 0, 0, 16, 78, 1, 0, 0, 0, 18, 88, 1, 0, 0, 0, 20, 99, 1, 0, 0, 0, 22,
		103, 1, 0, 0, 0, 24, 105, 1, 0, 0, 0, 26, 119, 1, 0, 0, 0, 28, 30, 3, 2,
		1, 0, 29, 28, 1, 0, 0, 0, 30, 33, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 32,
		1, 0, 0, 0, 32, 34, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 35, 5, 0, 0, 1,
		35, 1, 1, 0, 0, 0, 36, 46, 3, 4, 2, 0, 37, 46, 3, 6, 3, 0, 38, 46, 3, 8,
		4, 0, 39, 46, 3, 10, 5, 0, 40, 46, 3, 12, 6, 0, 41, 46, 3, 14, 7, 0, 42,
		46, 3, 16, 8, 0, 43, 46, 3, 18, 9, 0, 44, 46, 3, 24, 12, 0, 45, 36, 1,
		0, 0, 0, 45, 37, 1, 0, 0, 0, 45, 38, 1, 0, 0, 0, 45, 39, 1, 0, 0, 0, 45,
		40, 1, 0, 0, 0, 45, 41, 1, 0, 0, 0, 45, 42, 1, 0, 0, 0, 45, 43, 1, 0, 0,
		0, 45, 44, 1, 0, 0, 0, 46, 3, 1, 0, 0, 0, 47, 48, 5, 1, 0, 0, 48, 49, 5,
		38, 0, 0, 49, 50, 5, 2, 0, 0, 50, 51, 3, 26, 13, 0, 51, 5, 1, 0, 0, 0,
		52, 53, 5, 38, 0, 0, 53, 54, 5, 2, 0, 0, 54, 55, 3, 26, 13, 0, 55, 7, 1,
		0, 0, 0, 56, 57, 5, 38, 0, 0, 57, 58, 7, 0, 0, 0, 58, 59, 3, 26, 13, 0,
		59, 9, 1, 0, 0, 0, 60, 61, 5, 32, 0, 0, 61, 62, 3, 26, 13, 0, 62, 11, 1,
		0, 0, 0, 63, 64, 5, 17, 0, 0, 64, 65, 5, 14, 0, 0, 65, 66, 3, 26, 13, 0,
		66, 67, 5, 15, 0, 0, 67, 70, 3, 18, 9, 0, 68, 69, 5, 18, 0, 0, 69, 71,
		3, 18, 9, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 13, 1, 0, 0, 0,
		72, 73, 5, 19, 0, 0, 73, 74, 5, 14, 0, 0, 74, 75, 3, 26, 13, 0, 75, 76,
		5, 15, 0, 0, 76, 77, 3, 18, 9, 0, 77, 15, 1, 0, 0, 0, 78, 79, 5, 20, 0,
		0, 79, 80, 5, 14, 0, 0, 80, 81, 3, 20, 10, 0, 81, 82, 5, 21, 0, 0, 82,
		83, 3, 26, 13, 0, 83, 84, 5, 21, 0, 0, 84, 85, 3, 22, 11, 0, 85, 86, 5,
		15, 0, 0, 86, 87, 3, 18, 9, 0, 87, 17, 1, 0, 0, 0, 88, 92, 5, 13, 0, 0,
		89, 91, 3, 2, 1, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1,
		0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95,
		96, 5, 16, 0, 0, 96, 19, 1, 0, 0, 0, 97, 100, 3, 4, 2, 0, 98, 100, 3, 6,
		3, 0, 99, 97, 1, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 21, 1, 0, 0, 0, 101,
		104, 3, 6, 3, 0, 102, 104, 3, 24, 12, 0, 103, 101, 1, 0, 0, 0, 103, 102,
		1, 0, 0, 0, 104, 23, 1, 0, 0, 0, 105, 106, 5, 38, 0, 0, 106, 107, 7, 1,
		0, 0, 107, 25, 1, 0, 0, 0, 108, 109, 6, 13, -1, 0, 109, 110, 5, 37, 0,
		0, 110, 120, 3, 26, 13, 12, 111, 120, 5, 38, 0, 0, 112, 120, 5, 39, 0,
		0, 113, 120, 7, 2, 0, 0, 114, 120, 5, 40, 0, 0, 115, 116, 5, 14, 0, 0,
		116, 117, 3, 26, 13, 0, 117, 118, 5, 15, 0, 0, 118, 120, 1, 0, 0, 0, 119,
		108, 1, 0, 0, 0, 119, 111, 1, 0, 0, 0, 119, 112, 1, 0, 0, 0, 119, 113,
		1, 0, 0, 0, 119, 114, 1, 0, 0, 0, 119, 115, 1, 0, 0, 0, 120, 141, 1, 0,
		0, 0, 121, 122, 10, 11, 0, 0, 122, 123, 5, 25, 0, 0, 123, 140, 3, 26, 13,
		12, 124, 125, 10, 10, 0, 0, 125, 126, 7, 3, 0, 0, 126, 140, 3, 26, 13,
		11, 127, 128, 10, 9, 0, 0, 128, 129, 7, 4, 0, 0, 129, 140, 3, 26, 13, 10,
		130, 131, 10, 8, 0, 0, 131, 132, 7, 5, 0, 0, 132, 140, 3, 26, 13, 9, 133,
		134, 10, 7, 0, 0, 134, 135, 5, 35, 0, 0, 135, 140, 3, 26, 13, 8, 136, 137,
		10, 6, 0, 0, 137, 138, 5, 36, 0, 0, 138, 140, 3, 26, 13, 7, 139, 121, 1,
		0, 0, 0, 139, 124, 1, 0, 0, 0, 139, 127, 1, 0, 0, 0, 139, 130, 1, 0, 0,
		0, 139, 133, 1, 0, 0, 0, 139, 136, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141,
		139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 27, 1, 0, 0, 0, 143, 141, 1,
		0, 0, 0, 9, 31, 45, 70, 92, 99, 103, 119, 139, 141,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GrammarParserInit initializes any static state used to implement GrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GrammarParserInit() {
	staticData := &GrammarParserStaticData
	staticData.once.Do(grammarParserInit)
}

// NewGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGrammarParser(input antlr.TokenStream) *GrammarParser {
	GrammarParserInit()
	this := new(GrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Grammar.g4"

	return this
}

// GrammarParser tokens.
const (
	GrammarParserEOF              = antlr.TokenEOF
	GrammarParserVAR              = 1
	GrammarParserEQUALS           = 2
	GrammarParserPLUS             = 3
	GrammarParserMINUS            = 4
	GrammarParserSTAR             = 5
	GrammarParserSLASH            = 6
	GrammarParserLESS             = 7
	GrammarParserGREATER          = 8
	GrammarParserEQUALEQUAL       = 9
	GrammarParserBANGEQUAL        = 10
	GrammarParserLESSEQUAL        = 11
	GrammarParserGREATEREQUAL     = 12
	GrammarParserLBRACE           = 13
	GrammarParserLPAREN           = 14
	GrammarParserRPAREN           = 15
	GrammarParserRBRACE           = 16
	GrammarParserIF               = 17
	GrammarParserELSE             = 18
	GrammarParserWHILE            = 19
	GrammarParserFOR              = 20
	GrammarParserSEMICOLON        = 21
	GrammarParserINC              = 22
	GrammarParserDEC              = 23
	GrammarParserMODULO           = 24
	GrammarParserEXPONENTIAL      = 25
	GrammarParserPLUSEQUAL        = 26
	GrammarParserMINUSEQUAL       = 27
	GrammarParserSTAREQUAL        = 28
	GrammarParserSLASHEQUAL       = 29
	GrammarParserMODEQUAL         = 30
	GrammarParserEXPONENTIALEQUAL = 31
	GrammarParserPRINT            = 32
	GrammarParserTRUE             = 33
	GrammarParserFALSE            = 34
	GrammarParserAND              = 35
	GrammarParserOR               = 36
	GrammarParserNOT              = 37
	GrammarParserIDENTIFIER       = 38
	GrammarParserNUMBER           = 39
	GrammarParserSTRING           = 40
	GrammarParserWS               = 41
)

// GrammarParser rules.
const (
	GrammarParserRULE_program            = 0
	GrammarParserRULE_statement          = 1
	GrammarParserRULE_varDecl            = 2
	GrammarParserRULE_assignStmt         = 3
	GrammarParserRULE_compoundAssignStmt = 4
	GrammarParserRULE_printStmt          = 5
	GrammarParserRULE_ifStmt             = 6
	GrammarParserRULE_whileStmt          = 7
	GrammarParserRULE_forStmt            = 8
	GrammarParserRULE_blockStmt          = 9
	GrammarParserRULE_forInit            = 10
	GrammarParserRULE_forPost            = 11
	GrammarParserRULE_postfixStmt        = 12
	GrammarParserRULE_expr               = 13
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GrammarParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrammarParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&279174586370) != 0 {
		{
			p.SetState(28)
			p.Statement()
		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(34)
		p.Match(GrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDecl() IVarDeclContext
	AssignStmt() IAssignStmtContext
	CompoundAssignStmt() ICompoundAssignStmtContext
	PrintStmt() IPrintStmtContext
	IfStmt() IIfStmtContext
	WhileStmt() IWhileStmtContext
	ForStmt() IForStmtContext
	BlockStmt() IBlockStmtContext
	PostfixStmt() IPostfixStmtContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VarDecl() IVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *StatementContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *StatementContext) CompoundAssignStmt() ICompoundAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompoundAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompoundAssignStmtContext)
}

func (s *StatementContext) PrintStmt() IPrintStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStmtContext)
}

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) WhileStmt() IWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *StatementContext) ForStmt() IForStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *StatementContext) BlockStmt() IBlockStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStmtContext)
}

func (s *StatementContext) PostfixStmt() IPostfixStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixStmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrammarParserRULE_statement)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.VarDecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.AssignStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(38)
			p.CompoundAssignStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(39)
			p.PrintStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(40)
			p.IfStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(41)
			p.WhileStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(42)
			p.ForStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(43)
			p.BlockStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(44)
			p.PostfixStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	Expr() IExprContext

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_varDecl
	return p
}

func InitEmptyVarDeclContext(p *VarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_varDecl
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *VarDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *VarDeclContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUALS, 0)
}

func (s *VarDeclContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (s *VarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrammarParserRULE_varDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(GrammarParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Match(GrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(GrammarParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignStmtContext is an interface to support dynamic dispatch.
type IAssignStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	Expr() IExprContext

	// IsAssignStmtContext differentiates from other interfaces.
	IsAssignStmtContext()
}

type AssignStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignStmtContext() *AssignStmtContext {
	var p = new(AssignStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_assignStmt
	return p
}

func InitEmptyAssignStmtContext(p *AssignStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_assignStmt
}

func (*AssignStmtContext) IsAssignStmtContext() {}

func NewAssignStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignStmtContext {
	var p = new(AssignStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_assignStmt

	return p
}

func (s *AssignStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *AssignStmtContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUALS, 0)
}

func (s *AssignStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterAssignStmt(s)
	}
}

func (s *AssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitAssignStmt(s)
	}
}

func (s *AssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) AssignStmt() (localctx IAssignStmtContext) {
	localctx = NewAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrammarParserRULE_assignStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(GrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(GrammarParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompoundAssignStmtContext is an interface to support dynamic dispatch.
type ICompoundAssignStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Expr() IExprContext
	PLUSEQUAL() antlr.TerminalNode
	MINUSEQUAL() antlr.TerminalNode
	STAREQUAL() antlr.TerminalNode
	SLASHEQUAL() antlr.TerminalNode
	MODEQUAL() antlr.TerminalNode
	EXPONENTIALEQUAL() antlr.TerminalNode

	// IsCompoundAssignStmtContext differentiates from other interfaces.
	IsCompoundAssignStmtContext()
}

type CompoundAssignStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyCompoundAssignStmtContext() *CompoundAssignStmtContext {
	var p = new(CompoundAssignStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_compoundAssignStmt
	return p
}

func InitEmptyCompoundAssignStmtContext(p *CompoundAssignStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_compoundAssignStmt
}

func (*CompoundAssignStmtContext) IsCompoundAssignStmtContext() {}

func NewCompoundAssignStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompoundAssignStmtContext {
	var p = new(CompoundAssignStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_compoundAssignStmt

	return p
}

func (s *CompoundAssignStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CompoundAssignStmtContext) GetOp() antlr.Token { return s.op }

func (s *CompoundAssignStmtContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompoundAssignStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *CompoundAssignStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CompoundAssignStmtContext) PLUSEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserPLUSEQUAL, 0)
}

func (s *CompoundAssignStmtContext) MINUSEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserMINUSEQUAL, 0)
}

func (s *CompoundAssignStmtContext) STAREQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserSTAREQUAL, 0)
}

func (s *CompoundAssignStmtContext) SLASHEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserSLASHEQUAL, 0)
}

func (s *CompoundAssignStmtContext) MODEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserMODEQUAL, 0)
}

func (s *CompoundAssignStmtContext) EXPONENTIALEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEXPONENTIALEQUAL, 0)
}

func (s *CompoundAssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundAssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompoundAssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterCompoundAssignStmt(s)
	}
}

func (s *CompoundAssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitCompoundAssignStmt(s)
	}
}

func (s *CompoundAssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitCompoundAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) CompoundAssignStmt() (localctx ICompoundAssignStmtContext) {
	localctx = NewCompoundAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GrammarParserRULE_compoundAssignStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(GrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*CompoundAssignStmtContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4227858432) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*CompoundAssignStmtContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(58)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintStmtContext is an interface to support dynamic dispatch.
type IPrintStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	Expr() IExprContext

	// IsPrintStmtContext differentiates from other interfaces.
	IsPrintStmtContext()
}

type PrintStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStmtContext() *PrintStmtContext {
	var p = new(PrintStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_printStmt
	return p
}

func InitEmptyPrintStmtContext(p *PrintStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_printStmt
}

func (*PrintStmtContext) IsPrintStmtContext() {}

func NewPrintStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStmtContext {
	var p = new(PrintStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_printStmt

	return p
}

func (s *PrintStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(GrammarParserPRINT, 0)
}

func (s *PrintStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterPrintStmt(s)
	}
}

func (s *PrintStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitPrintStmt(s)
	}
}

func (s *PrintStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitPrintStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) PrintStmt() (localctx IPrintStmtContext) {
	localctx = NewPrintStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GrammarParserRULE_printStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(GrammarParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetThenBranch returns the thenBranch rule contexts.
	GetThenBranch() IBlockStmtContext

	// GetElseBranch returns the elseBranch rule contexts.
	GetElseBranch() IBlockStmtContext

	// SetThenBranch sets the thenBranch rule contexts.
	SetThenBranch(IBlockStmtContext)

	// SetElseBranch sets the elseBranch rule contexts.
	SetElseBranch(IBlockStmtContext)

	// Getter signatures
	IF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	AllBlockStmt() []IBlockStmtContext
	BlockStmt(i int) IBlockStmtContext
	ELSE() antlr.TerminalNode

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	thenBranch IBlockStmtContext
	elseBranch IBlockStmtContext
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) GetThenBranch() IBlockStmtContext { return s.thenBranch }

func (s *IfStmtContext) GetElseBranch() IBlockStmtContext { return s.elseBranch }

func (s *IfStmtContext) SetThenBranch(v IBlockStmtContext) { s.thenBranch = v }

func (s *IfStmtContext) SetElseBranch(v IBlockStmtContext) { s.elseBranch = v }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(GrammarParserIF, 0)
}

func (s *IfStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *IfStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *IfStmtContext) AllBlockStmt() []IBlockStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockStmtContext); ok {
			len++
		}
	}

	tst := make([]IBlockStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockStmtContext); ok {
			tst[i] = t.(IBlockStmtContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) BlockStmt(i int) IBlockStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStmtContext)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserELSE, 0)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GrammarParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(GrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.expr(0)
	}
	{
		p.SetState(66)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)

		var _x = p.BlockStmt()

		localctx.(*IfStmtContext).thenBranch = _x
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserELSE {
		{
			p.SetState(68)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)

			var _x = p.BlockStmt()

			localctx.(*IfStmtContext).elseBranch = _x
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBody returns the body rule contexts.
	GetBody() IBlockStmtContext

	// SetBody sets the body rule contexts.
	SetBody(IBlockStmtContext)

	// Getter signatures
	WHILE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	BlockStmt() IBlockStmtContext

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	body   IBlockStmtContext
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_whileStmt
	return p
}

func InitEmptyWhileStmtContext(p *WhileStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_whileStmt
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) GetBody() IBlockStmtContext { return s.body }

func (s *WhileStmtContext) SetBody(v IBlockStmtContext) { s.body = v }

func (s *WhileStmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(GrammarParserWHILE, 0)
}

func (s *WhileStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *WhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *WhileStmtContext) BlockStmt() IBlockStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStmtContext)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) WhileStmt() (localctx IWhileStmtContext) {
	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GrammarParserRULE_whileStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(GrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.expr(0)
	}
	{
		p.SetState(75)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)

		var _x = p.BlockStmt()

		localctx.(*WhileStmtContext).body = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInit returns the init rule contexts.
	GetInit() IForInitContext

	// GetCond returns the cond rule contexts.
	GetCond() IExprContext

	// GetPost returns the post rule contexts.
	GetPost() IForPostContext

	// GetBody returns the body rule contexts.
	GetBody() IBlockStmtContext

	// SetInit sets the init rule contexts.
	SetInit(IForInitContext)

	// SetCond sets the cond rule contexts.
	SetCond(IExprContext)

	// SetPost sets the post rule contexts.
	SetPost(IForPostContext)

	// SetBody sets the body rule contexts.
	SetBody(IBlockStmtContext)

	// Getter signatures
	FOR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ForInit() IForInitContext
	Expr() IExprContext
	ForPost() IForPostContext
	BlockStmt() IBlockStmtContext

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	init   IForInitContext
	cond   IExprContext
	post   IForPostContext
	body   IBlockStmtContext
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forStmt
	return p
}

func InitEmptyForStmtContext(p *ForStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forStmt
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) GetInit() IForInitContext { return s.init }

func (s *ForStmtContext) GetCond() IExprContext { return s.cond }

func (s *ForStmtContext) GetPost() IForPostContext { return s.post }

func (s *ForStmtContext) GetBody() IBlockStmtContext { return s.body }

func (s *ForStmtContext) SetInit(v IForInitContext) { s.init = v }

func (s *ForStmtContext) SetCond(v IExprContext) { s.cond = v }

func (s *ForStmtContext) SetPost(v IForPostContext) { s.post = v }

func (s *ForStmtContext) SetBody(v IBlockStmtContext) { s.body = v }

func (s *ForStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserFOR, 0)
}

func (s *ForStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *ForStmtContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserSEMICOLON)
}

func (s *ForStmtContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserSEMICOLON, i)
}

func (s *ForStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *ForStmtContext) ForInit() IForInitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForInitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForInitContext)
}

func (s *ForStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForStmtContext) ForPost() IForPostContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForPostContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForPostContext)
}

func (s *ForStmtContext) BlockStmt() IBlockStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStmtContext)
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ForStmt() (localctx IForStmtContext) {
	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GrammarParserRULE_forStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(GrammarParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)

		var _x = p.ForInit()

		localctx.(*ForStmtContext).init = _x
	}
	{
		p.SetState(81)
		p.Match(GrammarParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)

		var _x = p.expr(0)

		localctx.(*ForStmtContext).cond = _x
	}
	{
		p.SetState(83)
		p.Match(GrammarParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)

		var _x = p.ForPost()

		localctx.(*ForStmtContext).post = _x
	}
	{
		p.SetState(85)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)

		var _x = p.BlockStmt()

		localctx.(*ForStmtContext).body = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockStmtContext is an interface to support dynamic dispatch.
type IBlockStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockStmtContext differentiates from other interfaces.
	IsBlockStmtContext()
}

type BlockStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStmtContext() *BlockStmtContext {
	var p = new(BlockStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_blockStmt
	return p
}

func InitEmptyBlockStmtContext(p *BlockStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_blockStmt
}

func (*BlockStmtContext) IsBlockStmtContext() {}

func NewBlockStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStmtContext {
	var p = new(BlockStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_blockStmt

	return p
}

func (s *BlockStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACE, 0)
}

func (s *BlockStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACE, 0)
}

func (s *BlockStmtContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockStmtContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBlockStmt(s)
	}
}

func (s *BlockStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBlockStmt(s)
	}
}

func (s *BlockStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBlockStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) BlockStmt() (localctx IBlockStmtContext) {
	localctx = NewBlockStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GrammarParserRULE_blockStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(GrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&279174586370) != 0 {
		{
			p.SetState(89)
			p.Statement()
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
		p.Match(GrammarParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForInitContext is an interface to support dynamic dispatch.
type IForInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDecl() IVarDeclContext
	AssignStmt() IAssignStmtContext

	// IsForInitContext differentiates from other interfaces.
	IsForInitContext()
}

type ForInitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForInitContext() *ForInitContext {
	var p = new(ForInitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forInit
	return p
}

func InitEmptyForInitContext(p *ForInitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forInit
}

func (*ForInitContext) IsForInitContext() {}

func NewForInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInitContext {
	var p = new(ForInitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_forInit

	return p
}

func (s *ForInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInitContext) VarDecl() IVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *ForInitContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *ForInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterForInit(s)
	}
}

func (s *ForInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitForInit(s)
	}
}

func (s *ForInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitForInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ForInit() (localctx IForInitContext) {
	localctx = NewForInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GrammarParserRULE_forInit)
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.VarDecl()
		}

	case GrammarParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.AssignStmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForPostContext is an interface to support dynamic dispatch.
type IForPostContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AssignStmt() IAssignStmtContext
	PostfixStmt() IPostfixStmtContext

	// IsForPostContext differentiates from other interfaces.
	IsForPostContext()
}

type ForPostContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForPostContext() *ForPostContext {
	var p = new(ForPostContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forPost
	return p
}

func InitEmptyForPostContext(p *ForPostContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forPost
}

func (*ForPostContext) IsForPostContext() {}

func NewForPostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForPostContext {
	var p = new(ForPostContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_forPost

	return p
}

func (s *ForPostContext) GetParser() antlr.Parser { return s.parser }

func (s *ForPostContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *ForPostContext) PostfixStmt() IPostfixStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixStmtContext)
}

func (s *ForPostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForPostContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForPostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterForPost(s)
	}
}

func (s *ForPostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitForPost(s)
	}
}

func (s *ForPostContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitForPost(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ForPost() (localctx IForPostContext) {
	localctx = NewForPostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GrammarParserRULE_forPost)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.AssignStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.PostfixStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPostfixStmtContext is an interface to support dynamic dispatch.
type IPostfixStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	INC() antlr.TerminalNode
	DEC() antlr.TerminalNode

	// IsPostfixStmtContext differentiates from other interfaces.
	IsPostfixStmtContext()
}

type PostfixStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyPostfixStmtContext() *PostfixStmtContext {
	var p = new(PostfixStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_postfixStmt
	return p
}

func InitEmptyPostfixStmtContext(p *PostfixStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_postfixStmt
}

func (*PostfixStmtContext) IsPostfixStmtContext() {}

func NewPostfixStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixStmtContext {
	var p = new(PostfixStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_postfixStmt

	return p
}

func (s *PostfixStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixStmtContext) GetOp() antlr.Token { return s.op }

func (s *PostfixStmtContext) SetOp(v antlr.Token) { s.op = v }

func (s *PostfixStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *PostfixStmtContext) INC() antlr.TerminalNode {
	return s.GetToken(GrammarParserINC, 0)
}

func (s *PostfixStmtContext) DEC() antlr.TerminalNode {
	return s.GetToken(GrammarParserDEC, 0)
}

func (s *PostfixStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostfixStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterPostfixStmt(s)
	}
}

func (s *PostfixStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitPostfixStmt(s)
	}
}

func (s *PostfixStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitPostfixStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) PostfixStmt() (localctx IPostfixStmtContext) {
	localctx = NewPostfixStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GrammarParserRULE_postfixStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(GrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*PostfixStmtContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GrammarParserINC || _la == GrammarParserDEC) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*PostfixStmtContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NotContext struct {
	ExprContext
	op antlr.Token
}

func NewNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotContext {
	var p = new(NotContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotContext) GetOp() antlr.Token { return s.op }

func (s *NotContext) SetOp(v antlr.Token) { s.op = v }

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotContext) NOT() antlr.TerminalNode {
	return s.GetToken(GrammarParserNOT, 0)
}

func (s *NotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterNot(s)
	}
}

func (s *NotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitNot(s)
	}
}

func (s *NotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivModContext struct {
	ExprContext
	op antlr.Token
}

func NewMulDivModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivModContext {
	var p = new(MulDivModContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulDivModContext) GetOp() antlr.Token { return s.op }

func (s *MulDivModContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivModContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulDivModContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivModContext) STAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserSTAR, 0)
}

func (s *MulDivModContext) SLASH() antlr.TerminalNode {
	return s.GetToken(GrammarParserSLASH, 0)
}

func (s *MulDivModContext) MODULO() antlr.TerminalNode {
	return s.GetToken(GrammarParserMODULO, 0)
}

func (s *MulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMulDivMod(s)
	}
}

func (s *MulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMulDivMod(s)
	}
}

func (s *MulDivModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitMulDivMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierContext struct {
	ExprContext
}

func NewIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierContext {
	var p = new(IdentifierContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrContext struct {
	ExprContext
	op antlr.Token
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OrContext) GetOp() antlr.Token { return s.op }

func (s *OrContext) SetOp(v antlr.Token) { s.op = v }

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OrContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrContext) OR() antlr.TerminalNode {
	return s.GetToken(GrammarParserOR, 0)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitOr(s)
	}
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberContext struct {
	ExprContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GrammarParserNUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	ExprContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GrammarParserPLUS, 0)
}

func (s *AddSubContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GrammarParserMINUS, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparisonContext struct {
	ExprContext
	op antlr.Token
}

func NewComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonContext {
	var p = new(ComparisonContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ComparisonContext) GetOp() antlr.Token { return s.op }

func (s *ComparisonContext) SetOp(v antlr.Token) { s.op = v }

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ComparisonContext) LESS() antlr.TerminalNode {
	return s.GetToken(GrammarParserLESS, 0)
}

func (s *ComparisonContext) GREATER() antlr.TerminalNode {
	return s.GetToken(GrammarParserGREATER, 0)
}

func (s *ComparisonContext) EQUALEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUALEQUAL, 0)
}

func (s *ComparisonContext) BANGEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserBANGEQUAL, 0)
}

func (s *ComparisonContext) LESSEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserLESSEQUAL, 0)
}

func (s *ComparisonContext) GREATEREQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserGREATEREQUAL, 0)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndContext struct {
	ExprContext
	op antlr.Token
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AndContext) GetOp() antlr.Token { return s.op }

func (s *AndContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AndContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndContext) AND() antlr.TerminalNode {
	return s.GetToken(GrammarParserAND, 0)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExponentialContext struct {
	ExprContext
	op antlr.Token
}

func NewExponentialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExponentialContext {
	var p = new(ExponentialContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExponentialContext) GetOp() antlr.Token { return s.op }

func (s *ExponentialContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExponentialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExponentialContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExponentialContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExponentialContext) EXPONENTIAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEXPONENTIAL, 0)
}

func (s *ExponentialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExponential(s)
	}
}

func (s *ExponentialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExponential(s)
	}
}

func (s *ExponentialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitExponential(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	ExprContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(GrammarParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanContext struct {
	ExprContext
	val antlr.Token
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BooleanContext) GetVal() antlr.Token { return s.val }

func (s *BooleanContext) SetVal(v antlr.Token) { s.val = v }

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) TRUE() antlr.TerminalNode {
	return s.GetToken(GrammarParserTRUE, 0)
}

func (s *BooleanContext) FALSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserFALSE, 0)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenthesesContext struct {
	ExprContext
}

func NewParenthesesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesesContext {
	var p = new(ParenthesesContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenthesesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesesContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *ParenthesesContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenthesesContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *ParenthesesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterParentheses(s)
	}
}

func (s *ParenthesesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitParentheses(s)
	}
}

func (s *ParenthesesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitParentheses(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *GrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, GrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserNOT:
		localctx = NewNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(109)

			var _m = p.Match(GrammarParserNOT)

			localctx.(*NotContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.expr(12)
		}

	case GrammarParserIDENTIFIER:
		localctx = NewIdentifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(111)
			p.Match(GrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(112)
			p.Match(GrammarParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserTRUE, GrammarParserFALSE:
		localctx = NewBooleanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(113)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*BooleanContext).val = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GrammarParserTRUE || _la == GrammarParserFALSE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*BooleanContext).val = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case GrammarParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(114)
			p.Match(GrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserLPAREN:
		localctx = NewParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(115)
			p.Match(GrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.expr(0)
		}
		{
			p.SetState(117)
			p.Match(GrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(139)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExponentialContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(121)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(122)

					var _m = p.Match(GrammarParserEXPONENTIAL)

					localctx.(*ExponentialContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(123)
					p.expr(12)
				}

			case 2:
				localctx = NewMulDivModContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(124)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(125)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16777312) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(126)
					p.expr(11)
				}

			case 3:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(127)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(128)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserPLUS || _la == GrammarParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(129)
					p.expr(10)
				}

			case 4:
				localctx = NewComparisonContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(130)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(131)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparisonContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8064) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparisonContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(132)
					p.expr(9)
				}

			case 5:
				localctx = NewAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(133)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(134)

					var _m = p.Match(GrammarParserAND)

					localctx.(*AndContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(135)
					p.expr(8)
				}

			case 6:
				localctx = NewOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(136)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(137)

					var _m = p.Match(GrammarParserOR)

					localctx.(*OrContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(138)
					p.expr(7)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *GrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 13:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
