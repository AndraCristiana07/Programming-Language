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
		"", "'null'", "'switch'", "'case'", "'default'", "'try'", "'catch'",
		"'throw'", "'in'", "'.'", "':'", "'var'", "'='", "'+'", "'-'", "'*'",
		"'/'", "'<'", "'>'", "'=='", "'!='", "'<='", "'>='", "'{'", "'('", "')'",
		"'}'", "'['", "']'", "'if'", "'else'", "'elif'", "'while'", "'for'",
		"'break'", "'continue'", "';'", "','", "'++'", "'--'", "'%'", "'**'",
		"'+='", "'-='", "'*='", "'/='", "'%='", "'**='", "'&'", "'|'", "'^'",
		"'<<'", "'>>'", "'~'", "'&='", "'|='", "'^='", "'<<='", "'>>='", "'func'",
		"'return'", "'print'", "'true'", "'false'", "'and'", "'or'", "'not'",
	}
	staticData.SymbolicNames = []string{
		"", "", "SWITCH", "CASE", "DEFAULT", "TRY", "CATCH", "THROW", "IN",
		"DOT", "COLON", "VAR", "EQUALS", "PLUS", "MINUS", "STAR", "SLASH", "LESS",
		"GREATER", "EQUALEQUAL", "BANGEQUAL", "LESSEQUAL", "GREATEREQUAL", "LBRACE",
		"LPAREN", "RPAREN", "RBRACE", "LBRACKET", "RBRACKET", "IF", "ELSE",
		"ELIF", "WHILE", "FOR", "BREAK", "CONTINUE", "SEMICOLON", "COMMA", "INC",
		"DEC", "MODULO", "EXPONENTIAL", "PLUSEQUAL", "MINUSEQUAL", "STAREQUAL",
		"SLASHEQUAL", "MODEQUAL", "EXPONENTIALEQUAL", "BITAND", "BITOR", "BITXOR",
		"BITLSHIFT", "BITRSHIFT", "BITNOT", "BITANDEQUAL", "BITOREQUAL", "BITXOREQAUL",
		"BITLSHIFTEQUAL", "BITRSHIFTEQUAL", "FUNC", "RETURN", "PRINT", "TRUE",
		"FALSE", "AND", "OR", "NOT", "IDENTIFIER", "NUMBER", "STRING", "LINE_COMMENT",
		"BLOCK_COMMENT", "NL", "WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "returnStmt", "terminator", "funcStmt", "exprStmt",
		"tryCatchStmt", "throwStmt", "breakStmt", "continueStmt", "switchStmt",
		"caseBlock", "defaultBlock", "forInStmt", "varDecl", "assignStmt", "arrayAssignStmt",
		"compoundAssignStmt", "printStmt", "ifStmt", "whileStmt", "forStmt",
		"blockStmt", "ifInit", "forInit", "forPost", "postfixStmt", "expr",
		"mapEntry", "membershipOp", "arrayLit",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 73, 411, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 1,
		0, 1, 0, 5, 0, 66, 8, 0, 10, 0, 12, 0, 69, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 91, 8, 1, 1, 1, 3, 1, 94, 8, 1, 1, 2, 1,
		2, 3, 2, 98, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4,
		108, 8, 4, 10, 4, 12, 4, 111, 9, 4, 3, 4, 113, 8, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		5, 10, 141, 8, 10, 10, 10, 12, 10, 144, 9, 10, 1, 10, 5, 10, 147, 8, 10,
		10, 10, 12, 10, 150, 9, 10, 1, 10, 3, 10, 153, 8, 10, 1, 10, 5, 10, 156,
		8, 10, 10, 10, 12, 10, 159, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 5, 11, 168, 8, 11, 10, 11, 12, 11, 171, 9, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 5, 12, 177, 8, 12, 10, 12, 12, 12, 180, 9, 12, 1, 13, 1,
		13, 1, 13, 3, 13, 185, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 221, 8, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 232,
		8, 19, 10, 19, 12, 19, 235, 9, 19, 1, 19, 1, 19, 3, 19, 239, 8, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 261,
		8, 22, 10, 22, 12, 22, 264, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 3, 23, 270,
		8, 23, 1, 24, 1, 24, 3, 24, 274, 8, 24, 1, 25, 1, 25, 3, 25, 278, 8, 25,
		1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 289,
		8, 27, 10, 27, 12, 27, 292, 9, 27, 3, 27, 294, 8, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 304, 8, 27, 10, 27, 12, 27,
		307, 9, 27, 3, 27, 309, 8, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 321, 8, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 5, 27, 371, 8, 27, 10, 27, 12, 27, 374, 9, 27, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 3, 29, 383, 8, 29, 1, 30, 1,
		30, 1, 30, 1, 30, 5, 30, 389, 8, 30, 10, 30, 12, 30, 392, 9, 30, 3, 30,
		394, 8, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 3, 30, 405, 8, 30, 1, 30, 1, 30, 3, 30, 409, 8, 30, 1, 30, 0, 1, 54,
		31, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 0, 9, 2, 0, 36, 36,
		72, 72, 2, 0, 42, 47, 54, 58, 1, 0, 38, 39, 3, 0, 14, 14, 53, 53, 66, 66,
		1, 0, 62, 63, 2, 0, 15, 16, 40, 40, 1, 0, 13, 14, 1, 0, 51, 52, 1, 0, 17,
		22, 453, 0, 67, 1, 0, 0, 0, 2, 90, 1, 0, 0, 0, 4, 95, 1, 0, 0, 0, 6, 99,
		1, 0, 0, 0, 8, 101, 1, 0, 0, 0, 10, 117, 1, 0, 0, 0, 12, 119, 1, 0, 0,
		0, 14, 127, 1, 0, 0, 0, 16, 130, 1, 0, 0, 0, 18, 132, 1, 0, 0, 0, 20, 134,
		1, 0, 0, 0, 22, 162, 1, 0, 0, 0, 24, 172, 1, 0, 0, 0, 26, 181, 1, 0, 0,
		0, 28, 192, 1, 0, 0, 0, 30, 197, 1, 0, 0, 0, 32, 201, 1, 0, 0, 0, 34, 208,
		1, 0, 0, 0, 36, 212, 1, 0, 0, 0, 38, 215, 1, 0, 0, 0, 40, 240, 1, 0, 0,
		0, 42, 246, 1, 0, 0, 0, 44, 256, 1, 0, 0, 0, 46, 269, 1, 0, 0, 0, 48, 273,
		1, 0, 0, 0, 50, 277, 1, 0, 0, 0, 52, 279, 1, 0, 0, 0, 54, 320, 1, 0, 0,
		0, 56, 375, 1, 0, 0, 0, 58, 382, 1, 0, 0, 0, 60, 408, 1, 0, 0, 0, 62, 66,
		3, 2, 1, 0, 63, 66, 3, 8, 4, 0, 64, 66, 3, 6, 3, 0, 65, 62, 1, 0, 0, 0,
		65, 63, 1, 0, 0, 0, 65, 64, 1, 0, 0, 0, 66, 69, 1, 0, 0, 0, 67, 65, 1,
		0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 70, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 70,
		71, 5, 0, 0, 1, 71, 1, 1, 0, 0, 0, 72, 91, 3, 28, 14, 0, 73, 91, 3, 30,
		15, 0, 74, 91, 3, 32, 16, 0, 75, 91, 3, 34, 17, 0, 76, 91, 3, 10, 5, 0,
		77, 91, 3, 36, 18, 0, 78, 91, 3, 38, 19, 0, 79, 91, 3, 40, 20, 0, 80, 91,
		3, 20, 10, 0, 81, 91, 3, 26, 13, 0, 82, 91, 3, 42, 21, 0, 83, 91, 3, 44,
		22, 0, 84, 91, 3, 52, 26, 0, 85, 91, 3, 12, 6, 0, 86, 91, 3, 14, 7, 0,
		87, 91, 3, 4, 2, 0, 88, 91, 3, 16, 8, 0, 89, 91, 3, 18, 9, 0, 90, 72, 1,
		0, 0, 0, 90, 73, 1, 0, 0, 0, 90, 74, 1, 0, 0, 0, 90, 75, 1, 0, 0, 0, 90,
		76, 1, 0, 0, 0, 90, 77, 1, 0, 0, 0, 90, 78, 1, 0, 0, 0, 90, 79, 1, 0, 0,
		0, 90, 80, 1, 0, 0, 0, 90, 81, 1, 0, 0, 0, 90, 82, 1, 0, 0, 0, 90, 83,
		1, 0, 0, 0, 90, 84, 1, 0, 0, 0, 90, 85, 1, 0, 0, 0, 90, 86, 1, 0, 0, 0,
		90, 87, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 93, 1,
		0, 0, 0, 92, 94, 3, 6, 3, 0, 93, 92, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94,
		3, 1, 0, 0, 0, 95, 97, 5, 60, 0, 0, 96, 98, 3, 54, 27, 0, 97, 96, 1, 0,
		0, 0, 97, 98, 1, 0, 0, 0, 98, 5, 1, 0, 0, 0, 99, 100, 7, 0, 0, 0, 100,
		7, 1, 0, 0, 0, 101, 102, 5, 59, 0, 0, 102, 103, 5, 67, 0, 0, 103, 112,
		5, 24, 0, 0, 104, 109, 5, 67, 0, 0, 105, 106, 5, 37, 0, 0, 106, 108, 5,
		67, 0, 0, 107, 105, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0,
		0, 109, 110, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112,
		104, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115,
		5, 25, 0, 0, 115, 116, 3, 44, 22, 0, 116, 9, 1, 0, 0, 0, 117, 118, 3, 54,
		27, 0, 118, 11, 1, 0, 0, 0, 119, 120, 5, 5, 0, 0, 120, 121, 3, 44, 22,
		0, 121, 122, 5, 6, 0, 0, 122, 123, 5, 24, 0, 0, 123, 124, 5, 67, 0, 0,
		124, 125, 5, 25, 0, 0, 125, 126, 3, 44, 22, 0, 126, 13, 1, 0, 0, 0, 127,
		128, 5, 7, 0, 0, 128, 129, 3, 54, 27, 0, 129, 15, 1, 0, 0, 0, 130, 131,
		5, 34, 0, 0, 131, 17, 1, 0, 0, 0, 132, 133, 5, 35, 0, 0, 133, 19, 1, 0,
		0, 0, 134, 135, 5, 2, 0, 0, 135, 136, 5, 24, 0, 0, 136, 137, 3, 54, 27,
		0, 137, 138, 5, 25, 0, 0, 138, 142, 5, 23, 0, 0, 139, 141, 3, 6, 3, 0,
		140, 139, 1, 0, 0, 0, 141, 144, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142,
		143, 1, 0, 0, 0, 143, 148, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 145, 147,
		3, 22, 11, 0, 146, 145, 1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1,
		0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0,
		0, 151, 153, 3, 24, 12, 0, 152, 151, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0,
		153, 157, 1, 0, 0, 0, 154, 156, 3, 6, 3, 0, 155, 154, 1, 0, 0, 0, 156,
		159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 160,
		1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 161, 5, 26, 0, 0, 161, 21, 1, 0,
		0, 0, 162, 163, 5, 3, 0, 0, 163, 164, 3, 54, 27, 0, 164, 169, 5, 10, 0,
		0, 165, 168, 3, 2, 1, 0, 166, 168, 3, 6, 3, 0, 167, 165, 1, 0, 0, 0, 167,
		166, 1, 0, 0, 0, 168, 171, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170,
		1, 0, 0, 0, 170, 23, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 172, 173, 5, 4,
		0, 0, 173, 178, 5, 10, 0, 0, 174, 177, 3, 2, 1, 0, 175, 177, 3, 6, 3, 0,
		176, 174, 1, 0, 0, 0, 176, 175, 1, 0, 0, 0, 177, 180, 1, 0, 0, 0, 178,
		176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 25, 1, 0, 0, 0, 180, 178, 1,
		0, 0, 0, 181, 182, 5, 33, 0, 0, 182, 184, 5, 24, 0, 0, 183, 185, 5, 11,
		0, 0, 184, 183, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0,
		186, 187, 5, 67, 0, 0, 187, 188, 5, 8, 0, 0, 188, 189, 3, 54, 27, 0, 189,
		190, 5, 25, 0, 0, 190, 191, 3, 44, 22, 0, 191, 27, 1, 0, 0, 0, 192, 193,
		5, 11, 0, 0, 193, 194, 5, 67, 0, 0, 194, 195, 5, 12, 0, 0, 195, 196, 3,
		54, 27, 0, 196, 29, 1, 0, 0, 0, 197, 198, 3, 54, 27, 0, 198, 199, 5, 12,
		0, 0, 199, 200, 3, 54, 27, 0, 200, 31, 1, 0, 0, 0, 201, 202, 5, 67, 0,
		0, 202, 203, 5, 27, 0, 0, 203, 204, 3, 54, 27, 0, 204, 205, 5, 28, 0, 0,
		205, 206, 5, 12, 0, 0, 206, 207, 3, 54, 27, 0, 207, 33, 1, 0, 0, 0, 208,
		209, 3, 54, 27, 0, 209, 210, 7, 1, 0, 0, 210, 211, 3, 54, 27, 0, 211, 35,
		1, 0, 0, 0, 212, 213, 5, 61, 0, 0, 213, 214, 3, 54, 27, 0, 214, 37, 1,
		0, 0, 0, 215, 216, 5, 29, 0, 0, 216, 220, 5, 24, 0, 0, 217, 218, 3, 46,
		23, 0, 218, 219, 5, 36, 0, 0, 219, 221, 1, 0, 0, 0, 220, 217, 1, 0, 0,
		0, 220, 221, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223, 3, 54, 27, 0,
		223, 224, 5, 25, 0, 0, 224, 233, 3, 44, 22, 0, 225, 226, 5, 31, 0, 0, 226,
		227, 5, 24, 0, 0, 227, 228, 3, 54, 27, 0, 228, 229, 5, 25, 0, 0, 229, 230,
		3, 44, 22, 0, 230, 232, 1, 0, 0, 0, 231, 225, 1, 0, 0, 0, 232, 235, 1,
		0, 0, 0, 233, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 238, 1, 0, 0,
		0, 235, 233, 1, 0, 0, 0, 236, 237, 5, 30, 0, 0, 237, 239, 3, 44, 22, 0,
		238, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 39, 1, 0, 0, 0, 240, 241,
		5, 32, 0, 0, 241, 242, 5, 24, 0, 0, 242, 243, 3, 54, 27, 0, 243, 244, 5,
		25, 0, 0, 244, 245, 3, 44, 22, 0, 245, 41, 1, 0, 0, 0, 246, 247, 5, 33,
		0, 0, 247, 248, 5, 24, 0, 0, 248, 249, 3, 48, 24, 0, 249, 250, 5, 36, 0,
		0, 250, 251, 3, 54, 27, 0, 251, 252, 5, 36, 0, 0, 252, 253, 3, 50, 25,
		0, 253, 254, 5, 25, 0, 0, 254, 255, 3, 44, 22, 0, 255, 43, 1, 0, 0, 0,
		256, 262, 5, 23, 0, 0, 257, 261, 3, 2, 1, 0, 258, 261, 3, 8, 4, 0, 259,
		261, 3, 6, 3, 0, 260, 257, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 260, 259,
		1, 0, 0, 0, 261, 264, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0,
		0, 0, 263, 265, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 265, 266, 5, 26, 0, 0,
		266, 45, 1, 0, 0, 0, 267, 270, 3, 28, 14, 0, 268, 270, 3, 30, 15, 0, 269,
		267, 1, 0, 0, 0, 269, 268, 1, 0, 0, 0, 270, 47, 1, 0, 0, 0, 271, 274, 3,
		28, 14, 0, 272, 274, 3, 30, 15, 0, 273, 271, 1, 0, 0, 0, 273, 272, 1, 0,
		0, 0, 274, 49, 1, 0, 0, 0, 275, 278, 3, 30, 15, 0, 276, 278, 3, 52, 26,
		0, 277, 275, 1, 0, 0, 0, 277, 276, 1, 0, 0, 0, 278, 51, 1, 0, 0, 0, 279,
		280, 3, 54, 27, 0, 280, 281, 7, 2, 0, 0, 281, 53, 1, 0, 0, 0, 282, 283,
		6, 27, -1, 0, 283, 284, 5, 67, 0, 0, 284, 293, 5, 24, 0, 0, 285, 290, 3,
		54, 27, 0, 286, 287, 5, 37, 0, 0, 287, 289, 3, 54, 27, 0, 288, 286, 1,
		0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0,
		0, 291, 294, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293, 285, 1, 0, 0, 0, 293,
		294, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 321, 5, 25, 0, 0, 296, 297,
		7, 3, 0, 0, 297, 321, 3, 54, 27, 21, 298, 321, 3, 60, 30, 0, 299, 308,
		5, 23, 0, 0, 300, 305, 3, 56, 28, 0, 301, 302, 5, 37, 0, 0, 302, 304, 3,
		56, 28, 0, 303, 301, 1, 0, 0, 0, 304, 307, 1, 0, 0, 0, 305, 303, 1, 0,
		0, 0, 305, 306, 1, 0, 0, 0, 306, 309, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0,
		308, 300, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310,
		321, 5, 26, 0, 0, 311, 321, 5, 67, 0, 0, 312, 321, 5, 68, 0, 0, 313, 321,
		7, 4, 0, 0, 314, 321, 5, 69, 0, 0, 315, 321, 5, 1, 0, 0, 316, 317, 5, 24,
		0, 0, 317, 318, 3, 54, 27, 0, 318, 319, 5, 25, 0, 0, 319, 321, 1, 0, 0,
		0, 320, 282, 1, 0, 0, 0, 320, 296, 1, 0, 0, 0, 320, 298, 1, 0, 0, 0, 320,
		299, 1, 0, 0, 0, 320, 311, 1, 0, 0, 0, 320, 312, 1, 0, 0, 0, 320, 313,
		1, 0, 0, 0, 320, 314, 1, 0, 0, 0, 320, 315, 1, 0, 0, 0, 320, 316, 1, 0,
		0, 0, 321, 372, 1, 0, 0, 0, 322, 323, 10, 20, 0, 0, 323, 324, 5, 41, 0,
		0, 324, 371, 3, 54, 27, 21, 325, 326, 10, 19, 0, 0, 326, 327, 7, 5, 0,
		0, 327, 371, 3, 54, 27, 20, 328, 329, 10, 18, 0, 0, 329, 330, 7, 6, 0,
		0, 330, 371, 3, 54, 27, 19, 331, 332, 10, 17, 0, 0, 332, 333, 7, 7, 0,
		0, 333, 371, 3, 54, 27, 18, 334, 335, 10, 16, 0, 0, 335, 336, 5, 48, 0,
		0, 336, 371, 3, 54, 27, 17, 337, 338, 10, 15, 0, 0, 338, 339, 5, 50, 0,
		0, 339, 371, 3, 54, 27, 16, 340, 341, 10, 14, 0, 0, 341, 342, 5, 49, 0,
		0, 342, 371, 3, 54, 27, 15, 343, 344, 10, 13, 0, 0, 344, 345, 7, 8, 0,
		0, 345, 371, 3, 54, 27, 14, 346, 347, 10, 12, 0, 0, 347, 348, 3, 58, 29,
		0, 348, 349, 3, 54, 27, 13, 349, 371, 1, 0, 0, 0, 350, 351, 10, 11, 0,
		0, 351, 352, 5, 64, 0, 0, 352, 371, 3, 54, 27, 12, 353, 354, 10, 10, 0,
		0, 354, 355, 5, 65, 0, 0, 355, 371, 3, 54, 27, 11, 356, 357, 10, 9, 0,
		0, 357, 358, 5, 29, 0, 0, 358, 359, 3, 54, 27, 0, 359, 360, 5, 30, 0, 0,
		360, 361, 3, 54, 27, 10, 361, 371, 1, 0, 0, 0, 362, 363, 10, 24, 0, 0,
		363, 364, 5, 27, 0, 0, 364, 365, 3, 54, 27, 0, 365, 366, 5, 28, 0, 0, 366,
		371, 1, 0, 0, 0, 367, 368, 10, 23, 0, 0, 368, 369, 5, 9, 0, 0, 369, 371,
		5, 67, 0, 0, 370, 322, 1, 0, 0, 0, 370, 325, 1, 0, 0, 0, 370, 328, 1, 0,
		0, 0, 370, 331, 1, 0, 0, 0, 370, 334, 1, 0, 0, 0, 370, 337, 1, 0, 0, 0,
		370, 340, 1, 0, 0, 0, 370, 343, 1, 0, 0, 0, 370, 346, 1, 0, 0, 0, 370,
		350, 1, 0, 0, 0, 370, 353, 1, 0, 0, 0, 370, 356, 1, 0, 0, 0, 370, 362,
		1, 0, 0, 0, 370, 367, 1, 0, 0, 0, 371, 374, 1, 0, 0, 0, 372, 370, 1, 0,
		0, 0, 372, 373, 1, 0, 0, 0, 373, 55, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0,
		375, 376, 3, 54, 27, 0, 376, 377, 5, 10, 0, 0, 377, 378, 3, 54, 27, 0,
		378, 57, 1, 0, 0, 0, 379, 383, 5, 8, 0, 0, 380, 381, 5, 66, 0, 0, 381,
		383, 5, 8, 0, 0, 382, 379, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 383, 59, 1,
		0, 0, 0, 384, 393, 5, 27, 0, 0, 385, 390, 3, 54, 27, 0, 386, 387, 5, 37,
		0, 0, 387, 389, 3, 54, 27, 0, 388, 386, 1, 0, 0, 0, 389, 392, 1, 0, 0,
		0, 390, 388, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391, 394, 1, 0, 0, 0, 392,
		390, 1, 0, 0, 0, 393, 385, 1, 0, 0, 0, 393, 394, 1, 0, 0, 0, 394, 395,
		1, 0, 0, 0, 395, 409, 5, 28, 0, 0, 396, 397, 5, 27, 0, 0, 397, 398, 3,
		54, 27, 0, 398, 399, 5, 33, 0, 0, 399, 400, 5, 67, 0, 0, 400, 401, 5, 8,
		0, 0, 401, 404, 3, 54, 27, 0, 402, 403, 5, 29, 0, 0, 403, 405, 3, 54, 27,
		0, 404, 402, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406,
		407, 5, 28, 0, 0, 407, 409, 1, 0, 0, 0, 408, 384, 1, 0, 0, 0, 408, 396,
		1, 0, 0, 0, 409, 61, 1, 0, 0, 0, 36, 65, 67, 90, 93, 97, 109, 112, 142,
		148, 152, 157, 167, 169, 176, 178, 184, 220, 233, 238, 260, 262, 269, 273,
		277, 290, 293, 305, 308, 320, 370, 372, 382, 390, 393, 404, 408,
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
	GrammarParserT__0             = 1
	GrammarParserSWITCH           = 2
	GrammarParserCASE             = 3
	GrammarParserDEFAULT          = 4
	GrammarParserTRY              = 5
	GrammarParserCATCH            = 6
	GrammarParserTHROW            = 7
	GrammarParserIN               = 8
	GrammarParserDOT              = 9
	GrammarParserCOLON            = 10
	GrammarParserVAR              = 11
	GrammarParserEQUALS           = 12
	GrammarParserPLUS             = 13
	GrammarParserMINUS            = 14
	GrammarParserSTAR             = 15
	GrammarParserSLASH            = 16
	GrammarParserLESS             = 17
	GrammarParserGREATER          = 18
	GrammarParserEQUALEQUAL       = 19
	GrammarParserBANGEQUAL        = 20
	GrammarParserLESSEQUAL        = 21
	GrammarParserGREATEREQUAL     = 22
	GrammarParserLBRACE           = 23
	GrammarParserLPAREN           = 24
	GrammarParserRPAREN           = 25
	GrammarParserRBRACE           = 26
	GrammarParserLBRACKET         = 27
	GrammarParserRBRACKET         = 28
	GrammarParserIF               = 29
	GrammarParserELSE             = 30
	GrammarParserELIF             = 31
	GrammarParserWHILE            = 32
	GrammarParserFOR              = 33
	GrammarParserBREAK            = 34
	GrammarParserCONTINUE         = 35
	GrammarParserSEMICOLON        = 36
	GrammarParserCOMMA            = 37
	GrammarParserINC              = 38
	GrammarParserDEC              = 39
	GrammarParserMODULO           = 40
	GrammarParserEXPONENTIAL      = 41
	GrammarParserPLUSEQUAL        = 42
	GrammarParserMINUSEQUAL       = 43
	GrammarParserSTAREQUAL        = 44
	GrammarParserSLASHEQUAL       = 45
	GrammarParserMODEQUAL         = 46
	GrammarParserEXPONENTIALEQUAL = 47
	GrammarParserBITAND           = 48
	GrammarParserBITOR            = 49
	GrammarParserBITXOR           = 50
	GrammarParserBITLSHIFT        = 51
	GrammarParserBITRSHIFT        = 52
	GrammarParserBITNOT           = 53
	GrammarParserBITANDEQUAL      = 54
	GrammarParserBITOREQUAL       = 55
	GrammarParserBITXOREQAUL      = 56
	GrammarParserBITLSHIFTEQUAL   = 57
	GrammarParserBITRSHIFTEQUAL   = 58
	GrammarParserFUNC             = 59
	GrammarParserRETURN           = 60
	GrammarParserPRINT            = 61
	GrammarParserTRUE             = 62
	GrammarParserFALSE            = 63
	GrammarParserAND              = 64
	GrammarParserOR               = 65
	GrammarParserNOT              = 66
	GrammarParserIDENTIFIER       = 67
	GrammarParserNUMBER           = 68
	GrammarParserSTRING           = 69
	GrammarParserLINE_COMMENT     = 70
	GrammarParserBLOCK_COMMENT    = 71
	GrammarParserNL               = 72
	GrammarParserWS               = 73
)

