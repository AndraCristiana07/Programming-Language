// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type GrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GrammarLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func grammarlexerLexerInit() {
	staticData := &GrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'var'", "'='", "'+'", "'-'", "'*'", "'/'", "'<'", "'>'", "'=='",
		"'!='", "'<='", "'>='", "'{'", "'('", "')'", "'}'", "'if'", "'else'",
		"'while'", "'for'", "';'", "'++'", "'--'", "'print'", "'true'", "'false'",
	}
	staticData.SymbolicNames = []string{
		"", "VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "LBRACE", "LPAREN",
		"RPAREN", "RBRACE", "IF", "ELSE", "WHILE", "FOR", "SEMICOLON", "INC",
		"DEC", "PRINT", "TRUE", "FALSE", "IDENTIFIER", "NUMBER", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "LBRACE", "LPAREN",
		"RPAREN", "RBRACE", "IF", "ELSE", "WHILE", "FOR", "SEMICOLON", "INC",
		"DEC", "PRINT", "TRUE", "FALSE", "IDENTIFIER", "NUMBER", "STRING", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 172, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 26, 1, 26, 5, 26, 145, 8, 26, 10, 26, 12, 26, 148,
		9, 26, 1, 27, 4, 27, 151, 8, 27, 11, 27, 12, 27, 152, 1, 28, 1, 28, 1,
		28, 1, 28, 5, 28, 159, 8, 28, 10, 28, 12, 28, 162, 9, 28, 1, 28, 1, 28,
		1, 29, 4, 29, 167, 8, 29, 11, 29, 12, 29, 168, 1, 29, 1, 29, 0, 0, 30,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 1, 0, 5, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 1, 0, 48, 57, 1, 0, 34, 34, 3, 0, 9, 10, 13, 13, 32, 32, 176,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 1, 61, 1, 0, 0,
		0, 3, 65, 1, 0, 0, 0, 5, 67, 1, 0, 0, 0, 7, 69, 1, 0, 0, 0, 9, 71, 1, 0,
		0, 0, 11, 73, 1, 0, 0, 0, 13, 75, 1, 0, 0, 0, 15, 77, 1, 0, 0, 0, 17, 79,
		1, 0, 0, 0, 19, 82, 1, 0, 0, 0, 21, 85, 1, 0, 0, 0, 23, 88, 1, 0, 0, 0,
		25, 91, 1, 0, 0, 0, 27, 93, 1, 0, 0, 0, 29, 95, 1, 0, 0, 0, 31, 97, 1,
		0, 0, 0, 33, 99, 1, 0, 0, 0, 35, 102, 1, 0, 0, 0, 37, 107, 1, 0, 0, 0,
		39, 113, 1, 0, 0, 0, 41, 117, 1, 0, 0, 0, 43, 119, 1, 0, 0, 0, 45, 122,
		1, 0, 0, 0, 47, 125, 1, 0, 0, 0, 49, 131, 1, 0, 0, 0, 51, 136, 1, 0, 0,
		0, 53, 142, 1, 0, 0, 0, 55, 150, 1, 0, 0, 0, 57, 154, 1, 0, 0, 0, 59, 166,
		1, 0, 0, 0, 61, 62, 5, 118, 0, 0, 62, 63, 5, 97, 0, 0, 63, 64, 5, 114,
		0, 0, 64, 2, 1, 0, 0, 0, 65, 66, 5, 61, 0, 0, 66, 4, 1, 0, 0, 0, 67, 68,
		5, 43, 0, 0, 68, 6, 1, 0, 0, 0, 69, 70, 5, 45, 0, 0, 70, 8, 1, 0, 0, 0,
		71, 72, 5, 42, 0, 0, 72, 10, 1, 0, 0, 0, 73, 74, 5, 47, 0, 0, 74, 12, 1,
		0, 0, 0, 75, 76, 5, 60, 0, 0, 76, 14, 1, 0, 0, 0, 77, 78, 5, 62, 0, 0,
		78, 16, 1, 0, 0, 0, 79, 80, 5, 61, 0, 0, 80, 81, 5, 61, 0, 0, 81, 18, 1,
		0, 0, 0, 82, 83, 5, 33, 0, 0, 83, 84, 5, 61, 0, 0, 84, 20, 1, 0, 0, 0,
		85, 86, 5, 60, 0, 0, 86, 87, 5, 61, 0, 0, 87, 22, 1, 0, 0, 0, 88, 89, 5,
		62, 0, 0, 89, 90, 5, 61, 0, 0, 90, 24, 1, 0, 0, 0, 91, 92, 5, 123, 0, 0,
		92, 26, 1, 0, 0, 0, 93, 94, 5, 40, 0, 0, 94, 28, 1, 0, 0, 0, 95, 96, 5,
		41, 0, 0, 96, 30, 1, 0, 0, 0, 97, 98, 5, 125, 0, 0, 98, 32, 1, 0, 0, 0,
		99, 100, 5, 105, 0, 0, 100, 101, 5, 102, 0, 0, 101, 34, 1, 0, 0, 0, 102,
		103, 5, 101, 0, 0, 103, 104, 5, 108, 0, 0, 104, 105, 5, 115, 0, 0, 105,
		106, 5, 101, 0, 0, 106, 36, 1, 0, 0, 0, 107, 108, 5, 119, 0, 0, 108, 109,
		5, 104, 0, 0, 109, 110, 5, 105, 0, 0, 110, 111, 5, 108, 0, 0, 111, 112,
		5, 101, 0, 0, 112, 38, 1, 0, 0, 0, 113, 114, 5, 102, 0, 0, 114, 115, 5,
		111, 0, 0, 115, 116, 5, 114, 0, 0, 116, 40, 1, 0, 0, 0, 117, 118, 5, 59,
		0, 0, 118, 42, 1, 0, 0, 0, 119, 120, 5, 43, 0, 0, 120, 121, 5, 43, 0, 0,
		121, 44, 1, 0, 0, 0, 122, 123, 5, 45, 0, 0, 123, 124, 5, 45, 0, 0, 124,
		46, 1, 0, 0, 0, 125, 126, 5, 112, 0, 0, 126, 127, 5, 114, 0, 0, 127, 128,
		5, 105, 0, 0, 128, 129, 5, 110, 0, 0, 129, 130, 5, 116, 0, 0, 130, 48,
		1, 0, 0, 0, 131, 132, 5, 116, 0, 0, 132, 133, 5, 114, 0, 0, 133, 134, 5,
		117, 0, 0, 134, 135, 5, 101, 0, 0, 135, 50, 1, 0, 0, 0, 136, 137, 5, 102,
		0, 0, 137, 138, 5, 97, 0, 0, 138, 139, 5, 108, 0, 0, 139, 140, 5, 115,
		0, 0, 140, 141, 5, 101, 0, 0, 141, 52, 1, 0, 0, 0, 142, 146, 7, 0, 0, 0,
		143, 145, 7, 1, 0, 0, 144, 143, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146,
		144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 54, 1, 0, 0, 0, 148, 146, 1,
		0, 0, 0, 149, 151, 7, 2, 0, 0, 150, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0,
		0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 56, 1, 0, 0, 0, 154,
		160, 5, 34, 0, 0, 155, 159, 8, 3, 0, 0, 156, 157, 5, 34, 0, 0, 157, 159,
		5, 34, 0, 0, 158, 155, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 159, 162, 1, 0,
		0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 163, 1, 0, 0, 0,
		162, 160, 1, 0, 0, 0, 163, 164, 5, 34, 0, 0, 164, 58, 1, 0, 0, 0, 165,
		167, 7, 4, 0, 0, 166, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 166,
		1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 171, 6, 29,
		0, 0, 171, 60, 1, 0, 0, 0, 6, 0, 146, 152, 158, 160, 168, 1, 6, 0, 0,
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

// GrammarLexerInit initializes any static state used to implement GrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GrammarLexerInit() {
	staticData := &GrammarLexerLexerStaticData
	staticData.once.Do(grammarlexerLexerInit)
}

// NewGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGrammarLexer(input antlr.CharStream) *GrammarLexer {
	GrammarLexerInit()
	l := new(GrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Grammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GrammarLexer tokens.
const (
	GrammarLexerVAR          = 1
	GrammarLexerEQUALS       = 2
	GrammarLexerPLUS         = 3
	GrammarLexerMINUS        = 4
	GrammarLexerSTAR         = 5
	GrammarLexerSLASH        = 6
	GrammarLexerLESS         = 7
	GrammarLexerGREATER      = 8
	GrammarLexerEQUALEQUAL   = 9
	GrammarLexerBANGEQUAL    = 10
	GrammarLexerLESSEQUAL    = 11
	GrammarLexerGREATEREQUAL = 12
	GrammarLexerLBRACE       = 13
	GrammarLexerLPAREN       = 14
	GrammarLexerRPAREN       = 15
	GrammarLexerRBRACE       = 16
	GrammarLexerIF           = 17
	GrammarLexerELSE         = 18
	GrammarLexerWHILE        = 19
	GrammarLexerFOR          = 20
	GrammarLexerSEMICOLON    = 21
	GrammarLexerINC          = 22
	GrammarLexerDEC          = 23
	GrammarLexerPRINT        = 24
	GrammarLexerTRUE         = 25
	GrammarLexerFALSE        = 26
	GrammarLexerIDENTIFIER   = 27
	GrammarLexerNUMBER       = 28
	GrammarLexerSTRING       = 29
	GrammarLexerWS           = 30
)
