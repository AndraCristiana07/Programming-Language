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
		"'!='", "'true'", "'false'",
	}
	staticData.SymbolicNames = []string{
		"", "VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "TRUE", "FALSE", "IDENTIFIER", "NUMBER",
		"WS",
	}
	staticData.RuleNames = []string{
		"VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "TRUE", "FALSE", "IDENTIFIER", "NUMBER",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 85, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 5, 12, 69, 8, 12, 10, 12, 12, 12, 72, 9, 12, 1, 13, 4, 13, 75,
		8, 13, 11, 13, 12, 13, 76, 1, 14, 4, 14, 80, 8, 14, 11, 14, 12, 14, 81,
		1, 14, 1, 14, 0, 0, 15, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 1, 0, 4, 3, 0,
		65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48,
		57, 3, 0, 9, 10, 13, 13, 32, 32, 87, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0,
		0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0,
		0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0,
		0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0,
		0, 0, 0, 29, 1, 0, 0, 0, 1, 31, 1, 0, 0, 0, 3, 35, 1, 0, 0, 0, 5, 37, 1,
		0, 0, 0, 7, 39, 1, 0, 0, 0, 9, 41, 1, 0, 0, 0, 11, 43, 1, 0, 0, 0, 13,
		45, 1, 0, 0, 0, 15, 47, 1, 0, 0, 0, 17, 49, 1, 0, 0, 0, 19, 52, 1, 0, 0,
		0, 21, 55, 1, 0, 0, 0, 23, 60, 1, 0, 0, 0, 25, 66, 1, 0, 0, 0, 27, 74,
		1, 0, 0, 0, 29, 79, 1, 0, 0, 0, 31, 32, 5, 118, 0, 0, 32, 33, 5, 97, 0,
		0, 33, 34, 5, 114, 0, 0, 34, 2, 1, 0, 0, 0, 35, 36, 5, 61, 0, 0, 36, 4,
		1, 0, 0, 0, 37, 38, 5, 43, 0, 0, 38, 6, 1, 0, 0, 0, 39, 40, 5, 45, 0, 0,
		40, 8, 1, 0, 0, 0, 41, 42, 5, 42, 0, 0, 42, 10, 1, 0, 0, 0, 43, 44, 5,
		47, 0, 0, 44, 12, 1, 0, 0, 0, 45, 46, 5, 60, 0, 0, 46, 14, 1, 0, 0, 0,
		47, 48, 5, 62, 0, 0, 48, 16, 1, 0, 0, 0, 49, 50, 5, 61, 0, 0, 50, 51, 5,
		61, 0, 0, 51, 18, 1, 0, 0, 0, 52, 53, 5, 33, 0, 0, 53, 54, 5, 61, 0, 0,
		54, 20, 1, 0, 0, 0, 55, 56, 5, 116, 0, 0, 56, 57, 5, 114, 0, 0, 57, 58,
		5, 117, 0, 0, 58, 59, 5, 101, 0, 0, 59, 22, 1, 0, 0, 0, 60, 61, 5, 102,
		0, 0, 61, 62, 5, 97, 0, 0, 62, 63, 5, 108, 0, 0, 63, 64, 5, 115, 0, 0,
		64, 65, 5, 101, 0, 0, 65, 24, 1, 0, 0, 0, 66, 70, 7, 0, 0, 0, 67, 69, 7,
		1, 0, 0, 68, 67, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70,
		71, 1, 0, 0, 0, 71, 26, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 75, 7, 2, 0,
		0, 74, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77,
		1, 0, 0, 0, 77, 28, 1, 0, 0, 0, 78, 80, 7, 3, 0, 0, 79, 78, 1, 0, 0, 0,
		80, 81, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 83, 1,
		0, 0, 0, 83, 84, 6, 14, 0, 0, 84, 30, 1, 0, 0, 0, 4, 0, 70, 76, 81, 1,
		6, 0, 0,
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
	GrammarLexerVAR        = 1
	GrammarLexerEQUALS     = 2
	GrammarLexerPLUS       = 3
	GrammarLexerMINUS      = 4
	GrammarLexerSTAR       = 5
	GrammarLexerSLASH      = 6
	GrammarLexerLESS       = 7
	GrammarLexerGREATER    = 8
	GrammarLexerEQUALEQUAL = 9
	GrammarLexerBANGEQUAL  = 10
	GrammarLexerTRUE       = 11
	GrammarLexerFALSE      = 12
	GrammarLexerIDENTIFIER = 13
	GrammarLexerNUMBER     = 14
	GrammarLexerWS         = 15
)