// GrammarParser rules.
const (
	GrammarParserRULE_program            = 0
	GrammarParserRULE_statement          = 1
	GrammarParserRULE_returnStmt         = 2
	GrammarParserRULE_terminator         = 3
	GrammarParserRULE_funcStmt           = 4
	GrammarParserRULE_exprStmt           = 5
	GrammarParserRULE_tryCatchStmt       = 6
	GrammarParserRULE_throwStmt          = 7
	GrammarParserRULE_breakStmt          = 8
	GrammarParserRULE_continueStmt       = 9
	GrammarParserRULE_switchStmt         = 10
	GrammarParserRULE_caseBlock          = 11
	GrammarParserRULE_defaultBlock       = 12
	GrammarParserRULE_forInStmt          = 13
	GrammarParserRULE_varDecl            = 14
	GrammarParserRULE_assignStmt         = 15
	GrammarParserRULE_arrayAssignStmt    = 16
	GrammarParserRULE_compoundAssignStmt = 17
	GrammarParserRULE_printStmt          = 18
	GrammarParserRULE_ifStmt             = 19
	GrammarParserRULE_whileStmt          = 20
	GrammarParserRULE_forStmt            = 21
	GrammarParserRULE_blockStmt          = 22
	GrammarParserRULE_ifInit             = 23
	GrammarParserRULE_forInit            = 24
	GrammarParserRULE_forPost            = 25
	GrammarParserRULE_postfixStmt        = 26
	GrammarParserRULE_expr               = 27
	GrammarParserRULE_mapEntry           = 28
	GrammarParserRULE_membershipOp       = 29
	GrammarParserRULE_arrayLit           = 30
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
	AllFuncStmt() []IFuncStmtContext
	FuncStmt(i int) IFuncStmtContext
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext

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

func (s *ProgramContext) AllFuncStmt() []IFuncStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncStmtContext); ok {
			len++
		}
	}

	tst := make([]IFuncStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncStmtContext); ok {
			tst[i] = t.(IFuncStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) FuncStmt(i int) IFuncStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncStmtContext); ok {
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

	return t.(IFuncStmtContext)
}

