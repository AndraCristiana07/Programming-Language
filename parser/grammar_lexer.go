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
		"'*='", "'/='", "'%='", "'**='", "'&'", "'|'", "'^'", "'<<'", "'>>'",
		"'~'", "'print'", "'true'", "'false'", "'and'", "'or'", "'not'",
	}
	staticData.SymbolicNames = []string{
		"", "VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "LBRACE", "LPAREN",
		"RPAREN", "RBRACE", "IF", "ELSE", "WHILE", "FOR", "SEMICOLON", "INC",
		"DEC", "MODULO", "EXPONENTIAL", "PLUSEQUAL", "MINUSEQUAL", "STAREQUAL",
		"SLASHEQUAL", "MODEQUAL", "EXPONENTIALEQUAL", "BITAND", "BITOR", "BITXOR",
		"BITLSHIFT", "BITRSHIFT", "BITNOT", "PRINT", "TRUE", "FALSE", "AND",
		"OR", "NOT", "IDENTIFIER", "NUMBER", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "LBRACE", "LPAREN",
		"RPAREN", "RBRACE", "IF", "ELSE", "WHILE", "FOR", "SEMICOLON", "INC",
		"DEC", "MODULO", "EXPONENTIAL", "PLUSEQUAL", "MINUSEQUAL", "STAREQUAL",
		"SLASHEQUAL", "MODEQUAL", "EXPONENTIALEQUAL", "BITAND", "BITOR", "BITXOR",
		"BITLSHIFT", "BITRSHIFT", "BITNOT", "PRINT", "TRUE", "FALSE", "AND",
		"OR", "NOT", "IDENTIFIER", "NUMBER", "STRING", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 47, 255, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42,
		1, 42, 1, 43, 1, 43, 5, 43, 228, 8, 43, 10, 43, 12, 43, 231, 9, 43, 1,
		44, 4, 44, 234, 8, 44, 11, 44, 12, 44, 235, 1, 45, 1, 45, 1, 45, 1, 45,
		5, 45, 242, 8, 45, 10, 45, 12, 45, 245, 9, 45, 1, 45, 1, 45, 1, 46, 4,
		46, 250, 8, 46, 11, 46, 12, 46, 251, 1, 46, 1, 46, 0, 0, 47, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79,
		40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 1, 0, 5, 3,
		0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0,
		48, 57, 1, 0, 34, 34, 3, 0, 9, 10, 13, 13, 32, 32, 259, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0,
		41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0,
		0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0,
		0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0,
		0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1,
		0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79,
		1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0,
		87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0,
		1, 95, 1, 0, 0, 0, 3, 99, 1, 0, 0, 0, 5, 101, 1, 0, 0, 0, 7, 103, 1, 0,
		0, 0, 9, 105, 1, 0, 0, 0, 11, 107, 1, 0, 0, 0, 13, 109, 1, 0, 0, 0, 15,
		111, 1, 0, 0, 0, 17, 113, 1, 0, 0, 0, 19, 116, 1, 0, 0, 0, 21, 119, 1,
		0, 0, 0, 23, 122, 1, 0, 0, 0, 25, 125, 1, 0, 0, 0, 27, 127, 1, 0, 0, 0,
		29, 129, 1, 0, 0, 0, 31, 131, 1, 0, 0, 0, 33, 133, 1, 0, 0, 0, 35, 136,
		1, 0, 0, 0, 37, 141, 1, 0, 0, 0, 39, 147, 1, 0, 0, 0, 41, 151, 1, 0, 0,
		0, 43, 153, 1, 0, 0, 0, 45, 156, 1, 0, 0, 0, 47, 159, 1, 0, 0, 0, 49, 161,
		1, 0, 0, 0, 51, 164, 1, 0, 0, 0, 53, 167, 1, 0, 0, 0, 55, 170, 1, 0, 0,
		0, 57, 173, 1, 0, 0, 0, 59, 176, 1, 0, 0, 0, 61, 179, 1, 0, 0, 0, 63, 183,
		1, 0, 0, 0, 65, 185, 1, 0, 0, 0, 67, 187, 1, 0, 0, 0, 69, 189, 1, 0, 0,
		0, 71, 192, 1, 0, 0, 0, 73, 195, 1, 0, 0, 0, 75, 197, 1, 0, 0, 0, 77, 203,
		1, 0, 0, 0, 79, 208, 1, 0, 0, 0, 81, 214, 1, 0, 0, 0, 83, 218, 1, 0, 0,
		0, 85, 221, 1, 0, 0, 0, 87, 225, 1, 0, 0, 0, 89, 233, 1, 0, 0, 0, 91, 237,
		1, 0, 0, 0, 93, 249, 1, 0, 0, 0, 95, 96, 5, 118, 0, 0, 96, 97, 5, 97, 0,
		0, 97, 98, 5, 114, 0, 0, 98, 2, 1, 0, 0, 0, 99, 100, 5, 61, 0, 0, 100,
		4, 1, 0, 0, 0, 101, 102, 5, 43, 0, 0, 102, 6, 1, 0, 0, 0, 103, 104, 5,
		45, 0, 0, 104, 8, 1, 0, 0, 0, 105, 106, 5, 42, 0, 0, 106, 10, 1, 0, 0,
		0, 107, 108, 5, 47, 0, 0, 108, 12, 1, 0, 0, 0, 109, 110, 5, 60, 0, 0, 110,
		14, 1, 0, 0, 0, 111, 112, 5, 62, 0, 0, 112, 16, 1, 0, 0, 0, 113, 114, 5,
		61, 0, 0, 114, 115, 5, 61, 0, 0, 115, 18, 1, 0, 0, 0, 116, 117, 5, 33,
		0, 0, 117, 118, 5, 61, 0, 0, 118, 20, 1, 0, 0, 0, 119, 120, 5, 60, 0, 0,
		120, 121, 5, 61, 0, 0, 121, 22, 1, 0, 0, 0, 122, 123, 5, 62, 0, 0, 123,
		124, 5, 61, 0, 0, 124, 24, 1, 0, 0, 0, 125, 126, 5, 123, 0, 0, 126, 26,
		1, 0, 0, 0, 127, 128, 5, 40, 0, 0, 128, 28, 1, 0, 0, 0, 129, 130, 5, 41,
		0, 0, 130, 30, 1, 0, 0, 0, 131, 132, 5, 125, 0, 0, 132, 32, 1, 0, 0, 0,
		133, 134, 5, 105, 0, 0, 134, 135, 5, 102, 0, 0, 135, 34, 1, 0, 0, 0, 136,
		137, 5, 101, 0, 0, 137, 138, 5, 108, 0, 0, 138, 139, 5, 115, 0, 0, 139,
		140, 5, 101, 0, 0, 140, 36, 1, 0, 0, 0, 141, 142, 5, 119, 0, 0, 142, 143,
		5, 104, 0, 0, 143, 144, 5, 105, 0, 0, 144, 145, 5, 108, 0, 0, 145, 146,
		5, 101, 0, 0, 146, 38, 1, 0, 0, 0, 147, 148, 5, 102, 0, 0, 148, 149, 5,
		111, 0, 0, 149, 150, 5, 114, 0, 0, 150, 40, 1, 0, 0, 0, 151, 152, 5, 59,
		0, 0, 152, 42, 1, 0, 0, 0, 153, 154, 5, 43, 0, 0, 154, 155, 5, 43, 0, 0,
		155, 44, 1, 0, 0, 0, 156, 157, 5, 45, 0, 0, 157, 158, 5, 45, 0, 0, 158,
		46, 1, 0, 0, 0, 159, 160, 5, 37, 0, 0, 160, 48, 1, 0, 0, 0, 161, 162, 5,
		42, 0, 0, 162, 163, 5, 42, 0, 0, 163, 50, 1, 0, 0, 0, 164, 165, 5, 43,
		0, 0, 165, 166, 5, 61, 0, 0, 166, 52, 1, 0, 0, 0, 167, 168, 5, 45, 0, 0,
		168, 169, 5, 61, 0, 0, 169, 54, 1, 0, 0, 0, 170, 171, 5, 42, 0, 0, 171,
		172, 5, 61, 0, 0, 172, 56, 1, 0, 0, 0, 173, 174, 5, 47, 0, 0, 174, 175,
		5, 61, 0, 0, 175, 58, 1, 0, 0, 0, 176, 177, 5, 37, 0, 0, 177, 178, 5, 61,
		0, 0, 178, 60, 1, 0, 0, 0, 179, 180, 5, 42, 0, 0, 180, 181, 5, 42, 0, 0,
		181, 182, 5, 61, 0, 0, 182, 62, 1, 0, 0, 0, 183, 184, 5, 38, 0, 0, 184,
		64, 1, 0, 0, 0, 185, 186, 5, 124, 0, 0, 186, 66, 1, 0, 0, 0, 187, 188,
		5, 94, 0, 0, 188, 68, 1, 0, 0, 0, 189, 190, 5, 60, 0, 0, 190, 191, 5, 60,
		0, 0, 191, 70, 1, 0, 0, 0, 192, 193, 5, 62, 0, 0, 193, 194, 5, 62, 0, 0,
		194, 72, 1, 0, 0, 0, 195, 196, 5, 126, 0, 0, 196, 74, 1, 0, 0, 0, 197,
		198, 5, 112, 0, 0, 198, 199, 5, 114, 0, 0, 199, 200, 5, 105, 0, 0, 200,
		201, 5, 110, 0, 0, 201, 202, 5, 116, 0, 0, 202, 76, 1, 0, 0, 0, 203, 204,
		5, 116, 0, 0, 204, 205, 5, 114, 0, 0, 205, 206, 5, 117, 0, 0, 206, 207,
		5, 101, 0, 0, 207, 78, 1, 0, 0, 0, 208, 209, 5, 102, 0, 0, 209, 210, 5,
		97, 0, 0, 210, 211, 5, 108, 0, 0, 211, 212, 5, 115, 0, 0, 212, 213, 5,
		101, 0, 0, 213, 80, 1, 0, 0, 0, 214, 215, 5, 97, 0, 0, 215, 216, 5, 110,
		0, 0, 216, 217, 5, 100, 0, 0, 217, 82, 1, 0, 0, 0, 218, 219, 5, 111, 0,
		0, 219, 220, 5, 114, 0, 0, 220, 84, 1, 0, 0, 0, 221, 222, 5, 110, 0, 0,
		222, 223, 5, 111, 0, 0, 223, 224, 5, 116, 0, 0, 224, 86, 1, 0, 0, 0, 225,
		229, 7, 0, 0, 0, 226, 228, 7, 1, 0, 0, 227, 226, 1, 0, 0, 0, 228, 231,
		1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 88, 1, 0,
		0, 0, 231, 229, 1, 0, 0, 0, 232, 234, 7, 2, 0, 0, 233, 232, 1, 0, 0, 0,
		234, 235, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236,
		90, 1, 0, 0, 0, 237, 243, 5, 34, 0, 0, 238, 242, 8, 3, 0, 0, 239, 240,
		5, 34, 0, 0, 240, 242, 5, 34, 0, 0, 241, 238, 1, 0, 0, 0, 241, 239, 1,
		0, 0, 0, 242, 245, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0,
		0, 244, 246, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 246, 247, 5, 34, 0, 0, 247,
		92, 1, 0, 0, 0, 248, 250, 7, 4, 0, 0, 249, 248, 1, 0, 0, 0, 250, 251, 1,
		0, 0, 0, 251, 249, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 253, 1, 0, 0,
		0, 253, 254, 6, 46, 0, 0, 254, 94, 1, 0, 0, 0, 6, 0, 229, 235, 241, 243,
		251, 1, 6, 0, 0,
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
	GrammarLexerBITAND           = 32
	GrammarLexerBITOR            = 33
	GrammarLexerBITXOR           = 34
	GrammarLexerBITLSHIFT        = 35
	GrammarLexerBITRSHIFT        = 36
	GrammarLexerBITNOT           = 37
	GrammarLexerPRINT            = 38
	GrammarLexerTRUE             = 39
	GrammarLexerFALSE            = 40
	GrammarLexerAND              = 41
	GrammarLexerOR               = 42
	GrammarLexerNOT              = 43
	GrammarLexerIDENTIFIER       = 44
	GrammarLexerNUMBER           = 45
	GrammarLexerSTRING           = 46
	GrammarLexerWS               = 47
)
