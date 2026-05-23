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
		"'while'", "'for'", "';'", "'print'", "'true'", "'false'",
	}
	staticData.SymbolicNames = []string{
		"", "VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "LBRACE", "LPAREN",
		"RPAREN", "RBRACE", "IF", "ELSE", "WHILE", "FOR", "SEMICOLON", "PRINT",
		"TRUE", "FALSE", "IDENTIFIER", "NUMBER", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "LBRACE", "LPAREN",
		"RPAREN", "RBRACE", "IF", "ELSE", "WHILE", "FOR", "SEMICOLON", "PRINT",
		"TRUE", "FALSE", "IDENTIFIER", "NUMBER", "STRING", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 28, 162, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 24, 1, 24, 5, 24, 135, 8, 24, 10, 24, 12, 24, 138,
		9, 24, 1, 25, 4, 25, 141, 8, 25, 11, 25, 12, 25, 142, 1, 26, 1, 26, 1,
		26, 1, 26, 5, 26, 149, 8, 26, 10, 26, 12, 26, 152, 9, 26, 1, 26, 1, 26,
		1, 27, 4, 27, 157, 8, 27, 11, 27, 12, 27, 158, 1, 27, 1, 27, 0, 0, 28,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 1, 0, 5,
		3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1,
		0, 48, 57, 1, 0, 34, 34, 3, 0, 9, 10, 13, 13, 32, 32, 166, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 1, 57, 1, 0, 0, 0, 3, 61, 1, 0, 0, 0, 5, 63, 1, 0, 0, 0, 7, 65, 1,
		0, 0, 0, 9, 67, 1, 0, 0, 0, 11, 69, 1, 0, 0, 0, 13, 71, 1, 0, 0, 0, 15,
		73, 1, 0, 0, 0, 17, 75, 1, 0, 0, 0, 19, 78, 1, 0, 0, 0, 21, 81, 1, 0, 0,
		0, 23, 84, 1, 0, 0, 0, 25, 87, 1, 0, 0, 0, 27, 89, 1, 0, 0, 0, 29, 91,
		1, 0, 0, 0, 31, 93, 1, 0, 0, 0, 33, 95, 1, 0, 0, 0, 35, 98, 1, 0, 0, 0,
		37, 103, 1, 0, 0, 0, 39, 109, 1, 0, 0, 0, 41, 113, 1, 0, 0, 0, 43, 115,
		1, 0, 0, 0, 45, 121, 1, 0, 0, 0, 47, 126, 1, 0, 0, 0, 49, 132, 1, 0, 0,
		0, 51, 140, 1, 0, 0, 0, 53, 144, 1, 0, 0, 0, 55, 156, 1, 0, 0, 0, 57, 58,
		5, 118, 0, 0, 58, 59, 5, 97, 0, 0, 59, 60, 5, 114, 0, 0, 60, 2, 1, 0, 0,
		0, 61, 62, 5, 61, 0, 0, 62, 4, 1, 0, 0, 0, 63, 64, 5, 43, 0, 0, 64, 6,
		1, 0, 0, 0, 65, 66, 5, 45, 0, 0, 66, 8, 1, 0, 0, 0, 67, 68, 5, 42, 0, 0,
		68, 10, 1, 0, 0, 0, 69, 70, 5, 47, 0, 0, 70, 12, 1, 0, 0, 0, 71, 72, 5,
		60, 0, 0, 72, 14, 1, 0, 0, 0, 73, 74, 5, 62, 0, 0, 74, 16, 1, 0, 0, 0,
		75, 76, 5, 61, 0, 0, 76, 77, 5, 61, 0, 0, 77, 18, 1, 0, 0, 0, 78, 79, 5,
		33, 0, 0, 79, 80, 5, 61, 0, 0, 80, 20, 1, 0, 0, 0, 81, 82, 5, 60, 0, 0,
		82, 83, 5, 61, 0, 0, 83, 22, 1, 0, 0, 0, 84, 85, 5, 62, 0, 0, 85, 86, 5,
		61, 0, 0, 86, 24, 1, 0, 0, 0, 87, 88, 5, 123, 0, 0, 88, 26, 1, 0, 0, 0,
		89, 90, 5, 40, 0, 0, 90, 28, 1, 0, 0, 0, 91, 92, 5, 41, 0, 0, 92, 30, 1,
		0, 0, 0, 93, 94, 5, 125, 0, 0, 94, 32, 1, 0, 0, 0, 95, 96, 5, 105, 0, 0,
		96, 97, 5, 102, 0, 0, 97, 34, 1, 0, 0, 0, 98, 99, 5, 101, 0, 0, 99, 100,
		5, 108, 0, 0, 100, 101, 5, 115, 0, 0, 101, 102, 5, 101, 0, 0, 102, 36,
		1, 0, 0, 0, 103, 104, 5, 119, 0, 0, 104, 105, 5, 104, 0, 0, 105, 106, 5,
		105, 0, 0, 106, 107, 5, 108, 0, 0, 107, 108, 5, 101, 0, 0, 108, 38, 1,
		0, 0, 0, 109, 110, 5, 102, 0, 0, 110, 111, 5, 111, 0, 0, 111, 112, 5, 114,
		0, 0, 112, 40, 1, 0, 0, 0, 113, 114, 5, 59, 0, 0, 114, 42, 1, 0, 0, 0,
		115, 116, 5, 112, 0, 0, 116, 117, 5, 114, 0, 0, 117, 118, 5, 105, 0, 0,
		118, 119, 5, 110, 0, 0, 119, 120, 5, 116, 0, 0, 120, 44, 1, 0, 0, 0, 121,
		122, 5, 116, 0, 0, 122, 123, 5, 114, 0, 0, 123, 124, 5, 117, 0, 0, 124,
		125, 5, 101, 0, 0, 125, 46, 1, 0, 0, 0, 126, 127, 5, 102, 0, 0, 127, 128,
		5, 97, 0, 0, 128, 129, 5, 108, 0, 0, 129, 130, 5, 115, 0, 0, 130, 131,
		5, 101, 0, 0, 131, 48, 1, 0, 0, 0, 132, 136, 7, 0, 0, 0, 133, 135, 7, 1,
		0, 0, 134, 133, 1, 0, 0, 0, 135, 138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0,
		136, 137, 1, 0, 0, 0, 137, 50, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 141,
		7, 2, 0, 0, 140, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 140, 1, 0,
		0, 0, 142, 143, 1, 0, 0, 0, 143, 52, 1, 0, 0, 0, 144, 150, 5, 34, 0, 0,
		145, 149, 8, 3, 0, 0, 146, 147, 5, 34, 0, 0, 147, 149, 5, 34, 0, 0, 148,
		145, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148,
		1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 153, 1, 0, 0, 0, 152, 150, 1, 0,
		0, 0, 153, 154, 5, 34, 0, 0, 154, 54, 1, 0, 0, 0, 155, 157, 7, 4, 0, 0,
		156, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 158,
		159, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 161, 6, 27, 0, 0, 161, 56,
		1, 0, 0, 0, 6, 0, 136, 142, 148, 150, 158, 1, 6, 0, 0,
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
	GrammarLexerPRINT        = 22
	GrammarLexerTRUE         = 23
	GrammarLexerFALSE        = 24
	GrammarLexerIDENTIFIER   = 25
	GrammarLexerNUMBER       = 26
	GrammarLexerSTRING       = 27
	GrammarLexerWS           = 28
)