func (s *ProgramContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
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

	return t.(ITerminatorContext)
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
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-567453419208423258) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&79) != 0) {
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case GrammarParserT__0, GrammarParserSWITCH, GrammarParserTRY, GrammarParserTHROW, GrammarParserVAR, GrammarParserMINUS, GrammarParserLBRACE, GrammarParserLPAREN, GrammarParserLBRACKET, GrammarParserIF, GrammarParserWHILE, GrammarParserFOR, GrammarParserBREAK, GrammarParserCONTINUE, GrammarParserBITNOT, GrammarParserRETURN, GrammarParserPRINT, GrammarParserTRUE, GrammarParserFALSE, GrammarParserNOT, GrammarParserIDENTIFIER, GrammarParserNUMBER, GrammarParserSTRING:
			{
				p.SetState(62)
				p.Statement()
			}

		case GrammarParserFUNC:
			{
				p.SetState(63)
				p.FuncStmt()
			}

		case GrammarParserSEMICOLON, GrammarParserNL:
			{
				p.SetState(64)
				p.Terminator()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
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
	ArrayAssignStmt() IArrayAssignStmtContext
	CompoundAssignStmt() ICompoundAssignStmtContext
	ExprStmt() IExprStmtContext
	PrintStmt() IPrintStmtContext
	IfStmt() IIfStmtContext
	WhileStmt() IWhileStmtContext
	SwitchStmt() ISwitchStmtContext
	ForInStmt() IForInStmtContext
	ForStmt() IForStmtContext
	BlockStmt() IBlockStmtContext
	PostfixStmt() IPostfixStmtContext
	TryCatchStmt() ITryCatchStmtContext
	ThrowStmt() IThrowStmtContext
	ReturnStmt() IReturnStmtContext
	BreakStmt() IBreakStmtContext
	ContinueStmt() IContinueStmtContext
	Terminator() ITerminatorContext

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

func (s *StatementContext) ArrayAssignStmt() IArrayAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayAssignStmtContext)
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

func (s *StatementContext) ExprStmt() IExprStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprStmtContext)
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

func (s *StatementContext) SwitchStmt() ISwitchStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStmtContext)
}

func (s *StatementContext) ForInStmt() IForInStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForInStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForInStmtContext)
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

func (s *StatementContext) TryCatchStmt() ITryCatchStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITryCatchStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITryCatchStmtContext)
}

func (s *StatementContext) ThrowStmt() IThrowStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThrowStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThrowStmtContext)
}

func (s *StatementContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StatementContext) BreakStmt() IBreakStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStmtContext)
}

func (s *StatementContext) ContinueStmt() IContinueStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStmtContext)
}

