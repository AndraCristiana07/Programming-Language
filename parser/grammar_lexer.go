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
		"'<<'", "'>>'", "'~'", "'&='", "'|='", "'^='", "'<<='", "'>>='", "'func'",
		"'return'", "'print'", "'true'", "'false'", "'and'", "'or'", "'not'",
	}
	staticData.SymbolicNames = []string{
		"", "VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "LBRACE", "LPAREN",
		"RPAREN", "RBRACE", "LBRACKET", "RBRACKET", "IF", "ELSE", "WHILE", "FOR",
		"SEMICOLON", "COMMA", "INC", "DEC", "MODULO", "EXPONENTIAL", "PLUSEQUAL",
		"MINUSEQUAL", "STAREQUAL", "SLASHEQUAL", "MODEQUAL", "EXPONENTIALEQUAL",
		"BITAND", "BITOR", "BITXOR", "BITLSHIFT", "BITRSHIFT", "BITNOT", "BITANDEQUAL",
		"BITOREQUAL", "BITXOREQAUL", "BITLSHIFTEQUAL", "BITRSHIFTEQUAL", "FUNC",
		"RETURN", "PRINT", "TRUE", "FALSE", "AND", "OR", "NOT", "IDENTIFIER",
		"NUMBER", "STRING", "NL", "WS",
	}
	staticData.RuleNames = []string{
		"VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS", "GREATER",
		"EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "LBRACE", "LPAREN",
		"RPAREN", "RBRACE", "LBRACKET", "RBRACKET", "IF", "ELSE", "WHILE", "FOR",
		"SEMICOLON", "COMMA", "INC", "DEC", "MODULO", "EXPONENTIAL", "PLUSEQUAL",
		"MINUSEQUAL", "STAREQUAL", "SLASHEQUAL", "MODEQUAL", "EXPONENTIALEQUAL",
		"BITAND", "BITOR", "BITXOR", "BITLSHIFT", "BITRSHIFT", "BITNOT", "BITANDEQUAL",
		"BITOREQUAL", "BITXOREQAUL", "BITLSHIFTEQUAL", "BITRSHIFTEQUAL", "FUNC",
		"RETURN", "PRINT", "TRUE", "FALSE", "AND", "OR", "NOT", "IDENTIFIER",
		"NUMBER", "STRING", "NL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 58, 325, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39,
		1, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1,
		43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45,
		1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1,
		51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 5, 53, 285, 8, 53,
		10, 53, 12, 53, 288, 9, 53, 1, 54, 4, 54, 291, 8, 54, 11, 54, 12, 54, 292,
		1, 54, 1, 54, 4, 54, 297, 8, 54, 11, 54, 12, 54, 298, 3, 54, 301, 8, 54,
		1, 55, 1, 55, 1, 55, 1, 55, 5, 55, 307, 8, 55, 10, 55, 12, 55, 310, 9,
		55, 1, 55, 1, 55, 1, 56, 4, 56, 315, 8, 56, 11, 56, 12, 56, 316, 1, 57,
		4, 57, 320, 8, 57, 11, 57, 12, 57, 321, 1, 57, 1, 57, 0, 0, 58, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30,
		61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39,
		79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48,
		97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 56, 113,
		57, 115, 58, 1, 0, 6, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 1, 0, 48, 57, 1, 0, 34, 34, 2, 0, 10, 10, 13, 13,
		2, 0, 9, 9, 32, 32, 332, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0,
		0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1,
		0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67,
		1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0,
		75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0,
		0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0,
		0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0,
		0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105,
		1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0,
		0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 1, 117, 1, 0, 0, 0, 3, 121, 1,
		0, 0, 0, 5, 123, 1, 0, 0, 0, 7, 125, 1, 0, 0, 0, 9, 127, 1, 0, 0, 0, 11,
		129, 1, 0, 0, 0, 13, 131, 1, 0, 0, 0, 15, 133, 1, 0, 0, 0, 17, 135, 1,
		0, 0, 0, 19, 138, 1, 0, 0, 0, 21, 141, 1, 0, 0, 0, 23, 144, 1, 0, 0, 0,
		25, 147, 1, 0, 0, 0, 27, 149, 1, 0, 0, 0, 29, 151, 1, 0, 0, 0, 31, 153,
		1, 0, 0, 0, 33, 155, 1, 0, 0, 0, 35, 157, 1, 0, 0, 0, 37, 159, 1, 0, 0,
		0, 39, 162, 1, 0, 0, 0, 41, 167, 1, 0, 0, 0, 43, 173, 1, 0, 0, 0, 45, 177,
		1, 0, 0, 0, 47, 179, 1, 0, 0, 0, 49, 181, 1, 0, 0, 0, 51, 184, 1, 0, 0,
		0, 53, 187, 1, 0, 0, 0, 55, 189, 1, 0, 0, 0, 57, 192, 1, 0, 0, 0, 59, 195,
		1, 0, 0, 0, 61, 198, 1, 0, 0, 0, 63, 201, 1, 0, 0, 0, 65, 204, 1, 0, 0,
		0, 67, 207, 1, 0, 0, 0, 69, 211, 1, 0, 0, 0, 71, 213, 1, 0, 0, 0, 73, 215,
		1, 0, 0, 0, 75, 217, 1, 0, 0, 0, 77, 220, 1, 0, 0, 0, 79, 223, 1, 0, 0,
		0, 81, 225, 1, 0, 0, 0, 83, 228, 1, 0, 0, 0, 85, 231, 1, 0, 0, 0, 87, 234,
		1, 0, 0, 0, 89, 238, 1, 0, 0, 0, 91, 242, 1, 0, 0, 0, 93, 247, 1, 0, 0,
		0, 95, 254, 1, 0, 0, 0, 97, 260, 1, 0, 0, 0, 99, 265, 1, 0, 0, 0, 101,
		271, 1, 0, 0, 0, 103, 275, 1, 0, 0, 0, 105, 278, 1, 0, 0, 0, 107, 282,
		1, 0, 0, 0, 109, 290, 1, 0, 0, 0, 111, 302, 1, 0, 0, 0, 113, 314, 1, 0,
		0, 0, 115, 319, 1, 0, 0, 0, 117, 118, 5, 118, 0, 0, 118, 119, 5, 97, 0,
		0, 119, 120, 5, 114, 0, 0, 120, 2, 1, 0, 0, 0, 121, 122, 5, 61, 0, 0, 122,
		4, 1, 0, 0, 0, 123, 124, 5, 43, 0, 0, 124, 6, 1, 0, 0, 0, 125, 126, 5,
		45, 0, 0, 126, 8, 1, 0, 0, 0, 127, 128, 5, 42, 0, 0, 128, 10, 1, 0, 0,
		0, 129, 130, 5, 47, 0, 0, 130, 12, 1, 0, 0, 0, 131, 132, 5, 60, 0, 0, 132,
		14, 1, 0, 0, 0, 133, 134, 5, 62, 0, 0, 134, 16, 1, 0, 0, 0, 135, 136, 5,
		61, 0, 0, 136, 137, 5, 61, 0, 0, 137, 18, 1, 0, 0, 0, 138, 139, 5, 33,
		0, 0, 139, 140, 5, 61, 0, 0, 140, 20, 1, 0, 0, 0, 141, 142, 5, 60, 0, 0,
		142, 143, 5, 61, 0, 0, 143, 22, 1, 0, 0, 0, 144, 145, 5, 62, 0, 0, 145,
		146, 5, 61, 0, 0, 146, 24, 1, 0, 0, 0, 147, 148, 5, 123, 0, 0, 148, 26,
		1, 0, 0, 0, 149, 150, 5, 40, 0, 0, 150, 28, 1, 0, 0, 0, 151, 152, 5, 41,
		0, 0, 152, 30, 1, 0, 0, 0, 153, 154, 5, 125, 0, 0, 154, 32, 1, 0, 0, 0,
		155, 156, 5, 91, 0, 0, 156, 34, 1, 0, 0, 0, 157, 158, 5, 93, 0, 0, 158,
		36, 1, 0, 0, 0, 159, 160, 5, 105, 0, 0, 160, 161, 5, 102, 0, 0, 161, 38,
		1, 0, 0, 0, 162, 163, 5, 101, 0, 0, 163, 164, 5, 108, 0, 0, 164, 165, 5,
		115, 0, 0, 165, 166, 5, 101, 0, 0, 166, 40, 1, 0, 0, 0, 167, 168, 5, 119,
		0, 0, 168, 169, 5, 104, 0, 0, 169, 170, 5, 105, 0, 0, 170, 171, 5, 108,
		0, 0, 171, 172, 5, 101, 0, 0, 172, 42, 1, 0, 0, 0, 173, 174, 5, 102, 0,
		0, 174, 175, 5, 111, 0, 0, 175, 176, 5, 114, 0, 0, 176, 44, 1, 0, 0, 0,
		177, 178, 5, 59, 0, 0, 178, 46, 1, 0, 0, 0, 179, 180, 5, 44, 0, 0, 180,
		48, 1, 0, 0, 0, 181, 182, 5, 43, 0, 0, 182, 183, 5, 43, 0, 0, 183, 50,
		1, 0, 0, 0, 184, 185, 5, 45, 0, 0, 185, 186, 5, 45, 0, 0, 186, 52, 1, 0,
		0, 0, 187, 188, 5, 37, 0, 0, 188, 54, 1, 0, 0, 0, 189, 190, 5, 42, 0, 0,
		190, 191, 5, 42, 0, 0, 191, 56, 1, 0, 0, 0, 192, 193, 5, 43, 0, 0, 193,
		194, 5, 61, 0, 0, 194, 58, 1, 0, 0, 0, 195, 196, 5, 45, 0, 0, 196, 197,
		5, 61, 0, 0, 197, 60, 1, 0, 0, 0, 198, 199, 5, 42, 0, 0, 199, 200, 5, 61,
		0, 0, 200, 62, 1, 0, 0, 0, 201, 202, 5, 47, 0, 0, 202, 203, 5, 61, 0, 0,
		203, 64, 1, 0, 0, 0, 204, 205, 5, 37, 0, 0, 205, 206, 5, 61, 0, 0, 206,
		66, 1, 0, 0, 0, 207, 208, 5, 42, 0, 0, 208, 209, 5, 42, 0, 0, 209, 210,
		5, 61, 0, 0, 210, 68, 1, 0, 0, 0, 211, 212, 5, 38, 0, 0, 212, 70, 1, 0,
		0, 0, 213, 214, 5, 124, 0, 0, 214, 72, 1, 0, 0, 0, 215, 216, 5, 94, 0,
		0, 216, 74, 1, 0, 0, 0, 217, 218, 5, 60, 0, 0, 218, 219, 5, 60, 0, 0, 219,
		76, 1, 0, 0, 0, 220, 221, 5, 62, 0, 0, 221, 222, 5, 62, 0, 0, 222, 78,
		1, 0, 0, 0, 223, 224, 5, 126, 0, 0, 224, 80, 1, 0, 0, 0, 225, 226, 5, 38,
		0, 0, 226, 227, 5, 61, 0, 0, 227, 82, 1, 0, 0, 0, 228, 229, 5, 124, 0,
		0, 229, 230, 5, 61, 0, 0, 230, 84, 1, 0, 0, 0, 231, 232, 5, 94, 0, 0, 232,
		233, 5, 61, 0, 0, 233, 86, 1, 0, 0, 0, 234, 235, 5, 60, 0, 0, 235, 236,
		5, 60, 0, 0, 236, 237, 5, 61, 0, 0, 237, 88, 1, 0, 0, 0, 238, 239, 5, 62,
		0, 0, 239, 240, 5, 62, 0, 0, 240, 241, 5, 61, 0, 0, 241, 90, 1, 0, 0, 0,
		242, 243, 5, 102, 0, 0, 243, 244, 5, 117, 0, 0, 244, 245, 5, 110, 0, 0,
		245, 246, 5, 99, 0, 0, 246, 92, 1, 0, 0, 0, 247, 248, 5, 114, 0, 0, 248,
		249, 5, 101, 0, 0, 249, 250, 5, 116, 0, 0, 250, 251, 5, 117, 0, 0, 251,
		252, 5, 114, 0, 0, 252, 253, 5, 110, 0, 0, 253, 94, 1, 0, 0, 0, 254, 255,
		5, 112, 0, 0, 255, 256, 5, 114, 0, 0, 256, 257, 5, 105, 0, 0, 257, 258,
		5, 110, 0, 0, 258, 259, 5, 116, 0, 0, 259, 96, 1, 0, 0, 0, 260, 261, 5,
		116, 0, 0, 261, 262, 5, 114, 0, 0, 262, 263, 5, 117, 0, 0, 263, 264, 5,
		101, 0, 0, 264, 98, 1, 0, 0, 0, 265, 266, 5, 102, 0, 0, 266, 267, 5, 97,
		0, 0, 267, 268, 5, 108, 0, 0, 268, 269, 5, 115, 0, 0, 269, 270, 5, 101,
		0, 0, 270, 100, 1, 0, 0, 0, 271, 272, 5, 97, 0, 0, 272, 273, 5, 110, 0,
		0, 273, 274, 5, 100, 0, 0, 274, 102, 1, 0, 0, 0, 275, 276, 5, 111, 0, 0,
		276, 277, 5, 114, 0, 0, 277, 104, 1, 0, 0, 0, 278, 279, 5, 110, 0, 0, 279,
		280, 5, 111, 0, 0, 280, 281, 5, 116, 0, 0, 281, 106, 1, 0, 0, 0, 282, 286,
		7, 0, 0, 0, 283, 285, 7, 1, 0, 0, 284, 283, 1, 0, 0, 0, 285, 288, 1, 0,
		0, 0, 286, 284, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 108, 1, 0, 0, 0,
		288, 286, 1, 0, 0, 0, 289, 291, 7, 2, 0, 0, 290, 289, 1, 0, 0, 0, 291,
		292, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 300,
		1, 0, 0, 0, 294, 296, 5, 46, 0, 0, 295, 297, 7, 2, 0, 0, 296, 295, 1, 0,
		0, 0, 297, 298, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0,
		299, 301, 1, 0, 0, 0, 300, 294, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301,
		110, 1, 0, 0, 0, 302, 308, 5, 34, 0, 0, 303, 307, 8, 3, 0, 0, 304, 305,
		5, 34, 0, 0, 305, 307, 5, 34, 0, 0, 306, 303, 1, 0, 0, 0, 306, 304, 1,
		0, 0, 0, 307, 310, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 308, 309, 1, 0, 0,
		0, 309, 311, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 311, 312, 5, 34, 0, 0, 312,
		112, 1, 0, 0, 0, 313, 315, 7, 4, 0, 0, 314, 313, 1, 0, 0, 0, 315, 316,
		1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 114, 1, 0,
		0, 0, 318, 320, 7, 5, 0, 0, 319, 318, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0,
		321, 319, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323,
		324, 6, 57, 0, 0, 324, 116, 1, 0, 0, 0, 9, 0, 286, 292, 298, 300, 306,
		308, 316, 321, 1, 6, 0, 0,
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
	GrammarLexerBITANDEQUAL      = 41
	GrammarLexerBITOREQUAL       = 42
	GrammarLexerBITXOREQAUL      = 43
	GrammarLexerBITLSHIFTEQUAL   = 44
	GrammarLexerBITRSHIFTEQUAL   = 45
	GrammarLexerFUNC             = 46
	GrammarLexerRETURN           = 47
	GrammarLexerPRINT            = 48
	GrammarLexerTRUE             = 49
	GrammarLexerFALSE            = 50
	GrammarLexerAND              = 51
	GrammarLexerOR               = 52
	GrammarLexerNOT              = 53
	GrammarLexerIDENTIFIER       = 54
	GrammarLexerNUMBER           = 55
	GrammarLexerSTRING           = 56
	GrammarLexerNL               = 57
	GrammarLexerWS               = 58
)
