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
		"'!='", "'<='", "'>='", "'{'", "'('", "')'", "'}'", "'['", "']'", "'if'",
		"'else'", "'while'", "'for'", "';'", "','", "'++'", "'--'", "'%'", "'**'",
		"'+='", "'-='", "'*='", "'/='", "'%='", "'**='", "'&'", "'|'", "'^'",
		"'<<'", "'>>'", "'~'", "'func'", "'return'", "'print'", "'true'", "'false'",
		"'and'", "'or'", "'not'",
	}
	staticData.SymbolicNames = []string{
		"", "VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "LBRACE", "LPAREN",
		"RPAREN", "RBRACE", "LBRACKET", "RBRACKET", "IF", "ELSE", "WHILE", "FOR",
		"SEMICOLON", "COMMA", "INC", "DEC", "MODULO", "EXPONENTIAL", "PLUSEQUAL",
		"MINUSEQUAL", "STAREQUAL", "SLASHEQUAL", "MODEQUAL", "EXPONENTIALEQUAL",
		"BITAND", "BITOR", "BITXOR", "BITLSHIFT", "BITRSHIFT", "BITNOT", "FUNC",
		"RETURN", "PRINT", "TRUE", "FALSE", "AND", "OR", "NOT", "IDENTIFIER",
		"NUMBER", "STRING", "NL", "WS",
	}
	staticData.RuleNames = []string{
		"VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "LBRACE", "LPAREN",
		"RPAREN", "RBRACE", "LBRACKET", "RBRACKET", "IF", "ELSE", "WHILE", "FOR",
		"SEMICOLON", "COMMA", "INC", "DEC", "MODULO", "EXPONENTIAL", "PLUSEQUAL",
		"MINUSEQUAL", "STAREQUAL", "SLASHEQUAL", "MODEQUAL", "EXPONENTIALEQUAL",
		"BITAND", "BITOR", "BITXOR", "BITLSHIFT", "BITRSHIFT", "BITNOT", "FUNC",
		"RETURN", "PRINT", "TRUE", "FALSE", "AND", "OR", "NOT", "IDENTIFIER",
		"NUMBER", "STRING", "NL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 53, 290, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1,
		27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1,
		34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38,
		1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43,
		1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47,
		1, 48, 1, 48, 5, 48, 258, 8, 48, 10, 48, 12, 48, 261, 9, 48, 1, 49, 4,
		49, 264, 8, 49, 11, 49, 12, 49, 265, 1, 50, 1, 50, 1, 50, 1, 50, 5, 50,
		272, 8, 50, 10, 50, 12, 50, 275, 9, 50, 1, 50, 1, 50, 1, 51, 4, 51, 280,
		8, 51, 11, 51, 12, 51, 281, 1, 52, 4, 52, 285, 8, 52, 11, 52, 12, 52, 286,
		1, 52, 1, 52, 0, 0, 53, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44,
		89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105,
		53, 1, 0, 6, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 1, 0, 48, 57, 1, 0, 34, 34, 2, 0, 10, 10, 13, 13, 2, 0, 9, 9,
		32, 32, 295, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53,
		1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0,
		61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0,
		0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0,
		0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0,
		0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1,
		0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99,
		1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0,
		1, 107, 1, 0, 0, 0, 3, 111, 1, 0, 0, 0, 5, 113, 1, 0, 0, 0, 7, 115, 1,
		0, 0, 0, 9, 117, 1, 0, 0, 0, 11, 119, 1, 0, 0, 0, 13, 121, 1, 0, 0, 0,
		15, 123, 1, 0, 0, 0, 17, 125, 1, 0, 0, 0, 19, 128, 1, 0, 0, 0, 21, 131,
		1, 0, 0, 0, 23, 134, 1, 0, 0, 0, 25, 137, 1, 0, 0, 0, 27, 139, 1, 0, 0,
		0, 29, 141, 1, 0, 0, 0, 31, 143, 1, 0, 0, 0, 33, 145, 1, 0, 0, 0, 35, 147,
		1, 0, 0, 0, 37, 149, 1, 0, 0, 0, 39, 152, 1, 0, 0, 0, 41, 157, 1, 0, 0,
		0, 43, 163, 1, 0, 0, 0, 45, 167, 1, 0, 0, 0, 47, 169, 1, 0, 0, 0, 49, 171,
		1, 0, 0, 0, 51, 174, 1, 0, 0, 0, 53, 177, 1, 0, 0, 0, 55, 179, 1, 0, 0,
		0, 57, 182, 1, 0, 0, 0, 59, 185, 1, 0, 0, 0, 61, 188, 1, 0, 0, 0, 63, 191,
		1, 0, 0, 0, 65, 194, 1, 0, 0, 0, 67, 197, 1, 0, 0, 0, 69, 201, 1, 0, 0,
		0, 71, 203, 1, 0, 0, 0, 73, 205, 1, 0, 0, 0, 75, 207, 1, 0, 0, 0, 77, 210,
		1, 0, 0, 0, 79, 213, 1, 0, 0, 0, 81, 215, 1, 0, 0, 0, 83, 220, 1, 0, 0,
		0, 85, 227, 1, 0, 0, 0, 87, 233, 1, 0, 0, 0, 89, 238, 1, 0, 0, 0, 91, 244,
		1, 0, 0, 0, 93, 248, 1, 0, 0, 0, 95, 251, 1, 0, 0, 0, 97, 255, 1, 0, 0,
		0, 99, 263, 1, 0, 0, 0, 101, 267, 1, 0, 0, 0, 103, 279, 1, 0, 0, 0, 105,
		284, 1, 0, 0, 0, 107, 108, 5, 118, 0, 0, 108, 109, 5, 97, 0, 0, 109, 110,
		5, 114, 0, 0, 110, 2, 1, 0, 0, 0, 111, 112, 5, 61, 0, 0, 112, 4, 1, 0,
		0, 0, 113, 114, 5, 43, 0, 0, 114, 6, 1, 0, 0, 0, 115, 116, 5, 45, 0, 0,
		116, 8, 1, 0, 0, 0, 117, 118, 5, 42, 0, 0, 118, 10, 1, 0, 0, 0, 119, 120,
		5, 47, 0, 0, 120, 12, 1, 0, 0, 0, 121, 122, 5, 60, 0, 0, 122, 14, 1, 0,
		0, 0, 123, 124, 5, 62, 0, 0, 124, 16, 1, 0, 0, 0, 125, 126, 5, 61, 0, 0,
		126, 127, 5, 61, 0, 0, 127, 18, 1, 0, 0, 0, 128, 129, 5, 33, 0, 0, 129,
		130, 5, 61, 0, 0, 130, 20, 1, 0, 0, 0, 131, 132, 5, 60, 0, 0, 132, 133,
		5, 61, 0, 0, 133, 22, 1, 0, 0, 0, 134, 135, 5, 62, 0, 0, 135, 136, 5, 61,
		0, 0, 136, 24, 1, 0, 0, 0, 137, 138, 5, 123, 0, 0, 138, 26, 1, 0, 0, 0,
		139, 140, 5, 40, 0, 0, 140, 28, 1, 0, 0, 0, 141, 142, 5, 41, 0, 0, 142,
		30, 1, 0, 0, 0, 143, 144, 5, 125, 0, 0, 144, 32, 1, 0, 0, 0, 145, 146,
		5, 91, 0, 0, 146, 34, 1, 0, 0, 0, 147, 148, 5, 93, 0, 0, 148, 36, 1, 0,
		0, 0, 149, 150, 5, 105, 0, 0, 150, 151, 5, 102, 0, 0, 151, 38, 1, 0, 0,
		0, 152, 153, 5, 101, 0, 0, 153, 154, 5, 108, 0, 0, 154, 155, 5, 115, 0,
		0, 155, 156, 5, 101, 0, 0, 156, 40, 1, 0, 0, 0, 157, 158, 5, 119, 0, 0,
		158, 159, 5, 104, 0, 0, 159, 160, 5, 105, 0, 0, 160, 161, 5, 108, 0, 0,
		161, 162, 5, 101, 0, 0, 162, 42, 1, 0, 0, 0, 163, 164, 5, 102, 0, 0, 164,
		165, 5, 111, 0, 0, 165, 166, 5, 114, 0, 0, 166, 44, 1, 0, 0, 0, 167, 168,
		5, 59, 0, 0, 168, 46, 1, 0, 0, 0, 169, 170, 5, 44, 0, 0, 170, 48, 1, 0,
		0, 0, 171, 172, 5, 43, 0, 0, 172, 173, 5, 43, 0, 0, 173, 50, 1, 0, 0, 0,
		174, 175, 5, 45, 0, 0, 175, 176, 5, 45, 0, 0, 176, 52, 1, 0, 0, 0, 177,
		178, 5, 37, 0, 0, 178, 54, 1, 0, 0, 0, 179, 180, 5, 42, 0, 0, 180, 181,
		5, 42, 0, 0, 181, 56, 1, 0, 0, 0, 182, 183, 5, 43, 0, 0, 183, 184, 5, 61,
		0, 0, 184, 58, 1, 0, 0, 0, 185, 186, 5, 45, 0, 0, 186, 187, 5, 61, 0, 0,
		187, 60, 1, 0, 0, 0, 188, 189, 5, 42, 0, 0, 189, 190, 5, 61, 0, 0, 190,
		62, 1, 0, 0, 0, 191, 192, 5, 47, 0, 0, 192, 193, 5, 61, 0, 0, 193, 64,
		1, 0, 0, 0, 194, 195, 5, 37, 0, 0, 195, 196, 5, 61, 0, 0, 196, 66, 1, 0,
		0, 0, 197, 198, 5, 42, 0, 0, 198, 199, 5, 42, 0, 0, 199, 200, 5, 61, 0,
		0, 200, 68, 1, 0, 0, 0, 201, 202, 5, 38, 0, 0, 202, 70, 1, 0, 0, 0, 203,
		204, 5, 124, 0, 0, 204, 72, 1, 0, 0, 0, 205, 206, 5, 94, 0, 0, 206, 74,
		1, 0, 0, 0, 207, 208, 5, 60, 0, 0, 208, 209, 5, 60, 0, 0, 209, 76, 1, 0,
		0, 0, 210, 211, 5, 62, 0, 0, 211, 212, 5, 62, 0, 0, 212, 78, 1, 0, 0, 0,
		213, 214, 5, 126, 0, 0, 214, 80, 1, 0, 0, 0, 215, 216, 5, 102, 0, 0, 216,
		217, 5, 117, 0, 0, 217, 218, 5, 110, 0, 0, 218, 219, 5, 99, 0, 0, 219,
		82, 1, 0, 0, 0, 220, 221, 5, 114, 0, 0, 221, 222, 5, 101, 0, 0, 222, 223,
		5, 116, 0, 0, 223, 224, 5, 117, 0, 0, 224, 225, 5, 114, 0, 0, 225, 226,
		5, 110, 0, 0, 226, 84, 1, 0, 0, 0, 227, 228, 5, 112, 0, 0, 228, 229, 5,
		114, 0, 0, 229, 230, 5, 105, 0, 0, 230, 231, 5, 110, 0, 0, 231, 232, 5,
		116, 0, 0, 232, 86, 1, 0, 0, 0, 233, 234, 5, 116, 0, 0, 234, 235, 5, 114,
		0, 0, 235, 236, 5, 117, 0, 0, 236, 237, 5, 101, 0, 0, 237, 88, 1, 0, 0,
		0, 238, 239, 5, 102, 0, 0, 239, 240, 5, 97, 0, 0, 240, 241, 5, 108, 0,
		0, 241, 242, 5, 115, 0, 0, 242, 243, 5, 101, 0, 0, 243, 90, 1, 0, 0, 0,
		244, 245, 5, 97, 0, 0, 245, 246, 5, 110, 0, 0, 246, 247, 5, 100, 0, 0,
		247, 92, 1, 0, 0, 0, 248, 249, 5, 111, 0, 0, 249, 250, 5, 114, 0, 0, 250,
		94, 1, 0, 0, 0, 251, 252, 5, 110, 0, 0, 252, 253, 5, 111, 0, 0, 253, 254,
		5, 116, 0, 0, 254, 96, 1, 0, 0, 0, 255, 259, 7, 0, 0, 0, 256, 258, 7, 1,
		0, 0, 257, 256, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0,
		259, 260, 1, 0, 0, 0, 260, 98, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 264,
		7, 2, 0, 0, 263, 262, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 263, 1, 0,
		0, 0, 265, 266, 1, 0, 0, 0, 266, 100, 1, 0, 0, 0, 267, 273, 5, 34, 0, 0,
		268, 272, 8, 3, 0, 0, 269, 270, 5, 34, 0, 0, 270, 272, 5, 34, 0, 0, 271,
		268, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 272, 275, 1, 0, 0, 0, 273, 271,
		1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 276, 1, 0, 0, 0, 275, 273, 1, 0,
		0, 0, 276, 277, 5, 34, 0, 0, 277, 102, 1, 0, 0, 0, 278, 280, 7, 4, 0, 0,
		279, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 281,
		282, 1, 0, 0, 0, 282, 104, 1, 0, 0, 0, 283, 285, 7, 5, 0, 0, 284, 283,
		1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0, 286, 287, 1, 0,
		0, 0, 287, 288, 1, 0, 0, 0, 288, 289, 6, 52, 0, 0, 289, 106, 1, 0, 0, 0,
		7, 0, 259, 265, 271, 273, 281, 286, 1, 6, 0, 0,
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
	GrammarLexerLBRACKET         = 17
	GrammarLexerRBRACKET         = 18
	GrammarLexerIF               = 19
	GrammarLexerELSE             = 20
	GrammarLexerWHILE            = 21
	GrammarLexerFOR              = 22
	GrammarLexerSEMICOLON        = 23
	GrammarLexerCOMMA            = 24
	GrammarLexerINC              = 25
	GrammarLexerDEC              = 26
	GrammarLexerMODULO           = 27
	GrammarLexerEXPONENTIAL      = 28
	GrammarLexerPLUSEQUAL        = 29
	GrammarLexerMINUSEQUAL       = 30
	GrammarLexerSTAREQUAL        = 31
	GrammarLexerSLASHEQUAL       = 32
	GrammarLexerMODEQUAL         = 33
	GrammarLexerEXPONENTIALEQUAL = 34
	GrammarLexerBITAND           = 35
	GrammarLexerBITOR            = 36
	GrammarLexerBITXOR           = 37
	GrammarLexerBITLSHIFT        = 38
	GrammarLexerBITRSHIFT        = 39
	GrammarLexerBITNOT           = 40
	GrammarLexerFUNC             = 41
	GrammarLexerRETURN           = 42
	GrammarLexerPRINT            = 43
	GrammarLexerTRUE             = 44
	GrammarLexerFALSE            = 45
	GrammarLexerAND              = 46
	GrammarLexerOR               = 47
	GrammarLexerNOT              = 48
	GrammarLexerIDENTIFIER       = 49
	GrammarLexerNUMBER           = 50
	GrammarLexerSTRING           = 51
	GrammarLexerNL               = 52
	GrammarLexerWS               = 53
)