func (s *StatementContext) Terminator() ITerminatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
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
	p.EnterOuterAlt(localctx, 1)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(72)
			p.VarDecl()
		}

	case 2:
		{
			p.SetState(73)
			p.AssignStmt()
		}

	case 3:
		{
			p.SetState(74)
			p.ArrayAssignStmt()
		}

	case 4:
		{
			p.SetState(75)
			p.CompoundAssignStmt()
		}

	case 5:
		{
			p.SetState(76)
			p.ExprStmt()
		}

	case 6:
		{
			p.SetState(77)
			p.PrintStmt()
		}

	case 7:
		{
			p.SetState(78)
			p.IfStmt()
		}

	case 8:
		{
			p.SetState(79)
			p.WhileStmt()
		}

	case 9:
		{
			p.SetState(80)
			p.SwitchStmt()
		}

	case 10:
		{
			p.SetState(81)
			p.ForInStmt()
		}

	case 11:
		{
			p.SetState(82)
			p.ForStmt()
		}

	case 12:
		{
			p.SetState(83)
			p.BlockStmt()
		}

	case 13:
		{
			p.SetState(84)
			p.PostfixStmt()
		}

	case 14:
		{
			p.SetState(85)
			p.TryCatchStmt()
		}

	case 15:
		{
			p.SetState(86)
			p.ThrowStmt()
		}

	case 16:
		{
			p.SetState(87)
			p.ReturnStmt()
		}

	case 17:
		{
			p.SetState(88)
			p.BreakStmt()
		}

	case 18:
		{
			p.SetState(89)
			p.ContinueStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(92)
			p.Terminator()
		}

	} else if p.HasError() { // JIM
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

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRETURN, 0)
}

func (s *ReturnStmtContext) Expr() IExprContext {
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

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrammarParserRULE_returnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(GrammarParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(96)
			p.expr(0)
		}

	} else if p.HasError() { // JIM
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

// ITerminatorContext is an interface to support dynamic dispatch.
type ITerminatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEMICOLON() antlr.TerminalNode
	NL() antlr.TerminalNode

	// IsTerminatorContext differentiates from other interfaces.
	IsTerminatorContext()
}

type TerminatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerminatorContext() *TerminatorContext {
	var p = new(TerminatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_terminator
	return p
}

func InitEmptyTerminatorContext(p *TerminatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_terminator
}

func (*TerminatorContext) IsTerminatorContext() {}

func NewTerminatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TerminatorContext {
	var p = new(TerminatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_terminator

	return p
}

func (s *TerminatorContext) GetParser() antlr.Parser { return s.parser }

func (s *TerminatorContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserSEMICOLON, 0)
}

func (s *TerminatorContext) NL() antlr.TerminalNode {
	return s.GetToken(GrammarParserNL, 0)
}

func (s *TerminatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TerminatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TerminatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTerminator(s)
	}
}

func (s *TerminatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTerminator(s)
	}
}

func (s *TerminatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitTerminator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Terminator() (localctx ITerminatorContext) {
	localctx = NewTerminatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrammarParserRULE_terminator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GrammarParserSEMICOLON || _la == GrammarParserNL) {
			p.GetErrorHandler().RecoverInline(p)
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

// IFuncStmtContext is an interface to support dynamic dispatch.
type IFuncStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	BlockStmt() IBlockStmtContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncStmtContext differentiates from other interfaces.
	IsFuncStmtContext()
}

type FuncStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncStmtContext() *FuncStmtContext {
	var p = new(FuncStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_funcStmt
	return p
}

func InitEmptyFuncStmtContext(p *FuncStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_funcStmt
}

func (*FuncStmtContext) IsFuncStmtContext() {}

func NewFuncStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncStmtContext {
	var p = new(FuncStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_funcStmt

	return p
}

func (s *FuncStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncStmtContext) FUNC() antlr.TerminalNode {
	return s.GetToken(GrammarParserFUNC, 0)
}

func (s *FuncStmtContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserIDENTIFIER)
}

func (s *FuncStmtContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, i)
}

func (s *FuncStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *FuncStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *FuncStmtContext) BlockStmt() IBlockStmtContext {
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

func (s *FuncStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOMMA)
}

func (s *FuncStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMMA, i)
}

func (s *FuncStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterFuncStmt(s)
	}
}

func (s *FuncStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitFuncStmt(s)
	}
}

