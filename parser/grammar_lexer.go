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
		"'!='", "'<='", "'>='", "'print'", "'true'", "'false'",
	}
	staticData.SymbolicNames = []string{
		"", "VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "PRINT", "TRUE",
		"FALSE", "IDENTIFIER", "NUMBER", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "PRINT", "TRUE",
		"FALSE", "IDENTIFIER", "NUMBER", "STRING", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 116, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 5, 15, 89, 8, 15, 10, 15, 12, 15, 92, 9, 15, 1, 16, 4, 16, 95, 8, 16,
		11, 16, 12, 16, 96, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 103, 8, 17, 10,
		17, 12, 17, 106, 9, 17, 1, 17, 1, 17, 1, 18, 4, 18, 111, 8, 18, 11, 18,
		12, 18, 112, 1, 18, 1, 18, 0, 0, 19, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15,
		31, 16, 33, 17, 35, 18, 37, 19, 1, 0, 5, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 1, 0, 34, 34, 3, 0,
		9, 10, 13, 13, 32, 32, 120, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 1, 39, 1, 0, 0, 0, 3, 43, 1, 0, 0, 0, 5, 45, 1, 0,
		0, 0, 7, 47, 1, 0, 0, 0, 9, 49, 1, 0, 0, 0, 11, 51, 1, 0, 0, 0, 13, 53,
		1, 0, 0, 0, 15, 55, 1, 0, 0, 0, 17, 57, 1, 0, 0, 0, 19, 60, 1, 0, 0, 0,
		21, 63, 1, 0, 0, 0, 23, 66, 1, 0, 0, 0, 25, 69, 1, 0, 0, 0, 27, 75, 1,
		0, 0, 0, 29, 80, 1, 0, 0, 0, 31, 86, 1, 0, 0, 0, 33, 94, 1, 0, 0, 0, 35,
		98, 1, 0, 0, 0, 37, 110, 1, 0, 0, 0, 39, 40, 5, 118, 0, 0, 40, 41, 5, 97,
		0, 0, 41, 42, 5, 114, 0, 0, 42, 2, 1, 0, 0, 0, 43, 44, 5, 61, 0, 0, 44,
		4, 1, 0, 0, 0, 45, 46, 5, 43, 0, 0, 46, 6, 1, 0, 0, 0, 47, 48, 5, 45, 0,
		0, 48, 8, 1, 0, 0, 0, 49, 50, 5, 42, 0, 0, 50, 10, 1, 0, 0, 0, 51, 52,
		5, 47, 0, 0, 52, 12, 1, 0, 0, 0, 53, 54, 5, 60, 0, 0, 54, 14, 1, 0, 0,
		0, 55, 56, 5, 62, 0, 0, 56, 16, 1, 0, 0, 0, 57, 58, 5, 61, 0, 0, 58, 59,
		5, 61, 0, 0, 59, 18, 1, 0, 0, 0, 60, 61, 5, 33, 0, 0, 61, 62, 5, 61, 0,
		0, 62, 20, 1, 0, 0, 0, 63, 64, 5, 60, 0, 0, 64, 65, 5, 61, 0, 0, 65, 22,
		1, 0, 0, 0, 66, 67, 5, 62, 0, 0, 67, 68, 5, 61, 0, 0, 68, 24, 1, 0, 0,
		0, 69, 70, 5, 112, 0, 0, 70, 71, 5, 114, 0, 0, 71, 72, 5, 105, 0, 0, 72,
		73, 5, 110, 0, 0, 73, 74, 5, 116, 0, 0, 74, 26, 1, 0, 0, 0, 75, 76, 5,
		116, 0, 0, 76, 77, 5, 114, 0, 0, 77, 78, 5, 117, 0, 0, 78, 79, 5, 101,
		0, 0, 79, 28, 1, 0, 0, 0, 80, 81, 5, 102, 0, 0, 81, 82, 5, 97, 0, 0, 82,
		83, 5, 108, 0, 0, 83, 84, 5, 115, 0, 0, 84, 85, 5, 101, 0, 0, 85, 30, 1,
		0, 0, 0, 86, 90, 7, 0, 0, 0, 87, 89, 7, 1, 0, 0, 88, 87, 1, 0, 0, 0, 89,
		92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 32, 1, 0, 0,
		0, 92, 90, 1, 0, 0, 0, 93, 95, 7, 2, 0, 0, 94, 93, 1, 0, 0, 0, 95, 96,
		1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 34, 1, 0, 0, 0,
		98, 104, 5, 34, 0, 0, 99, 103, 8, 3, 0, 0, 100, 101, 5, 34, 0, 0, 101,
		103, 5, 34, 0, 0, 102, 99, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 103, 106,
		1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 107, 1, 0,
		0, 0, 106, 104, 1, 0, 0, 0, 107, 108, 5, 34, 0, 0, 108, 36, 1, 0, 0, 0,
		109, 111, 7, 4, 0, 0, 110, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112,
		110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115,
		6, 18, 0, 0, 115, 38, 1, 0, 0, 0, 6, 0, 90, 96, 102, 104, 112, 1, 6, 0,
		0,
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
	GrammarLexerPRINT        = 13
	GrammarLexerTRUE         = 14
	GrammarLexerFALSE        = 15
	GrammarLexerIDENTIFIER   = 16
	GrammarLexerNUMBER       = 17
	GrammarLexerSTRING       = 18
	GrammarLexerWS           = 19
)
