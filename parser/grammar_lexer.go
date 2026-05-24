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
		"'while'", "'for'", "';'", "'++'", "'--'", "'%'", "'**'", "'+='", "'-='",
		"'*='", "'/='", "'%='", "'**='", "'print'", "'true'", "'false'",
	}
	staticData.SymbolicNames = []string{
		"", "VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "LBRACE", "LPAREN",
		"RPAREN", "RBRACE", "IF", "ELSE", "WHILE", "FOR", "SEMICOLON", "INC",
		"DEC", "MODULO", "EXPONENTIAL", "PLUSEQUAL", "MINUSEQUAL", "STAREQUAL",
		"SLASHEQUAL", "MODEQUAL", "EXPONENTIALEQUAL", "PRINT", "TRUE", "FALSE",
		"IDENTIFIER", "NUMBER", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "LBRACE", "LPAREN",
		"RPAREN", "RBRACE", "IF", "ELSE", "WHILE", "FOR", "SEMICOLON", "INC",
		"DEC", "MODULO", "EXPONENTIAL", "PLUSEQUAL", "MINUSEQUAL", "STAREQUAL",
		"SLASHEQUAL", "MODEQUAL", "EXPONENTIALEQUAL", "PRINT", "TRUE", "FALSE",
		"IDENTIFIER", "NUMBER", "STRING", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 38, 212, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 34, 1, 34, 5, 34, 185, 8, 34, 10, 34, 12, 34, 188, 9, 34, 1, 35, 4,
		35, 191, 8, 35, 11, 35, 12, 35, 192, 1, 36, 1, 36, 1, 36, 1, 36, 5, 36,
		199, 8, 36, 10, 36, 12, 36, 202, 9, 36, 1, 36, 1, 36, 1, 37, 4, 37, 207,
		8, 37, 11, 37, 12, 37, 208, 1, 37, 1, 37, 0, 0, 38, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63,
		32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 1, 0, 5, 3, 0, 65,
		90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57,
		1, 0, 34, 34, 3, 0, 9, 10, 13, 13, 32, 32, 216, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1,
		0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57,
		1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0,
		65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0,
		0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 1, 77, 1, 0, 0, 0, 3, 81, 1, 0, 0,
		0, 5, 83, 1, 0, 0, 0, 7, 85, 1, 0, 0, 0, 9, 87, 1, 0, 0, 0, 11, 89, 1,
		0, 0, 0, 13, 91, 1, 0, 0, 0, 15, 93, 1, 0, 0, 0, 17, 95, 1, 0, 0, 0, 19,
		98, 1, 0, 0, 0, 21, 101, 1, 0, 0, 0, 23, 104, 1, 0, 0, 0, 25, 107, 1, 0,
		0, 0, 27, 109, 1, 0, 0, 0, 29, 111, 1, 0, 0, 0, 31, 113, 1, 0, 0, 0, 33,
		115, 1, 0, 0, 0, 35, 118, 1, 0, 0, 0, 37, 123, 1, 0, 0, 0, 39, 129, 1,
		0, 0, 0, 41, 133, 1, 0, 0, 0, 43, 135, 1, 0, 0, 0, 45, 138, 1, 0, 0, 0,
		47, 141, 1, 0, 0, 0, 49, 143, 1, 0, 0, 0, 51, 146, 1, 0, 0, 0, 53, 149,
		1, 0, 0, 0, 55, 152, 1, 0, 0, 0, 57, 155, 1, 0, 0, 0, 59, 158, 1, 0, 0,
		0, 61, 161, 1, 0, 0, 0, 63, 165, 1, 0, 0, 0, 65, 171, 1, 0, 0, 0, 67, 176,
		1, 0, 0, 0, 69, 182, 1, 0, 0, 0, 71, 190, 1, 0, 0, 0, 73, 194, 1, 0, 0,
		0, 75, 206, 1, 0, 0, 0, 77, 78, 5, 118, 0, 0, 78, 79, 5, 97, 0, 0, 79,
		80, 5, 114, 0, 0, 80, 2, 1, 0, 0, 0, 81, 82, 5, 61, 0, 0, 82, 4, 1, 0,
		0, 0, 83, 84, 5, 43, 0, 0, 84, 6, 1, 0, 0, 0, 85, 86, 5, 45, 0, 0, 86,
		8, 1, 0, 0, 0, 87, 88, 5, 42, 0, 0, 88, 10, 1, 0, 0, 0, 89, 90, 5, 47,
		0, 0, 90, 12, 1, 0, 0, 0, 91, 92, 5, 60, 0, 0, 92, 14, 1, 0, 0, 0, 93,
		94, 5, 62, 0, 0, 94, 16, 1, 0, 0, 0, 95, 96, 5, 61, 0, 0, 96, 97, 5, 61,
		0, 0, 97, 18, 1, 0, 0, 0, 98, 99, 5, 33, 0, 0, 99, 100, 5, 61, 0, 0, 100,
		20, 1, 0, 0, 0, 101, 102, 5, 60, 0, 0, 102, 103, 5, 61, 0, 0, 103, 22,
		1, 0, 0, 0, 104, 105, 5, 62, 0, 0, 105, 106, 5, 61, 0, 0, 106, 24, 1, 0,
		0, 0, 107, 108, 5, 123, 0, 0, 108, 26, 1, 0, 0, 0, 109, 110, 5, 40, 0,
		0, 110, 28, 1, 0, 0, 0, 111, 112, 5, 41, 0, 0, 112, 30, 1, 0, 0, 0, 113,
		114, 5, 125, 0, 0, 114, 32, 1, 0, 0, 0, 115, 116, 5, 105, 0, 0, 116, 117,
		5, 102, 0, 0, 117, 34, 1, 0, 0, 0, 118, 119, 5, 101, 0, 0, 119, 120, 5,
		108, 0, 0, 120, 121, 5, 115, 0, 0, 121, 122, 5, 101, 0, 0, 122, 36, 1,
		0, 0, 0, 123, 124, 5, 119, 0, 0, 124, 125, 5, 104, 0, 0, 125, 126, 5, 105,
		0, 0, 126, 127, 5, 108, 0, 0, 127, 128, 5, 101, 0, 0, 128, 38, 1, 0, 0,
		0, 129, 130, 5, 102, 0, 0, 130, 131, 5, 111, 0, 0, 131, 132, 5, 114, 0,
		0, 132, 40, 1, 0, 0, 0, 133, 134, 5, 59, 0, 0, 134, 42, 1, 0, 0, 0, 135,
		136, 5, 43, 0, 0, 136, 137, 5, 43, 0, 0, 137, 44, 1, 0, 0, 0, 138, 139,
		5, 45, 0, 0, 139, 140, 5, 45, 0, 0, 140, 46, 1, 0, 0, 0, 141, 142, 5, 37,
		0, 0, 142, 48, 1, 0, 0, 0, 143, 144, 5, 42, 0, 0, 144, 145, 5, 42, 0, 0,
		145, 50, 1, 0, 0, 0, 146, 147, 5, 43, 0, 0, 147, 148, 5, 61, 0, 0, 148,
		52, 1, 0, 0, 0, 149, 150, 5, 45, 0, 0, 150, 151, 5, 61, 0, 0, 151, 54,
		1, 0, 0, 0, 152, 153, 5, 42, 0, 0, 153, 154, 5, 61, 0, 0, 154, 56, 1, 0,
		0, 0, 155, 156, 5, 47, 0, 0, 156, 157, 5, 61, 0, 0, 157, 58, 1, 0, 0, 0,
		158, 159, 5, 37, 0, 0, 159, 160, 5, 61, 0, 0, 160, 60, 1, 0, 0, 0, 161,
		162, 5, 42, 0, 0, 162, 163, 5, 42, 0, 0, 163, 164, 5, 61, 0, 0, 164, 62,
		1, 0, 0, 0, 165, 166, 5, 112, 0, 0, 166, 167, 5, 114, 0, 0, 167, 168, 5,
		105, 0, 0, 168, 169, 5, 110, 0, 0, 169, 170, 5, 116, 0, 0, 170, 64, 1,
		0, 0, 0, 171, 172, 5, 116, 0, 0, 172, 173, 5, 114, 0, 0, 173, 174, 5, 117,
		0, 0, 174, 175, 5, 101, 0, 0, 175, 66, 1, 0, 0, 0, 176, 177, 5, 102, 0,
		0, 177, 178, 5, 97, 0, 0, 178, 179, 5, 108, 0, 0, 179, 180, 5, 115, 0,
		0, 180, 181, 5, 101, 0, 0, 181, 68, 1, 0, 0, 0, 182, 186, 7, 0, 0, 0, 183,
		185, 7, 1, 0, 0, 184, 183, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186, 184,
		1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 70, 1, 0, 0, 0, 188, 186, 1, 0,
		0, 0, 189, 191, 7, 2, 0, 0, 190, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0,
		192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 72, 1, 0, 0, 0, 194, 200,
		5, 34, 0, 0, 195, 199, 8, 3, 0, 0, 196, 197, 5, 34, 0, 0, 197, 199, 5,
		34, 0, 0, 198, 195, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 202, 1, 0, 0,
		0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 203, 1, 0, 0, 0, 202,
		200, 1, 0, 0, 0, 203, 204, 5, 34, 0, 0, 204, 74, 1, 0, 0, 0, 205, 207,
		7, 4, 0, 0, 206, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 206, 1, 0,
		0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 6, 37, 0, 0,
		211, 76, 1, 0, 0, 0, 6, 0, 186, 192, 198, 200, 208, 1, 6, 0, 0,
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
	GrammarLexerVAR              = 1
	GrammarLexerEQUALS           = 2
	GrammarLexerPLUS             = 3
	GrammarLexerMINUS            = 4
	GrammarLexerSTAR             = 5
	GrammarLexerSLASH            = 6
	GrammarLexerLESS             = 7
	GrammarLexerGREATER          = 8
	GrammarLexerEQUALEQUAL       = 9
	GrammarLexerBANGEQUAL        = 10
	GrammarLexerLESSEQUAL        = 11
	GrammarLexerGREATEREQUAL     = 12
	GrammarLexerLBRACE           = 13
	GrammarLexerLPAREN           = 14
	GrammarLexerRPAREN           = 15
	GrammarLexerRBRACE           = 16
	GrammarLexerIF               = 17
	GrammarLexerELSE             = 18
	GrammarLexerWHILE            = 19
	GrammarLexerFOR              = 20
	GrammarLexerSEMICOLON        = 21
	GrammarLexerINC              = 22
	GrammarLexerDEC              = 23
	GrammarLexerMODULO           = 24
	GrammarLexerEXPONENTIAL      = 25
	GrammarLexerPLUSEQUAL        = 26
	GrammarLexerMINUSEQUAL       = 27
	GrammarLexerSTAREQUAL        = 28
	GrammarLexerSLASHEQUAL       = 29
	GrammarLexerMODEQUAL         = 30
	GrammarLexerEXPONENTIALEQUAL = 31
	GrammarLexerPRINT            = 32
	GrammarLexerTRUE             = 33
	GrammarLexerFALSE            = 34
	GrammarLexerIDENTIFIER       = 35
	GrammarLexerNUMBER           = 36
	GrammarLexerSTRING           = 37
	GrammarLexerWS               = 38
)