func (s *FuncStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitFuncStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) FuncStmt() (localctx IFuncStmtContext) {
	localctx = NewFuncStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GrammarParserRULE_funcStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(GrammarParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(GrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserIDENTIFIER {
		{
			p.SetState(104)
			p.Match(GrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GrammarParserCOMMA {
			{
				p.SetState(105)
				p.Match(GrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(106)
				p.Match(GrammarParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(114)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.BlockStmt()
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

// IExprStmtContext is an interface to support dynamic dispatch.
type IExprStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsExprStmtContext differentiates from other interfaces.
	IsExprStmtContext()
}

type ExprStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprStmtContext() *ExprStmtContext {
	var p = new(ExprStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprStmt
	return p
}

func InitEmptyExprStmtContext(p *ExprStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprStmt
}

func (*ExprStmtContext) IsExprStmtContext() {}

func NewExprStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprStmtContext {
	var p = new(ExprStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_exprStmt

	return p
}

func (s *ExprStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprStmtContext) Expr() IExprContext {
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

func (s *ExprStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExprStmt(s)
	}
}

func (s *ExprStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExprStmt(s)
	}
}

func (s *ExprStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitExprStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ExprStmt() (localctx IExprStmtContext) {
	localctx = NewExprStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GrammarParserRULE_exprStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
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

// ITryCatchStmtContext is an interface to support dynamic dispatch.
type ITryCatchStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTryBody returns the tryBody rule contexts.
	GetTryBody() IBlockStmtContext

	// GetCatchBody returns the catchBody rule contexts.
	GetCatchBody() IBlockStmtContext

	// SetTryBody sets the tryBody rule contexts.
	SetTryBody(IBlockStmtContext)

	// SetCatchBody sets the catchBody rule contexts.
	SetCatchBody(IBlockStmtContext)

	// Getter signatures
	TRY() antlr.TerminalNode
	CATCH() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllBlockStmt() []IBlockStmtContext
	BlockStmt(i int) IBlockStmtContext

	// IsTryCatchStmtContext differentiates from other interfaces.
	IsTryCatchStmtContext()
}

type TryCatchStmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	tryBody   IBlockStmtContext
	catchBody IBlockStmtContext
}

func NewEmptyTryCatchStmtContext() *TryCatchStmtContext {
	var p = new(TryCatchStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_tryCatchStmt
	return p
}

func InitEmptyTryCatchStmtContext(p *TryCatchStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_tryCatchStmt
}

func (*TryCatchStmtContext) IsTryCatchStmtContext() {}

func NewTryCatchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TryCatchStmtContext {
	var p = new(TryCatchStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_tryCatchStmt

	return p
}

func (s *TryCatchStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *TryCatchStmtContext) GetTryBody() IBlockStmtContext { return s.tryBody }

func (s *TryCatchStmtContext) GetCatchBody() IBlockStmtContext { return s.catchBody }

func (s *TryCatchStmtContext) SetTryBody(v IBlockStmtContext) { s.tryBody = v }

func (s *TryCatchStmtContext) SetCatchBody(v IBlockStmtContext) { s.catchBody = v }

func (s *TryCatchStmtContext) TRY() antlr.TerminalNode {
	return s.GetToken(GrammarParserTRY, 0)
}

func (s *TryCatchStmtContext) CATCH() antlr.TerminalNode {
	return s.GetToken(GrammarParserCATCH, 0)
}

func (s *TryCatchStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *TryCatchStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *TryCatchStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *TryCatchStmtContext) AllBlockStmt() []IBlockStmtContext {
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

func (s *TryCatchStmtContext) BlockStmt(i int) IBlockStmtContext {
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

func (s *TryCatchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TryCatchStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TryCatchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTryCatchStmt(s)
	}
}

func (s *TryCatchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTryCatchStmt(s)
	}
}

func (s *TryCatchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitTryCatchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) TryCatchStmt() (localctx ITryCatchStmtContext) {
	localctx = NewTryCatchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GrammarParserRULE_tryCatchStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(GrammarParserTRY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)

		var _x = p.BlockStmt()

		localctx.(*TryCatchStmtContext).tryBody = _x
	}
	{
		p.SetState(121)
		p.Match(GrammarParserCATCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(GrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)

		var _x = p.BlockStmt()

		localctx.(*TryCatchStmtContext).catchBody = _x
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

// IThrowStmtContext is an interface to support dynamic dispatch.
type IThrowStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	THROW() antlr.TerminalNode
	Expr() IExprContext

	// IsThrowStmtContext differentiates from other interfaces.
	IsThrowStmtContext()
}

type ThrowStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThrowStmtContext() *ThrowStmtContext {
	var p = new(ThrowStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_throwStmt
	return p
}

func InitEmptyThrowStmtContext(p *ThrowStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_throwStmt
}

func (*ThrowStmtContext) IsThrowStmtContext() {}

func NewThrowStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThrowStmtContext {
	var p = new(ThrowStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_throwStmt

	return p
}

func (s *ThrowStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ThrowStmtContext) THROW() antlr.TerminalNode {
	return s.GetToken(GrammarParserTHROW, 0)
}

func (s *ThrowStmtContext) Expr() IExprContext {
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

func (s *ThrowStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThrowStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThrowStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterThrowStmt(s)
	}
}

func (s *ThrowStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitThrowStmt(s)
	}
}

func (s *ThrowStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitThrowStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ThrowStmt() (localctx IThrowStmtContext) {
	localctx = NewThrowStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GrammarParserRULE_throwStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(GrammarParserTHROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
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

// IBreakStmtContext is an interface to support dynamic dispatch.
type IBreakStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreakStmtContext differentiates from other interfaces.
	IsBreakStmtContext()
}

type BreakStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStmtContext() *BreakStmtContext {
	var p = new(BreakStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_breakStmt
	return p
}

func InitEmptyBreakStmtContext(p *BreakStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_breakStmt
}

func (*BreakStmtContext) IsBreakStmtContext() {}

func NewBreakStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStmtContext {
	var p = new(BreakStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_breakStmt

	return p
}

func (s *BreakStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(GrammarParserBREAK, 0)
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) BreakStmt() (localctx IBreakStmtContext) {
	localctx = NewBreakStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GrammarParserRULE_breakStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(GrammarParserBREAK)
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

// IContinueStmtContext is an interface to support dynamic dispatch.
type IContinueStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinueStmtContext differentiates from other interfaces.
	IsContinueStmtContext()
}

type ContinueStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStmtContext() *ContinueStmtContext {
	var p = new(ContinueStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_continueStmt
	return p
}

func InitEmptyContinueStmtContext(p *ContinueStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_continueStmt
}

func (*ContinueStmtContext) IsContinueStmtContext() {}

func NewContinueStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_continueStmt

	return p
}

func (s *ContinueStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(GrammarParserCONTINUE, 0)
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ContinueStmt() (localctx IContinueStmtContext) {
	localctx = NewContinueStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GrammarParserRULE_continueStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(GrammarParserCONTINUE)
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

// ISwitchStmtContext is an interface to support dynamic dispatch.
type ISwitchStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext
	AllCaseBlock() []ICaseBlockContext
	CaseBlock(i int) ICaseBlockContext
	DefaultBlock() IDefaultBlockContext

	// IsSwitchStmtContext differentiates from other interfaces.
	IsSwitchStmtContext()
}

type SwitchStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStmtContext() *SwitchStmtContext {
	var p = new(SwitchStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_switchStmt
	return p
}

func InitEmptySwitchStmtContext(p *SwitchStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_switchStmt
}

func (*SwitchStmtContext) IsSwitchStmtContext() {}

func NewSwitchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_switchStmt

	return p
}

func (s *SwitchStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStmtContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(GrammarParserSWITCH, 0)
}

func (s *SwitchStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *SwitchStmtContext) Expr() IExprContext {
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

func (s *SwitchStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *SwitchStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACE, 0)
}

func (s *SwitchStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACE, 0)
}

func (s *SwitchStmtContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStmtContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
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

	return t.(ITerminatorContext)
}

func (s *SwitchStmtContext) AllCaseBlock() []ICaseBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseBlockContext); ok {
			len++
		}
	}

	tst := make([]ICaseBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseBlockContext); ok {
			tst[i] = t.(ICaseBlockContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStmtContext) CaseBlock(i int) ICaseBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseBlockContext); ok {
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

	return t.(ICaseBlockContext)
}

func (s *SwitchStmtContext) DefaultBlock() IDefaultBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultBlockContext)
}

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitSwitchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) SwitchStmt() (localctx ISwitchStmtContext) {
	localctx = NewSwitchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GrammarParserRULE_switchStmt)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(GrammarParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.expr(0)
	}
	{
		p.SetState(137)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(GrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(139)
				p.Terminator()
			}

		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserCASE {
		{
			p.SetState(145)
			p.CaseBlock()
		}

		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserDEFAULT {
		{
			p.SetState(151)
			p.DefaultBlock()
		}

	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserSEMICOLON || _la == GrammarParserNL {
		{
			p.SetState(154)
			p.Terminator()
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(160)
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

// ICaseBlockContext is an interface to support dynamic dispatch.
type ICaseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expr() IExprContext
	COLON() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext

	// IsCaseBlockContext differentiates from other interfaces.
	IsCaseBlockContext()
}

type CaseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseBlockContext() *CaseBlockContext {
	var p = new(CaseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_caseBlock
	return p
}

func InitEmptyCaseBlockContext(p *CaseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_caseBlock
}

func (*CaseBlockContext) IsCaseBlockContext() {}

func NewCaseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseBlockContext {
	var p = new(CaseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_caseBlock

	return p
}

func (s *CaseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseBlockContext) CASE() antlr.TerminalNode {
	return s.GetToken(GrammarParserCASE, 0)
}

func (s *CaseBlockContext) Expr() IExprContext {
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

func (s *CaseBlockContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *CaseBlockContext) AllStatement() []IStatementContext {
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

func (s *CaseBlockContext) Statement(i int) IStatementContext {
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

func (s *CaseBlockContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *CaseBlockContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
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

	return t.(ITerminatorContext)
}

func (s *CaseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterCaseBlock(s)
	}
}

func (s *CaseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitCaseBlock(s)
	}
}

func (s *CaseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitCaseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) CaseBlock() (localctx ICaseBlockContext) {
	localctx = NewCaseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GrammarParserRULE_caseBlock)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(GrammarParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.expr(0)
	}
	{
		p.SetState(164)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case GrammarParserT__0, GrammarParserSWITCH, GrammarParserTRY, GrammarParserTHROW, GrammarParserVAR, GrammarParserMINUS, GrammarParserLBRACE, GrammarParserLPAREN, GrammarParserLBRACKET, GrammarParserIF, GrammarParserWHILE, GrammarParserFOR, GrammarParserBREAK, GrammarParserCONTINUE, GrammarParserBITNOT, GrammarParserRETURN, GrammarParserPRINT, GrammarParserTRUE, GrammarParserFALSE, GrammarParserNOT, GrammarParserIDENTIFIER, GrammarParserNUMBER, GrammarParserSTRING:
				{
					p.SetState(165)
					p.Statement()
				}

			case GrammarParserSEMICOLON, GrammarParserNL:
				{
					p.SetState(166)
					p.Terminator()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultBlockContext is an interface to support dynamic dispatch.
type IDefaultBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext

	// IsDefaultBlockContext differentiates from other interfaces.
	IsDefaultBlockContext()
}

type DefaultBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultBlockContext() *DefaultBlockContext {
	var p = new(DefaultBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_defaultBlock
	return p
}

func InitEmptyDefaultBlockContext(p *DefaultBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_defaultBlock
}

func (*DefaultBlockContext) IsDefaultBlockContext() {}

func NewDefaultBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultBlockContext {
	var p = new(DefaultBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_defaultBlock

	return p
}

func (s *DefaultBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultBlockContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(GrammarParserDEFAULT, 0)
}

func (s *DefaultBlockContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *DefaultBlockContext) AllStatement() []IStatementContext {
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

func (s *DefaultBlockContext) Statement(i int) IStatementContext {
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

func (s *DefaultBlockContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *DefaultBlockContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
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

	return t.(ITerminatorContext)
}

func (s *DefaultBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterDefaultBlock(s)
	}
}

func (s *DefaultBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitDefaultBlock(s)
	}
}

func (s *DefaultBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitDefaultBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) DefaultBlock() (localctx IDefaultBlockContext) {
	localctx = NewDefaultBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GrammarParserRULE_defaultBlock)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(GrammarParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(176)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case GrammarParserT__0, GrammarParserSWITCH, GrammarParserTRY, GrammarParserTHROW, GrammarParserVAR, GrammarParserMINUS, GrammarParserLBRACE, GrammarParserLPAREN, GrammarParserLBRACKET, GrammarParserIF, GrammarParserWHILE, GrammarParserFOR, GrammarParserBREAK, GrammarParserCONTINUE, GrammarParserBITNOT, GrammarParserRETURN, GrammarParserPRINT, GrammarParserTRUE, GrammarParserFALSE, GrammarParserNOT, GrammarParserIDENTIFIER, GrammarParserNUMBER, GrammarParserSTRING:
				{
					p.SetState(174)
					p.Statement()
				}

			case GrammarParserSEMICOLON, GrammarParserNL:
				{
					p.SetState(175)
					p.Terminator()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForInStmtContext is an interface to support dynamic dispatch.
type IForInStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetBody returns the body rule contexts.
	GetBody() IBlockStmtContext

	// SetBody sets the body rule contexts.
	SetBody(IBlockStmtContext)

	// Getter signatures
	FOR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	IN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	BlockStmt() IBlockStmtContext
	VAR() antlr.TerminalNode

	// IsForInStmtContext differentiates from other interfaces.
	IsForInStmtContext()
}

type ForInStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	body   IBlockStmtContext
}

func NewEmptyForInStmtContext() *ForInStmtContext {
	var p = new(ForInStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forInStmt
	return p
}

func InitEmptyForInStmtContext(p *ForInStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forInStmt
}

func (*ForInStmtContext) IsForInStmtContext() {}

func NewForInStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInStmtContext {
	var p = new(ForInStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_forInStmt

	return p
}

func (s *ForInStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInStmtContext) GetId() antlr.Token { return s.id }

func (s *ForInStmtContext) SetId(v antlr.Token) { s.id = v }

func (s *ForInStmtContext) GetBody() IBlockStmtContext { return s.body }

func (s *ForInStmtContext) SetBody(v IBlockStmtContext) { s.body = v }

func (s *ForInStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserFOR, 0)
}

func (s *ForInStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *ForInStmtContext) IN() antlr.TerminalNode {
	return s.GetToken(GrammarParserIN, 0)
}

func (s *ForInStmtContext) Expr() IExprContext {
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

func (s *ForInStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *ForInStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *ForInStmtContext) BlockStmt() IBlockStmtContext {
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

func (s *ForInStmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *ForInStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterForInStmt(s)
	}
}

func (s *ForInStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitForInStmt(s)
	}
}

func (s *ForInStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitForInStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ForInStmt() (localctx IForInStmtContext) {
	localctx = NewForInStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GrammarParserRULE_forInStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(GrammarParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserVAR {
		{
			p.SetState(183)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(186)

		var _m = p.Match(GrammarParserIDENTIFIER)

		localctx.(*ForInStmtContext).id = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.Match(GrammarParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.expr(0)
	}
	{
		p.SetState(189)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)

		var _x = p.BlockStmt()

		localctx.(*ForInStmtContext).body = _x
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
	p.EnterRule(localctx, 28, GrammarParserRULE_varDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(GrammarParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.Match(GrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.Match(GrammarParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
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
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	EQUALS() antlr.TerminalNode

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

func (s *AssignStmtContext) AllExpr() []IExprContext {
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

func (s *AssignStmtContext) Expr(i int) IExprContext {
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

func (s *AssignStmtContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUALS, 0)
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
	p.EnterRule(localctx, 30, GrammarParserRULE_assignStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.expr(0)
	}
	{
		p.SetState(198)
		p.Match(GrammarParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
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

// IArrayAssignStmtContext is an interface to support dynamic dispatch.
type IArrayAssignStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LBRACKET() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	RBRACKET() antlr.TerminalNode
	EQUALS() antlr.TerminalNode

	// IsArrayAssignStmtContext differentiates from other interfaces.
	IsArrayAssignStmtContext()
}

type ArrayAssignStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayAssignStmtContext() *ArrayAssignStmtContext {
	var p = new(ArrayAssignStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_arrayAssignStmt
	return p
}

func InitEmptyArrayAssignStmtContext(p *ArrayAssignStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_arrayAssignStmt
}

func (*ArrayAssignStmtContext) IsArrayAssignStmtContext() {}

func NewArrayAssignStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayAssignStmtContext {
	var p = new(ArrayAssignStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_arrayAssignStmt

	return p
}

func (s *ArrayAssignStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayAssignStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *ArrayAssignStmtContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACKET, 0)
}

func (s *ArrayAssignStmtContext) AllExpr() []IExprContext {
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

func (s *ArrayAssignStmtContext) Expr(i int) IExprContext {
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

func (s *ArrayAssignStmtContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACKET, 0)
}

func (s *ArrayAssignStmtContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUALS, 0)
}

func (s *ArrayAssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayAssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterArrayAssignStmt(s)
	}
}

func (s *ArrayAssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitArrayAssignStmt(s)
	}
}

func (s *ArrayAssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitArrayAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ArrayAssignStmt() (localctx IArrayAssignStmtContext) {
	localctx = NewArrayAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GrammarParserRULE_arrayAssignStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(GrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Match(GrammarParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.expr(0)
	}
	{
		p.SetState(204)
		p.Match(GrammarParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(GrammarParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
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
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	PLUSEQUAL() antlr.TerminalNode
	MINUSEQUAL() antlr.TerminalNode
	STAREQUAL() antlr.TerminalNode
	SLASHEQUAL() antlr.TerminalNode
	MODEQUAL() antlr.TerminalNode
	EXPONENTIALEQUAL() antlr.TerminalNode
	BITANDEQUAL() antlr.TerminalNode
	BITOREQUAL() antlr.TerminalNode
	BITXOREQAUL() antlr.TerminalNode
	BITLSHIFTEQUAL() antlr.TerminalNode
	BITRSHIFTEQUAL() antlr.TerminalNode

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

func (s *CompoundAssignStmtContext) AllExpr() []IExprContext {
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

func (s *CompoundAssignStmtContext) Expr(i int) IExprContext {
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

func (s *CompoundAssignStmtContext) BITANDEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITANDEQUAL, 0)
}

func (s *CompoundAssignStmtContext) BITOREQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITOREQUAL, 0)
}

func (s *CompoundAssignStmtContext) BITXOREQAUL() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITXOREQAUL, 0)
}

func (s *CompoundAssignStmtContext) BITLSHIFTEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITLSHIFTEQUAL, 0)
}

func (s *CompoundAssignStmtContext) BITRSHIFTEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITRSHIFTEQUAL, 0)
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
	p.EnterRule(localctx, 34, GrammarParserRULE_compoundAssignStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.expr(0)
	}
	{
		p.SetState(209)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*CompoundAssignStmtContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&558723430724141056) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*CompoundAssignStmtContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(210)
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
	p.EnterRule(localctx, 36, GrammarParserRULE_printStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(GrammarParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
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

	// GetInit returns the init rule contexts.
	GetInit() IIfInitContext

	// GetThenBranch returns the thenBranch rule contexts.
	GetThenBranch() IBlockStmtContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_blockStmt returns the _blockStmt rule contexts.
	Get_blockStmt() IBlockStmtContext

	// GetElseBranch returns the elseBranch rule contexts.
	GetElseBranch() IBlockStmtContext

	// SetInit sets the init rule contexts.
	SetInit(IIfInitContext)

	// SetThenBranch sets the thenBranch rule contexts.
	SetThenBranch(IBlockStmtContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_blockStmt sets the _blockStmt rule contexts.
	Set_blockStmt(IBlockStmtContext)

	// SetElseBranch sets the elseBranch rule contexts.
	SetElseBranch(IBlockStmtContext)

	// GetElifCond returns the elifCond rule context list.
	GetElifCond() []IExprContext

	// GetElifBranch returns the elifBranch rule context list.
	GetElifBranch() []IBlockStmtContext

	// SetElifCond sets the elifCond rule context list.
	SetElifCond([]IExprContext)

	// SetElifBranch sets the elifBranch rule context list.
	SetElifBranch([]IBlockStmtContext)

	// Getter signatures
	IF() antlr.TerminalNode
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	AllBlockStmt() []IBlockStmtContext
	BlockStmt(i int) IBlockStmtContext
	SEMICOLON() antlr.TerminalNode
	AllELIF() []antlr.TerminalNode
	ELIF(i int) antlr.TerminalNode
	ELSE() antlr.TerminalNode
	IfInit() IIfInitContext

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	init       IIfInitContext
	thenBranch IBlockStmtContext
	_expr      IExprContext
	elifCond   []IExprContext
	_blockStmt IBlockStmtContext
	elifBranch []IBlockStmtContext
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

func (s *IfStmtContext) GetInit() IIfInitContext { return s.init }

func (s *IfStmtContext) GetThenBranch() IBlockStmtContext { return s.thenBranch }

func (s *IfStmtContext) Get_expr() IExprContext { return s._expr }

func (s *IfStmtContext) Get_blockStmt() IBlockStmtContext { return s._blockStmt }

func (s *IfStmtContext) GetElseBranch() IBlockStmtContext { return s.elseBranch }

func (s *IfStmtContext) SetInit(v IIfInitContext) { s.init = v }

func (s *IfStmtContext) SetThenBranch(v IBlockStmtContext) { s.thenBranch = v }

func (s *IfStmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *IfStmtContext) Set_blockStmt(v IBlockStmtContext) { s._blockStmt = v }

func (s *IfStmtContext) SetElseBranch(v IBlockStmtContext) { s.elseBranch = v }

func (s *IfStmtContext) GetElifCond() []IExprContext { return s.elifCond }

func (s *IfStmtContext) GetElifBranch() []IBlockStmtContext { return s.elifBranch }

func (s *IfStmtContext) SetElifCond(v []IExprContext) { s.elifCond = v }

func (s *IfStmtContext) SetElifBranch(v []IBlockStmtContext) { s.elifBranch = v }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(GrammarParserIF, 0)
}

func (s *IfStmtContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserLPAREN)
}

func (s *IfStmtContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, i)
}

func (s *IfStmtContext) AllExpr() []IExprContext {
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

func (s *IfStmtContext) Expr(i int) IExprContext {
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

func (s *IfStmtContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserRPAREN)
}

func (s *IfStmtContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, i)
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

func (s *IfStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserSEMICOLON, 0)
}

func (s *IfStmtContext) AllELIF() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserELIF)
}

func (s *IfStmtContext) ELIF(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserELIF, i)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserELSE, 0)
}

func (s *IfStmtContext) IfInit() IIfInitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfInitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfInitContext)
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
	p.EnterRule(localctx, 38, GrammarParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(GrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(217)

			var _x = p.IfInit()

			localctx.(*IfStmtContext).init = _x
		}
		{
			p.SetState(218)
			p.Match(GrammarParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(222)
		p.expr(0)
	}
	{
		p.SetState(223)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)

		var _x = p.BlockStmt()

		localctx.(*IfStmtContext).thenBranch = _x
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserELIF {
		{
			p.SetState(225)
			p.Match(GrammarParserELIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.Match(GrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)

			var _x = p.expr(0)

			localctx.(*IfStmtContext)._expr = _x
		}
		localctx.(*IfStmtContext).elifCond = append(localctx.(*IfStmtContext).elifCond, localctx.(*IfStmtContext)._expr)
		{
			p.SetState(228)
			p.Match(GrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)

			var _x = p.BlockStmt()

			localctx.(*IfStmtContext)._blockStmt = _x
		}
		localctx.(*IfStmtContext).elifBranch = append(localctx.(*IfStmtContext).elifBranch, localctx.(*IfStmtContext)._blockStmt)

		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserELSE {
		{
			p.SetState(236)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)

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
	p.EnterRule(localctx, 40, GrammarParserRULE_whileStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(GrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.expr(0)
	}
	{
		p.SetState(243)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)

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
	p.EnterRule(localctx, 42, GrammarParserRULE_forStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(GrammarParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)

		var _x = p.ForInit()

		localctx.(*ForStmtContext).init = _x
	}
	{
		p.SetState(249)
		p.Match(GrammarParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)

		var _x = p.expr(0)

		localctx.(*ForStmtContext).cond = _x
	}
	{
		p.SetState(251)
		p.Match(GrammarParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(252)

		var _x = p.ForPost()

		localctx.(*ForStmtContext).post = _x
	}
	{
		p.SetState(253)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)

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
	AllFuncStmt() []IFuncStmtContext
	FuncStmt(i int) IFuncStmtContext
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext

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

func (s *BlockStmtContext) AllFuncStmt() []IFuncStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncStmtContext); ok {
			len++
		}
	}

	tst := make([]IFuncStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncStmtContext); ok {
			tst[i] = t.(IFuncStmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockStmtContext) FuncStmt(i int) IFuncStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncStmtContext); ok {
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

	return t.(IFuncStmtContext)
}

func (s *BlockStmtContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *BlockStmtContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
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

	return t.(ITerminatorContext)
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
	p.EnterRule(localctx, 44, GrammarParserRULE_blockStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Match(GrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-567453419208423258) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&79) != 0) {
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case GrammarParserT__0, GrammarParserSWITCH, GrammarParserTRY, GrammarParserTHROW, GrammarParserVAR, GrammarParserMINUS, GrammarParserLBRACE, GrammarParserLPAREN, GrammarParserLBRACKET, GrammarParserIF, GrammarParserWHILE, GrammarParserFOR, GrammarParserBREAK, GrammarParserCONTINUE, GrammarParserBITNOT, GrammarParserRETURN, GrammarParserPRINT, GrammarParserTRUE, GrammarParserFALSE, GrammarParserNOT, GrammarParserIDENTIFIER, GrammarParserNUMBER, GrammarParserSTRING:
			{
				p.SetState(257)
				p.Statement()
			}

		case GrammarParserFUNC:
			{
				p.SetState(258)
				p.FuncStmt()
			}

		case GrammarParserSEMICOLON, GrammarParserNL:
			{
				p.SetState(259)
				p.Terminator()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(265)
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

// IIfInitContext is an interface to support dynamic dispatch.
type IIfInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDecl() IVarDeclContext
	AssignStmt() IAssignStmtContext

	// IsIfInitContext differentiates from other interfaces.
	IsIfInitContext()
}

type IfInitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfInitContext() *IfInitContext {
	var p = new(IfInitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_ifInit
	return p
}

func InitEmptyIfInitContext(p *IfInitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_ifInit
}

func (*IfInitContext) IsIfInitContext() {}

func NewIfInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfInitContext {
	var p = new(IfInitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_ifInit

	return p
}

func (s *IfInitContext) GetParser() antlr.Parser { return s.parser }

func (s *IfInitContext) VarDecl() IVarDeclContext {
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

func (s *IfInitContext) AssignStmt() IAssignStmtContext {
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

func (s *IfInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterIfInit(s)
	}
}

func (s *IfInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitIfInit(s)
	}
}

func (s *IfInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIfInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) IfInit() (localctx IIfInitContext) {
	localctx = NewIfInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GrammarParserRULE_ifInit)
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(267)
			p.VarDecl()
		}

	case GrammarParserT__0, GrammarParserMINUS, GrammarParserLBRACE, GrammarParserLPAREN, GrammarParserLBRACKET, GrammarParserBITNOT, GrammarParserTRUE, GrammarParserFALSE, GrammarParserNOT, GrammarParserIDENTIFIER, GrammarParserNUMBER, GrammarParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(268)
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
	p.EnterRule(localctx, 48, GrammarParserRULE_forInit)
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.VarDecl()
		}

	case GrammarParserT__0, GrammarParserMINUS, GrammarParserLBRACE, GrammarParserLPAREN, GrammarParserLBRACKET, GrammarParserBITNOT, GrammarParserTRUE, GrammarParserFALSE, GrammarParserNOT, GrammarParserIDENTIFIER, GrammarParserNUMBER, GrammarParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(272)
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
	p.EnterRule(localctx, 50, GrammarParserRULE_forPost)
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.AssignStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(276)
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
	Expr() IExprContext
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

func (s *PostfixStmtContext) Expr() IExprContext {
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
	p.EnterRule(localctx, 52, GrammarParserRULE_postfixStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.expr(0)
	}
	{
		p.SetState(280)

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

type NullContext struct {
	ExprContext
}

func NewNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullContext {
	var p = new(NullContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterNull(s)
	}
}

func (s *NullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitNull(s)
	}
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitOrContext struct {
	ExprContext
	op antlr.Token
}

func NewBitOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitOrContext {
	var p = new(BitOrContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BitOrContext) GetOp() antlr.Token { return s.op }

func (s *BitOrContext) SetOp(v antlr.Token) { s.op = v }

func (s *BitOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitOrContext) AllExpr() []IExprContext {
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

func (s *BitOrContext) Expr(i int) IExprContext {
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

func (s *BitOrContext) BITOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITOR, 0)
}

func (s *BitOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBitOr(s)
	}
}

func (s *BitOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBitOr(s)
	}
}

func (s *BitOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBitOr(s)

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

type BitShiftContext struct {
	ExprContext
	op antlr.Token
}

func NewBitShiftContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitShiftContext {
	var p = new(BitShiftContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BitShiftContext) GetOp() antlr.Token { return s.op }

func (s *BitShiftContext) SetOp(v antlr.Token) { s.op = v }

func (s *BitShiftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitShiftContext) AllExpr() []IExprContext {
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

func (s *BitShiftContext) Expr(i int) IExprContext {
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

func (s *BitShiftContext) BITLSHIFT() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITLSHIFT, 0)
}

func (s *BitShiftContext) BITRSHIFT() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITRSHIFT, 0)
}

func (s *BitShiftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBitShift(s)
	}
}

func (s *BitShiftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBitShift(s)
	}
}

func (s *BitShiftContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBitShift(s)

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

type ArrayLiteralContext struct {
	ExprContext
}

func NewArrayLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ArrayLit() IArrayLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLitContext)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayIndexContext struct {
	ExprContext
}

func NewArrayIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayIndexContext {
	var p = new(ArrayIndexContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayIndexContext) AllExpr() []IExprContext {
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

func (s *ArrayIndexContext) Expr(i int) IExprContext {
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

func (s *ArrayIndexContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACKET, 0)
}

func (s *ArrayIndexContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACKET, 0)
}

func (s *ArrayIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterArrayIndex(s)
	}
}

func (s *ArrayIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitArrayIndex(s)
	}
}

func (s *ArrayIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitArrayIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryContext struct {
	ExprContext
	op antlr.Token
}

func NewUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryContext {
	var p = new(UnaryContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnaryContext) GetOp() antlr.Token { return s.op }

func (s *UnaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) Expr() IExprContext {
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

func (s *UnaryContext) NOT() antlr.TerminalNode {
	return s.GetToken(GrammarParserNOT, 0)
}

func (s *UnaryContext) BITNOT() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITNOT, 0)
}

func (s *UnaryContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GrammarParserMINUS, 0)
}

func (s *UnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterUnary(s)
	}
}

func (s *UnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitUnary(s)
	}
}

func (s *UnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

type TernaryOpContext struct {
	ExprContext
	trueExpr  IExprContext
	condExpr  IExprContext
	falseExpr IExprContext
}

func NewTernaryOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryOpContext {
	var p = new(TernaryOpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *TernaryOpContext) GetTrueExpr() IExprContext { return s.trueExpr }

func (s *TernaryOpContext) GetCondExpr() IExprContext { return s.condExpr }

func (s *TernaryOpContext) GetFalseExpr() IExprContext { return s.falseExpr }

func (s *TernaryOpContext) SetTrueExpr(v IExprContext) { s.trueExpr = v }

func (s *TernaryOpContext) SetCondExpr(v IExprContext) { s.condExpr = v }

func (s *TernaryOpContext) SetFalseExpr(v IExprContext) { s.falseExpr = v }

func (s *TernaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryOpContext) IF() antlr.TerminalNode {
	return s.GetToken(GrammarParserIF, 0)
}

func (s *TernaryOpContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserELSE, 0)
}

func (s *TernaryOpContext) AllExpr() []IExprContext {
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

func (s *TernaryOpContext) Expr(i int) IExprContext {
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

func (s *TernaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTernaryOp(s)
	}
}

func (s *TernaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTernaryOp(s)
	}
}

func (s *TernaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitTernaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type MapLiteralContext struct {
	ExprContext
}

func NewMapLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapLiteralContext {
	var p = new(MapLiteralContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MapLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACE, 0)
}

func (s *MapLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACE, 0)
}

func (s *MapLiteralContext) AllMapEntry() []IMapEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMapEntryContext); ok {
			len++
		}
	}

	tst := make([]IMapEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMapEntryContext); ok {
			tst[i] = t.(IMapEntryContext)
			i++
		}
	}

	return tst
}

func (s *MapLiteralContext) MapEntry(i int) IMapEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapEntryContext); ok {
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

	return t.(IMapEntryContext)
}

func (s *MapLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOMMA)
}

func (s *MapLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMMA, i)
}

func (s *MapLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMapLiteral(s)
	}
}

func (s *MapLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMapLiteral(s)
	}
}

func (s *MapLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitMapLiteral(s)

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

type BitXorContext struct {
	ExprContext
	op antlr.Token
}

func NewBitXorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitXorContext {
	var p = new(BitXorContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BitXorContext) GetOp() antlr.Token { return s.op }

func (s *BitXorContext) SetOp(v antlr.Token) { s.op = v }

func (s *BitXorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitXorContext) AllExpr() []IExprContext {
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

func (s *BitXorContext) Expr(i int) IExprContext {
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

func (s *BitXorContext) BITXOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITXOR, 0)
}

func (s *BitXorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBitXor(s)
	}
}

func (s *BitXorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBitXor(s)
	}
}

func (s *BitXorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBitXor(s)

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

type BitAndContext struct {
	ExprContext
	op antlr.Token
}

func NewBitAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitAndContext {
	var p = new(BitAndContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BitAndContext) GetOp() antlr.Token { return s.op }

func (s *BitAndContext) SetOp(v antlr.Token) { s.op = v }

func (s *BitAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitAndContext) AllExpr() []IExprContext {
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

func (s *BitAndContext) Expr(i int) IExprContext {
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

func (s *BitAndContext) BITAND() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITAND, 0)
}

func (s *BitAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBitAnd(s)
	}
}

func (s *BitAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBitAnd(s)
	}
}

func (s *BitAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type MembershipContext struct {
	ExprContext
	leftExpr  IExprContext
	op        IMembershipOpContext
	rightExpr IExprContext
}

func NewMembershipContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MembershipContext {
	var p = new(MembershipContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MembershipContext) GetLeftExpr() IExprContext { return s.leftExpr }

func (s *MembershipContext) GetOp() IMembershipOpContext { return s.op }

func (s *MembershipContext) GetRightExpr() IExprContext { return s.rightExpr }

func (s *MembershipContext) SetLeftExpr(v IExprContext) { s.leftExpr = v }

func (s *MembershipContext) SetOp(v IMembershipOpContext) { s.op = v }

func (s *MembershipContext) SetRightExpr(v IExprContext) { s.rightExpr = v }

func (s *MembershipContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembershipContext) AllExpr() []IExprContext {
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

func (s *MembershipContext) Expr(i int) IExprContext {
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

func (s *MembershipContext) MembershipOp() IMembershipOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMembershipOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMembershipOpContext)
}

func (s *MembershipContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMembership(s)
	}
}

func (s *MembershipContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMembership(s)
	}
}

func (s *MembershipContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitMembership(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallContext struct {
	ExprContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *FunctionCallContext) AllExpr() []IExprContext {
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

func (s *FunctionCallContext) Expr(i int) IExprContext {
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

func (s *FunctionCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOMMA)
}

func (s *FunctionCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMMA, i)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitFunctionCall(s)

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

type FieldAccessContext struct {
	ExprContext
}

func NewFieldAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAccessContext {
	var p = new(FieldAccessContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FieldAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAccessContext) Expr() IExprContext {
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

func (s *FieldAccessContext) DOT() antlr.TerminalNode {
	return s.GetToken(GrammarParserDOT, 0)
}

func (s *FieldAccessContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *FieldAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterFieldAccess(s)
	}
}

func (s *FieldAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitFieldAccess(s)
	}
}

func (s *FieldAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitFieldAccess(s)

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
	_startState := 54
	p.EnterRecursionRule(localctx, 54, GrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(283)
			p.Match(GrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.Match(GrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-4602678819013246974) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&15) != 0) {
			{
				p.SetState(285)
				p.expr(0)
			}
			p.SetState(290)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == GrammarParserCOMMA {
				{
					p.SetState(286)
					p.Match(GrammarParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(287)
					p.expr(0)
				}

				p.SetState(292)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(295)
			p.Match(GrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(296)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-14)) & ^0x3f) == 0 && ((int64(1)<<(_la-14))&4504149383184385) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(297)
			p.expr(21)
		}

	case 3:
		localctx = NewArrayLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(298)
			p.ArrayLit()
		}

	case 4:
		localctx = NewMapLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(299)
			p.Match(GrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-4602678819013246974) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&15) != 0) {
			{
				p.SetState(300)
				p.MapEntry()
			}
			p.SetState(305)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == GrammarParserCOMMA {
				{
					p.SetState(301)
					p.Match(GrammarParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(302)
					p.MapEntry()
				}

				p.SetState(307)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(310)
			p.Match(GrammarParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewIdentifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(311)
			p.Match(GrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(312)
			p.Match(GrammarParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewBooleanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(313)

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

	case 8:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(314)
			p.Match(GrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewNullContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(315)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(316)
			p.Match(GrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(317)
			p.expr(0)
		}
		{
			p.SetState(318)
			p.Match(GrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(370)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExponentialContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(322)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(323)

					var _m = p.Match(GrammarParserEXPONENTIAL)

					localctx.(*ExponentialContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(324)
					p.expr(21)
				}

			case 2:
				localctx = NewMulDivModContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(325)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(326)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1099511726080) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(327)
					p.expr(20)
				}

			case 3:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(328)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(329)

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
					p.SetState(330)
					p.expr(19)
				}

			case 4:
				localctx = NewBitShiftContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(331)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(332)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BitShiftContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserBITLSHIFT || _la == GrammarParserBITRSHIFT) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BitShiftContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(333)
					p.expr(18)
				}

			case 5:
				localctx = NewBitAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(334)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(335)

					var _m = p.Match(GrammarParserBITAND)

					localctx.(*BitAndContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(336)
					p.expr(17)
				}

			case 6:
				localctx = NewBitXorContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(337)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(338)

					var _m = p.Match(GrammarParserBITXOR)

					localctx.(*BitXorContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(339)
					p.expr(16)
				}

			case 7:
				localctx = NewBitOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(340)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(341)

					var _m = p.Match(GrammarParserBITOR)

					localctx.(*BitOrContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(342)
					p.expr(15)
				}

			case 8:
				localctx = NewComparisonContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(343)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(344)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparisonContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8257536) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparisonContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(345)
					p.expr(14)
				}

			case 9:
				localctx = NewMembershipContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MembershipContext).leftExpr = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(346)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(347)

					var _x = p.MembershipOp()

					localctx.(*MembershipContext).op = _x
				}
				{
					p.SetState(348)

					var _x = p.expr(13)

					localctx.(*MembershipContext).rightExpr = _x
				}

			case 10:
				localctx = NewAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(350)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(351)

					var _m = p.Match(GrammarParserAND)

					localctx.(*AndContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(352)
					p.expr(12)
				}

			case 11:
				localctx = NewOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(353)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(354)

					var _m = p.Match(GrammarParserOR)

					localctx.(*OrContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(355)
					p.expr(11)
				}

			case 12:
				localctx = NewTernaryOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*TernaryOpContext).trueExpr = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(356)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(357)
					p.Match(GrammarParserIF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(358)

					var _x = p.expr(0)

					localctx.(*TernaryOpContext).condExpr = _x
				}
				{
					p.SetState(359)
					p.Match(GrammarParserELSE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(360)

					var _x = p.expr(10)

					localctx.(*TernaryOpContext).falseExpr = _x
				}

			case 13:
				localctx = NewArrayIndexContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(362)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
					goto errorExit
				}
				{
					p.SetState(363)
					p.Match(GrammarParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(364)
					p.expr(0)
				}
				{
					p.SetState(365)
					p.Match(GrammarParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 14:
				localctx = NewFieldAccessContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(367)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(368)
					p.Match(GrammarParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(369)
					p.Match(GrammarParserIDENTIFIER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(374)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
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

// IMapEntryContext is an interface to support dynamic dispatch.
type IMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	COLON() antlr.TerminalNode

	// IsMapEntryContext differentiates from other interfaces.
	IsMapEntryContext()
}

type MapEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapEntryContext() *MapEntryContext {
	var p = new(MapEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_mapEntry
	return p
}

func InitEmptyMapEntryContext(p *MapEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_mapEntry
}

func (*MapEntryContext) IsMapEntryContext() {}

func NewMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapEntryContext {
	var p = new(MapEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_mapEntry

	return p
}

func (s *MapEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *MapEntryContext) AllExpr() []IExprContext {
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

func (s *MapEntryContext) Expr(i int) IExprContext {
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

func (s *MapEntryContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *MapEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMapEntry(s)
	}
}

func (s *MapEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMapEntry(s)
	}
}

func (s *MapEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitMapEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) MapEntry() (localctx IMapEntryContext) {
	localctx = NewMapEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GrammarParserRULE_mapEntry)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(375)
		p.expr(0)
	}
	{
		p.SetState(376)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)
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

// IMembershipOpContext is an interface to support dynamic dispatch.
type IMembershipOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IN() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsMembershipOpContext differentiates from other interfaces.
	IsMembershipOpContext()
}

type MembershipOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMembershipOpContext() *MembershipOpContext {
	var p = new(MembershipOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_membershipOp
	return p
}

func InitEmptyMembershipOpContext(p *MembershipOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_membershipOp
}

func (*MembershipOpContext) IsMembershipOpContext() {}

func NewMembershipOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MembershipOpContext {
	var p = new(MembershipOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_membershipOp

	return p
}

func (s *MembershipOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MembershipOpContext) IN() antlr.TerminalNode {
	return s.GetToken(GrammarParserIN, 0)
}

func (s *MembershipOpContext) NOT() antlr.TerminalNode {
	return s.GetToken(GrammarParserNOT, 0)
}

func (s *MembershipOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembershipOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MembershipOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMembershipOp(s)
	}
}

func (s *MembershipOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMembershipOp(s)
	}
}

func (s *MembershipOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitMembershipOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) MembershipOp() (localctx IMembershipOpContext) {
	localctx = NewMembershipOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GrammarParserRULE_membershipOp)
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserIN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(379)
			p.Match(GrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserNOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(380)
			p.Match(GrammarParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.Match(GrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IArrayLitContext is an interface to support dynamic dispatch.
type IArrayLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArrayLitContext differentiates from other interfaces.
	IsArrayLitContext()
}

type ArrayLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLitContext() *ArrayLitContext {
	var p = new(ArrayLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_arrayLit
	return p
}

func InitEmptyArrayLitContext(p *ArrayLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_arrayLit
}

func (*ArrayLitContext) IsArrayLitContext() {}

func NewArrayLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLitContext {
	var p = new(ArrayLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_arrayLit

	return p
}

func (s *ArrayLitContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLitContext) CopyAll(ctx *ArrayLitContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ArrayLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListComprehensionContext struct {
	ArrayLitContext
	transformExpr IExprContext
	id            antlr.Token
	srcExpr       IExprContext
	filterExpr    IExprContext
}

func NewListComprehensionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListComprehensionContext {
	var p = new(ListComprehensionContext)

	InitEmptyArrayLitContext(&p.ArrayLitContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArrayLitContext))

	return p
}

func (s *ListComprehensionContext) GetId() antlr.Token { return s.id }

func (s *ListComprehensionContext) SetId(v antlr.Token) { s.id = v }

func (s *ListComprehensionContext) GetTransformExpr() IExprContext { return s.transformExpr }

func (s *ListComprehensionContext) GetSrcExpr() IExprContext { return s.srcExpr }

func (s *ListComprehensionContext) GetFilterExpr() IExprContext { return s.filterExpr }

func (s *ListComprehensionContext) SetTransformExpr(v IExprContext) { s.transformExpr = v }

func (s *ListComprehensionContext) SetSrcExpr(v IExprContext) { s.srcExpr = v }

func (s *ListComprehensionContext) SetFilterExpr(v IExprContext) { s.filterExpr = v }

func (s *ListComprehensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListComprehensionContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACKET, 0)
}

func (s *ListComprehensionContext) FOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserFOR, 0)
}

func (s *ListComprehensionContext) IN() antlr.TerminalNode {
	return s.GetToken(GrammarParserIN, 0)
}

func (s *ListComprehensionContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACKET, 0)
}

func (s *ListComprehensionContext) AllExpr() []IExprContext {
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

func (s *ListComprehensionContext) Expr(i int) IExprContext {
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

func (s *ListComprehensionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *ListComprehensionContext) IF() antlr.TerminalNode {
	return s.GetToken(GrammarParserIF, 0)
}

func (s *ListComprehensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterListComprehension(s)
	}
}

func (s *ListComprehensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitListComprehension(s)
	}
}

func (s *ListComprehensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitListComprehension(s)

	default:
		return t.VisitChildren(s)
	}
}

type StandardArrayContext struct {
	ArrayLitContext
}

func NewStandardArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StandardArrayContext {
	var p = new(StandardArrayContext)

	InitEmptyArrayLitContext(&p.ArrayLitContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArrayLitContext))

	return p
}

func (s *StandardArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StandardArrayContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACKET, 0)
}

func (s *StandardArrayContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACKET, 0)
}

func (s *StandardArrayContext) AllExpr() []IExprContext {
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

func (s *StandardArrayContext) Expr(i int) IExprContext {
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

func (s *StandardArrayContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOMMA)
}

func (s *StandardArrayContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMMA, i)
}

func (s *StandardArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterStandardArray(s)
	}
}

func (s *StandardArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitStandardArray(s)
	}
}

func (s *StandardArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStandardArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ArrayLit() (localctx IArrayLitContext) {
	localctx = NewArrayLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GrammarParserRULE_arrayLit)
	var _la int

	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStandardArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(384)
			p.Match(GrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-4602678819013246974) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&15) != 0) {
			{
				p.SetState(385)
				p.expr(0)
			}
			p.SetState(390)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == GrammarParserCOMMA {
				{
					p.SetState(386)
					p.Match(GrammarParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(387)
					p.expr(0)
				}

				p.SetState(392)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(395)
			p.Match(GrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewListComprehensionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(396)
			p.Match(GrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)

			var _x = p.expr(0)

			localctx.(*ListComprehensionContext).transformExpr = _x
		}
		{
			p.SetState(398)
			p.Match(GrammarParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)

			var _m = p.Match(GrammarParserIDENTIFIER)

			localctx.(*ListComprehensionContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)
			p.Match(GrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)

			var _x = p.expr(0)

			localctx.(*ListComprehensionContext).srcExpr = _x
		}
		p.SetState(404)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserIF {
			{
				p.SetState(402)
				p.Match(GrammarParserIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(403)

				var _x = p.expr(0)

				localctx.(*ListComprehensionContext).filterExpr = _x
			}

		}
		{
			p.SetState(406)
			p.Match(GrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

func (p *GrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 27:
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
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 23)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
